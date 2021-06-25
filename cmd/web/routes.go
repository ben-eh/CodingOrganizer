package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)
	// mux.HandleFunc("/indexSearch", indexSearchHandler)
	// mux.HandleFunc("/addEntry", addEntryHandler)
	// mux.HandleFunc("/saveEntry", saveEntryHandler)
	// mux.HandleFunc("/showEntry/{entry_id}", showEntryHandler)
	// mux.HandleFunc("/editEntry/{entry_id}", editEntryHandler)
	// mux.HandleFunc("/updateEntry/{entry_id}", updateEntryHandler)
	// mux.HandleFunc("/deleteEntry/{entry_id}", deleteEntryHandler)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
