package day_2

import (
	"fmt"
	"math"
)

type ValidatorType int

const (
	ValidatorPart1 ValidatorType = iota
	ValidatorPart2
)

type IDRange struct {
	begin int
	end   int
}

func MakeIDRange(begin int, end int) (IDRange, error) {
	if begin > end {
		return IDRange{}, fmt.Errorf("invalid id range, begin (%d) is greater than end (%d)", begin, end)
	}

	return IDRange{begin: begin, end: end}, nil
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	return int(math.Log10(float64(n))) + 1
}

func isValidIDPart1(id int) bool {
	numDigits := countDigits(id)

	if numDigits%2 != 0 {
		return true
	}

	divisor := int(math.Pow10(numDigits / 2))
	return id/divisor != id%divisor
}

func generateRepeatedPattern(pattern int, patternLen, repeats int) int {
	result := 0
	multiplier := int(math.Pow10(patternLen))

	for i := 0; i < repeats; i++ {
		result = result*multiplier + pattern
	}

	return result
}

func isValidIDPart2(id int) bool {
	numDigits := countDigits(id)

	for patternLen := numDigits / 2; patternLen > 0; patternLen-- {
		if numDigits%patternLen != 0 {
			// Pattern doesn't divide id
			continue
		}

		divisor := int(math.Pow10(numDigits - patternLen))
		pattern := id / divisor

		if generateRepeatedPattern(pattern, patternLen, numDigits/patternLen) == id {
			// Matches generated pattern. ID is not valid
			return false
		}
	}

	return true
}

func isValidID(id int, validatorType ValidatorType) bool {
	if validatorType == ValidatorPart1 {
		return isValidIDPart1(id)
	} else {
		return isValidIDPart2(id)
	}
}

func FindInvalidIDs(idRanges []IDRange, validatorType ValidatorType) []int {
	invalidIDs := make([]int, 0)

	for _, idRange := range idRanges {
		for id := idRange.begin; id <= idRange.end; id += 1 {
			if !isValidID(id, validatorType) {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}

	return invalidIDs
}

func SumInvalidIDs(idRanges []IDRange, validatorType ValidatorType) int {
	invalidIDs := FindInvalidIDs(idRanges, validatorType)

	sum := 0
	for _, id := range invalidIDs {
		sum += id
	}

	return sum
}
