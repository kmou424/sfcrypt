package v1

import . "github.com/kmou424/sfcrypt/app/common"

func Panic(format string, a ...any) {
	err := Errorf(format, a...)
	panic(err)
}
