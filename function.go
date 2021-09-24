package models

import "net/http"

type Function struct {
	ID string `json:"id"`

	Name    string  `json:"name"`
	Doc     string  `json:"doc"`
	Inputs  []Field `json:"inputs"`
	Outputs []Field `json:"outputs"`
}

func (f *Function) WriteListItem(w http.ResponseWriter) {
}
