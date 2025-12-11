package day_1

import (
	"aoc-2025-go/utils"
	"slices"
	"testing"
)

func TestParsing(t *testing.T) {
	expectedSeq := []Rotation{
		{DirLeft, 68},
		{DirLeft, 30},
		{DirRight, 48},
		{DirLeft, 5},
		{DirRight, 60},
		{DirLeft, 55},
		{DirLeft, 1},
		{DirLeft, 99},
		{DirRight, 14},
		{DirLeft, 82},
	}
	seq, err := ReadRotationSequence(utils.GetTestDataPath("small_example.txt"))

	if err != nil {
		t.Fail()
	}

	if !slices.Equal(expectedSeq, seq) {
		t.Fail()
	}
}

func TestApply(t *testing.T) {
	type testCase struct {
		dialValue        int
		rot              Rotation
		expectedDial     int
		expectedZeroPass int
	}

	executeTestCases := func(testCases []testCase) {
		for _, tc := range testCases {
			newDialValue, zeroPasses := tc.rot.apply(tc.dialValue)

			if newDialValue != tc.expectedDial {
				t.Errorf("Expected new dial value %d, actual %d", tc.expectedDial, newDialValue)
			}

			if zeroPasses != tc.expectedZeroPass {
				t.Errorf("Expected %d zero passes, actual %d", tc.expectedZeroPass, zeroPasses)
			}
		}
	}

	rightTestCases := []testCase{
		{50, Rotation{Dir: DirRight, NumClicks: 20}, 70, 0},
		{80, Rotation{Dir: DirRight, NumClicks: 20}, 0, 1},
		{50, Rotation{Dir: DirRight, NumClicks: 1000}, 50, 10},
		{0, Rotation{Dir: DirRight, NumClicks: 100}, 0, 1},
		{70, Rotation{Dir: DirRight, NumClicks: 30}, 0, 1},
		{70, Rotation{Dir: DirRight, NumClicks: 29}, 99, 0},
		{99, Rotation{Dir: DirRight, NumClicks: 1}, 0, 1},
	}

	executeTestCases(rightTestCases)

	leftTestCases := []testCase{
		{0, Rotation{Dir: DirLeft, NumClicks: 5}, 95, 0},
		{0, Rotation{Dir: DirLeft, NumClicks: 100}, 0, 1},
		{50, Rotation{Dir: DirLeft, NumClicks: 50}, 0, 1},
		{1, Rotation{Dir: DirLeft, NumClicks: 1000}, 1, 10},
		{1, Rotation{Dir: DirLeft, NumClicks: 1001}, 0, 11},
		{99, Rotation{Dir: DirLeft, NumClicks: 99}, 0, 1},
		{0, Rotation{Dir: DirLeft, NumClicks: 1000}, 0, 10},
		{50, Rotation{Dir: DirLeft, NumClicks: 82}, 68, 1},
		{50, Rotation{Dir: DirLeft, NumClicks: 49}, 1, 0},
		{1, Rotation{Dir: DirLeft, NumClicks: 99}, 2, 1},
		{0, Rotation{Dir: DirLeft, NumClicks: 201}, 99, 2},
	}

	executeTestCases(leftTestCases)
}
