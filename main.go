package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ben-eh/CodingOrganizer/database"
	"github.com/ben-eh/CodingOrganizer/entry"
)

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
// }

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// htmlPage, err := ioutil.ReadFile("index.html")
	// if err != nil {
	// 	log.Fatal("Could not read index.html")
	// }

	var entries []entry.Entry
	entries = database.GetEntries()
	// str, err := json.Marshal(entries)
	// if err != nil {
	// 	fmt.Fprintf(w, "there is an error")
	// 	return
	// }

	// str := entries

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, entries)
	// tpl.ExecuteTemplate(w, "index.gohtml", ".", entries)
	// fmt.Fprintf(w, string(htmlPage), str)
}

func showEntryHandler(w http.ResponseWriter, r *http.Request) {
	entry := database.ShowEntry()

	t, _ := template.ParseFiles("show.html")
	t.Execute(w, entry)
}

func addEntryHandler(w http.ResponseWriter, r *http.Request) {
	htmlPage, err := ioutil.ReadFile("addEntry.html")
	if err != nil {
		log.Fatal("Could not read addEntry.html")
	}
	fmt.Fprintf(w, string(htmlPage))
}

func saveEntryHandler(w http.ResponseWriter, r *http.Request) {
	postedName := r.FormValue("name")
	postedURL := r.FormValue("url")
	postedBlock := r.FormValue("codeblock")
	postedNotes := r.FormValue("notes")
	if postedName != "" {
		e := &entry.Entry{
			Name:      postedName,
			URL:       postedURL,
			CodeBlock: postedBlock,
			Notes:     postedNotes,
		}
		database.SaveEntry(*e)
	}
	http.Redirect(w, r, "/", 301)
}

func initWebServer() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/addEntry", addEntryHandler)
	http.HandleFunc("/saveEntry", saveEntryHandler)
	http.HandleFunc("/showEntry", showEntryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	initWebServer()
}
