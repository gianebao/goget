package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Exec runs an os executable `name` with parameters args ands returns the string response
func Exec(name string, args ...string) string {
	log.Print(strings.Join(append([]string{name}, args...), " "))
	c := exec.Command(name, args...)
	c.Env = os.Environ()

	out, err := c.Output()

	if err != nil {
		panic("Output: " + err.Error())
	}

	return string(out)
}
