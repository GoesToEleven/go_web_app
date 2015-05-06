/*
we can use conditional logic in templates
*/

package main

import (
    "net/http"
    "text/template"
)

type Context struct {
    FirstName string
    Message string
    URL string
}

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{"Todd", "more beer, please", req.URL.Path}
        tmpl.Execute(w, context)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>First Template</title>
</head>
<body>
    {{if eq .URL "/nobeer"}}
        <h1>We're out of beer. Sorry!</h1>
    {{else}}
        <h1>Yes, grab another beer, {{.FirstName}}</h1>
    {{end}}

    <hr>

    <h2>Here's all the data:</h2>
    <p>{{.}}</p>
</body>
</html>
`

/*
conditionals
if
if / else
if / else if

testing
eq - equal
---- an unlimited number of conditions can be tested against the first term
------ eq 1 (0+1) (2-1)
------ if they all evaluate to be the same, the test evaluates to TRUE
------ operator is listed first, followed by operands
ne - not equal
lt - less than
---- first arg compared to second
---- first condition less than second - evals to true
gt - greater than
le - less than or equal to
ge - greater than or equal to



*/