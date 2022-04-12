package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hideaki10/web-application-tool/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "404 not found.", http.StatusNotFound)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v\n", snippet)
	// }
	// data := &templateData{Snippets: s}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	//log.Println(err.Error())
	// 	app.errorLog.Println(err.Error())
	// 	//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// 	return
	// }

	// w.Write([]byte("Hello World"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		app.errorLog.Println(err.Error())
		//http.NotFound(w, r)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
	// }
	// data := &templateData{Snippet: s}
	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	// fmt.Fprintf(w, "%v", s)
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

	title := " O snail"
	content := "O snail\nClimb mount fuji\n But slowly,slowly!\n\n kobayashi Issa"
	expires := "7"
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	w.Write([]byte("create a new snippet..."))
}
