package solver

import (
	"context"
	"errors"
	"learning-go-sudoku/internal/solver/dataset"
	"time"
)

func SolveAsync(grid dataset.Grid, workers int) (dataset.Grid, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	solutions := make(chan dataset.Grid)
	// true - success
	// false - failure
	attempts := make(chan bool)
	jobs := make(chan dataset.Grid)

	for i := 0; i < workers; i++ {
		go worker(ctx, jobs, attempts, solutions)
	}
	jobs <- grid

	var attemptCount int
	var failedCount int

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
				return grid, errors.New("failed to solve the board")
			}
		case <-time.After(time.Second * 5):
			return grid, errors.New("timeout. Failed to solve the board")
		}
	}
}

func worker(ctx context.Context, jobs chan dataset.Grid, attempts chan<- bool, solutions chan dataset.Grid) {

	for {
		select {
		case j := <-jobs:
			attempts <- true
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
	attempts <- true

	for y, row := range grid {
		for x, cell := range row {
			if dataset.IsEmpty(cell) {
				for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
					copied, err := grid.Rebuild(dataset.Position{X: x, Y: y}, int64(guess))
					if err != nil {
						attempts <- false
						return
					}

					if err := dataset.Validate(copied); err == nil {
						if copied.IsFilled() {
							solutions <- copied
							return
						} else {
							go func() {
								jobs <- copied
							}()
						}
					}
				}
				attempts <- false
				// no more possible solutions for given grid
				return
			}
		}
	}

	if err := dataset.Validate(grid); err == nil {
		solutions <- grid
	} else {
		attempts <- false
	}
}
