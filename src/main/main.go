package main

import (
    "net/http"
    "os"
    "html/template"
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


    http.ListenAndServe(":8080", nil)
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


steps:
(1) create a template "cache"
-- one template to hold other templates
-- all of the "held" templates will be siblings of each other
--- this means the "held" templates can call/include each other
--- templates can call/include sibling & descendent templates

test it here
http://localhost:8080/home


*/