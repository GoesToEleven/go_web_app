package main

import (
    "net/http"
    "io/ioutil"
    "log"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
//    path := "templates" + r.URL.Path
    log.Println("PATH: ", path)

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

/*
structuring go apps

*****************
SUPER IMPORTANT
*****************

ROOT is determined by ...
WHERE YOU ARE IN YOUR DIRECTORY WHEN YOU RUN "GO RUN"
wherever you are in your dir when you run "go run" - that's your root

for example
if this is your file structure
go_web_app
-- bin
-- pkg
-- src
---- main
------ main.go
-- templates
---- home.html

if you were in "go_web_app / src / main"
and you ran "go run main.go"
then the root of your files would be "go_web_app / src / main"
and you would not be able to access "templates" nor "templates/home.html"

if you were in "go_web_app"
and you ran "go run main.go"
then the root of your files would be "go_web_app"
and you would be able to access "templates" and "templates/home.html"

this is a good article - though advanced (so don't freak out if you're just getting started)
https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091

this can also be useful in debugging:
-- import "os"
---- log.Println("ENV: ", os.Environ())
*/