package models

type Function struct {
	ID ID

	Name    string
	Doc     string
	Inputs  []Field
	Outputs []Field
}

func (f *Function) WriteListItem(w http.ResponseWriter) {
}
