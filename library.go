package main

import "strings"

// Library represents the dependencies of the Go package
type Library struct {
	Import  string   `json:"import"`
	Command []string `json:"command"`
}

// GetLibraries returns a lists of libraries required by the Go package
func GetLibraries() []string {
	return strings.Split(strings.TrimSpace(Exec("go", "list", "-f", "{{ join .Imports \"\\n\" }}")), "\n")
}

// NewLibrary creates a new Library instance
func NewLibrary(path string) Library {
	return Library{
		Import:  path,
		Command: []string{"go", "get", "-u", "-v", path},
	}
}

// Get executes the []Command
func (l Library) Get() string {
	return Exec(l.Command[0], l.Command[1:len(l.Command)]...)
}
