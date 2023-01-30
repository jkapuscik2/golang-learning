package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"learning-go-sudoku/internal/solver"
	"learning-go-sudoku/internal/solver/dataset"
	"log"
	"os"
	"runtime"
)

func main() {
	filePath := flag.String("path", "data/1.txt", "Path to fle with sudoku")
	prof := flag.Bool("profile", false, "If application should be profiled")
	workers := flag.Int("workers", runtime.GOMAXPROCS(0), "Number of workers")
	help := flag.Bool("help", false, "Display help ")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *prof {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}

	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	data, err := dataset.Load(file)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Initial grid")
	dataset.PrettyPrint(data)

	sync(dataset.CopyGrid(data))

	async(dataset.CopyGrid(data), *workers)
}

func sync(grid dataset.Grid) {
	fmt.Println("Solving sudoku synchronously")

	solution, err := solver.SolveBacktrace(grid)
	if err != nil {
		fmt.Println("Failed to solve grid", err)
	} else {
		dataset.PrettyPrint(solution)
	}
}

func async(grid dataset.Grid, workers int) {
	fmt.Println("Solving sudoku asynchronously")

	solution, err := solver.SolveAsync(grid, workers)
	if err != nil {
		fmt.Println("Failed to solve grid", err)
	} else {
		dataset.PrettyPrint(solution)
	}
}
