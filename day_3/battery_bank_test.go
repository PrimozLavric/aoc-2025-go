package day_3

import (
	"aoc-2025-go/utils"
	"fmt"
	"testing"
)

func TestReadBatteryBanks(t *testing.T) {
	testPath := utils.GetTestDataPath("input.txt")

	banks, err := ReadBatteryBanks(testPath)

	if err != nil {
		fmt.Printf("Failed to read battery banks, error: %v", err)
		t.Fatal()
	}

	part1_result := BatteryBanksMaxJolts(banks, 2)
	expected_part1_result := 16993
	fmt.Printf("Part 1 total output: %d\n", part1_result)
	if part1_result != 16993 {
		fmt.Printf("Wrong part 1 result. Expected %d, actual %d", expected_part1_result, part1_result)
	}

	part2_result := BatteryBanksMaxJolts(banks, 12)
	expected_part2_result := 168617068915447
	fmt.Printf("Part 2 total output: %d\n", part2_result)
	if part2_result != 168617068915447 {
		fmt.Printf("Wrong part 1 result. Expected %d, actual %d", expected_part2_result, part2_result)
	}
}
