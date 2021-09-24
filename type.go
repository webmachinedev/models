package models

import "net/http"

type Type struct {
	ID string `json:"id"`

	Name   string  `json:"name"`
	Doc    string  `json:"doc"`
	Fields []Field `json:"fields"`
}

func (t *Type) WriteListItem(w http.ResponseWriter) {
}
