package spec

// OIDC configures MONGODB-OIDC Workload Identity Federation authentication using
// one of the driver's built-in providers. When set, it overrides any credentials
// in the connection_string.
type OIDC struct {
	// OIDC built-in provider environment. One of `gcp`, `azure` or `k8s`.
	Environment string `json:"environment" jsonschema:"required,enum=gcp,enum=azure,enum=k8s"`

	// Token resource (audience) configured on your Atlas deployment.
	// Required for the `gcp` and `azure` environments.
	TokenResource string `json:"token_resource,omitempty"`
}
