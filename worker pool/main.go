package wp

import (
	"fmt"
	"golang-concurrency-patterns/pipelines"
	"time"
)

func makeIndex(id int) int {
	index := id * 2
	return index
}

func taskWorker(id int, jobs <-chan pipelines.Task, result chan<- pipelines.Task) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %v\n", id, j)
		j.Index = makeIndex(j.Id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %v\n", id, j)
		result <- j
	}
}

func main() {
	inputData := []int{1, 2, 3, 4, 5}
	maxBuffer := 100 // could be in config
	workerCount := 3 // could be in config

	tasks := make(chan pipelines.Task, maxBuffer) // jobs
	result := make(chan pipelines.Task, maxBuffer)

	for w := 0; w < workerCount; w++ {
		go taskWorker(w, tasks, result)
	}

	for _, id := range inputData {
		tasks <- pipelines.Task{Id: id}
	}
	close(tasks)

	for i := 0; i < len(inputData); i++ {
		<-result
	}
}
