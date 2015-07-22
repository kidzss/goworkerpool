package main

import "fmt"

type Payload struct {
}

type Job struct {
	Payload Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	WorkerID   int
	quit       chan bool
}

func NewWorker(workerPool chan chan Job, workerID int) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		WorkerID:   workerID,
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			println("worker working start")
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Println(job)
				fmt.Printf("worker id is %v\n", w.WorkerID)
				fmt.Println("job come job come")
			case <-w.quit:
				return
			}
			println("worker working end")
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}