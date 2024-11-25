package v2

import (
	"github.com/kmou424/ero"
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

type (
	preprocessFunc      func() error
	resolveFragmentFunc func(int64) (bool, error)
)

func (c *SFCipher) doFileCiphering(preprocess preprocessFunc, resolveFragment resolveFragmentFunc) (err error) {
	if c.opt.Input == c.opt.Output {
		return ero.New("input file is same as output file")
	}

	c.fIn, err = os.OpenFile(c.opt.Input, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = c.fIn.Close() }()

	if _, err = os.Stat(c.opt.Output); err == nil {
		return ero.Newf("output file already exists: %s", c.opt.Output)
	}

	c.fOut, err = os.OpenFile(c.opt.Output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = c.fOut.Close()
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

	syncResolveFragment := func(offset int64) {
		defer func() {
			<-c.maxRoutineCtrl
			c.wg.Done()
		}()

		isEof, err := resolveFragment(offset)
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
		go syncResolveFragment(blk * c.blkSize)
	}

	c.wg.Wait()

	return nil
}

func (c *SFCipher) Encrypt() error {
	err := c.doFileCiphering(
		c.encryptPreprocess,
		c.encryptFragment,
	)
	if err != nil {
		return ero.Wrap(err, "encryption error")
	}
	return nil
}

func (c *SFCipher) Decrypt() error {
	err := c.doFileCiphering(
		c.decryptPreprocess,
		c.decryptFragment,
	)
	if err != nil {
		return ero.Wrap(err, "decryption error")
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
