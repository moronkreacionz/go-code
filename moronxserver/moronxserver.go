package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/go-code/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })

    http.HandleFunc("/go-code/greet/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/greet/"):]
        fmt.Fprintf(w, "Hello %s\n", name)
    })

    http.ListenAndServe(":8001", nil)
}
