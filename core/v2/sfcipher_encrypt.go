package v2

import (
	"bytes"
	"github.com/kmou424/ero"
	. "github.com/kmou424/sfcrypt/app/common"
	"io"
)

func (c *SFCipher) encryptPreprocess() (err error) {
	// all preprocess errors can be force ignored on encrypting
	defer func() {
		if err != nil && c.opt.Force {
			Logger.Warn(ero.LineTrace(err))
			err = nil
		}
	}()

	header := &SFHeader{}
	_, err = header.ReadFromFile(c.fIn, false)
	// can't parse file header, not encrypted file
	if err != nil {
		// fallback offset to ensure read from start next time
		// check fallback error here to avoid missing hit
		_, err := c.fIn.Seek(0, io.SeekStart)
		if err != nil {
			return ero.Wrap(err, "unable to seek file header: %v")
		}
		return nil
	}

	defer func() {
		if err == nil || (err != nil && c.opt.Force) {
			// write encrypted file header
			var n int
			n, err = DefHeader.WriteToFile(c.fOut)
			if err != nil {
				err = ero.Wrap(err, "write file header error")
			}
			c.headerSize = int64(n)
		}
	}()

	// not encrypted file, skip
	if bytes.Equal(DefHeader.Magic[:], header.Magic[:]) {
		return ero.Wrap(err, "input file is already encrypted")
	}

	return nil
}

func (c *SFCipher) encryptFragment(offset int64) (eof bool, err error) {
	buf := make([]byte, c.blkSize)
	length, err := c.fIn.ReadAt(buf, offset)
	if err != nil && err != io.EOF {
		err = ero.Wrap(err, "unexpected error while reading file")
		return
	}
	if err == io.EOF {
		eof = true
	}

	err = c.opt.Cipher.Encrypt(buf)
	if err != nil {
		err = ero.Newf("process bytes %d - %d error: %v", offset, offset+int64(length), err)
		return
	}

	_, err = c.fOut.WriteAt(buf[:length], offset+c.headerSize)
	if err != nil {
		err = ero.Wrap(err, "unexpected error while writing file")
		return
	}
	return
}
