package models

import "net/http"

type Type struct {
	ID ID

	Name   string
	Doc    string
	Fields []Field
}

func (t *Type) WriteListItem(w http.ResponseWriter) {
}
