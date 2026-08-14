package spec

// OIDCCredentials configures MongoDB Atlas [Workload Identity Federation](https://www.mongodb.com/docs/atlas/workload-oidc/)
// authentication using the built-in `MONGODB-OIDC` machine (workload) flows of the MongoDB driver.
//
// Requires an Atlas M10+ dedicated cluster running MongoDB 7.0.11 or later with Workload
// Identity Federation configured for the matching identity provider.
type OIDCCredentials struct {
	// The workload environment the driver obtains an OIDC token from.
	// One of `k8s` (Kubernetes / EKS service account), `azure` (Azure managed identity)
	// or `gcp` (Google service account).
	Environment string `json:"environment" jsonschema:"required,enum=k8s,enum=azure,enum=gcp,example=k8s"`

	// The audience configured on the MongoDB deployment.
	// Required when `environment` is `azure` or `gcp`; must not be set for `k8s`.
	TokenResource string `json:"token_resource,omitempty" jsonschema:"example=api://my-audience"`

	// The Azure managed-identity client ID (or application ID).
	// Only valid when `environment` is `azure`; may be omitted if the VM has a single managed identity.
	Username string `json:"username,omitempty" jsonschema:"example=00000000-0000-0000-0000-000000000000"`
}

const (
	OIDCEnvironmentK8s   = "k8s"
	OIDCEnvironmentAzure = "azure"
	OIDCEnvironmentGCP   = "gcp"
)
