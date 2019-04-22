package main

import (
    "fmt"
    "net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        handleRouted(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func handleRouted(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Custom Router!")
}

func main() {
    mux := &MyMux{}
    http.ListenAndServe(":9090", mux)
}