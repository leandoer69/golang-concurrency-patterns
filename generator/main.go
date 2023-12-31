package main

import (
	"fmt"
	"golang-concurrency-patterns/pipelines"
)

func makeIndex(id int) int {
	index := id * 2
	return index
}

func taskGenerator(start, end int) <-chan pipelines.Task {
	ch := make(chan pipelines.Task)

	go func(ch chan<- pipelines.Task) {
		for i := start; i <= end; i++ {
			ch <- pipelines.Task{
				Id:    i,
				Index: makeIndex(i),
			}
		}
		close(ch)
	}(ch)

	return ch
}

func main() {
	for task := range taskGenerator(1, 5) {
		fmt.Println(task.Id, task.Index)
	}
}
