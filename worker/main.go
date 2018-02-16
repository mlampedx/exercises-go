package main

import "github.com/mlampedx/exercises-go/worker/worker"

func main() {
	jobs := make(chan int, 100)
	defer close(jobs)

	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker.Create(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	for a := 1; a <= 5; a++ {
		<-results
	}
}
