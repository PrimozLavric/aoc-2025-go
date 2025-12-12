package day_4

import (
	"aoc-2025-go/utils"
	"testing"
)

func TestDiagram_RunRemoval_SmallExample(t *testing.T) {
	testPath := utils.GetTestDataPath("small_example.txt")

	diagram, err := ReadDiagram(testPath)

	if err != nil {
		t.Fatalf("Failed to parse diagram. Error: %v", err)
	}

	count := diagram.RunRemoval()
	if count != 13 {
		t.Fatalf("Expected RunRemoval to return 13, actual %d", count)
	}
}

func TestDiagram_RunRemoval_Input(t *testing.T) {
	testPath := utils.GetTestDataPath("input.txt")

	diagram, err := ReadDiagram(testPath)

	if err != nil {
		t.Fatalf("Failed to parse diagram. Error: %v", err)
	}

	count := diagram.RunRemoval()
	if count != 1478 {
		t.Fatalf("Expected RunRemoval to return 1478, actual %d", count)
	}
}

func TestDiagram_RunRepeatedRemoval_SmallExample(t *testing.T) {
	testPath := utils.GetTestDataPath("small_example.txt")

	diagram, err := ReadDiagram(testPath)

	if err != nil {
		t.Fatalf("Failed to parse diagram. Error: %v", err)
	}

	count := diagram.RunRepeatedRemoval()
	if count != 43 {
		t.Fatalf("Expected RunRepeatedRemoval to return 43, actual %d", count)
	}
}

func TestDiagram_RunRepeatedRemoval_Input(t *testing.T) {
	testPath := utils.GetTestDataPath("input.txt")

	diagram, err := ReadDiagram(testPath)

	if err != nil {
		t.Fatalf("Failed to parse diagram. Error: %v", err)
	}

	count := diagram.RunRepeatedRemoval()
	if count != 9120 {
		t.Fatalf("Expected RunRepeatedRemoval to return 9120, actual %d", count)
	}
}
