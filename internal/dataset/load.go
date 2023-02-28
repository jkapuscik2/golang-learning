package dataset

import (
	"bufio"
	"io"
	"strconv"
)

func Load(file io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rowNum int
	var data Grid

	for scanner.Scan() {
		if rowNum >= GridSize {
			return data, ErrTooManyRows
		}

		row := scanner.Text()
		ints := []int64{}

		for i, s := range row {
			val, err := strconv.ParseInt(string(s), 10, 0)

			if err != nil {
				return data, ErrInvalidData
			}

			if i > GridSize {
				return data, ErrTooManyCols
			}
			ints = append(ints, val)
		}

		if len(ints) != GridSize {
			return data, ErrIncompleteData
		}

		data[rowNum] = *(*[GridSize]int64)(ints)
		rowNum += 1
	}

	if rowNum < GridSize {
		return data, ErrIncompleteData
	}

	if err := Validate(data); err != nil {
		return data, err
	}

	return data, nil
}
