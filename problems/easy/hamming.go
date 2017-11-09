package easy

import (
	"strconv"
	"strings"
)

// HammingDistance calculates the Hamming distace.
// The Hamming distance between two integers is the number of positions
// at which the corresponding bits are different.
func HammingDistance(x, y uint64) uint64 {
	var distance uint64
	// XOR of two numbers
	// 6(110) and 0(000) gives 6(110)
	value := x ^ y
	// count the number of bits set
	for value != 0 {
		// A bit is set, so increment the count and clear the bit
		distance++
		// counting bits by algorihm of Peter Wegner
		// http://dl.acm.org/citation.cfm?id=367236.367286
		// step1: 6(110) & 5(101) = 4(100) (distance=1)
		// step2: 4(100) & 3(011) = 0(000) (distance=2)
		value &= value - 1
	}
	// Return the number of differing bits
	return distance
}

// HammingDistanceDummy is unoptimised dummy version
// to calculate Hamming distance
func HammingDistanceDummy(x, y uint64) uint64 {
	// binary representation of 1st number
	s1 := strconv.FormatUint(x, 2)
	// binary representation of 2nd number
	s2 := strconv.FormatUint(y, 2)
	// completes strings if they have not equal lengths
	// for example, s1="1001", s2="111000"
	// then after func s1=001001, s2=111000
	s1, s2 = completeStrings(s1, s2)
	// algorithm is easy, compare each symbol in both strings
	// if they are not equal then increment distance
	var distance uint64
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return distance
}

// completeStrings adds some "0" (difference between length of
// two numbers) before number that has less digits in binary string represantion
func completeStrings(s1, s2 string) (string, string) {
	diff := len(s1) - len(s2)
	if diff > 0 {
		s2 = strings.Repeat("0", diff) + s2
	} else {
		s1 = strings.Repeat("0", -diff) + s1
	}

	return s1, s2
}
