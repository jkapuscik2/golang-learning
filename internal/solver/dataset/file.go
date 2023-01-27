package dataset

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

func Load(file io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var data Grid

	for scanner.Scan() {
		row := scanner.Text()
		ints := make([]int64, len(row))

		for i, s := range row {
			val, err := strconv.ParseInt(string(s), 0, 0)

			if err != nil {
				return data, errors.New("invalid dataset provided")
			}

			ints[i] = val
		}

		data = append(data, ints)
	}

	return data, nil
}
