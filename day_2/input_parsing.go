package day_2

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIDRanges(path string) ([]IDRange, error) {
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
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the delimiter
		if i := bytes.IndexByte(data, ','); i >= 0 {
			// Return token up to (but not including) the delimiter
			return i + 1, data[0:i], nil
		}

		// If at EOF, return remaining data
		if atEOF {
			return len(data), data, nil
		}

		// Request more data
		return 0, nil, nil
	})

	idRanges := make([]IDRange, 0)

	for scanner.Scan() {
		entry := scanner.Text()

		rawRange := strings.Split(entry, "-")

		if len(rawRange) != 2 {
			return nil, fmt.Errorf("bad id range entry \"%s\"", entry)
		}

		begin, err := strconv.Atoi(rawRange[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse begin of id range \"%s\", error: %v", entry, err)
		}

		end, err := strconv.Atoi(rawRange[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse end of id range \"%s\", error: %v", entry, err)
		}

		idRange, err := MakeIDRange(begin, end)
		if err != nil {
			return nil, fmt.Errorf("failed to parse id range \"%s\", error: %v", entry, err)
		}

		idRanges = append(idRanges, idRange)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return idRanges, nil
}
