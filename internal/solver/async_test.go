package solver

import (
	"errors"
	"github.com/jkapuscik2/sudoku-solver/internal/solver/dataset"
	"reflect"
	"runtime"
	"testing"
)

var emptyGrid = dataset.Grid{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var sampleGrid = dataset.Grid{
	{0, 2, 0, 8, 1, 0, 7, 4, 0},
	{7, 0, 0, 0, 0, 3, 1, 0, 0},
	{0, 9, 0, 0, 0, 2, 8, 0, 5},
	{0, 0, 9, 0, 4, 0, 0, 8, 7},
	{4, 0, 0, 2, 0, 8, 0, 0, 3},
	{1, 6, 0, 0, 3, 0, 2, 0, 0},
	{3, 0, 2, 7, 0, 0, 0, 6, 0},
	{0, 0, 5, 6, 0, 0, 0, 0, 8},
	{0, 7, 6, 0, 5, 1, 0, 9, 0},
}

var sampleInvalidGrid = dataset.Grid{
	{0, 2, 0, 8, 1, 0, 7, 4, 0},
	{7, 0, 0, 0, 0, 3, 1, 0, 0},
	{0, 9, 0, 0, 0, 2, 8, 0, 5},
	{0, 0, 9, 0, 4, 0, 0, 8, 7},
	{4, 0, 0, 2, 0, 8, 0, 0, 3},
	{1, 6, 0, 0, 3, 0, 2, 0, 0},
	{3, 0, 2, 7, 0, 0, 0, 6, 0},
	{0, 0, 5, 6, 0, 0, 0, 0, 8},
	{0, 7, 6, 0, 5, 1, 0, 9, 9},
}

var sampleGridSolved = dataset.Grid{
	{5, 2, 3, 8, 1, 6, 7, 4, 9},
	{7, 8, 4, 5, 9, 3, 1, 2, 6},
	{6, 9, 1, 4, 7, 2, 8, 3, 5},
	{2, 3, 9, 1, 4, 5, 6, 8, 7},
	{4, 5, 7, 2, 6, 8, 9, 1, 3},
	{1, 6, 8, 9, 3, 7, 2, 5, 4},
	{3, 4, 2, 7, 8, 9, 5, 6, 1},
	{9, 1, 5, 6, 2, 4, 3, 7, 8},
	{8, 7, 6, 3, 5, 1, 4, 9, 2},
}

var sampleGridSolvedInvalid = dataset.Grid{
	{5, 5, 5, 5, 5, 6, 7, 4, 9},
	{7, 8, 4, 5, 9, 3, 1, 2, 6},
	{6, 9, 1, 4, 7, 2, 8, 3, 5},
	{2, 3, 9, 1, 4, 5, 6, 8, 7},
	{4, 5, 7, 2, 6, 8, 9, 1, 3},
	{1, 6, 8, 9, 3, 7, 2, 5, 4},
	{3, 4, 2, 7, 8, 9, 5, 6, 1},
	{9, 1, 5, 6, 2, 4, 3, 7, 8},
	{8, 7, 6, 3, 5, 1, 4, 9, 2},
}

func TestSolveAsync(t *testing.T) {
	type args struct {
		grid    dataset.Grid
		workers int
	}
	tests := []struct {
		name    string
		args    args
		want    dataset.Grid
		wantErr bool
		errType error
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
			errType: ErrNoSolutions,
		},
		{
			name:    "solved dataset",
			args:    args{grid: sampleGridSolved, workers: runtime.NumCPU()},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "wrongly solved dataset",
			args:    args{grid: sampleGridSolvedInvalid, workers: runtime.NumCPU()},
			want:    sampleGridSolvedInvalid,
			wantErr: true,
			errType: ErrNoSolutions,
		},
		//{
		//	name:    "empty dataset",
		//	args:    args{grid: emptyGrid, workers: 1},
		//	want:    emptyGrid,
		//	wantErr: false,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveAsync(tt.args.grid, tt.args.workers)
			if tt.wantErr && err != nil && !errors.Is(err, tt.errType) {
				t.Errorf("SolveAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveAsync() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolveAsync(b *testing.B) {
	b.ReportAllocs()

	workers := runtime.GOMAXPROCS(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		SolveAsync(grid, workers)
	}
}
