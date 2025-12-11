package day_1

import (
	"aoc-2025-go/utils"
	"log"
	"testing"
)

func TestSafeDialZeroHits(t *testing.T) {
	dial := NewSafeDial()
	seq, err := ReadRotationSequence(utils.GetTestDataPath("input.txt"))

	if err != nil {
		log.Fatal(err)
	}

	zeroHits, _ := dial.ExecuteRotationSequence(seq)

	log.Printf("Zero hits: %d \n", zeroHits)
	if zeroHits != 1026 {
		log.Println("Wrong result!")
		t.Fail()
	}
}

func TestSafeDialZeroHitsAndPasses(t *testing.T) {
	dial := NewSafeDial()
	seq, err := ReadRotationSequence(utils.GetTestDataPath("input.txt"))

	if err != nil {
		log.Fatal(err)
	}

	_, zeroHitsAndPasses := dial.ExecuteRotationSequence(seq)

	log.Printf("Zero hits and passes: %d \n", zeroHitsAndPasses)
	if zeroHitsAndPasses != 5923 {
		log.Println("Wrong result!")
		t.Fail()
	}
}
