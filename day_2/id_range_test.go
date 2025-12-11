package day_2

import (
	"aoc-2025-go/utils"
	"log"
	"slices"
	"testing"
)

func TestReadIDRanges(t *testing.T) {
	testPath := utils.GetTestDataPath("small_example.txt")

	ranges, err := ReadIDRanges(testPath)

	if err != nil {
		log.Printf("Failed to read ID ranges, error: %v\n", err)
		t.Fail()
	}

	expectedRanges := []IDRange{
		{11, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},
		{1698522, 1698528},
		{446443, 446449},
		{38593856, 38593862},
		{565653, 565659},
		{824824821, 824824827},
		{2121212118, 2121212124},
	}

	if !slices.Equal(expectedRanges, ranges) {
		log.Println("Incorrectly read ranges.")
		t.Fail()
	}
}

func TestSumInvalidIDsPart1(t *testing.T) {
	// Small Example
	{
		idRanges, err := ReadIDRanges(utils.GetTestDataPath("small_example.txt"))

		if err != nil {
			log.Printf("Failed to read ID ranges, error: %v\n", err)
			t.Fail()
		}

		invalidIDsSum := SumInvalidIDs(idRanges, ValidatorPart1)
		expectedInvalidIDsSum := 1227775554

		log.Printf("Small example invalid ID sum: %d", invalidIDsSum)
		if invalidIDsSum != expectedInvalidIDsSum {
			log.Printf("Expected invalid id sum to be %d, actual: %d", invalidIDsSum, expectedInvalidIDsSum)
			t.Fail()
		}
	}

	// Puzzle input
	{
		idRanges, err := ReadIDRanges(utils.GetTestDataPath("input.txt"))

		if err != nil {
			log.Printf("Failed to read ID ranges, error: %v\n", err)
			t.Fail()
		}

		invalidIDsSum := SumInvalidIDs(idRanges, ValidatorPart1)
		expectedInvalidIDsSum := 13919717792

		log.Printf("Puzzle input invalid ID sum: %d", invalidIDsSum)
		if invalidIDsSum != expectedInvalidIDsSum {
			log.Printf("Expected invalid id sum to be %d, actual: %d", invalidIDsSum, expectedInvalidIDsSum)
			t.Fail()
		}
	}
}

func TestSumInvalidIDsPart2(t *testing.T) {
	// Small Example
	{
		idRanges, err := ReadIDRanges(utils.GetTestDataPath("small_example.txt"))

		if err != nil {
			log.Printf("Failed to read ID ranges, error: %v\n", err)
			t.Fail()
		}

		invalidIDsSum := SumInvalidIDs(idRanges, ValidatorPart2)
		expectedInvalidIDsSum := 4174379265

		log.Printf("Small example invalid ID sum: %d", invalidIDsSum)
		if invalidIDsSum != expectedInvalidIDsSum {
			log.Printf("Expected invalid id sum to be %d, actual: %d", invalidIDsSum, expectedInvalidIDsSum)
			t.Fail()
		}
	}

	// Puzzle input
	{
		idRanges, err := ReadIDRanges(utils.GetTestDataPath("input.txt"))

		if err != nil {
			log.Printf("Failed to read ID ranges, error: %v\n", err)
			t.Fail()
		}

		invalidIDsSum := SumInvalidIDs(idRanges, ValidatorPart2)
		expectedInvalidIDsSum := 14582313461

		log.Printf("Puzzle input invalid ID sum: %d", invalidIDsSum)
		if invalidIDsSum != expectedInvalidIDsSum {
			log.Printf("Expected invalid id sum to be %d, actual: %d", invalidIDsSum, expectedInvalidIDsSum)
			t.Fail()
		}
	}
}
