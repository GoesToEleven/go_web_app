/*
we can use flow control in templates
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
    Beers []string
    Title string

}

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        context := Context{
            "Todd",
            "more beer, please",
            req.URL.Path,
            []string{"New Belgium", "La Fin Du Monde", "The Alchemist"},
            "Favorite Beers",
        }
        tmpl.Execute(w, context)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>

    <h1>{{.FirstName}} says, "{{.Message}}"</h1>

    {{if eq .URL "/nobeer"}}
        <h2>We're out of beer, {{.FirstName}}. Sorry!</h2>
    {{else}}
        <h2>Yes, grab another beer, {{.FirstName}}</h2>
        <ul>
            {{range .Beers}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    {{end}}

    <hr>

    <h2>Here's all the data:</h2>
    <p>{{.}}</p>
</body>
</html>
`

/*
range
allows you to loop over data with many items
array, slice, map, channel
when you "range" loop over data
the pipeline {{.}} gets set to the current item in the data
another way to say this: "the range operator resets the pipeline
to be the individual item in the collection"
range / else
same as range
however, if the data is len == 0, then
the else block gets executed
(eg, empty shopping cart)

sub-templates
include templates in templates
a view can include many different templates
-- call this, call that, call another thing

*/