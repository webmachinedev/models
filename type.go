package models

import (
	"regexp"
	"strings"
)

type Type struct {
	Name   string  `json:"name"`
	Doc    string  `json:"doc"`
	Fields []Field `json:"fields"`
}

var allButLettersRegex regexp.Regexp

func init() {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	allButLettersRegex = *reg
}

func filterString(str string) string {
	return allButLettersRegex.ReplaceAllString(str, "")
}

func createIDFromName(name string) string {
	id := strings.ToLower(name)
	id = filterString(id)
	return id
}

func (t *Type) ID() string {
	id := createIDFromName(t.Name)
	id += "-"
	for _, field := range t.Fields {
		id += field.ID()
	}
	return id
}
