package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "404 not found.", http.StatusNotFound)
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		//log.Println(err.Error())
		app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Hello World"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		app.errorLog.Println(err.Error())
		//http.NotFound(w, r)
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// == w.WriteHeader(405) + w.Write([]byte("Method not allowed"))
		app.errorLog.Println("Method not allowed")
		//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create a new snippet..."))
}
