package solver

import (
	"errors"
	"github.com/jkapuscik2/sudoku-solver/internal/solver/dataset"
	"reflect"
	"testing"
)

func TestSolveBacktrace(t *testing.T) {
	type args struct {
		grid dataset.Grid
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
			args:    args{grid: sampleGrid},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "unsolvable dataset",
			args:    args{grid: sampleInvalidGrid},
			want:    sampleInvalidGrid,
			wantErr: true,
			errType: ErrNoSolutions,
		},
		{
			name:    "solved dataset",
			args:    args{grid: sampleGridSolved},
			want:    sampleGridSolved,
			wantErr: false,
		},
		{
			name:    "wrongly solved dataset",
			args:    args{grid: sampleGridSolvedInvalid},
			want:    sampleGridSolvedInvalid,
			wantErr: true,
			errType: ErrNoSolutions,
		},
		{
			name:              "empty dataset",
			args:              args{grid: emptyGrid},
			want:              emptyGrid,
			multipleSolutions: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveBacktrace(tt.args.grid)
			if tt.wantErr && err != nil && !errors.Is(err, tt.errType) {
				t.Errorf("SolveBacktrace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.multipleSolutions && !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveBacktrace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTestSolveBacktrace(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		grid := dataset.CopyGrid(sampleGrid)
		SolveBacktrace(grid)
	}
}
