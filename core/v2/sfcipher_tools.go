package v2

import (
	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/kmou424/sfcrypt/app/version"
)

func isHeaderVersionMatched(header *SFHeader) error {
	if ok, err := version.Equal(string(header.Version[:]), version.VMajor|version.VMinor); !ok {
		if err != nil {
			return Errorf(
				"the file is already encrypted with unknown version, maybe not compatible with this version %s",
				string(DefHeader.Version[:]),
			)
		}
		return Errorf(
			"the file is already encrypted with version %s, maybe not compatible with this version %s",
			string(header.Version[:]), string(DefHeader.Version[:]),
		)
	}
	return nil
}
