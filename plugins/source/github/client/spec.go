package client

import (
	_ "embed"
)

//go:embed spec.example.yaml
var exampleSpec string

func ExampleSpec() string {
	return exampleSpec
}

type Spec struct {
	AccessToken string   `json:"access_token"`
	Orgs        []string `json:"orgs"`
}
