# gowiki
### WHAT
- Implement web application to getting used to code with golang.
- Following [this](https://go.dev/doc/articles/wiki/) official doc.

### HOW
1. Clone [this](https://github.com/snkzt/gowiki) repo to your local
2. Run below in your teminal
   ```
   $ go build wiki.go
   $ ./wiki
   ```
3. Visit http://localhost:8080/view/ANewPage will show you the page edit form.
   You can edit and save the text, and will be redirected to the newly created page.

  ##### NOTE
  - The text will be saved to your local repository where this repo exists.
  - You can view / edit other saved pages simply by visiting `http://localhost:8080/${view|edit}/${page title}`.
