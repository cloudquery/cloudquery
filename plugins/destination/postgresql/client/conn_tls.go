package client

import (
	"crypto/tls"

	"github.com/jackc/pgx/v5"
)

// tlsRequirement is the weakest TLS configuration a connection may use. Auth
// methods that put a credential on the wire (a Lakebase OAuth token, an AWS IAM
// authentication token) reject anything below their requirement.
type tlsRequirement int

const (
	// tlsEncrypted requires TLS, i.e. an sslmode that always encrypts: `require`,
	// `verify-ca` or `verify-full`.
	tlsEncrypted tlsRequirement = iota
	// tlsVerified additionally requires the server certificate chain and hostname
	// to be verified, i.e. `sslmode=verify-full`.
	tlsVerified
)

// hint describes how to satisfy the requirement, for use in error messages.
func (r tlsRequirement) hint() string {
	if r == tlsVerified {
		return "`sslmode=verify-full`"
	}
	return "`sslmode=require` (or `verify-ca`/`verify-full`)"
}

// checkTLS reports whether every connection path in the pgx config satisfies req.
// pgx represents sslmode=prefer (the default) and sslmode=allow via Fallbacks that
// can silently downgrade to a plaintext connection, so every path (the primary
// config plus all fallbacks) is checked.
func checkTLS(connConfig *pgx.ConnConfig, req tlsRequirement) bool {
	if !satisfiesTLS(connConfig.TLSConfig, req) {
		return false
	}
	for _, fb := range connConfig.Fallbacks {
		if !satisfiesTLS(fb.TLSConfig, req) {
			return false
		}
	}
	return true
}

// satisfiesTLS reports whether a single connection path satisfies req. A nil
// TLSConfig is a plaintext connection. pgx only leaves InsecureSkipVerify unset
// for sslmode=verify-full: `require` and `verify-ca` both set it, as verify-ca
// checks the certificate chain itself in VerifyPeerCertificate but deliberately
// skips hostname verification.
func satisfiesTLS(tlsConfig *tls.Config, req tlsRequirement) bool {
	if tlsConfig == nil {
		return false
	}
	if req == tlsVerified {
		return !tlsConfig.InsecureSkipVerify
	}
	return true
}
