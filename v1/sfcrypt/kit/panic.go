package kit

import "fmt"

func Panic(format string, a ...any) {
	err := fmt.Errorf(format, a...)
	panic(err)
}
