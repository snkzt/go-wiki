package main

import (
  "html/template"
  "os"
  "log"
  "net/http"
)

type Page struct {
  Title string
  Body []byte  
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  // 0600 read-write permission for the current user only
  return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := os.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  page, _ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  page, err := loadPage(title)
  if err != nil {
    page = &Page{Title: title}
  }
  // ParseFiles function read the contents of edit.html and return a *teplate.Template, which to generate HTML
  editHtmlTemplate, _ := template.ParseFiles("edit.html")
  // Execute function executes the template, writing the generated HTML to the w
  editHtmlTemplate.Execute(w, page)
}

func main() {
  // Initialise http: HandlerFunc telles the http package to handle all request to the /view/ with viewHandler func
  http.HandleFunc("/view/", viewHandler)
  // ListenAndServe specified that it should listen on port 8080 on any interface
  // ListenAndServe always returns an error because it's the only occation it returns smt, and to log that error it should be wrapped with log.Fatal
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

