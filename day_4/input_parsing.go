package day_4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadDiagram(path string) (*Diagram, error) {
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

	fields := make([][]Field, 0)

	for scanner.Scan() {
		entry := scanner.Text()

		rowFieldsRaw := strings.Split(entry, "")
		rowFields := make([]Field, 0)

		for _, fieldChar := range rowFieldsRaw {
			field := FieldEmpty

			if fieldChar == "@" {
				field = FieldRoll
			}

			rowFields = append(rowFields, field)
		}

		fields = append(fields, rowFields)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return NewDiagram(fields), nil
}
