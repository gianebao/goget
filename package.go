package main

import (
	"log"
	"strings"
)

// Package represents the a Go package
type Package struct {
	Libraries map[string]Library `json:"libraries"`
	Archives  map[string]Library `json:"archives"`
}

const (
	// PkgName defines the file to be created that will store the installation instructions
	PkgName = "go-package.json"
)

// NewPackage creates a new Package instance
func NewPackage() *Package {
	var (
		p = &Package{
			Libraries: map[string]Library{},
			Archives:  map[string]Library{},
		}

		ll = GetLibraries()
	)

	p.putLibraries(ll)
	return p
}

// putLibraries inserts all libraries from ll to the package instance
func (p *Package) putLibraries(ll []string) (bool, *Package) {
	changed := false

	for _, l := range ll {
		if _, ok := p.Libraries[l]; !ok {
			p.Libraries[l] = NewLibrary(l)
			changed = true
		}
	}
	return changed, p
}

// Verify makes sure the package file contains the relevant libraries
func (p *Package) Verify() (bool, error) {
	ll := GetLibraries()

	changed, _ := p.putLibraries(ll)

	for k, l := range p.Libraries {
		found := false
		for _, i := range ll {
			if i == k {
				found = true
				break
			}
		}
		if !found {
			p.Archives[k] = l
			delete(p.Libraries, k)
			changed = true
		}
	}

	return changed, nil
}

// LoadAll executes all library `Get` command
func (p *Package) LoadAll() {
	for i, l := range p.Libraries {
		log.Println(">> Installing `", i, "`")
		log.Println(">>> Output:", strings.TrimSpace(l.Get()))
	}
}
