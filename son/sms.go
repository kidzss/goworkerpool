package main

import "net/http"
import "strings"

//import "fmt"
import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		var str = "hello world"
		str += martini.Env
		return str
	})

	m.Post("/send_sms", func(res http.ResponseWriter, req *http.Request) string {
		var str = "post hello world"
		req.ParseForm()
		for k, v := range req.Form {
			str += string(k)
			str += strings.Join(v, "")
		}
		return str
	})

	m.RunOnAddr(":8080")
}
