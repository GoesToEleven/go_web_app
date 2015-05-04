/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

SERVEMUX
ServeMux = HTTP request router = multiplexor = mux
compares incoming requests against a list of predefined URL paths,
and calls the associated handler for the path whenever a match is found.

HANDLERS
responsible for writing response headers and bodies.
Almost any type ("object") can be a handler, so long as it satisfies the http.Handler interface.
In lay terms, that simply means it must have a ServeHTTP method with the following signature:
ServeHTTP(http.ResponseWriter, *http.Request)

*/

package main

import (
    "net/http"
)

func main() {
    myMux := http.NewServeMux()
    myMux.HandleFunc("/", someFunc)
    http.ListenAndServe(":8080", myMux)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello Los Angeles"))
}

/*
from:
http://www.alexedwards.net/blog/a-recap-of-request-handling

before:
func main() {
    http.HandleFunc("/", someFunc)
    http.ListenAndServe(":8080", nil)
}
*/