package sub

import (
	"time"
)

// CheckPrime check weather given input is prime number
func CheckPrime(num int) bool {
	// for illustrate power of parallel test.
	time.Sleep(1 * time.Millisecond)

	if num > 1 {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
