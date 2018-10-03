package main

import (
	"fmt"
)

func main() {
	fmt.Println(size(10))
}

// size convert size into string based on conditions.
func size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}
