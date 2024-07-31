package common

import "fmt"

func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

func ErrorfCaused(format string, cause error, args ...any) error {
	if cause != nil {
		return Errorf(fmt.Sprintf("%s \n\tCaused by: %s", Errorf(format, args...), cause))
	}
	return Errorf(format, args...)
}
