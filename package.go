package models

type Package struct {
	Types     map[string]Type     `json:"types"`
	Functions map[string]Function `json:"functions"`
}
