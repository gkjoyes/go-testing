package sub

const failed = "\u2717"

// test cases
var tests = []struct {
	input   int
	isPrime bool
}{
	{
		input:   23,
		isPrime: true,
	},
	{
		input:   50,
		isPrime: false,
	},
	{
		input:   25,
		isPrime: false,
	},
	{
		input:   13,
		isPrime: true,
	},
}
