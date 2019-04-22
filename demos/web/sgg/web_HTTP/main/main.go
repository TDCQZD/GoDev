package main

import (
	"fmt"
	"net/http"
)

func handler(w  http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"HTTP协议测试")
}

func main()  {
	http.HandleFunc("/http",handler)


	http.ListenAndServe(":8080",nil)
}