package views

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml",
	)
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
