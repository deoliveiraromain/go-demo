package main

import "strconv"

const fizz = "FIZZ"

const buzz = "BUZZ"

//Compute the fizzbuzz value of an integer :
// Fizz if %3=0, Buzz if %5=0, FizzBuzz if both
func fizzbuzz(value int) string {
	isfizz := value%3 == 0
	isbuzz := value%5 == 0

	if isfizz && isbuzz {
		return fizz + buzz
	}
	if isfizz {
		return fizz
	}
	if isbuzz {
		return buzz
	}
	return strconv.Itoa(value)
}

func main() {
}
