package day_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadBatteryBanks(path string) ([]BatteryBank, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file %s: %v", path, err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	batteryBanks := make([]BatteryBank, 0)

	for scanner.Scan() {
		entry := scanner.Text()

		batteriesRaw := strings.Split(entry, "")
		batteries := make([]int, len(batteriesRaw))

		for i, batteryRaw := range batteriesRaw {
			battery, err := strconv.Atoi(batteryRaw)
			if err != nil {
				return nil, fmt.Errorf("failed to parse batteries values \"%s\", error: %v", entry, err)
			}

			batteries[i] = battery
		}

		batteryBanks = append(batteryBanks, BatteryBank{batteries: batteries})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return batteryBanks, nil
}
