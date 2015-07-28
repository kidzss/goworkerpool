package main

import "fmt"

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	println(maxWorkers)
	pool := make(chan chan Job, maxWorkers)
	fmt.Printf("make worker pool of dispatcher :%v workers\n", maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		fmt.Printf("new worker, worker id: %v\n", i)
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		println("dispatch start")
		println(JobQueue)
		select {
		case job := <-JobQueue:
			go func(job Job) {
				//从全局的workerpool中拿到一条jobChannel
				//然后把当前的job塞到这条jobChannel
				//然后这条jobChannel对应的worker就可以拿到job去处理了
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
		println("dispatch end")
	}
}
