package app 

import (
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

func WebIntro() {
    http.HandleFunc("/", defaultHandler)
    err := http.ListenAndServe(":9999", nil)
    fmt.Println(err)
}

