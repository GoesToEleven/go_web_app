package main

import (
    "net/http"
    "os"
    "html/template"
    "bufio"
    "strings"
    "log"
)

func main() {
    templates := populateTemplates()

    http.HandleFunc("/",
    func(w http.ResponseWriter, req *http.Request) {
        requestedFile := req.URL.Path[1:]
        template := templates.Lookup(requestedFile + ".html")

        if template != nil {
            template.Execute(w, nil)
        } else {
            w.WriteHeader(404)
        }
    })
    http.HandleFunc("/img/", serveResource)
    http.HandleFunc("/css/", serveResource)
    http.HandleFunc("/scripts/", serveResource)
    http.ListenAndServe(":8080", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
    path := "public" + req.URL.Path
    var contentType string
    if strings.HasSuffix(path, ".css") {
        contentType = "text/css"
    } else if strings.HasSuffix(path, ".png") {
        contentType = "image/png"
    } else if strings.HasSuffix(path, ".jpg") {
        contentType = "image/jpg"
    } else if strings.HasSuffix(path, ".svg") {
        contentType = "image/svg+xml"
    } else if strings.HasSuffix(path, ".js") {
        contentType = "application/javascript"
    } else {
        contentType = "text/plain"
    }

    log.Println(path)
    log.Println(contentType)

    f, err := os.Open(path)

    if err == nil {
        defer f.Close()
        w.Header().Add("Content Type", contentType)
        br := bufio.NewReader(f)
        br.WriteTo(w)
    } else {
        w.WriteHeader(404)
    }
}

func populateTemplates() *template.Template {
    result := template.New("templates")

    basePath := "templates"
    templateFolder, _ := os.Open(basePath)
    defer templateFolder.Close()

    templatePathsRaw, _ := templateFolder.Readdir(-1)
    // -1 means all of the contents
    templatePaths := new([]string)
    for _, pathInfo := range templatePathsRaw {
        log.Println(pathInfo.Name())
        if !pathInfo.IsDir() {
            *templatePaths = append(*templatePaths,
            basePath + "/" + pathInfo.Name())
        }
    }

    result.ParseFiles(*templatePaths...)

    return result
}


/*
MODEL
business logic & rules
data storage

VIEW
what the client sees

CONTROLLER
the glue between model & view
coordinates the model & view layers
determines how the model needs to be interacted with to meet a user's request
passes the results of the model layers work to the view layer
responsibilities:
- generate output and send it back to client
-- templates
-- bind data
- receive user actions
-- ajax
-- forms

change your references in your HTML files from this stuff ...
<link rel='stylesheet' href='../public/css/flyout_menu.css'>
... to this stuff ...
<link rel='stylesheet' href='/css/flyout_menu.css'>

test it here
http://localhost:8080/home


*/