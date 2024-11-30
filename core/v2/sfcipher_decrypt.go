package v2

import (
	"bytes"
	"io"

	"github.com/kmou424/ero"
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

			// assign error to nil if force option is enabled
			err = nil
		}
	}()

	if !bytes.Equal(DefHeader.Magic[:], header.Magic[:]) {
		return ero.Newf("file is not a sfcrypt encrypted file")
	}

	if err := isHeaderVersionMatched(header); err != nil {
		if !c.opt.Force {
			return ero.Wrap(err, "header version mismatch")
		}
	}

	return nil
}

func (c *SFCipher) decryptFragment(offset int64) (eof bool, err error) {
	buf := make([]byte, c.blkSize)

	length, err := c.fIn.ReadAt(buf, offset+c.headerSize)
	if err != nil && err != io.EOF {
		err = ero.Wrap(err, "unexpected error while reading file")
		return
	}
	if err == io.EOF {
		eof = true
	}

	err = c.opt.Cipher.Decrypt(buf)
	if err != nil {
		err = ero.Newf("process bytes %d - %d error: %v", offset, offset+int64(length), err)
		return
	}

	_, err = c.fOut.WriteAt(buf[:length], offset)
	if err != nil {
		err = ero.Wrap(err, "unexpected error while writing file")
		return
	}

	return
}
