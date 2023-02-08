package dataset

import (
	"bufio"
	"io"
	"strconv"
)

func Load(file io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var rowNum int
	var data Grid

	for scanner.Scan() {
		if rowNum >= GridSize {
			return data, ErrTooManyRows
		}

		row := scanner.Text()
		ints := [GridSize]int64{}

		for i, s := range row {
			val, err := strconv.ParseInt(string(s), 0, 0)

			if err != nil {
				return data, ErrInvalidData
			}

			if i > GridSize {
				return data, ErrTooManyCols
			}
			ints[i] = val
		}

		data[rowNum] = ints
		rowNum += 1
	}

	return data, nil
}
