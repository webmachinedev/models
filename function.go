package models

type Function struct {
	ID ID

	Name    string
	Doc     string
	Inputs  []Field
	Outputs []Field
}
