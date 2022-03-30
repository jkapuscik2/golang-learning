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

type Dataset struct {
	Grid [][]int64
}

func (dataset Dataset) Data() [][]int64 {
	return dataset.Grid
}

func (dataset Dataset) Validate() error {
	if len(dataset.Grid) != GridSize {
		return ErrTooManyRows
	}

	columns := make([][]int64, len(dataset.Grid))
	for y, row := range dataset.Grid {
		if len(row) != GridSize {
			return ErrTooManyCols
		}

		// check if there are duplicated values in rows
		if hasDuplicates(row) {
			return ErrDuplicatedValues
		}

		// check if cell vales are correct
		for _, item := range row {
			if item != EmptyVal && (item < MinVal || item > MaxVal) {
				return ErrInvalidData
			}

			columns[y] = append(columns[y], item)
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
		subset := dataset.Grid[x : x+BoxSize]

		for y := 0; y <= GridSize-1; y = y + BoxSize {
			var boxVars []int64

			for _, vals := range subset {
				dd := vals[y : y+BoxSize]
				boxVars = append(boxVars, dd...)
			}

			if hasDuplicates(boxVars) {
				return ErrDuplicatedValues
			}
		}
	}

	return nil
}

func (dataset *Dataset) SetValue(pos Position, val int64) error {
	if !has(dataset.Data(), pos) {
		return ErrorFieldNoExists
	}

	dataset.Grid[pos.Y][pos.X] = val

	return nil
}

func (dataset Dataset) GetValue(pos Position) (int64, error) {
	if !has(dataset.Data(), pos) {
		return 0, ErrorFieldNoExists
	}

	return dataset.Grid[pos.Y][pos.X], nil
}

func (dataset Dataset) PrettyPrint() {
	for _, row := range dataset.Grid {
		fmt.Println(row)
	}
}

func has(grid [][]int64, pos Position) bool {
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

func Create(data [][]int64) Dataset {
	return Dataset{data}
}
