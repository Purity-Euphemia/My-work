package main

import (
	"fmt"
	"sync"
	"time"
)

func multipleGoroutine(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += job * i
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Worker", id, "processed job", job)
		results <- sum
	}
}

func main() {
	var numJobs int
	fmt.Print("Enter number of jobs: ")
	fmt.Scan(&numJobs)

	numWorkers := 3

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go multipleGoroutine(w, jobs, results, &wg)
	}
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for result := range results {
		total += result
	}
	fmt.Println("Final Total:", total)
	// fmt.Println("All goroutines finished")
}
