package version

import (
	"fmt"
	"strings"
)

var (
	major       int16
	minor       int16
	patch       int16
	initialized bool
)

func InitVersion(m, n, p int16) {
	if initialized {
		return
	}
	defer func() {
		initialized = true
	}()
	major = m
	minor = n
	patch = p
}

type Flag int

const (
	VMajor Flag = 1 << iota
	VMinor
	VPatch
)

func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func Equal(version string, flags Flag) (bool, error) {
	embedV := strings.Split(GetVersion(), ".")
	givenV := strings.Split(version, ".")
	if len(embedV) != 3 || len(givenV) != 3 {
		return false, fmt.Errorf("invalid version format")
	}

	if flags&VMajor == VMajor {
		if embedV[0] != givenV[0] {
			return false, nil
		}
	}

	if flags&VMinor == VMinor {
		if embedV[1] != givenV[1] {
			return false, nil
		}
	}

	if flags&VPatch == VPatch {
		if embedV[2] != givenV[2] {
			return false, nil
		}
	}

	return true, nil
}
