package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "os"
)

type MyHandler struct {
    http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    log.Println("PATH: ", path)
    log.Println("ENV: ", os.Environ())
    data, err := ioutil.ReadFile(string(path))

    if err == nil {
        w.Write(data)
    } else {
        w.WriteHeader(404)
        w.Write([]byte("404 My Friend - " + http.StatusText(404)))
    }
}

func main() {
    http.Handle("/", new(MyHandler))
    http.ListenAndServe(":8080", nil)
}