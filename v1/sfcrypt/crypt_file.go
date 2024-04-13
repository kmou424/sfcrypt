package sfcrypt

import (
	"context"
	_ "context"
	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/sfcrypt/v1/sfcrypt/kit"
	"io"
	"os"
	"sync"
)

func doSFCryptBak(input string, output string, blockSize int, password string, threads int) {
	var wg sync.WaitGroup

	var (
		inputFile  *os.File
		outputFile *os.File
	)

	inputFile, err := fsutil.OpenReadFile(input)
	if err != nil {
		kit.Panic(err.Error())
	}
	defer inputFile.Close()

	if input == output {
		outputFile, err = fsutil.OpenFile(output, fsutil.FsCWFlags, 0755)
	} else {
		outputFile, err = fsutil.QuickOpenFile(output, 0755)
	}
	if err != nil {
		kit.Panic(err.Error())
	}
	defer outputFile.Close()

	var eof = false

	for block := 0; eof; block++ {
		wg.Add(1)
		go func(block int) {
			defer wg.Done()

			offset := block * blockSize

			buf := make([]byte, blockSize)
			length, err := inputFile.ReadAt(buf, int64(offset))
			if err != nil && err != io.EOF {
				kit.Panic(err.Error())
			}
			if err == io.EOF {
				eof = true
			}

			xorCryptBytes(buf, password)

			_, err = outputFile.WriteAt(buf[:length], int64(offset))
			if err != nil {
				kit.Panic(err.Error())
			}
		}(block)
	}

	wg.Wait()
}

func (s *SFCrypt) CryptFile(input string, output string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(error); ok && r != nil {
				err = r
			}
		}
	}()
	crypt := &fileCrypt{
		input:     input,
		output:    output,
		blockSize: BufferSize * s.blockRatio,
		password:  s.password,
		threads:   s.threads,
	}
	crypt.start()
	return
}

type fileCrypt struct {
	input     string
	output    string
	blockSize int
	password  string
	threads   int

	inputFile  *os.File
	outputFile *os.File
}

func (c *fileCrypt) open() {
	var err error
	c.inputFile, err = fsutil.OpenReadFile(c.input)
	if err != nil {
		kit.Panic(err.Error())
	}

	if c.input == c.output {
		c.outputFile, err = fsutil.OpenFile(c.output, fsutil.FsCWFlags, 0755)
	} else {
		c.outputFile, err = fsutil.QuickOpenFile(c.output, 0755)
	}
	if err != nil {
		_ = c.inputFile.Close()
		kit.Panic(err.Error())
	}
}

func (c *fileCrypt) close() {
	_ = c.inputFile.Close()
	_ = c.outputFile.Close()
}

func (c *fileCrypt) start() {
	c.open()
	defer c.close()
	var wg sync.WaitGroup

	var blocks = make(chan int, c.threads)

	ctx, cancel := context.WithCancel(context.Background())
	callback := make(chan bool, c.threads)
	defer close(callback)

	cryptBlock := func(block int) {
		offset := int64(block * c.blockSize)

		buf := make([]byte, c.blockSize)
		length, err := c.inputFile.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			kit.Panic(err.Error())
		}
		callback <- err != io.EOF

		xorCryptBytes(buf, c.password)

		_, err = c.outputFile.WriteAt(buf[:length], offset)
		if err != nil {
			kit.Panic(err.Error())
		}
	}

	for i := 0; i < c.threads; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				case block := <-blocks:
					cryptBlock(block)
				}
			}
		}()
	}

	for block := 0; ; block++ {
		blocks <- block
		if !<-callback {
			break
		}
	}

	cancel()

	wg.Wait()
}
