package main

import (
    "fmt"
    "net/http"
)

func example(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            fmt.Fprintf(w, "A GET request was received!\n")
        case "POST":
            fmt.Fprintf(w, "A POST request was received!\n")
        default:
            fmt.Fprintf(w, "Some other request was received: %s\n", r.Method)
    }
}

func main() {
    http.HandleFunc("/method", example)
    http.ListenAndServe(":8080", nil)
}


