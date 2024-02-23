package fakes

import (
	"time"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/okta/okta-sdk-golang/v4/okta"
)

func User() okta.User {
	var u okta.User
	if err := faker.FakeObject(&u); err != nil {
		panic(err)
	}
	tf := &okta.TestFactory{}
	up := tf.NewValidTestUserProfile()
	u.Profile = &up
	u.Credentials = tf.NewValidTestUserCredentialsWithPassword()
	u.Status = &okta.AllowedUserStatusEnumValues[0]
	u.TransitioningToStatus = &okta.AllowedUserStatusEnumValues[0]
	t := time.Now()
	u.Activated.Set(&t)
	u.LastLogin.Set(&t)
	u.PasswordChanged.Set(&t)
	u.StatusChanged.Set(&t)
	u.LastUpdated = &t

	u.AdditionalProperties = map[string]any{"key": "value"}
	u.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	u.Links = &okta.UserLinks{
		Self: &okta.HrefObject{Href: "#"},
	}
	u.Type.Links = &okta.UserTypeLinks{Self: &okta.HrefObjectSelfLink{Href: "#"}}

	return u
}
