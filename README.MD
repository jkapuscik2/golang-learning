### Learning go

The script solves sudoku grid provided in a form of .txt file

The script solves the provided sudoku in sync and async manner

```bash
go build main.go

./main
```

Possible flags:
```
  -help
        Display help 
  -path string
        Path to fle with sudoku (default "data/1.txt")
  -profile
        If application should be profiled
  -workers int
        Number of workers (default number of avaliable CPU)
```


Testing

```bash
go test  ./... -v -bench=.
```

Sources:
- https://go.dev/blog/context
- https://go.dev/blog/laws-of-reflection
- https://go.dev/doc/articles/race_detector
- https://go.dev/blog/slices-intro
- https://go.dev/security/fuzz/
