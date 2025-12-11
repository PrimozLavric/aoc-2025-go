package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
