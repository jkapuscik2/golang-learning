package dataset

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

const sampleLoadDataset = `003020600
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300`

const alphanumLoadDataset = `0030206AA
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300`

const tooShortLoadDataset = `123456789`

const tooLongDataset = `003020600
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300
800203009
005010300`

const tooManyColumns = `1234567891011`

const tooFewColumns = `1234`

var sampleLoadGrid = Grid{
	{0, 0, 3, 0, 2, 0, 6, 0, 0},
	{9, 0, 0, 3, 0, 5, 0, 0, 1},
	{0, 0, 1, 8, 0, 6, 4, 0, 0},
	{0, 0, 8, 1, 0, 2, 9, 0, 0},
	{7, 0, 0, 0, 0, 0, 0, 0, 8},
	{0, 0, 6, 7, 0, 8, 2, 0, 0},
	{0, 0, 2, 6, 0, 9, 5, 0, 0},
	{8, 0, 0, 2, 0, 3, 0, 0, 9},
	{0, 0, 5, 0, 1, 0, 3, 0, 0},
}

func TestLoad(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Grid
		wantErr bool
		errType error
	}{
		{
			name: "correct dataset",
			args: args{file: strings.NewReader(sampleLoadDataset)},
			want: sampleLoadGrid,
		},
		{
			name:    "dataset with letters",
			args:    args{file: strings.NewReader(alphanumLoadDataset)},
			wantErr: true,
			errType: ErrInvalidData,
		},
		{
			name:    "too few rows",
			args:    args{file: strings.NewReader(tooShortLoadDataset)},
			wantErr: true,
			errType: ErrIncompleteData,
		},
		{
			name:    "too many rows",
			args:    args{file: strings.NewReader(tooLongDataset)},
			wantErr: true,
			errType: ErrTooManyRows,
		},
		{
			name:    "too many columns",
			args:    args{file: strings.NewReader(tooManyColumns)},
			wantErr: true,
			errType: ErrTooManyCols,
		},
		{
			name:    "too little columns",
			args:    args{file: strings.NewReader(tooFewColumns)},
			wantErr: true,
			errType: ErrIncompleteData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.file)
			if tt.wantErr && err != nil && !errors.Is(err, tt.errType) {
				t.Errorf("Load() error = %v, want %v", err.Error(), tt.errType.Error())
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzLoad(f *testing.F) {
	f.Add(sampleLoadDataset)

	f.Fuzz(func(t *testing.T, dataset string) {
		r := strings.NewReader(dataset)
		res, err := Load(r)

		if err == nil {
			err := Validate(res)
			if err != nil {
				t.Error("loaded invalid dataset")
			}
		}
	})
}
