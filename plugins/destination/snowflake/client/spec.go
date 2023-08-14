package client

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	defaultBatchSize          = 1000
	defaultBatchSizeBytes     = 4 * 1024 * 1024
	defaultMigrateConcurrency = 1
)

type Spec struct {
	ConnectionString   string `json:"connection_string,omitempty"`
	PrivateKey         string `json:"private_key,omitempty"`
	BatchSize          int    `json:"batch_size,omitempty"`
	BatchSizeBytes     int    `json:"batch_size_bytes,omitempty"`
	MigrateConcurrency int    `json:"migrate_concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	// stub for any future defaults
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.MigrateConcurrency == 0 {
		s.MigrateConcurrency = defaultMigrateConcurrency
	}
}

func (s Spec) DSN() (string, error) {
	cs := s.ConnectionString
	if cs == "" {
		return "", fmt.Errorf("connection_string is required")
	}

	if s.PrivateKey != "" {
		pk, err := formatPrivateKey(s.PrivateKey)
		if err != nil {
			return "", fmt.Errorf("private_key: %w", err)
		}

		sep := "?"
		if strings.Contains(cs, "?") {
			sep = "&"
		}
		cs += sep + "authenticator=snowflake_jwt&privateKey=" + base64.URLEncoding.EncodeToString(pk)
	}
	return cs, nil
}

var whitespace = regexp.MustCompile(`\s+`)

func formatPrivateKey(blob string) ([]byte, error) {
	// Strip any PEM block headers.
	const (
		pemBegin = "-----BEGIN "
		pemSep   = "-----"
		pemEnd   = "-----END "
	)
	_, rest, hadBegin := strings.Cut(blob, pemBegin)
	head, rest, hadEnd := strings.Cut(rest, pemSep)
	key, rest, hadKey := strings.Cut(rest, pemEnd)
	tail, _, hadTail := strings.Cut(rest, pemSep)
	if !hadBegin || !hadEnd || !hadKey || !hadTail {
		return nil, fmt.Errorf("unable to find %s...%s...%s...%s in private key", pemBegin, pemSep, pemEnd, pemSep)
	}

	// Encrypted private keys aren't supported (TODO: Is this only because
	// pem.Decode doesn't support it? Does the underlying Snowflake Go SQL
	// Driver support it?)
	const pemPrivKey = "PRIVATE KEY"
	switch strings.ToUpper(head) {
	case pemPrivKey:
		// OK.
	case "ENCRYPTED PRIVATE KEY":
		return nil, errors.New("encrypted private keys are not supported, use decrypted private key")
	default:
		return nil, fmt.Errorf("unrecognised start block %s%s%s, expected %s%s%s", pemBegin, head, pemSep, pemBegin, pemPrivKey, pemSep)
	}

	// Rebuild the key with the correct line breaks.
	//
	// The expansion of ${file:./private.key} in our YAML specs doesn't retain
	// newlines at the time of writing (unless private.key contains valid JSON,
	// which it shouldn't here) so we're going to substitute all inner
	// whitespace with newlines.
	blob = pemBegin + head + pemSep + "\n" + strings.TrimSpace(whitespace.ReplaceAllString(key, "\n")) + "\n" + pemEnd + tail + pemSep

	// Parse and reformat.
	//
	// https://github.com/snowflakedb/gosnowflake/blob/7de6b8d13750ca70667f554335862f97a82720ea/cmd/keypair/keypair.go#L39-L52
	block, _ := pem.Decode([]byte(blob))
	if block == nil {
		return nil, errors.New("could not decode PEM block")
	}
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}
	rsaPrivateKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("expected *rsa.PrivateKey but got %T", privKey)
	}
	rsaBytes, err := x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("re-marshalling rsa private key: %w", err)
	}

	return rsaBytes, nil
}
