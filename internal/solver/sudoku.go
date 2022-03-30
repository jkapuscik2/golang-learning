package solver

import (
	"errors"
	"learning-go-d/internal/solver/dataset"
)

func SolveSync(d dataset.Dataset) (dataset.Dataset, error) {
	for y, row := range d.Data() {
		for x, cell := range row {
			if isEmpty(cell) {
				for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
					initVal, err := d.GetValue(dataset.Position{X: x, Y: y})

					if err != nil {
						return d, err
					}

					d.SetValue(dataset.Position{X: x, Y: y}, int64(guess))

					if err := d.Validate(); err == nil {
						sol, err := SolveSync(d)

						if err == nil {
							return sol, nil
						}
					}
					// Wrong guess, restoring initial value
					if err = d.SetValue(dataset.Position{X: x, Y: y}, initVal); err != nil {
						return d, err
					}
				}
				return d, errors.New("failed to solve the board")
			}
		}
	}

	return d, nil
}

type Response struct {
	D   dataset.Dataset
	Err error
}

func SolveAsync(d dataset.Dataset, ch chan<- Response) (dataset.Dataset, error) {
	for y, row := range d.Data() {
		for x, cell := range row {
			if isEmpty(cell) {
				for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
					initVal, err := d.GetValue(dataset.Position{X: x, Y: y})

					if err != nil {
						return d, err
					}

					d.SetValue(dataset.Position{X: x, Y: y}, int64(guess))

					if err := d.Validate(); err == nil {
						sol, err := SolveSync(d)

						if err == nil {
							ch <- Response{d, nil}
							return sol, nil
						}
					}
					// Wrong guess, restoring initial value
					if err = d.SetValue(dataset.Position{X: x, Y: y}, initVal); err != nil {
						return d, err
					}
				}
				return d, errors.New("failed to solve the board")
			}
		}
	}

	ch <- Response{d, nil}
	return d, nil
}

func isEmpty(cell int64) bool {
	return int(cell) == dataset.EmptyVal
}
