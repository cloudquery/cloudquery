package testing

import (
	"log"
	"testing"
)

func TestManager_Load(t *testing.T) {
	// Skip test for now since github is annoying
	files, err := FilePathWalkDir("../../../database-data")
	log.Println(files, err)

}

func TestManager_Group(t *testing.T) {
	// Skip test for now since github is annoying
	files, _ := FilePathWalkDir("../../../database-data")
	log.Printf("%+v", FilterFiles(files))

}
