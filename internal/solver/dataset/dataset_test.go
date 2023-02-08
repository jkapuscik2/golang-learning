package dataset

import (
	"fmt"
	"reflect"
	"testing"
)

var sampleGrid = Grid{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 6, 0, 1, 0, 7, 8},
	{0, 0, 7, 0, 4, 0, 2, 6, 0},
	{0, 0, 1, 0, 5, 0, 9, 3, 0},
	{9, 0, 4, 0, 6, 0, 0, 0, 5},
	{0, 7, 0, 3, 0, 0, 0, 1, 2},
	{1, 2, 0, 0, 0, 7, 4, 0, 0},
	{0, 4, 9, 2, 0, 6, 0, 0, 7},
}

func TestDataset_Validate(t *testing.T) {

	t.Run("No error when sampleGrid is correct", func(t *testing.T) {
		dataset := CopyGrid(sampleGrid)

		if Validate(dataset) != nil {
			t.Error("Validated correct Grid as invalid")
		}
	})

	t.Run("Grid has too low values provided", func(t *testing.T) {
		nGrid := CopyGrid(sampleGrid)

		nGrid[0][0] = -1

		if Validate(nGrid) == nil {
			t.Error("Validated Grid with negative values as correct")
		} else {
			if Validate(nGrid) != ErrInvalidData {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has too high values provided", func(t *testing.T) {
		nGrid := CopyGrid(sampleGrid)

		nGrid[0][0] = 10

		if Validate(nGrid) == nil {
			t.Error("Validated Grid with too high values as correct")
		} else {
			if Validate(nGrid) != ErrInvalidData {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a row", func(t *testing.T) {
		nGrid := CopyGrid(sampleGrid)

		nGrid[0][0] = 5
		nGrid[0][1] = 5

		if Validate(nGrid) == nil {
			t.Error("Validated Grid with duplicated values in a row as correct")
		} else {
			if Validate(nGrid) != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a column", func(t *testing.T) {
		nGrid := CopyGrid(sampleGrid)

		nGrid[8][0] = 1

		if Validate(nGrid) == nil {
			t.Error("Validated Grid with duplicated values in a column as correct")
		} else {
			if Validate(nGrid) != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a box", func(t *testing.T) {
		nGrid := CopyGrid(sampleGrid)

		nGrid[0][0] = 3
		nGrid[1][1] = 3

		if Validate(nGrid) == nil {
			t.Error("Validated Grid with duplicated values in a box as correct")
		} else {
			if Validate(nGrid) != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})
}

func TestDataset_GetValue(t *testing.T) {
	nGrid := CopyGrid(sampleGrid)

	_, err := nGrid.GetValue(Position{10, 10})

	if err != ErrorFieldNoExists {
		t.Error("Getting invalid value succeeded")
	}

	val, err := nGrid.GetValue(Position{8, 8})

	if val != sampleGrid[8][8] {
		t.Error("could not get value from the grid")
	}

	if err != nil {
		t.Errorf("error during getting value from the grid: %v", err.Error())
	}
}

func TestGrid_Rebuild(t *testing.T) {
	grid := CopyGrid(sampleGrid)
	copiedGrid := grid.Rebuild(Position{X: 0, Y: 0}, 0)

	if !reflect.DeepEqual(grid, copiedGrid) {
		t.Error("values of rebuilt grid should remain same")
	}

	if fmt.Sprintf("%p", &grid) == fmt.Sprintf("%p", &copiedGrid) {
		t.Error("memory address of grids are same")
	}

	if fmt.Sprintf("%p", &grid[0]) == fmt.Sprintf("%p", &copiedGrid[0]) {
		t.Error("memory address of grids elements are same")
	}
}
