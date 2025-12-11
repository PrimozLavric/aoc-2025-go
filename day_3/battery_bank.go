package day_3

import "math"

type BatteryBank struct {
	batteries []int
}

func (bb *BatteryBank) GetMaxJoltageV1() int {
	lhsMax := bb.batteries[0]
	lhsIdx := 1
	rhsMax := bb.batteries[len(bb.batteries)-1]
	rhsIdx := len(bb.batteries) - 2

	for lhsIdx <= rhsIdx {
		if bb.batteries[rhsIdx] <= lhsMax {
			if rhsMax < bb.batteries[rhsIdx] {
				rhsMax = bb.batteries[rhsIdx]
			}
			rhsIdx -= 1
			continue
		}

		if lhsMax < bb.batteries[lhsIdx] {
			lhsMax = bb.batteries[lhsIdx]
		}
		lhsIdx += 1
	}

	return lhsMax*10 + rhsMax
}

func (bb *BatteryBank) GetMaxJoltageV2(numBatteries int) int {
	selectedBatteries := make([]int, 0)
	lastSelectedIdx := -1

	for i := 0; i < numBatteries; i++ {
		remainingCount := numBatteries - i - 1
		selected := 0

		for j := lastSelectedIdx + 1; j < len(bb.batteries)-remainingCount; j++ {
			if selected < bb.batteries[j] {
				selected = bb.batteries[j]
				lastSelectedIdx = j
			}
		}

		selectedBatteries = append(selectedBatteries, selected)
	}

	total := 0
	for i, b := range selectedBatteries {
		total += b * int(math.Pow10(numBatteries-i-1))
	}

	return total
}

func BatteryBanksMaxJolts(bb []BatteryBank, numBatteries int) int {
	total := 0
	for _, bank := range bb {
		total += bank.GetMaxJoltageV2(numBatteries)
	}

	return total
}
