package client

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	sf "github.com/snowflakedb/gosnowflake"
)

const (
	defaultBatchSize          = 1000
	defaultBatchSizeBytes     = 4 * 1024 * 1024
	defaultMigrateConcurrency = 1
)

type Spec struct {
	// ConnectionString is a DSN the user can specify with all config options,
	// but we also support each attribute individually. See [sf.ParseDSN].
	ConnectionString string `json:"connection_string,omitempty"`

	// Account is [sf.Config.Account].
	Account string `json:"account,omitempty"`
	// User is [sf.Config.User].
	User string `json:"user,omitempty"`
	// Password is [sf.Config.Password].
	Password string `json:"password,omitempty"`
	// Database is [sf.Config.Database].
	Database string `json:"database,omitempty"`
	// Schema is [sf.Config.Schema].
	Schema string `json:"schema,omitempty"`
	// Warehouse is [sf.Config.Warehouse].
	Warehouse string `json:"warehouse,omitempty"`
	// Role is [sf.Config.Role].
	Role string `json:"role,omitempty"`
	// Region is [sf.Config.Region].
	Region string `json:"region,omitempty"`

	// Params is [sf.Config.Params] - other arbitrary connection
	// options.
	Params map[string]*string `json:"params,omitempty"`

	// ClientIP is [sf.Config.ClientIP].
	ClientIP net.IP `json:"client_ip,omitempty"`
	// Protocol is [sf.Config.Protocol].
	Protocol string `json:"protocol,omitempty"`
	// Host is [sf.Config.Host].
	Host string `json:"host,omitempty"`
	// Port is [sf.Config.Port].
	Port int `json:"port,omitempty"`

	// Authenticator maps to [sf.Config.Authenticator].
	Authenticator string `json:"authenticator,omitempty"`

	// Login is [sf.Config.LoginTimeout] - Login retry timeout EXCLUDING network
	// roundtrip and read out http response.
	LoginTimeout time.Duration `json:"login_timeout,omitempty"`
	// RequestTimeout is [sf.Config.RequestTimeout] - request retry timeout
	// EXCLUDING network roundtrip and read out http response.
	RequestTimeout time.Duration `json:"request_timeout,omitempty"`
	// JWTExpireTimeout is [sf.Config.JWTExpireTimeout] - JWT expire after
	// timeout.
	JWTExpireTimeout time.Duration `json:"jwt_expire_timeout,omitempty"`
	// ClientTimeout is [sf.Config.ClientTimeout] - Timeout for network round
	// trip + read out http response.
	ClientTimeout time.Duration `json:"client_timeout,omitempty"`

	// Application is [sf.Config.Application] - the application name.
	Application string `json:"application,omitempty"`
	// InsecureMode is [sf.Config.InsecureMode] - driver doesn't check
	// certificate revocation status.
	InsecureMode *bool `json:"insecure_mode,omitempty"`
	// OCSPFailOpen is [sf.Config.OCSPFailOpen].
	OCSPFailOpen string `json:"ocsp_fail_open,omitempty"` // FIXME: map to OCSPFailOpenMode?

	// Token is [sf.Config.Token] - Token to use for OAuth other forms of token
	// based auth.
	Token string `json:"token,omitempty"`
	// KeepSessionAlive is [sf.Config.KeepSessionAlive] - Enables the session to
	// persist even after the connection is closed.
	KeepSessionAlive *bool `json:"keep_session_alive,omitempty"`

	// PrivateKey is the PEM-encoded privateKey connection parameter. Implies
	// Authenticator: JWT. This is the contents of file rsa_key.p8 as generated
	// by a command like the following:
	//
	//  openssl genrsa -out rsa_key.p8 2048
	//
	// rsa_key.p8 is the private key and the reciprocal public key can be
	// generated with
	//
	//  openssl rsa -in rsa_key.p8 -pubout -out rsa_key.pub
	//
	// rsa_key.pub now contains the 2public key in PEM format. Omitting line
	// breaks and the === BEGIN === / === END === delimiters, you can attach
	// this public key to a user with:
	//
	//  ALTER USER username SET rsa_public_key='public-key';
	PrivateKey string `json:"private_key,omitempty"`

	// DisableTelemetry is [sf.Config.DisableTelemetry] - indicates whether to disable telemetry.
	DisableTelemetry *bool `json:"disable_telemetry,omitempty"`

	// Tracing is [sf.Config.Tracing] - sets logging level.
	Tracing string `json:"tracing,omitempty"`

	// ClientRequestMfaToken is [sf.Config.ClientRequestMfaToken] -
	// When true the MFA token is cached in the credential manager. True by
	// default in Windows/OSX. False for Linux..
	ClientRequestMfaToken *bool `json:"client_request_mfa_token,omitempty"`
	// ClientStoreTemporaryCredential is
	// [sf.Config.ClientStoreTemporaryCredential] - When true the ID
	// token is cached in the credential manager. True by default in
	// Windows/OSX. False for Linux..
	ClientStoreTemporaryCredential *bool `json:"client_store_temporary_credential,omitempty"`

	// CloudQuery destination config options.
	BatchSize          int `json:"batch_size,omitempty"`
	BatchSizeBytes     int `json:"batch_size_bytes,omitempty"`
	MigrateConcurrency int `json:"migrate_concurrency,omitempty"`
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

func (s *Spec) Config(extraParams map[string]*string) (*sf.Config, error) {
	var cfg *sf.Config

	// Use existing values from the connection string.
	if s.ConnectionString != "" {
		c, err := sf.ParseDSN(s.ConnectionString)
		if err != nil {
			return nil, err
		}
		cfg = c
	} else {
		cfg = &sf.Config{}
	}

	// Override with other values.
	if s.Account != "" {
		cfg.Account = s.Account
	}
	if s.User != "" {
		cfg.User = s.User
	}
	if s.Password != "" {
		cfg.Password = s.Password
	}
	if s.Database != "" {
		cfg.Database = s.Database
	}
	if s.Schema != "" {
		cfg.Schema = s.Schema
	}
	if s.Warehouse != "" {
		cfg.Warehouse = s.Warehouse
	}
	if s.Role != "" {
		cfg.Role = s.Role
	}
	if s.Region != "" {
		cfg.Region = s.Region
	}
	if s.ClientIP != nil {
		cfg.ClientIP = s.ClientIP
	}
	if s.Protocol != "" {
		cfg.Protocol = s.Protocol
	}
	if s.Host != "" {
		cfg.Host = s.Host
	}
	if s.Port != 0 {
		cfg.Port = s.Port
	}
	if s.Authenticator != "" {
		authType, ok := parseAuthenticator(s.Authenticator)
		if !ok {
			return nil, fmt.Errorf("authenticator: unknown value %q", s.Authenticator)
		}
		cfg.Authenticator = authType
	}

	if s.LoginTimeout != 0 {
		cfg.LoginTimeout = s.LoginTimeout
	}
	if s.RequestTimeout != 0 {
		cfg.RequestTimeout = s.RequestTimeout
	}
	if s.JWTExpireTimeout != 0 {
		cfg.JWTExpireTimeout = s.JWTExpireTimeout
	}
	if s.ClientTimeout != 0 {
		cfg.ClientTimeout = s.ClientTimeout
	}

	if s.Application != "" {
		cfg.Application = s.Application
	}
	if s.InsecureMode != nil {
		cfg.InsecureMode = *s.InsecureMode
	}
	if s.OCSPFailOpen != "" {
		failOpen, ok := parseOCSPFailOpen(s.OCSPFailOpen)
		if !ok {
			return nil, fmt.Errorf(`ocsp_fail_open: unknown mode %q, expected "FAIL_OPEN" or "FAIL_CLOSED"`, s.OCSPFailOpen)
		}
		cfg.OCSPFailOpen = failOpen
	}

	if s.Token != "" {
		cfg.Token = s.Token
	}
	if s.KeepSessionAlive != nil {
		cfg.KeepSessionAlive = *s.KeepSessionAlive
	}

	if s.PrivateKey != "" {
		rsaKey, err := parsePEMRSAKey(s.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("private_key: %s", err)
		}
		cfg.PrivateKey = rsaKey
		cfg.Authenticator = sf.AuthTypeJwt
	}

	if s.DisableTelemetry != nil {
		cfg.DisableTelemetry = *s.DisableTelemetry
	}

	if s.Tracing != "" {
		cfg.Tracing = s.Tracing
	}

	if s.ClientRequestMfaToken != nil {
		cfg.ClientRequestMfaToken = boolPtrToConfig(s.ClientRequestMfaToken)
	}
	if s.ClientStoreTemporaryCredential != nil {
		cfg.ClientStoreTemporaryCredential = boolPtrToConfig(s.ClientStoreTemporaryCredential)
	}

	if len(s.Params)+len(extraParams) > 0 && cfg.Params == nil {
		cfg.Params = make(map[string]*string, len(s.Params)+len(extraParams))
	}
	for k, v := range s.Params {
		cfg.Params[k] = v
	}
	for k, v := range extraParams {
		cfg.Params[k] = v
	}

	if cfg.Database == "" {
		return nil, errors.New("a snowflake database must be specified")
	}
	if cfg.Schema == "" {
		return nil, errors.New("a snowflake schema must be specified")
	}
	if cfg.Warehouse == "" {
		return nil, errors.New("a snowflake warehouse must be specified")
	}

	return cfg, nil
}

func boolPtrToConfig(b *bool) sf.ConfigBool {
	if b == nil {
		var unset sf.ConfigBool
		return unset
	}
	if *b {
		return sf.ConfigBoolTrue
	}
	return sf.ConfigBoolFalse
}

func parseOCSPFailOpen(v string) (sf.OCSPFailOpenMode, bool) {
	// https://github.com/snowflakedb/gosnowflake/blob/v1.6.23/ocsp.go#L60-L61
	switch strings.ToUpper(v) {
	case "FAIL_OPEN":
		return sf.OCSPFailOpenTrue, true
	case "FAIL_CLOSED":
		return sf.OCSPFailOpenFalse, true
	}

	var unknown sf.OCSPFailOpenMode
	return unknown, false
}

func parseAuthenticator(v string) (sf.AuthType, bool) {
	// https://github.com/snowflakedb/gosnowflake/blob/v1.6.23/auth.go#L106
	switch strings.ToUpper(v) {
	case "SNOWFLAKE":
		return sf.AuthTypeSnowflake, true
	case "OAUTH":
		return sf.AuthTypeOAuth, true
	case "EXTERNALBROWSER":
		return sf.AuthTypeExternalBrowser, true
	case "OKTA":
		return sf.AuthTypeOkta, true
	case "SNOWFLAKE_JWT":
		return sf.AuthTypeJwt, true
	case "TOKENACCESSOR":
		return sf.AuthTypeTokenAccessor, true
	case "USERNAME_PASSWORD_MFA":
		return sf.AuthTypeUsernamePasswordMFA, true
	}

	var unknown sf.AuthType
	return unknown, false
}

var whitespace = regexp.MustCompile(`\s+`)

func parsePEMRSAKey(blob string) (*rsa.PrivateKey, error) {
	// Chop out the relevant parts of the key.
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
	// Driver suport it?)
	const pemPrivKey = "PRIVATE KEY"
	switch strings.ToUpper(head) {
	case pemPrivKey:
		break // OK.
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
	return rsaPrivateKey, nil
}
