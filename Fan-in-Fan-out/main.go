package main

import (
	"fmt"
	"golang-concurrency-patterns/pipelines"
	"sync"
)

func mergeTasks(in ...<-chan pipelines.Task) <-chan pipelines.Task {
	wg := &sync.WaitGroup{}
	out := make(chan pipelines.Task)

	output := func(c <-chan pipelines.Task) {
		for task := range c {
			out <- task
		}
		wg.Done()
	}

	wg.Add(len(in))
	for _, v := range in {
		go output(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := pipelines.Build([]int{1, 2, 3})
	// fan-out
	out1 := pipelines.FillIndex(in)
	out2 := pipelines.FillIndex(in)

	// fan-in
	for task := range mergeTasks(out1, out2) {
		fmt.Println(task.Id, task.Index)
	}

}
