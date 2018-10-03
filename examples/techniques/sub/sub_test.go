package sub

import (
	"fmt"
	"testing"
)

// TestPrimeSequential tests given sets of the number are prime in sequential
func TestPrimeSequential(t *testing.T) {
	// run each test case
	for i, test := range tests {
		st := func(t *testing.T) {
			t.Logf("\tTest: %d\t Check weather given number %d is prime", i, test.input)
			{
				isPrime := CheckPrime(test.input)
				if isPrime != test.isPrime {
					t.Errorf("\t%s\tCould not get expected result", failed)
				}
			}
		}
		t.Run(fmt.Sprintf("check prime"), st)
	}
}
