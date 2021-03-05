package views

import "html/template"

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	// Creating a View with Template = t and get a pointer by &
	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
