package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dir int

const (
	DirLeft Dir = iota
	DirRight
)

type Rotation struct {
	Dir       Dir
	NumClicks int
}

func wraparoundMod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func (r Rotation) apply(dialPosition int) (int, int) {
	if r.Dir == DirLeft {
		if dialPosition == 0 {
			// [Special case]: Starting at 0
			return wraparoundMod(-r.NumClicks, 100), r.NumClicks / 100
		}

		if r.NumClicks < dialPosition {
			// [Special case]: Never wraparound
			return dialPosition - r.NumClicks, 0
		}

		clicksAfterFirstZero := r.NumClicks - dialPosition
		zeroHits := 1 + (clicksAfterFirstZero / 100)

		return wraparoundMod(dialPosition-r.NumClicks, 100), zeroHits
	}

	newPos := dialPosition + r.NumClicks
	return newPos % 100, newPos / 100
}

func ReadRotationSequence(path string) ([]Rotation, error) {
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

	rotationSequence := make([]Rotation, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var dir Dir

		switch line[0] {
		case 'L':
			dir = DirLeft
		case 'R':
			dir = DirRight
		default:
			return nil, fmt.Errorf("failed to parse direction '%c'", line[0])
		}

		numClicks, err := strconv.Atoi(line[1:])

		if err != nil {
			return nil, err
		}

		rotationSequence = append(rotationSequence, Rotation{Dir: dir, NumClicks: numClicks})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotationSequence, nil
}
