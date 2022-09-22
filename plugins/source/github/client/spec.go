package client

import (
	_ "embed"
)

type Spec struct {
	AccessToken string   `json:"access_token"`
	Orgs        []string `json:"orgs"`
}
