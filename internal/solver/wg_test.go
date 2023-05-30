package solver_test

import (
	"errors"
	"github.com/jkapuscik2/sudoku-solver/internal/dataset"
	"github.com/jkapuscik2/sudoku-solver/internal/solver"
	"reflect"
	"runtime"
	"testing"
)

func TestSolveWg(t *testing.T) {
	type args struct {
		grid    dataset.Grid
		workers int
	}
	tests := []struct {
		name              string
		args              args
		want              dataset.Grid
		wantErr           bool
		errType           error
		multipleSolutions bool
	}{
		{
			name:    "correct dataset",
			args:    args{grid: sampleGrid, workers: runtime.NumCPU()},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "unsolvable dataset",
			args:    args{grid: sampleInvalidGrid, workers: runtime.NumCPU()},
			want:    sampleInvalidGrid,
			wantErr: true,
			errType: solver.ErrNoSolutions,
		},
		{
			name:    "blocked dataset",
			args:    args{grid: sampleGridSolved, workers: runtime.NumCPU()},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "wrongly solved dataset",
			args:    args{grid: sampleGridSolvedInvalid, workers: runtime.NumCPU()},
			want:    sampleGridSolvedInvalid,
			wantErr: true,
			errType: solver.ErrNoSolutions,
		},
		{
			name:              "empty dataset",
			args:              args{grid: emptyGrid, workers: 1},
			want:              emptyGrid,
			wantErr:           false,
			multipleSolutions: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solver.SolveWg(tt.args.grid, tt.args.workers)
			if tt.wantErr && err != nil && !errors.Is(err, tt.errType) {
				t.Errorf("SolveWg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.multipleSolutions && !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveWg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolveWg(b *testing.B) {
	b.ReportAllocs()

	workers := runtime.GOMAXPROCS(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		solver.SolveWg(grid, workers)
	}
}

func BenchmarkSolveWgSimple(b *testing.B) {
	b.ReportAllocs()

	workers := runtime.GOMAXPROCS(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGridSimple)
		solver.SolveWg(grid, workers)
	}
}

//func BenchmarkSolveWgEmpty(b *testing.B) {
//	b.ReportAllocs()
//
//	workers := runtime.GOMAXPROCS(0)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		grid := dataset.CopyGrid(emptyGrid)
//		SolveWg(grid, workers)
//	}
//}

func BenchmarkSolveWgHard(b *testing.B) {
	b.ReportAllocs()

	workers := runtime.GOMAXPROCS(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGridHard)
		solver.SolveWg(grid, workers)
	}
}
