package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Function struct {
	Name    string  `json:"name"`
	Doc     string  `json:"doc"`
	Inputs  []Field `json:"inputs"`
	Outputs []Field `json:"outputs"`
	Body []Statement `json:"body"`
}

func lowercase(str string) string {
	return str
}

func removeSpaces(str string) string {
	return str
}

func hash(arr ...interface{}) string {
	hash := sha256.New()
	for _, val := range arr {
		bytes, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}
		hash.Write(bytes)
	}
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%x", hash.Sum(nil))
	return buf.String()
}

func removeAllCharsThatAreNotLowercaseEnglishLetters(str string) string {
	valid := func(char rune) bool {
		return 65 <= char && char <= 90
	}
	var res string
	for _, char := range str {
		if valid(char) {
			res += string(char)
		}
	}
	return res
}

func validateRunes(str string, allowlist []rune) string {
	var res string
	for _, char := range str {
		var valid bool
		for _, r := range allowlist {
			if char == r {
				valid = true
			}
		}
		if valid {
			res += string(char)
		}
	}
	return res
}

func (f *Function) ID() string {
	return hash(f.Body)
}

type Statement struct {
	Type string `json:"type"`
	Args map[string]Statement `json:"args"`
	Children []Statement `json:"children"`
}

func (s *Statement) Exec(ns Namespace) (Result, error)

type Result struct {
	Loading bool
	Error error
	Value interface{}
}

type Namespace map[string][]string
