package v2

import (
	"bytes"
	. "github.com/kmou424/sfcrypt/app/common"
	"io"
)

func (c *SFCipher) decryptPreprocess() (err error) {
	header := &SFHeader{}
	n, err := header.ReadFromFile(c.fIn, true)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil || (err != nil && c.opt.Force) {
			c.headerSize = int64(n)
		}
	}()

	if !bytes.Equal(DefHeader.Magic[:], header.Magic[:]) {
		return Errorf("file is not a sfcrypt encrypted file")
	}

	// only version mismatch error can be force ignored on decrypting
	defer c.ignorePreprocessErrorsOnForceEnabled(&err)

	err = isHeaderVersionMatched(header)
	if err != nil {
		return err
	}

	return nil
}

func (c *SFCipher) decryptDoWithOffset(offset int64) (eof bool, err error) {
	buf := make([]byte, c.blkSize)

	length, err := c.fIn.ReadAt(buf, offset+c.headerSize)
	if err != nil && err != io.EOF {
		err = ErrorfCaused("unexpected error while reading file", err)
		return
	}
	if err == io.EOF {
		eof = true
	}

	out, err := c.opt.Cipher.Decrypt(buf)
	if err != nil {
		err = Errorf("process bytes %d - %d error: %v", offset, offset+int64(length), err)
		return
	}
	copy(buf, out)

	_, err = c.fOut.WriteAt(buf[:length], offset)
	if err != nil {
		err = ErrorfCaused("unexpected error while writing file", err)
		return
	}

	return
}
