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

func TestManager_Execute(t *testing.T) {
	// Skip test for now since github is annoying
	TestPolicy(t, ".cq/policies/github.com/cloudquery-policies/aws", "ec2/EC2.18")

}
