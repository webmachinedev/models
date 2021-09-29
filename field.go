package models

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Doc  string `json:"doc"`
}

func (f *Field) ID() string {
	return createIDFromName(f.Name)+f.Type
}
