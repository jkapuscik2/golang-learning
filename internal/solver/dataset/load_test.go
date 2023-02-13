package dataset

import (
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
	r := strings.NewReader(sampleLoadDataset)

	res, err := Load(r)

	if err != nil {
		t.Error("could not load sample grid")
	}
	if !reflect.DeepEqual(res, sampleLoadGrid) {
		t.Error("invalid grid loaded")
	}
}

func FuzzLoad(f *testing.F) {
	f.Add(sampleLoadDataset)

	f.Fuzz(func(t *testing.T, orig string) {
		r := strings.NewReader(orig)

		_, err := Load(r)

		if err == nil {
			t.Error("loaded invalid dataset")
		}
	})
}
