package day_4

type Field int

const (
	FieldEmpty Field = iota
	FieldRoll
)

type Diagram struct {
	fields [][]Field
}

func NewDiagram(fields [][]Field) *Diagram {
	return &Diagram{fields: fields}
}

func (d *Diagram) CanBeRemoved(row int, col int) bool {
	numRows := len(d.fields)
	numCols := len(d.fields[0])

	if row < 0 || row >= numRows || col < 0 || col >= numCols {
		// Row and col is out of range.
		return false
	}

	if d.fields[row][col] == FieldEmpty {
		// Nothing to remove in empty field.
		return false
	}

	// Start at -1 because we'll also count the tested roll
	rollsCount := -1

	for i := max(row-1, 0); i < min(row+2, numRows); i++ {
		for j := max(col-1, 0); j < min(col+2, numCols); j++ {
			if d.fields[i][j] == FieldRoll {
				rollsCount++
			}
		}
	}

	return rollsCount < 4
}

func (d *Diagram) RunRemoval() int {
	numRows := len(d.fields)
	numCols := len(d.fields[0])

	type ToRemoveEntry = struct {
		row int
		col int
	}
	toRemove := make([]ToRemoveEntry, 0)

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if d.CanBeRemoved(i, j) {
				toRemove = append(toRemove, ToRemoveEntry{row: i, col: j})
			}
		}
	}

	for _, entry := range toRemove {
		d.fields[entry.row][entry.col] = FieldEmpty
	}

	return len(toRemove)
}

func (d *Diagram) RunRepeatedRemoval() int {
	count := 0

	for {
		diff := d.RunRemoval()

		if diff == 0 {
			break
		}

		count += diff
	}

	return count
}
