package solver

import (
	"context"
	"github.com/jkapuscik2/sudoku-solver/internal/dataset"
	"golang.org/x/sync/semaphore"
	"sync/atomic"
	"time"
)

const timeout = 5

type publisher struct {
	blocked atomic.Bool
}

func (p *publisher) solve(ch chan<- dataset.Grid, val dataset.Grid) {
	if !p.blocked.Load() {
		p.block()
		ch <- val
	}
}

func (p *publisher) block() {
	p.blocked.Store(true)
}

func (p *publisher) attempt(ch chan<- bool, val bool) {
	if !p.blocked.Load() {
		ch <- val
	}
}

func SolveAsync(grid dataset.Grid, maxWorkers int) (dataset.Grid, error) {
	sem := semaphore.NewWeighted(int64(maxWorkers))

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	solutions := make(chan dataset.Grid, 1)

	// true - success
	// false - failure
	attempts := make(chan bool, maxWorkers)
	//defer close(attempts)

	pub := publisher{}
	defer pub.block()

	go func() {
		sem.Acquire(ctx, 1)
		guessAsync(ctx, grid, solutions, attempts, sem, &pub)
		defer sem.Release(1)
	}()

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

func guessAsync(
	ctx context.Context,
	grid dataset.Grid,
	solutions chan<- dataset.Grid,
	attempts chan<- bool,
	sem *semaphore.Weighted,
	pub *publisher,
) {
	select {
	case <-ctx.Done():
		return
	default:
		for y, row := range grid {
			for x, cell := range row {
				if dataset.IsEmpty(cell) {
					for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
						copied := grid.Rebuild(dataset.Position{X: x, Y: y}, int64(guess))

						if err := dataset.Validate(copied); err == nil {
							if copied.IsFilled() {
								pub.solve(solutions, copied)

								return
							} else {
								pub.attempt(attempts, true)

								if sem.TryAcquire(1) {
									go func(
										ctx context.Context,
										grid dataset.Grid,
										solutions chan<- dataset.Grid,
										attempts chan<- bool,
										sem *semaphore.Weighted,
										pub *publisher,
									) {
										guessAsync(ctx, copied, solutions, attempts, sem, pub)
										defer func() {
											defer sem.Release(1)
										}()
									}(ctx, copied, solutions, attempts, sem, pub)
								} else {
									// No more async workers jobs possible, solving synchronously
									guessAsync(ctx, copied, solutions, attempts, sem, pub)
								}
							}
						}
					}
					// no more possible solutions for given empty cell
					pub.attempt(attempts, false)

					return
				}
			}
		}
		if err := dataset.Validate(grid); err == nil {
			pub.solve(solutions, grid)
		} else {
			pub.attempt(attempts, false)
		}
	}
}
