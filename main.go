package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ben-eh/CodingOrganizer/database"
	"github.com/ben-eh/CodingOrganizer/entry"
	"github.com/gorilla/mux"
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

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, entries)
	// tpl.ExecuteTemplate(w, "index.gohtml", ".", entries)
	// fmt.Fprintf(w, string(htmlPage), str)
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

func addEntryHandler(w http.ResponseWriter, r *http.Request) {
	htmlPage, err := ioutil.ReadFile("addEntry.html")
	if err != nil {
		log.Fatal("Could not read addEntry.html")
	}
	fmt.Fprintf(w, string(htmlPage))
}

func showEntryHandler(w http.ResponseWriter, r *http.Request) {
	entry := database.FetchEntry(r)

	t, _ := template.ParseFiles("templates/show.html")
	t.Execute(w, entry)
}

func editEntryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("this msg should pop up as soon as I hit edit entry")
	entry := database.FetchEntry(r)

	t, _ := template.ParseFiles("templates/editEntry.html")
	t.Execute(w, entry)
}

func updateEntryHandler(w http.ResponseWriter, r *http.Request) {
	// postedID, _ := strconv.Atoi(r.FormValue("id"))
	vars := mux.Vars(r)

	entryID := vars["entry_id"]
	log.Println("this is the entry_id in the updateEntryHandler: ", entryID)
	postedName := r.FormValue("name")
	postedURL := r.FormValue("url")
	postedBlock := r.FormValue("codeblock")
	postedNotes := r.FormValue("notes")
	if postedName != "" {
		e := &entry.Entry{
			// ID:        postedID,
			Name:      postedName,
			URL:       postedURL,
			CodeBlock: postedBlock,
			Notes:     postedNotes,
		}
		database.UpdateEntry(r, *e)
	}
	http.Redirect(w, r, "/", 301)
}

func deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	database.DeleteEntry(r)
	http.Redirect(w, r, "/", 301)
}

// func initWebServer() {
// r := mux.NewRouter()
// r.HandleFunc("/", indexHandler)
// r.HandleFunc("/addEntry", addEntryHandler)
// r.HandleFunc("/saveEntry", saveEntryHandler)
// r.HandleFunc("/showEntry/{entry_id}", showEntryHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	// initWebServer()
	r := mux.NewRouter()
	// r.Host("http://localhost/8080")
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/addEntry", addEntryHandler)
	r.HandleFunc("/saveEntry", saveEntryHandler)
	r.HandleFunc("/showEntry/{entry_id}", showEntryHandler)
	r.HandleFunc("/editEntry/{entry_id}", editEntryHandler)
	r.HandleFunc("/updateEntry/{entry_id}", updateEntryHandler)
	r.HandleFunc("/deleteEntry/{entry_id}", deleteEntryHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
	// http.Handle("/", r)
}
