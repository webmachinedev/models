package models

import "net/http"

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}

func (f *Field) WriteListItem(w http.ResponseWriter) {
}
