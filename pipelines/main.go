package pipelines

import "fmt"

type Task struct {
	Id    int
	Index int
	Data  int
}

// Build is needed to create input channel
func Build(in []int) <-chan Task {
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

// FillIndex is a first pipeline
func FillIndex(in <-chan Task) <-chan Task {
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

// FillData is a second pipeline
func FillData(in <-chan Task) <-chan Task {
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
	in := Build([]int{1, 2, 3})
	out := FillData(FillIndex(in))

	for task := range out {
		fmt.Println(task.Id, task.Index, task.Data)
	}
}
