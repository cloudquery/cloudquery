package client

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Spec struct {
	// Endpoint for the database. Supported schemes are `wss://` and `ws://`, the default port is `8182`.
	Endpoint string `json:"endpoint" jsonschema:"required,pattern=^ws(s?)://[^\n]+$"`

	// Whether to skip TLS verification. Defaults to `false`. This should be set on a macOS environment when connecting to an AWS Neptune endpoint.
	Insecure bool `json:"insecure" jsonschema:"default=false"`

	// Authentication mode to use. `basic` uses static credentials, `aws` uses AWS IAM authentication.
	AuthMode authMode `json:"auth_mode" jsonschema:"default=none"`

	// Username to connect to the database. Required when `auth_mode` is `basic`.
	Username string `json:"username"`

	// Password to connect to the database. Required when `auth_mode` is `basic`.
	Password string `json:"password"`

	// Number of retries on `ConcurrentModificationException` before giving up for each batch.
	// Retries are exponentially backed off.
	MaxRetries int `json:"max_retries" jsonschema:"minimum=1,default=5"`

	// AWS region to use for AWS IAM authentication. Required when `auth_mode` is `aws`.
	AWSRegion string `json:"aws_region"`

	// Maximum number of concurrent connections to the database. Defaults to the number of CPUs.
	MaxConcurrentConnections int `json:"max_concurrent_connections" jsonschema:"minimum=1"`

	// Whether to use all Gremlin types or just a basic subset.
	// Should remain `false` for Amazon Neptune compatibility.
	CompleteTypes bool `json:"complete_types" jsonschema:"default=false"`

	// Number of records to batch together before sending to the database.
	BatchSize int `json:"batch_size" jsonschema:"minimum=1,default=200"`

	// Number of bytes (as Arrow buffer size) to batch together before sending to the database.
	BatchSizeBytes int `json:"batch_size_bytes" jsonschema:"minimum=1,default=4194304"`
}

type authMode string

const (
	authModeNone  = authMode("none")
	authModeBasic = authMode("basic")
	authModeAWS   = authMode("aws")
)

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.Endpoint != "" {
		// Default to "wss://<endpoint>:8182" where "wss://" and ":8182" are optional
		e := strings.SplitN(s.Endpoint, "://", 2)
		if len(e) == 1 {
			e = []string{"wss", e[0]}
		}
		if !strings.Contains(e[1], ":") {
			e[1] += ":8182"
		}
		s.Endpoint = strings.Join(e, "://")
	}

	if s.AuthMode == "" {
		s.AuthMode = authModeNone
	} else {
		s.AuthMode = authMode(strings.ToLower(string(s.AuthMode)))
	}

	if s.MaxRetries < 1 {
		s.MaxRetries = 5 // 5 retries by default
	}

	if s.MaxConcurrentConnections <= 0 {
		s.MaxConcurrentConnections = runtime.NumCPU()
	}

	if s.BatchSize < 1 {
		s.BatchSize = 200
	}
	if s.BatchSizeBytes < 1 {
		s.BatchSizeBytes = 1024 * 1024 * 4
	}
}

func (s *Spec) Validate() error {
	if s.Endpoint == "" {
		return fmt.Errorf("endpoint is required")
	}
	if s.AuthMode != authModeNone && s.AuthMode != authModeBasic && s.AuthMode != authModeAWS {
		return fmt.Errorf("invalid auth_mode, valid values are %q, %q and %q", authModeNone, authModeBasic, authModeAWS)
	}
	if s.AuthMode == authModeAWS && s.AWSRegion == "" {
		return fmt.Errorf("aws_region is required when auth_mode is %q", authModeAWS)
	}
	if s.AuthMode == authModeNone && (s.Username != "" || s.Password != "") {
		return fmt.Errorf("username or password specified with auth_mode %q. Set auth mode to %q or remove username and password", authModeNone, authModeBasic)
	}

	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	authMode := sc.Properties.Value("auth_mode")
	authMode.Enum = []any{authModeNone, authModeBasic, authModeAWS}

	sc.OneOf = []*jsonschema.Schema{
		{
			// if auth_mode is aws, require aws_region to be set
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				one := uint64(1)
				properties := jsonschema.NewProperties()

				authMode := *sc.Properties.Value("auth_mode")
				authMode.Const = authModeAWS
				properties.Set("auth_mode", &authMode)

				awsRegion := *sc.Properties.Value("aws_region")
				awsRegion.MinLength = &one
				properties.Set("aws_region", &awsRegion)

				return properties
			}(),
			Required: []string{"auth_mode", "aws_region"},
		},
		{
			// If auth_mode is not aws, don't care about aws_region
			Not: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					authMode := *sc.Properties.Value("auth_mode")
					authMode.Const = authModeAWS
					properties.Set("auth_mode", &authMode)
					return properties
				}(),
				Required: []string{"auth_mode"},
			},
		},
	}

	sc.AllOf = []*jsonschema.Schema{
		{
			// if username is set, require password to be set and auth_mode to be basic
			If: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					one := uint64(1)
					properties := jsonschema.NewProperties()
					val := *sc.Properties.Value("username")
					val.MinLength = &one
					properties.Set("username", &val)
					return properties
				}(),
				Required: []string{"username"},
			},
			Then: &jsonschema.Schema{
				// require properties not to be empty or null
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					one := uint64(1)

					password := *sc.Properties.Value("password")
					password.MinLength = &one
					properties.Set("password", &password)

					authMode := *sc.Properties.Value("auth_mode")
					authMode.Enum = []any{authModeBasic}
					authMode.Pattern = string(authModeBasic)
					properties.Set("auth_mode", &authMode)

					return properties
				}(),
				Required: []string{"password", "auth_mode"},
			},
		},
		{
			// if password is set, require username to be set
			If: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					one := uint64(1)
					properties := jsonschema.NewProperties()
					val := *sc.Properties.Value("password")
					val.MinLength = &one
					properties.Set("password", &val)
					return properties
				}(),
				Required: []string{"password"},
			},
			Then: &jsonschema.Schema{
				// require properties not to be empty or null
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					one := uint64(1)

					val := *sc.Properties.Value("username")
					val.MinLength = &one
					properties.Set("username", &val)

					return properties
				}(),
				Required: []string{"username"},
			},
		},
	}
}
