package main

import "runtime"
import "net/http"
import "time"
import "fmt"

const (
	MAXWORKER = 3
)

func init() {
	JobQueue = make(chan Job)
	dispatcher := NewDispatcher(MAXWORKER)
	dispatcher.Run()
}

func handleFunc(w http.ResponseWriter, req *http.Request) {
	newPayload := Payload{}
	newJob := Job{Payload: newPayload}

	// Push the work onto the queue.
	go func() {
		println("push work to queue in another goroutine")
		JobQueue <- newJob
		println("push work to queue end")
	}()
	fmt.Println(time.Now().Unix())
	runtime.Gosched()
}

func main() {
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println("serve http error")
	}

}
