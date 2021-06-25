package main

import (
	"html/template"
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// var entries []entry.Entry
	s, err := app.entries.GetEntries()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// var tags []entry.Tag
	// tags = entry.GetAllTags()

	data := &templateData{
		Entries: s,
		// Tags:    tags,
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// w.Header().Set("Content-Type", "text/css")
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

}

// func indexSearchHandler(w http.ResponseWriter, r *http.Request) {

// 	tags := entry.GetAllTags()

// 	t, _ := template.ParseFiles("templates/indexSearch.html")
// 	t.Execute(w, tags)
// }

// func saveEntryHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	tags := r.Form["[]tags"]
// 	log.Println(tags)
// 	log.Println("pause")

// 	if r.FormValue("name") != "" {
// 		e := &entry.Entry{
// 			Name:      r.FormValue("name"),
// 			URL:       r.FormValue("url"),
// 			CodeBlock: r.FormValue("codeblock"),
// 			Notes:     r.FormValue("notes"),
// 		}
// 		entryID := entry.SaveEntry(*e)
// 		entryHasTag.CreateEntryHasTag(tags, entryID)
// 	}

// 	http.Redirect(w, r, "/", 301)
// }

// func addEntryHandler(w http.ResponseWriter, r *http.Request) {

// 	var allTags []entry.Tag
// 	allTags = entry.GetAllTags()

// 	t, _ := template.ParseFiles("templates/addEntry.html")
// 	t.Execute(w, allTags)

// }

// func showEntryHandler(w http.ResponseWriter, r *http.Request) {
// 	entry := entry.FetchEntry(r)

// 	t, _ := template.ParseFiles("templates/show.html")
// 	t.Execute(w, entry)
// }

// func editEntryHandler(w http.ResponseWriter, r *http.Request) {

// 	e := entry.FetchEntry(r)
// 	var allTags []entry.Tag
// 	allTags = entry.GetAllTags()

// 	data := SoloData{
// 		Entry:   e,
// 		AllTags: allTags,
// 	}

// 	t, _ := template.ParseFiles("templates/editEntry.html")
// 	t.Execute(w, data)
// }

// func updateEntryHandler(w http.ResponseWriter, r *http.Request) {
// 	// postedID, _ := strconv.Atoi(r.FormValue("id"))
// 	vars := mux.Vars(r)

// 	entryID := vars["entry_id"]
// 	log.Println("this is the entry_id in the updateEntryHandler: ", entryID)
// 	postedName := r.FormValue("name")
// 	postedURL := r.FormValue("url")
// 	postedBlock := r.FormValue("codeblock")
// 	postedNotes := r.FormValue("notes")
// 	if postedName != "" {
// 		e := &entry.Entry{
// 			// ID:        postedID,
// 			Name:      postedName,
// 			URL:       postedURL,
// 			CodeBlock: postedBlock,
// 			Notes:     postedNotes,
// 		}
// 		entry.UpdateEntry(r, *e)
// 	}
// 	http.Redirect(w, r, "/", 301)
// }

// func deleteEntryHandler(w http.ResponseWriter, r *http.Request) {
// 	entry.DeleteEntry(r)
// 	http.Redirect(w, r, "/", 301)
// }
