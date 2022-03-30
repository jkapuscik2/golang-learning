package main

import (
	"fmt"
	"learning-go-d/internal/solver"
	"learning-go-d/internal/solver/dataset"
	"log"
	"os"
)

func main() {

	fileName := "data/1.txt"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Dataset file does not exist")
	}

	data, err := dataset.Load(file)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = data.Validate()

	if err != nil {
		log.Fatalf("Invalid data provided %q", err.Error())
	}

	fmt.Println("Initial grid")
	data.PrettyPrint()

	sol, err := solver.SolveSync(data)
	fmt.Println("Solved grid", err)
	sol.PrettyPrint()

	// failed attempts to do same with multiple threads

	//ch := make(chan solver.Response)
	//ch := make(chan solver.Response, 1)

	//fmt.Println(fmt.Printf("Data address in main.go: %p \n", &data))
	//fmt.Println(fmt.Printf("Grid address in main.go: %p \n", &data.Grid))

	//go solver.SolveAsync(data, ch)
	//fmt.Println("Solved grid", err)
	//sol.PrettyPrint()

	//res := <-ch

	//fmt.Println(res)
	//select {
	//case msg1 := <-ch:
	//	fmt.Println("received", msg1)
	//case <-time.After(5 * time.Second):
	//	fmt.Println("timeout")
	//}
	//fmt.Println(res)
}
