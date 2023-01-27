package solver

import (
	"learning-go-sudoku/internal/solver/dataset"
	"reflect"
	"testing"
)

func TestSolveBacktrace(t *testing.T) {
	data := dataset.CopyGrid(sampleGrid)
	res, err := SolveBacktrace(data)

	if err != nil {
		t.Fatalf("Error during solving test grid: %q", err.Error())
	}

	if !reflect.DeepEqual(res, sampleGridSolved) {
		t.Error("Sudoku was not solved correctly")
	}
}

func TestSolveBacktraceInvalid(t *testing.T) {
	data := dataset.CopyGrid(sampleInvalidGrid)

	_, err := SolveBacktrace(data)
	if err == nil {
		t.Fatalf("Did not report invalid grid")
	}
}

func BenchmarkTestSolveBacktrace(b *testing.B) {

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		SolveBacktrace(grid)
	}
}
