/*
templates:
- New
- Parse
_ Execute
*/

package main

import (
    "net/http"
    "text/template"
)

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        tmpl.Execute(w, req.URL.Path)
//        tmpl.Execute(w, req.URL.Path[1:])
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
    <h1>Hello {{.}}</h1>
</body>
</html>
`

/*
{{.}}
the period is called a "pipeline" in go docs;
the path the template will take through the data
to get to the data that needs to be injected;
the {{.}} says, "use all the data"

try these url's
http://localhost:8080/
http://localhost:8080/taiwan
http://localhost:8080/china/taiwan/singapore/yourMama




































template.New
my description: create the template "object"
go doesn't have constructors like in object oriented languages
go is type based
go's struct type often used like an object
-- holds data
-- can have methods (the struct must be a "receiver" on a function)
if any initialization is needed, then create a method that does this and returns the instance:
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template/#New

template.Parse
my description: put your template into the template "object"
func (t *Template) Parse(text string) (*Template, error)
Parse parses a string into a template. Nested template definitions will be associated
with the top-level template t. Parse may be called multiple times to parse definitions
of templates to associate with t.
http://golang.org/pkg/text/template/#Template.Parse

template.Execute
my description: merge your template with data
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
Execute applies a parsed template to the specified data object, and writes the output to wr.
If an error occurs executing the template or writing its output, execution stops,
but partial results may already have been written to the output writer.
A template may be executed safely in parallel.
http://golang.org/pkg/text/template/#Template.Execute






















FUNCTION
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template/#New

METHOD
func (t *Template) New(name string) *Template
New allocates a new template associated with the given one and with the same delimiters.
The association, which is transitive, allows one template to invoke another with a {{template}} action.
http://golang.org/pkg/text/template/#Template.New

*/