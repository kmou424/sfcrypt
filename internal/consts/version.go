package consts

import "fmt"

var (
	MajorVersion = 1
	MinorVersion = 0
	PatchVersion = 0
)

func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", MajorVersion, MinorVersion, PatchVersion)
}
