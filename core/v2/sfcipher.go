package v2

import (
	"os"
	"sync"

	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/kmou424/sfcrypt/core/cipher"
)

type SFCipher struct {
	opt *cipher.FileCipherOptions

	wg             sync.WaitGroup
	maxRoutineCtrl chan any
	blkSize        int64
	headerSize     int64

	fIn  *os.File
	fOut *os.File
}

func (c *SFCipher) Init(opt *cipher.FileCipherOptions) {
	c.opt = opt
	c.maxRoutineCtrl = make(chan any, MaxRoutines)
	c.blkSize = calcBlockSize()
}

func (c *SFCipher) processFileInParallel(preprocess func() error, doWithOffset func(int64) (bool, error)) (err error) {
	if c.opt.Input == c.opt.Output {
		return Errorf("input file is same as output file")
	}

	c.fIn, err = os.OpenFile(c.opt.Input, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = c.fIn.Close() }()
	// todo: check exists
	// todo: try create file first
	c.fOut, err = os.OpenFile(c.opt.Output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = c.fOut.Close()
		// cleanup broken output file
		// todo: file can be deleted when it was created by ourselves
		if err != nil {
			_ = os.Remove(c.opt.Output)
			return
		}
	}()

	if preprocess != nil {
		err := preprocess()
		if err != nil {
			return err
		}
	}

	var (
		blk int64 = 0
		eof       = false
	)

	syncDoAlignedOffset := func(offset int64) {
		defer func() {
			<-c.maxRoutineCtrl
			c.wg.Done()
		}()

		isEof, err := doWithOffset(offset)
		if err != nil {
			eof = true
			Logger.Error(err.Error())
			return
		}
		if isEof {
			eof = true
		}
	}

	for ; !eof; blk++ {
		c.wg.Add(1)
		c.maxRoutineCtrl <- nil
		go syncDoAlignedOffset(blk * c.blkSize)
	}

	c.wg.Wait()

	return nil
}

func (c *SFCipher) Encrypt() error {
	err := c.processFileInParallel(
		c.encryptPreprocess,
		c.encryptDoWithOffset,
	)
	if err != nil {
		return ErrorfCaused("encryption error", err)
	}
	return nil
}

func (c *SFCipher) Decrypt() error {
	err := c.processFileInParallel(
		c.decryptPreprocess,
		c.decryptDoWithOffset,
	)
	if err != nil {
		return ErrorfCaused("decryption error", err)
	}
	return nil
}

func calcBlockSize() (blkSize int64) {
	blkSize = BufferSize * BlockRatio

	if blkSize%pbkdf2KeySize != 0 {
		// adjust block size to be the least common multiple of pbkdf2KeySize and BufferSize
		blkSize = blkSize * pbkdf2KeySize / GCD(blkSize, pbkdf2KeySize)
	}

	return
}
