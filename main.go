package main

import (
	"aoc-2025-go/day_1"
	"fmt"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	dial := day_1.NewSafeDial()
	seq, err := day_1.ReadRotationSequence("day_1/testdata/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	zeroHits, zeroHitsAndPasses := dial.ExecuteRotationSequence(seq)

	fmt.Printf("Zero hits: %d \n", zeroHits)
	fmt.Printf("Zero hits and passes: %d \n", zeroHitsAndPasses)
}
