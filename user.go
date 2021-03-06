package models

import "net/http"

// User models a user of an application. It may be an individual or a corperation
// or an automated entity.
type User struct {
	// Name is what they want to be called.
	Name string `json:"name"`

	// PublicKey can be used for encrypting messages for this user.
	PublicKey string

	// GithubID is their GitHub username.
	GithubID string
	// GithubKey is their GitHub token.
	GithubKey string

	// FSRoot is the global ID of the user's filesystem. Example: "mike.rybka.ca/fs".
	// This should be an instance of github.com/webmachinedev/userfs
	FSRoot string
}

func (u *User) ID() string {
	return hash(u.FSRoot)
}

func (f *User) WriteListItem(w http.ResponseWriter) {
}
