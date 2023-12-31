package main

import "fmt"

type Task struct {
	Id    int
	Index int
	Data  int
}

// creating input channel
func build(in []int) <-chan Task {
	out := make(chan Task)

	go func() {
		for _, n := range in {
			out <- Task{Id: n}
		}
		close(out)
	}()

	return out
}

// simple job
func makeIndex(id int) int {
	return id - 10
}

// pipeline
func fillIndex(in <-chan Task) <-chan Task {
	out := make(chan Task)

	go func() {
		for task := range in {
			task.Index = makeIndex(task.Id)
			out <- task
		}
		close(out)
	}()

	return out
}

// second pipeline
func fillData(in <-chan Task) <-chan Task {
	out := make(chan Task)

	go func() {
		for task := range in {
			// some job
			task.Data = task.Index * 2
			out <- task
		}
		close(out)
	}()

	return out
}

func main() {
	in := build([]int{1, 2, 3})
	out := fillData(fillIndex(in))

	for task := range out {
		fmt.Println(task.Id, task.Index, task.Data)
	}
}
