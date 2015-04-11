package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type repo struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Slug        string `json:"slug"`
}

type repos []repo

const indexTemplateHtml = `
<!doctype html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="x-ua-compatible" content="ie=edge">

        <title>Nylar</title>

        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <link href="http://fonts.googleapis.com/css?family=Open+Sans:400,600,700" rel="stylesheet" type="text/css">
        <link href="/public/css/master.css" rel="stylesheet" />
        <link href="/public/css/colours.css" rel="stylesheet" />
    </head>

    <body>
        <section id="wrap">
            <h1>Nylar</h1>
            <h2>A polyglot software engineer, based in the UK.</h2>
            <h3>Projects</h3>
            {{range .}}
            <section class="project">
                <p><a href="{{.Url}}">{{.Name}}</a></p> <span class="lang {{.Slug}}">{{.Language}}</span>
                <span class="description">{{.Description}}</span>
            </section>
            {{else}}
            <section>No projects found :(</section>
            {{end}}
        </section>
    </body>
</html>
`

var indexTemplate = template.Must(template.New("index").Parse(indexTemplateHtml))

func index(w http.ResponseWriter, r *http.Request) {
	repositories := new(repos)
	projects, err := ioutil.ReadFile("projects.json")
	if err != nil {
		log.Println(err.Error())
	}

	if err := json.Unmarshal(projects, repositories); err != nil {
		log.Println(err.Error())
	}

	if err := indexTemplate.Execute(w, repositories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	portFlag := flag.Int("port", 8080, "Port to serve on.")
	flag.Parse()

	port := fmt.Sprintf(":%d", *portFlag)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)

	http.ListenAndServe(port, nil)
}
