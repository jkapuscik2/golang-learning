package solver

import (
	"context"
	"github.com/jkapuscik2/sudoku-solver/internal/dataset"
	"golang.org/x/sync/semaphore"
	"sync"
	"sync/atomic"
)

func SolveWg(grid dataset.Grid, maxWorkers int) (dataset.Grid, error) {
	var solved atomic.Bool
	sem := semaphore.NewWeighted(int64(maxWorkers))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	solutions := make(chan dataset.Grid, 1)
	defer close(solutions)

	go func() {
		sem.TryAcquire(1)
		guessWg(ctx, grid, solutions, sem, &wg, &solved)
		defer sem.Release(1)
	}()

	wg.Wait()

	select {
	case solution := <-solutions:
		return solution, nil
	default:
		return grid, ErrNoSolutions
	}
}

func guessWg(
	ctx context.Context,
	grid dataset.Grid,
	solutions chan<- dataset.Grid,
	sem *semaphore.Weighted,
	wg *sync.WaitGroup,
	solved *atomic.Bool,
) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		if solved.Load() {
			return
		}
		for y, row := range grid {
			for x, cell := range row {
				if dataset.IsEmpty(cell) {
					for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
						copied := grid.Rebuild(dataset.Position{X: x, Y: y}, int64(guess))

						if err := dataset.Validate(copied); err == nil {
							if copied.IsFilled() {
								select {
								case <-ctx.Done():
									return
								default:
									if !solved.Load() {
										solutions <- copied
										solved.Store(true)
									}
								}

								return
							} else {
								wg.Add(1)
								if sem.TryAcquire(1) {
									go func(
										ctx context.Context,
										grid dataset.Grid,
										solutions chan<- dataset.Grid,
										sem *semaphore.Weighted,
										wg *sync.WaitGroup,
									) {
										guessWg(ctx, copied, solutions, sem, wg, solved)
										defer func() {
											defer sem.Release(1)
										}()
									}(ctx, copied, solutions, sem, wg)
								} else {
									// No more async workers jobs possible, solving synchronously
									guessWg(ctx, copied, solutions, sem, wg, solved)
								}
							}
						}
					}
					return
				}
			}
		}
		if err := dataset.Validate(grid); err == nil {
			select {
			case <-ctx.Done():
				return
			default:
				if !solved.Load() {
					solutions <- grid
					solved.Store(true)
				}
			}
		}
	}
}
