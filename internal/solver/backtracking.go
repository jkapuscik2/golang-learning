package solver

import (
	"github.com/jkapuscik2/sudoku-solver/internal/dataset"
)

func SolveBacktrace(grid dataset.Grid) (dataset.Grid, error) {
	sol, err := solveBacktrace(grid)
	if err == nil {
		return sol, nil
	} else {
		return grid, err
	}
}

func solveBacktrace(grid dataset.Grid) (dataset.Grid, error) {
	for y, row := range grid {
		for x, cell := range row {
			if dataset.IsEmpty(cell) {
				for guess := dataset.MinVal; guess <= dataset.MaxVal; guess++ {
					initVal, err := grid.GetValue(dataset.Position{X: x, Y: y})

					if err != nil {
						return grid, err
					}

					grid[y][x] = int64(guess)

					if err := dataset.Validate(grid); err == nil {
						sol, err := solveBacktrace(grid)

						if err == nil {
							return sol, nil
						}
					}

					// Wrong guess, restoring initial value
					grid[y][x] = initVal
				}
				return grid, ErrNoSolutions
			}
		}
	}
	if err := dataset.Validate(grid); err == nil {
		return grid, nil
	} else {
		return grid, ErrNoSolutions
	}
}
