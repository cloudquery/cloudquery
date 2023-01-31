package client

import "runtime"

type Spec struct {
	Addresses []string // A list of Elasticsearch nodes to use.
	Username  string   // Username for HTTP Basic Authentication.
	Password  string   // Password for HTTP Basic Authentication.

	CloudID                string // Endpoint for the Elastic Service (https://elastic.co/cloud).
	APIKey                 string // Base64-encoded token for authorization; if set, overrides username/password and service token.
	ServiceToken           string // Service token for authorization; if set, overrides username/password.
	CertificateFingerprint string // SHA256 hex fingerprint given by Elasticsearch on first launch.

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	CACert string

	Concurrency int // Number of concurrent worker goroutines to use for indexing. (Default: number of CPUs)
}

func (s *Spec) SetDefaults() {
	if len(s.Addresses) == 0 {
		s.Addresses = []string{"http://localhost:9200"}
	}
	if s.Concurrency == 0 {
		s.Concurrency = runtime.NumCPU()
	}
}

func (*Spec) Validate() error {
	return nil
}
