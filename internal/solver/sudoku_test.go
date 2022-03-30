package solver

import (
	"learning-go-d/internal/solver/dataset"
	"reflect"
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

var sampleGridSolved = [][]int64{
	{1, 2, 3, 4, 7, 8, 5, 6, 9},
	{6, 7, 8, 2, 5, 9, 1, 3, 4},
	{4, 5, 9, 6, 3, 1, 2, 7, 8},
	{3, 5, 7, 1, 4, 9, 2, 6, 8},
	{2, 6, 1, 7, 5, 8, 9, 3, 4},
	{9, 8, 4, 2, 6, 3, 1, 7, 5},
	{5, 7, 6, 3, 4, 8, 9, 1, 2},
	{1, 2, 3, 5, 9, 7, 4, 6, 8},
	{8, 4, 9, 2, 1, 6, 3, 5, 7},
}

func TestSudoku_SolveSync(t *testing.T) {
	data := dataset.Create(sampleGrid)
	res, err := SolveSync(data)

	if err != nil {
		t.Fatalf("Error during solving test grid: %q", err.Error())
	}

	if !reflect.DeepEqual(res, dataset.Create(sampleGridSolved)) {
		t.Error("Sudoku was not solved correctly")
	}
}

func BenchmarkSolveSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := dataset.Create(sampleGrid)
		SolveSync(data)
	}
}
