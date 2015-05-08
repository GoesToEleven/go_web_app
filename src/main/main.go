package main

import (
    "net/http"
//    "text/template"
    "html/template"
)

//var Message string = "more beer, please sir"
//var Message string = "alert('you have been pwned')"
var Message string = "<script>alert('you have been pwned, BIATCH')</script>"

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        tmpl.Execute(w, Message)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>Injection Safe</title>
</head>
<body>

    <p>{{.}}</p>

    <script>{{.}}</script>

</body>
</html>
`

/*
html/template
Package template (html/template) implements data-driven templates for generating
HTML output safe against code injection. It provides the same interface as package
text/template and should be used instead of text/template whenever the output is HTML.

HTML templates treat data values as plain text which should be encoded so they can be
safely embedded in an HTML document. The escaping is contextual, so actions can appear
within JavaScript, CSS, and URI contexts.

http://golang.org/pkg/html/template/

to run the above code ...
try the different Message variables with text/template import
... then ...
try the different Message variables with html/template import


*/