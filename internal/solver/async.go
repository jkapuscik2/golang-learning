package solver

import (
	"context"
	"github.com/jkapuscik2/sudoku-solver/internal/solver/dataset"
	"time"
)

func SolveAsync(grid dataset.Grid, workers int) (dataset.Grid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	solutions := make(chan dataset.Grid, 1)
	// true - success
	// false - failure
	attempts := make(chan bool, workers)
	jobs := make(chan dataset.Grid, workers)

	for i := 0; i < workers; i++ {
		go worker(ctx, jobs, attempts, solutions)
	}
	jobs <- grid

	attemptCount, failedCount := 1, 0

	for {
		select {
		case solution := <-solutions:
			return solution, nil
		case res := <-attempts:
			if res {
				attemptCount += 1
			} else {
				failedCount += 1
			}
			if failedCount >= attemptCount {
				return grid, ErrNoSolutions
			}
		case <-ctx.Done():
			return grid, ErrTimeout
		}
	}
}

func worker(ctx context.Context, jobs chan dataset.Grid, attempts chan<- bool, solutions chan dataset.Grid) {
	for {
		select {
		case j := <-jobs:
			guessAsync(j, solutions, attempts, jobs)
		case <-ctx.Done():
			return
		}
	}
}

func guessAsync(
	grid dataset.Grid,
	solutions chan<- dataset.Grid,
	attempts chan<- bool,
	jobs chan<- dataset.Grid,
) {
	for y, row := range grid {
		for x, cell := range row {
			if dataset.IsEmpty(cell) {
				for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
					copied := grid.Rebuild(dataset.Position{X: x, Y: y}, int64(guess))

					if err := dataset.Validate(copied); err == nil {
						if copied.IsFilled() {
							solutions <- copied
							return
						} else {
							attempts <- true
							go func() {
								jobs <- copied
							}()
						}
					}
				}
				// no more possible solutions for given empty cell
				attempts <- false
				return
			}
		}
	}

	if err := dataset.Validate(grid); err == nil {
		solutions <- grid
		return
	} else {
		attempts <- false
		return
	}
}
