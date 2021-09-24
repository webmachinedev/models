package models

import "net/http"

type Field struct {
	Name string
	Type string
	Doc  string
}

func (f *Field) WriteListItem(w http.ResponseWriter) {
}
