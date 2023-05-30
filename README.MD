## Sudoku solver

[![Go Report Card](https://goreportcard.com/badge/github.com/jkapuscik2/sudoku-solver)](https://goreportcard.com/report/github.com/jkapuscik2/sudoku-solver)

The script solves sudoku grid provided in a form of .txt file

The script was intended as a learning practise and solves the provided grid in a following ways:

- synchronously using a backtracking approach
- using multiple goroutines and communication via channels 
- using wait groups

```bash
go build main.go

./main
```

Available flags:
```
  -help
        Display help 
  -path string
        Path to fle with sudoku (default "data/1.txt")
  -profile
        If application should be profiled
  -maxWorkers int
        Number of maxWorkers (default number of avaliable CPU)
```


Testing

```bash
go test  ./... -v -bench=.
```


