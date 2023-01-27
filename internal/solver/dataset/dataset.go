package dataset

import "fmt"

const (
	EmptyVal            = 0
	BoxSize             = 3
	MinVal              = 1
	MaxVal              = 9
	GridSize            = 9
	ErrInvalidData      = GridError("Invalid data provided")
	ErrTooManyCols      = GridError("Too many columns in dataset")
	ErrTooManyRows      = GridError("Too many rows in dataset")
	ErrDuplicatedValues = GridError("There are duplicated values in the Grid")
	ErrorFieldNoExists  = GridError("Filed does not exists")
)

type GridError string

func (e GridError) Error() string {
	return string(e)
}

type Position struct {
	X int
	Y int
}

type Grid [][]int64

func (dataset Grid) IsFilled() bool {
	for _, arr := range dataset {
		for _, val := range arr {
			if val == EmptyVal {
				return false
			}
		}
	}
	return true
}

func (dataset Grid) Rebuild(pos Position, val int64) (Grid, error) {
	rebuilt := CopyGrid(dataset)
	rebuilt[pos.Y][pos.X] = val

	return rebuilt, nil
}

func CopyGrid(matrix Grid) Grid {
	duplicate := make(Grid, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int64, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}

func (dataset Grid) GetValue(pos Position) (int64, error) {
	if !has(dataset, pos) {
		return 0, ErrorFieldNoExists
	}

	return dataset[pos.Y][pos.X], nil
}

func Validate(dataset Grid) error {
	if len(dataset) != GridSize {
		return ErrTooManyRows
	}

	columns := make([][]int64, len(dataset))
	for _, row := range dataset {
		if len(row) != GridSize {
			return ErrTooManyCols
		}

		// check if there are duplicated values in rows
		if hasDuplicates(row) {
			return ErrTooManyCols
		}

		// check if cell vales are correct
		for x, item := range row {
			if item != EmptyVal && (item < MinVal || item > MaxVal) {
				return ErrInvalidData
			}

			columns[x] = append(columns[x], item)
		}
	}

	// check if there are duplicated values in columns
	for _, column := range columns {
		if hasDuplicates(column) {
			return ErrDuplicatedValues
		}
	}

	// check if there duplicated values in 3x3 boxes
	for x := 0; x <= GridSize-1; x = x + BoxSize {
		subset := dataset[x : x+BoxSize]

		for y := 0; y <= GridSize-1; y = y + BoxSize {
			var boxVars []int64

			for _, values := range subset {
				dd := values[y : y+BoxSize]
				boxVars = append(boxVars, dd...)
			}

			if hasDuplicates(boxVars) {
				return ErrDuplicatedValues
			}
		}
	}

	return nil
}

func PrettyPrint(dataset Grid) {
	for _, row := range dataset {
		fmt.Println(row)
	}
}

func IsEmpty(cell int64) bool {
	return int(cell) == EmptyVal
}

func has(grid Grid, pos Position) bool {
	if len(grid) < pos.Y {
		return false
	}

	for _, row := range grid {
		if len(row) < pos.X {
			return false
		}
	}

	return true
}

func hasDuplicates(row []int64) bool {
	checked := make(map[int64]int64, len(row))
	for _, val := range row {
		if _, ok := checked[val]; ok {
			return true
		}

		if val != EmptyVal {
			checked[val] = val
		}
	}

	return false
}
