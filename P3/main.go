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

func main() {

	// A whole number, x, is a factor of another number, y,
	// if x * z = y.
	// A prime factor requires that x also be a prime number.
	//secondAttempt()

	limit := 1000000000
	primeNumberList := findPrimeNumbers(limit)
	fmt.Printf("The prime numbers from 2 to %d are: %v\n", limit, primeNumberList)
}

//////////////////////////////////////////////////////////////
// FINAL ANSWER
//////////////////////////////////////////////////////////////

// TODO: Code is still work in progress.

//////////////////////////////////////////////////////////////
// WORKSPACE (Train of thought not used in the final answer)
//////////////////////////////////////////////////////////////

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

func firstAttempt() {
	inputNum := 100 //600851475143
	divCounter := inputNum / 2
	largestPrimeFactor := -1
	for largestPrimeFactor < 0 {
		if (inputNum % divCounter) == 0 {
			// We have found the highest factor so far.
			// Check if it is a prime number.
			if isPrimeNumber(divCounter) {
				largestPrimeFactor = divCounter
			}
		}
		divCounter--
	}

	fmt.Printf("The largest prime factor of the number %d is: %d", inputNum, largestPrimeFactor)
}

func isPrimeNumber(x int) bool {
	bRetFlag := true
	for i := 2; i < x; i++ {
		if (x % i) == 0 {
			bRetFlag = false
			break
		}
	}
	return bRetFlag
}

func isNumDivisibleByAnyItem(x int, items []int) bool {
	for _, e := range items {
		if (x % e) == 0 {
			return true
		}
	}
	return false
}

func doesSliceContains(slice []int, x int) bool {
	for _, e := range slice {
		if e == x {
			return true
		}
	}
	return false
}

func findPrimeNumbers(limit int) []int {
	// Warning: Sieve of Eratosthenes has a large memory complexity.
	return findPrimeNumbers_SOE(limit)
}

func findPrimeNumbers_segmentedSOE(limit int) []int {

	// Segmented version of the Sieve of Eratosthenes
	// in order to reduce space complexity for large Limits.

	// (Segment size in bytes. 32KB)
	segmentSize := 32768
	numRequiredSegments := int(math.Ceil(float64(limit) / float64(segmentSize)))
	segment := make([]bool, segmentSize)
	primeNumberList := make([]int, 0, 10)

	for s := 0; s < numRequiredSegments; s++ {

		// TODO: Continue here.

		// Clear the segment.
		for i := 0; i <= limit; i++ {
			segment[i] = true
		}

	}

	return primeNumberList
}

func findPrimeNumbers_SOE(limit int) []int {

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
