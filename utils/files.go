package utils

import (
	"log"
	"os"
)

/*
Files struct that contains util methods
*/
type Files struct{}

/*
Open open and return a file
*/
func (m *Files) Open(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
