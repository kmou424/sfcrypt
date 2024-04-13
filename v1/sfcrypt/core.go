package sfcrypt

import (
	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/sfcrypt/internal/kit"
	"io"
	"os"
	"runtime"
	"sync"
)

func doSFCrypt(input string, output string, blockSize int, password string, threads int) {
	var blockNum int
	var wg sync.WaitGroup
	var eof bool

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

	for !eof {
		for runtime.NumGoroutine()-1 >= threads {
			runtime.Gosched()
		}

		wg.Add(1)
		go func(blockNum int) {
			defer wg.Done()

			offset := blockNum * blockSize

			buf := make([]byte, blockSize)
			length, err := inputFile.ReadAt(buf, int64(offset))
			if err != nil && err != io.EOF {
				kit.Panic(err.Error())
			}
			if err == io.EOF {
				eof = true
			}

			var bufCrypt []byte
			bufCrypt = kit.XORCryptBytes(buf, length, password)

			_, err = outputFile.WriteAt(bufCrypt[:length], int64(offset))
			if err != nil {
				kit.Panic(err.Error())
			}
		}(blockNum)

		blockNum++
	}

	wg.Wait()
}
