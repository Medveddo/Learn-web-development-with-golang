package views

import (
	"bytes"
	"html/template"
	"io"
	"learn-web-dev-with-go/context"
	"net/http"
	"path/filepath"
)

var (
	TemplateDir string = "views/"
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {

	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	// Creating a View with Template = t and get a pointer by &
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

//Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
		// do nothing
	default:
		vd = Data{
			Yield: data,
		}
	}
	vd.User = context.User(r.Context())
	var buf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		http.Error(w, "Something went wrong. If the problem presists, please email support@futurecycling.space", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

// not capital letter because we dont wanna export this function out of package
// returns a slice of strings representing the layout used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

/* Return example
"views/layouts/bootstrap.gohtml"
"views/layouts/navbar.gohtml"
"views/layouts/footer.gohtml"
*/

// addTemplatePath takes an a slice of strings
// representing file paths for templates and it prepends
// the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir = "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplatePath takes an a slice of strings
// representing file paths for templates and it appends
// the TemplateExt extension to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"home.gohtml"} if TemplateExt = ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
