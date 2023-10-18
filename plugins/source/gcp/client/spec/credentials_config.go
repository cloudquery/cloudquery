package spec

import "google.golang.org/api/cloudresourcemanager/v3"

type CredentialsConfig struct {
	// The email address of the service account to impersonate.
	TargetPrincipal string `json:"target_principal" jsonschema:"required,format=email"`

	// Scopes that the impersonated credential should have.
	//
	// See available scopes in the [documentation](https://developers.google.com/identity/protocols/oauth2/scopes).
	Scopes []string `json:"scopes" jsonschema:"pattern=^https://www.googleapis.com/auth/(.)+$,default=https://www.googleapis.com/auth/cloud-platform"`

	// Delegates are the service account email addresses in a delegation chain.
	// Each service account must be granted `roles/iam.serviceAccountTokenCreator` on the next service account in the chain.
	Delegates []string `json:"delegates" jsonschema:"format=email"`

	// The subject field of a JWT (`sub`).
	// This field should only be set if you wish to impersonate a user.
	// This feature is useful when using domain wide delegation.
	Subject string `json:"subject" jsonschema:"minLength=1"`
}

func (c *CredentialsConfig) SetDefaults() {
	if c == nil {
		return
	}
	if len(c.Scopes) == 0 {
		// `https://www.googleapis.com/auth/cloud-platform`
		// We use this as some APIs don't utilize the read-only alternative `https://www.googleapis.com/auth/cloud-platform.read-only`
		// See https://developers.google.com/identity/protocols/oauth2/scopes for more details.
		c.Scopes = []string{cloudresourcemanager.CloudPlatformScope}
	}
}
