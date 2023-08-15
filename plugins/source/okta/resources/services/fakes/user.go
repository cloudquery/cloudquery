package fakes

import (
	"time"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func User() okta.User {
	var u okta.User
	if err := faker.FakeObject(&u); err != nil {
		panic(err)
	}
	tf := &okta.TestFactory{}
	up := tf.NewValidTestUserProfile()
	u.Profile = &up
	u.Credentials.Password.Hash.Algorithm = &okta.AllowedPasswordCredentialHashAlgorithmEnumValues[0]
	u.Credentials.Provider.Type = &okta.AllowedAuthenticationProviderTypeEnumValues[0]
	u.Status = &okta.AllowedUserStatusEnumValues[0]
	u.TransitioningToStatus = &okta.AllowedUserStatusEnumValues[0]
	t := time.Now()
	u.Activated.Set(&t)
	u.LastLogin.Set(&t)
	u.PasswordChanged.Set(&t)
	u.StatusChanged.Set(&t)
	u.LastUpdated = &t

	return u
}
