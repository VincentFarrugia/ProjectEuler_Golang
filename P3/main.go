/*------------------------------------------------------------------------------------------------*
PROJECT EULER:
Problem ID: 3
Title: Largest Prime Factor
Description:
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
Code Author: Vincent Farrugia
*------------------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math"
)

// ***********************
// ***********************
// ***********************
// ***********************
// WIP CODE. NOT READY YET.
// ***********************
// ***********************
// ***********************
// ***********************

func main() {
	findHighestCommonFactor(600851475143)
}

func findHighestCommonFactor(inputNum int) {
	factorsList := findFactorsOf(inputNum)
	fmt.Printf("Factors of %d are: %v\n", inputNum, factorsList)
}

func findFactorsOf(inputNum int) []int {
	retFactorList := make([]int, 0, 10)
	halfInputNum := int(inputNum / 2)
	for k := 2; k < halfInputNum; k++ {
		if inputNum%k == 0 {
			retFactorList = append(retFactorList, k)
		}
	}
	return retFactorList
}

func isPrimeNumber(inputNum int) bool {

	// Segmented version of the Sieve of Eratosthenes
	// in order to reduce space complexity for large Limits.

	limit := inputNum + 1
	cacheSize := 32768
	sqrt := int(math.Sqrt(float64(limit)))
	segmentSize := maxInt(sqrt, cacheSize)

	i := 3
	n := 3
	s := 3

	sieve := make([]bool, segmentSize)
	primeFlagList := make([]bool, sqrt+1)
	for k := 0; k < len(primeFlagList); k++ {
		primeFlagList[k] = true
	}
	primes := make([]int, 0, 10)
	multiples := make([]int, 0, 10)

	//retLargestPrimeFactor := -1

	//if limit > 2 {
	//	retPrimes = append(retPrimes, 2)
	//}

	for low := 0; low <= limit; low += segmentSize {

		for k := 0; k < len(sieve); k++ {
			sieve[k] = true
		}

		// Current segment = [low, high]
		high := low + segmentSize - 1
		high = minInt(high, limit)

		// Generate sieving primes using basic sieve of Eratosthenes
		for i*i <= high {
			if primeFlagList[i] {
				for j := i * i; j <= sqrt; j += i {
					primeFlagList[j] = false
				}
			}
			i += 2
		}

		// Initialise sieving primes for segmented sieve
		for s*s <= high {
			if primeFlagList[s] {
				primes = append(primes, s)
				multiples = append(multiples, ((s * s) - low))
			}
			s += 2
		}

		// Sieve the current
		numPrimesSoFar := len(primes)
		for k := 0; k < numPrimesSoFar; k++ {
			j := multiples[k]
			for m := primes[k] * 2; j < segmentSize; j += m {
				sieve[j] = false
			}
			multiples[k] = j - segmentSize
		}

		for n <= high {
			if sieve[n-low] {
				// n is prime
				//retPrimes = append(retPrimes, n)
				if n == inputNum {
					return true
				}
			}
			n += 2
		}
	}

	return false
}

///////////////////////////////////////////////////////////////////////////////
// OPTIMISED ANSWER V1 (Using Segmented Sieve of Eratosthenes)
// Still slow for 600851475143. I might consider using multiple threads next.
///////////////////////////////////////////////////////////////////////////////

func optimisedAttempt() {
	limit := 1000000001 //600851475143
	largestPrimeFactor := findLargestPrimeFactor(limit)
	fmt.Printf("The largest prime factor of %d is: %d\n", limit, largestPrimeFactor)
}

func findLargestPrimeFactor(inputNum int) int {

	// Segmented version of the Sieve of Eratosthenes
	// in order to reduce space complexity for large Limits.

	limit := inputNum
	cacheSize := 32768
	sqrt := int(math.Sqrt(float64(limit)))
	segmentSize := maxInt(sqrt, cacheSize)

	i := 3
	n := 3
	s := 3

	sieve := make([]bool, segmentSize)
	primeFlagList := make([]bool, sqrt+1)
	for k := 0; k < len(primeFlagList); k++ {
		primeFlagList[k] = true
	}
	primes := make([]int, 0, 10)
	multiples := make([]int, 0, 10)

	retLargestPrimeFactor := -1

	//if limit > 2 {
	//	retPrimes = append(retPrimes, 2)
	//}

	for low := 0; low <= limit; low += segmentSize {

		for k := 0; k < len(sieve); k++ {
			sieve[k] = true
		}

		// Current segment = [low, high]
		high := low + segmentSize - 1
		high = minInt(high, limit)

		// Generate sieving primes using basic sieve of Eratosthenes
		for i*i <= high {
			if primeFlagList[i] {
				for j := i * i; j <= sqrt; j += i {
					primeFlagList[j] = false
				}
			}
			i += 2
		}

		// Initialise sieving primes for segmented sieve
		for s*s <= high {
			if primeFlagList[s] {
				primes = append(primes, s)
				multiples = append(multiples, ((s * s) - low))
			}
			s += 2
		}

		// Sieve the current
		numPrimesSoFar := len(primes)
		for k := 0; k < numPrimesSoFar; k++ {
			j := multiples[k]
			for m := primes[k] * 2; j < segmentSize; j += m {
				sieve[j] = false
			}
			multiples[k] = j - segmentSize
		}

		for n <= high {
			if sieve[n-low] {
				// n is prime
				//retPrimes = append(retPrimes, n)
				if inputNum%n == 0 {
					retLargestPrimeFactor = n
				}
			}
			n += 2
		}
	}

	return retLargestPrimeFactor
}

func minInt(x int, y int) int {
	if x < y {
		return x
	} else if y < x {
		return y
	}
	return x
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	} else if y > x {
		return y
	}
	return x
}

//////////////////////////////////////////////////////////////
// WORKSPACE (Train of thought not used in the final answer)
//////////////////////////////////////////////////////////////

/////////////////////////////////
// FIRST ATTEMPT
// (The Naive approach
// Works but takes far too long)
/////////////////////////////////

func firstAttempt() {
	inputNum := 100 //600851475143
	divCounter := inputNum / 2
	largestPrimeFactor := -1
	for largestPrimeFactor < 0 {
		if (inputNum % divCounter) == 0 {
			// We have found the highest factor so far.
			// Check if it is a prime number.
			if isPrimeNumberSimpleMethod(divCounter) {
				largestPrimeFactor = divCounter
			}
		}
		divCounter--
	}

	fmt.Printf("The largest prime factor of the number %d is: %d", inputNum, largestPrimeFactor)
}

func isPrimeNumberSimpleMethod(x int) bool {
	bRetFlag := true
	for i := 2; i < x; i++ {
		if (x % i) == 0 {
			bRetFlag = false
			break
		}
	}
	return bRetFlag
}

/////////////////////////////////
// SECOND ATTEMPT
/////////////////////////////////

func secondAttempt() {

	// 1. Fill primes array.
	// 2. Search primes array for largest factor of the input num.
	// 3. Return result.

	inputNum := 600851475143

	primesList := make([]int, 0, 10)
	for i := 2; i < inputNum; i++ {
		if isNumDivisibleByAnyItem(i, primesList) == false {
			// New prime found.
			primesList = append(primesList, i)
			fmt.Println("Num Prime Check:", i, true)
		} else {
			fmt.Println("Num Prime Check:", i)
		}
	}

	largestPrimeFactor := -1
	for i := (len(primesList) - 1); i >= 0; i-- {
		if (inputNum % primesList[i]) == 0 {
			largestPrimeFactor = primesList[i]
			break
		}
	}

	fmt.Printf("The largest prime factor of the number %d is: %d", inputNum, largestPrimeFactor)
}

func isNumDivisibleByAnyItem(x int, items []int) bool {
	for _, e := range items {
		if (x % e) == 0 {
			return true
		}
	}
	return false
}

/////////////////////////////////
// THIRD ATTEMPT
// Basic Sieve of Eratosthenes
// (NOT_SEGMENTED)
/////////////////////////////////

func nonSegmentedSOEAttempt() {
	limit := 100
	primeNumberList := findPrimeNumbersSOE(limit)
	//fmt.Printf("The prime numbers from 2 to %d are: %v\n", limit, primeNumberList)

	largestPrimeFactor := -1
	for i := len(primeNumberList) - 1; i > 0; i-- {
		if limit%primeNumberList[i] == 0 {
			largestPrimeFactor = primeNumberList[i]
			break
		}
	}
	fmt.Printf("The largest prime factor of %d is: %d\n", limit, largestPrimeFactor)
}

func findPrimeNumbersSOE(limit int) []int {

	// Sieve of Eratosthenes method.

	// Initialise primeFlagList by assuming
	// that all numbers are prime.
	primeFlagList := make([]bool, limit+1)
	for i := 0; i <= limit; i++ {
		primeFlagList[i] = true
	}

	for p := 2; p <= limit; p++ {
		// Find next non-marked down number.
		// This is the next prime.
		if primeFlagList[p] == true {
			// Mark multiples of p as non-prime.
			for i := p * 2; i <= limit; i += p {
				primeFlagList[i] = false
			}
		}
	}

	// All marked items are prime numbers.
	retPrimeNumberList := make([]int, 0, 100)
	for i := 2; i <= limit; i++ {
		if primeFlagList[i] == true {
			retPrimeNumberList = append(retPrimeNumberList, i)
		}
	}

	return retPrimeNumberList
}

//////////////////////////////////////////////////////////////
