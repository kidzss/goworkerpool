package main

//import "fmt"
//import "github.com/go-martini/martini"
import "runtime"
import "net/http"
import "time"
import "fmt"

func handleFunc(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Unix()
	fmt.Println(t)
	payload := Payload{}
	work := Job{Payload: payload}

	// Push the work onto the queue.
	go func() {
		println("fuck start")
		println(JobQueue)
		JobQueue <- work
		println("fuck middle")
		println("fuck en")
	}()
	fmt.Println(time.Now().Unix())
	runtime.Gosched()
}

func main() {
	//m := martini.Classic()
	MaxWorker := 3
	JobQueue = make(chan Job)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	/*m.Get("/", func() string {
		payload := Payload{}
		work := Job{Payload: payload}

		// Push the work onto the queue.
		go func() {
			println("fuck start")
			println(JobQueue)
			JobQueue <- work
			println("fuck middle")
			println("fuck en")
		}()
		runtime.Gosched()
		return "hello world"
	})
	*/
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println("fuck!")
	}

	/*go func() {
		for {
			select {
			case x := <-JobQueue:
				println("job come")
				fmt.Println(x)
				println("job end")
			}
		}
	}()
	*/

	//m.RunOnAddr(":8080")
}
