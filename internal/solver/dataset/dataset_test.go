package dataset

import (
	"fmt"
	"testing"
)

var sampleGrid = [][]int64{
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
		dataset := Dataset{sampleGrid}

		if dataset.Validate() != nil {
			t.Error("Validated correct Grid as invalid")
		}
	})

	t.Run("Grid has too many rows", func(t *testing.T) {
		dataset := Dataset{append(sampleGrid, sampleGrid[0])}

		if dataset.Validate() == nil {
			t.Error("Validated Grid with too many rows as correct")
		} else {
			if dataset.Validate() != ErrTooManyRows {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has too many columns", func(t *testing.T) {
		colGrid := make([][]int64, len(sampleGrid))
		copy(colGrid, sampleGrid)
		colGrid[0] = append(colGrid[0], 1, 2, 3)

		dataset := Dataset{colGrid}

		if dataset.Validate() == nil {
			t.Error("Validated Grid with too many columns as correct")
		} else {
			if dataset.Validate() != ErrTooManyCols {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has too low values provided", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		dataset.SetValue(Position{0, 0}, -1)

		if dataset.Validate() == nil {
			t.Error("Validated Grid with negative values as correct")
		} else {
			if dataset.Validate() != ErrInvalidData {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has too high values provided", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		dataset.SetValue(Position{0, 0}, 10)

		if dataset.Validate() == nil {
			t.Error("Validated Grid with too high values as correct")
		} else {
			if dataset.Validate() != ErrInvalidData {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a row", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		dataset.SetValue(Position{0, 0}, 5)
		dataset.SetValue(Position{0, 1}, 5)

		if dataset.Validate() == nil {
			t.Error("Validated Grid with duplicated values in a row as correct")
		} else {
			if dataset.Validate() != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a column", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		dataset.SetValue(Position{0, 5}, 5)
		dataset.SetValue(Position{1, 5}, 5)

		if dataset.Validate() == nil {
			t.Error("Validated Grid with duplicated values in a column as correct")
		} else {
			if dataset.Validate() != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})

	t.Run("Grid has duplicated values in a box", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		dataset.SetValue(Position{5, 5}, 3)
		dataset.SetValue(Position{5, 6}, 3)

		if dataset.Validate() == nil {
			t.Error("Validated Grid with duplicated values in a box as correct")
		} else {
			if dataset.Validate() != ErrDuplicatedValues {
				t.Error("Invalid error provided")
			}
		}
	})
}

func TestDataset_SetValue(t *testing.T) {
	t.Run("Set correct value", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		err := dataset.SetValue(Position{1, 1}, 2)

		if err != nil {
			t.Fatal("Could not set value")
		}

		res, err := dataset.GetValue(Position{1, 1})
		if res != 2 && err != nil {
			t.Error("Could not fetch set value")
		}
	})

	t.Run("Set not existing position", func(t *testing.T) {
		nGrid := make([][]int64, len(sampleGrid))
		copy(nGrid, sampleGrid)

		dataset := Dataset{nGrid}
		err := dataset.SetValue(Position{1, 10}, 2)

		if err != ErrorFieldNoExists {
			t.Error("Setting invalid value succeeded")
		}
	})
}

func TestDataset_GetValue(t *testing.T) {
	nGrid := make([][]int64, len(sampleGrid))
	copy(nGrid, sampleGrid)

	dataset := Dataset{nGrid}
	d, err := dataset.GetValue(Position{10, 10})
	fmt.Println(d)
	if err != ErrorFieldNoExists {
		t.Error("Getting invalid value succeeded")
	}
}
