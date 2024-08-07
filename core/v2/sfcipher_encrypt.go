package v2

import (
	"bytes"
	. "github.com/kmou424/sfcrypt/app/common"
	"io"
)

func (c *SFCipher) encryptPreprocess() (err error) {
	// all preprocess errors can be force ignored on encrypting
	defer c.ignorePreprocessErrorsOnForceEnabled(&err)

	header := &SFHeader{}
	_, err = header.ReadFromFile(c.fIn, false)
	// can't parse file header, not encrypted file
	if err != nil {
		// fallback offset to ensure read from start next time
		// check fallback error here to avoid missing hit
		_, err := c.fIn.Seek(0, io.SeekStart)
		if err != nil {
			return ErrorfCaused("unable to seek file header: %v", err)
		}
		return nil
	}

	defer func() {
		if err == nil || (err != nil && c.opt.Force) {
			// write encrypted file header
			var n int
			n, err = DefHeader.WriteToFile(c.fOut)
			if err != nil {
				err = ErrorfCaused("write file header error", err)
			}
			c.headerSize = int64(n)
		}
	}()

	// not encrypted file, skip
	if !bytes.Equal(DefHeader.Magic[:], header.Magic[:]) {
		return nil
	}

	err = isHeaderVersionMatched(header)
	if err != nil {
		return err
	}

	return Errorf("input file is already encrypted")
}

func (c *SFCipher) encryptDoWithOffset(offset int64) (eof bool, err error) {
	buf := make([]byte, c.blkSize)
	length, err := c.fIn.ReadAt(buf, offset)
	if err != nil && err != io.EOF {
		err = ErrorfCaused("unexpected error while reading file", err)
		return
	}
	if err == io.EOF {
		eof = true
	}

	out, err := c.opt.Cipher.Encrypt(buf)
	if err != nil {
		err = Errorf("process bytes %d - %d error: %v", offset, offset+int64(length), err)
		return
	}
	copy(buf, out)

	_, err = c.fOut.WriteAt(buf[:length], offset+c.headerSize)
	if err != nil {
		err = ErrorfCaused("unexpected error while writing file", err)
		return
	}
	return
}
