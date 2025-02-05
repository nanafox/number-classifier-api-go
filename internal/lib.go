// This file contains useful functions used in the API responses

package internal

import (
	"math"
	"strconv"
)

// IsEven returns true if the `number` is even, else false.
func IsEven(number int) bool {
	return number&1 == 0
}

// IsPerfect returns true or false based on whether is perfect or not,
// respectively.
func IsPerfect(number int) bool {
	if !IsEven(number) {
		return false
	}

	return isPerfect(number)
}

// isPerfect is a helper function to compute and return true or false whether a
// number is perfect or not.
func isPerfect(number int) bool {
	total := 1
	divisor := 2

	for divisor*divisor <= number {
		if number%divisor == 0 {
			total += divisor + number/divisor
		}

		divisor += 1
	}

	return total == number
}

// DigitSum returns the digit sum of the provided `number`.
//
// The digit sum of a number is the sum of all its digits.
func DigitSum(number int) (sum int) {
	sum = 0

	for number != 0 {
		last := number % 10

		sum += last
		number /= 10
	}

	return
}

// IsPrime returns true if the `number` is prime, false otherwise.
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number == 2 || number == 3 {
		return true
	}

	if IsEven(number) || number%3 == 0 {
		return false
	}

	divisor := 5

	for divisor <= int(math.Sqrt(float64(number))) {
		if number%divisor == 0 || number%(divisor+2) == 0 {
			return false
		}

		divisor += 5
	}

	return true
}

// IsArmstrong returns true if a number an Armstrong, false otherwise.
func IsArmstrong(number int) bool {
	sum := 0
	numberLen := len(strconv.Itoa(number))

	tempNum := number

	for tempNum != 0 {
		last := tempNum % 10
		sum += int(math.Pow(float64(last), float64(numberLen)))
		tempNum = tempNum / 10
	}

	return sum == number
}

// NumberProperties returns the properites of a number.
//
// The properties included are whether the number is odd, even or armstrong. A
// number could be odd and armstrong.
//
// The result is returned an array of strings.
func NumberProperties(number int) (properites []string) {
	properites = make([]string, 0)

	if IsEven(number) {
		properites = append(properites, "even")
	} else {
		properites = append(properites, "odd")
	}

	if IsArmstrong(number) {
		properites = append(properites, "armstrong")
	}

	return
}
