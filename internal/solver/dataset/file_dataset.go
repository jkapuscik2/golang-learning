package dataset

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

func Load(file io.Reader) (Dataset, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var data [][]int64

	for scanner.Scan() {
		row := scanner.Text()
		ints := make([]int64, len(row))

		for i, s := range row {
			val, err := strconv.ParseInt(string(s), 0, 0)

			if err != nil {
				return Dataset{}, errors.New("invalid dataset provided")
			}

			ints[i] = val
		}

		data = append(data, ints)
	}

	return Dataset{data}, nil
}
