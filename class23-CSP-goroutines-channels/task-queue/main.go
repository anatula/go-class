package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id       int
	Duration int
}

// takes a worker id and a receive-only channel of Task
func worker(wId int, tasks <-chan Task, done chan<- bool) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d (%d seconds) \n", wId, task.Id, task.Duration)
		time.Sleep(time.Duration(task.Duration) * time.Second)
		fmt.Printf("Worker %d finished task %d \n", wId, task.Id)
	}
	// Signal that this worker is done
	done <- true
	fmt.Printf("Worker %d exiting\n", wId)
}

func main() {
	// number of worker
	wNum := 3
	// channel buffer size
	chanBufSize := 5
	// unbuffered channel to signal done to exit main and finish the program
	done := make(chan bool)
	// create a channel with buffer 5
	tasksChannel := make(chan Task, chanBufSize)

	// start 3 workers
	for i := 1; i <= wNum; i++ {
		//  whole channel with a usage restriction
		// you can only RECEIVE from it
		go worker(i, tasksChannel, done)
	}

	// create 5 tasks with different durations
	taskList := []Task{
		{Id: 1, Duration: 2},
		{Id: 2, Duration: 1},
		{Id: 3, Duration: 3},
		{Id: 4, Duration: 2},
		{Id: 5, Duration: 1},
	}

	//send task to channel
	for _, task := range taskList {
		tasksChannel <- task
		fmt.Printf("Sent task %d to queue \n", task.Id)
	}

	// Close channel - signals workers to exit when done
	// otherwise will waits FOREVER in range taskChannel inside worker
	// for range exits when channel is closed and empty
	close(tasksChannel)

	// Sleep to let workers finish (simple alternative to WaitGroup)
	fmt.Println("All tasks sent. Waiting for workers to finish...")

	// Wait for all 3 workers to signal they're done
	for i := 0; i < wNum; i++ {
		<-done // Blocks until a worker sends a signal
		fmt.Printf("Received done signal %d/%d \n", i+1, wNum)
	}

	fmt.Println("All workers finished! Done!")

	//time.Sleep(10 * time.Second)
	//fmt.Println("Done!")

}
