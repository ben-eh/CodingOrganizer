package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ben-eh/CodingOrganizer/entry"
	"github.com/ben-eh/CodingOrganizer/entryHasTag"
	"github.com/gorilla/mux"
)

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
// }

type IndexData struct {
	Entries []entry.Entry
	Tags    []entry.Tag
}

type SoloData struct {
	Entry   entry.Entry
	AllTags []entry.Tag
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var entries []entry.Entry
	entries = entry.GetEntries()
	var tags []entry.Tag
	tags = entry.GetAllTags()

	data := IndexData{
		Entries: entries,
		Tags:    tags,
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, data)
}

func saveEntryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	tags := r.Form["[]tags"]
	log.Println(tags)
	log.Println("pause")

	if r.FormValue("name") != "" {
		e := &entry.Entry{
			Name:      r.FormValue("name"),
			URL:       r.FormValue("url"),
			CodeBlock: r.FormValue("codeblock"),
			Notes:     r.FormValue("notes"),
		}
		entryID := entry.SaveEntry(*e)
		entryHasTag.CreateEntryHasTag(tags, entryID)
	}

	http.Redirect(w, r, "/", 301)
}

func addEntryHandler(w http.ResponseWriter, r *http.Request) {

	var allTags []entry.Tag
	allTags = entry.GetAllTags()

	t, _ := template.ParseFiles("templates/addEntry.html")
	t.Execute(w, allTags)

}

func showEntryHandler(w http.ResponseWriter, r *http.Request) {
	entry := entry.FetchEntry(r)

	t, _ := template.ParseFiles("templates/show.html")
	t.Execute(w, entry)
}

func editEntryHandler(w http.ResponseWriter, r *http.Request) {

	e := entry.FetchEntry(r)
	var allTags []entry.Tag
	allTags = entry.GetAllTags()

	data := SoloData{
		Entry:   e,
		AllTags: allTags,
	}

	t, _ := template.ParseFiles("templates/editEntry.html")
	t.Execute(w, data)
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
		entry.UpdateEntry(r, *e)
	}
	http.Redirect(w, r, "/", 301)
}

func deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	entry.DeleteEntry(r)
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
