package models

import "github.com/okta/okta-sdk-golang/v2/okta"

type ApplicationUser struct {
	*okta.AppUser
}

type GroupUser struct {
	Id string
}
