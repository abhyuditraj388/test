package main

import "fmt"

type Result struct {
	id        int
	afterwork int
}

func worker(id int, jobs <-chan int, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		// Simulate work
		results <- Result{id: job, afterwork: job * 2} // Send result back
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan Result, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		temp := <-results
		fmt.Printf("Id : %d, result : %d\n", temp.id, temp.afterwork)
	}
}
