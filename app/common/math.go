package common

func GCD[T int | int64 | int32 | int16 | int8](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
