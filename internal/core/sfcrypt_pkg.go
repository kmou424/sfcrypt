package core

import (
	"github.com/gookit/goutil/fsutil"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/sfcrypt/internal/types"
	"github.com/kmou424/sfcrypt/internal/utils"
	"io"
	"runtime"
	"sync"
)

func SFCryptFile(input string, output string, blockSize int, password string, threads int) {
	var blockNum int
	var wg sync.WaitGroup
	var eof bool

	inputFile, err := fsutil.OpenReadFile(input)
	exceptiongo.QuickThrow[types.IOException](err)
	defer inputFile.Close()

	outputFile, err := fsutil.QuickOpenFile(output, 0755)
	exceptiongo.QuickThrow[types.IOException](err)
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
				exceptiongo.QuickThrow[types.IOException](err)
			}
			if err == io.EOF {
				eof = true
			}

			var bufCrypt []byte
			bufCrypt = utils.XORCryptBytes(buf, length, password)

			_, err = outputFile.WriteAt(bufCrypt[:length], int64(offset))
			if err != nil {
				exceptiongo.QuickThrow[types.IOException](err)
			}
		}(blockNum)

		blockNum++
	}

	wg.Wait()
}
