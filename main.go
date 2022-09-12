package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600) // 0600 read-write permission for the current user only
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", page)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w, "edit", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	// Execute function executes the template, writing the generated HTML to the w
	// The template name is the template file name, we must append ".html" to the tmpl argument.
	err := templates.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		// http.Error function sends specified ("Internal Server Error") HTTP response code and error message
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Validate expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// The returned function is called a closure because it encloses values defined outside of it.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		match := validPath.FindStringSubmatch(r.URL.Path)
		if match == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, match[2]) // The closure: extracts the title from the request path, and validates it with the validPath().
	}
}

func main() {
	// Initialise http: HandlerFunc telles the http package to handle all request to the /view/ with viewHandler func
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// ListenAndServe specified that it should listen on port 8080 on any interface
	// ListenAndServe always returns an error because it's the only occation it returns smt, and to log that error it should be wrapped with log.Fatal
	log.Fatal(http.ListenAndServe(":8080", nil))
}
