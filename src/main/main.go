/*
create a struct to hold your data
*/

package main

import (
    "net/http"
    "text/template"
)

type Context struct {
    FirstName string
    Message string
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>First Template</title>
</head>
<body>
    <h1>My name is {{.FirstName}}</h1>
    <p>{{.Message}}</p>
</body>
</html>
`

func toddFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"Todd", "more Go, please"}
        tmpl.Execute(w, context)
    }
}

func main() {
    http.HandleFunc("/todd", toddFunc)
    http.HandleFunc("/ming", mingFunc)
    http.HandleFunc("/rio", rioFunc)
    http.HandleFunc("/", jamesFunc)
    http.ListenAndServe(":8080", nil)
}














func mingFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"Ming", "I am a problem solver!"}
        tmpl.Execute(w, context)
    }
}

func rioFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"Rio", "I drank the google-aid"}
        tmpl.Execute(w, context)
    }
}

func jamesFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"James", "Another beer, please"}
        tmpl.Execute(w, context)
    }
}

/*
the pipeline starts at the data {{.}}
then finds the child specified:
-- fName
-- Message

pipelines can be as deep as you want

nested objects and members
accessed through drill-down using dot operator
*/