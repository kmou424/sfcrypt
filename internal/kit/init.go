package kit

import (
	"fmt"
)

var hashing = map[string]int{}

func init() {
	cnt := 0
	for i := 0; i <= 9; i++ {
		hashing[fmt.Sprintf("%d", i)] = cnt
		cnt++
	}
	for i := 'a'; i <= 'z'; i++ {
		hashing[string(i)] = cnt
		cnt++
	}
}
