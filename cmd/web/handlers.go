package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hideaki10/web-application-tool/pkg/forms"
	"github.com/hideaki10/web-application-tool/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	// http.Error(w, "404 not found.", http.StatusNotFound)
	// 	app.notFound(w)
	// 	return
	// }

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

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	// == w.WriteHeader(405) + w.Write([]byte("Method not allowed"))
	// 	app.errorLog.Println("Method not allowed")
	// 	//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	form := forms.New(r.PostForm)

	form.Required("title", "content", "expires")

	//
	form.MaxLength("title", 100)

	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}
	// errors := make(map[string]string)

	// if strings.TrimSpace(title) == "" {
	// 	errors["title"] = "title field cannot be blank"
	// } else if utf8.RuneCountInString(title) > 100 {
	// 	errors["title"] = "this field is too lang(maximum is 100 characters)"
	// }

	// if strings.TrimSpace(content) == "" {
	// 	errors["content"] = "this field cannot be blank"
	// }

	// if strings.TrimSpace(expires) == "" {
	// 	errors["expires"] = "This field cannot be blank"
	// } else if expires != "365" && expires != "7" && expires != "1" {
	// 	errors["expires"] = "This field is invalid"
	// }

	// if len(errors) > 0 {
	//
	// 	return
	// }

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "snippent successfully created! ")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
	// w.Write([]byte("create a new snippet..."))
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})

		return
	}

	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplidateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "singup.page.tmpl", &templateData{
				Form: form,
			})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the  user...")
}
