package main
 
import (
    "fmt"
    "sync"
)
 
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        // Perform the task associated with the job
        result := job * 2
 
        // Send the result back through the results channel
        results <- result
    }
}
 
func main() {

	fmt.Println("Hello ...!")

    numJobs := 10
    numWorkers := 3
 
    // Create the channels for communication
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
 
    // Create a wait group to wait for all workers to finish
    var wg sync.WaitGroup
 
    // Launch the worker goroutines
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            worker(workerID, jobs, results)
        }(i)
    }
 
    // Send jobs to the jobs channel
    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }
    close(jobs)
 
    // Collect the results from the results channel
    go func() {
        for i := 1; i <= numJobs; i++ {
            result := <-results
            fmt.Println("Result:", result)
        }
    }()
 
    // Wait for all workers to finish
    wg.Wait()
    close(results)
}