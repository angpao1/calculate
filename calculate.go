package fizzbuzz_tdd_go

import "strconv"

func plusOneAndOne(x int, y int) string {

	z := 0
	if x+y == 2 {
		z = x + y
	}
	return strconv.Itoa(z)
}
