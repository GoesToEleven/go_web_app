/*
RESOURCE SERVER
- listen on a TCP port
- handle requests: route a URL to a file

ServeMux = HTTP request router = multiplexor = mux
http://www.alexedwards.net/blog/a-recap-of-request-handling

Mutexes are something else
http://www.alexedwards.net/blog/understanding-mutexes

NICE TO KNOW
www.rawgit.com
https://cdn.rawgit.com/GoesToEleven/go_web_app/01_working_static_page/templates/home.html

NET/HTTP

- http.ListenAndServe
-- listens for, and responds to, http requests
-- handles each request using go routines
--- lightweight concurrency (eg, coroutines - processes --> threads --> coroutines)
--- this is multiplexing, thus, multiplexor ( = HTTP request router = ServeMux = mux )
--- blacks main thread (call after configuration of server complete)

- http.Handle
-- handles a URL request
-- maps a URL to any TYPE ("object") implementing the handler interface
--- http://golang.org/pkg/net/http/#Handler

- http.HandleFunc
-- handles a URL request
-- maps a URL to a FUNCTION
--- "wrapper" around a function
---- turns any function into a handler

Handle -> handler
HandleFunc -> handlerFunc


*/

package main

import "net/http"

func main() {
    http.HandleFunc("/", someFunc)
    http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello universe"))
}

/*
http.HandleFunc("/", someFunction)
Go matches requests to the most specific route registered
http://golang.org/pkg/net/http/#ServeMux
everything matches "/"

nil
- http.ListenAndServe(":8080", nil)
- meaning: use the DefaultServeMux
http://golang.org/pkg/net/http/#pkg-variables

behind the scenes:
- request comes in
- received on primary thread
- goroutine created
-- runs concurrently to main thread
-- lightweight
- request passed to goroutine
- handling multiple requests at same time: "multiplexing"

*/
