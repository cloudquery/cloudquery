package client

import (
	"crypto/tls"

	"github.com/jackc/pgx/v5"
)

// tlsRequirement is the weakest TLS configuration a connection may use.
type tlsRequirement int

const (
	// tlsEncrypted is satisfied by sslmode require, verify-ca and verify-full.
	tlsEncrypted tlsRequirement = iota
	// tlsVerified also requires hostname verification, i.e. sslmode=verify-full.
	tlsVerified
)

func (r tlsRequirement) hint() string {
	if r == tlsVerified {
		return "`sslmode=verify-full`"
	}
	return "`sslmode=require` (or `verify-ca`/`verify-full`)"
}

// pgx implements sslmode=prefer (the default) and sslmode=allow as Fallbacks that
// can silently downgrade to plaintext, so every path is checked.
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

// pgx only leaves InsecureSkipVerify unset for verify-full: verify-ca sets it and
// checks the chain itself in VerifyPeerCertificate, skipping the hostname.
func satisfiesTLS(tlsConfig *tls.Config, req tlsRequirement) bool {
	if tlsConfig == nil {
		return false
	}
	if req == tlsVerified {
		return !tlsConfig.InsecureSkipVerify
	}
	return true
}
