package client

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	"github.com/meilisearch/meilisearch-go"
	"github.com/valyala/fasthttp"
)

type Spec struct {
	// A Meilisearch instance host & port to use.
	// If your Meilisearch instance uses private SSL certificate, make sure to specify `ca_cert` option, too.
	Host string `json:"host" jsonschema:"required,minLength=1"`

	//   Meilisearch API key, granted the following actions:
	//
	//  - `documents.add`
	//  - `indexes.create`
	//  - `indexes.get`
	//  - `indexes.update`
	//  - `tasks.get`
	//  - `settings.get`
	//  - `settings.update`
	//  - `version`
	APIKey string `json:"api_key" jsonschema:"required,minLength=1"`

	// Meilisearch API client timeout.
	Timeout *configtype.Duration `json:"timeout,omitempty"`

	//  PEM-encoded certificate authorities.
	//  When set, a certificate pool will be created by appending the certificates to the system pool.
	//  See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example) for how to read this value from a file.
	CACert string `json:"ca_cert,omitempty"`

	// Maximum amount of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`

	// Timeout for writing a single batch.
	BatchTimeout *configtype.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) validate() error {
	switch {
	case len(s.Host) == 0:
		return errors.New("empty \"host\" value")
	case len(s.APIKey) == 0:
		return errors.New("empty \"api_key\" value")
	default:
		return nil
	}
}

func (s *Spec) setDefaults() {
	if s.Timeout == nil {
		d := configtype.NewDuration(5 * time.Minute) // 5m
		s.Timeout = &d
	}

	if s.BatchSize == 0 {
		s.BatchSize = 1000 // 1K
	}

	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = 4 << 20 // 4 MiB
	}

	if s.BatchTimeout == nil {
		d := configtype.NewDuration(20 * time.Second) // 20s
		s.BatchTimeout = &d
	}
}

func (s *Spec) getClient() (*meilisearch.Client, error) {
	config := meilisearch.ClientConfig{
		Host:    s.Host,
		APIKey:  s.APIKey,
		Timeout: s.Timeout.Duration(),
	}
	if len(s.CACert) == 0 {
		return meilisearch.NewClient(config), nil
	}

	// read file
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	if ok := certPool.AppendCertsFromPEM([]byte(s.CACert)); !ok {
		return nil, errors.New("failed to append \"ca_cert\" value")
	}

	httpClient := &fasthttp.Client{
		Name: "meilisearch-client-with-custom-tls",
		// Reuse the most recently-used idle connection.
		ConnPoolStrategy: fasthttp.LIFO,
		// Add tls config
		TLSConfig: &tls.Config{RootCAs: certPool},
	}

	return meilisearch.NewFastHTTPCustomClient(config, httpClient), nil
}

//go:embed schema.json
var JSONSchema string

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	timeout := sc.Properties.Value("timeout").OneOf[0] // 0 - val, 1 - null
	timeout.Default = "5m"

	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "20s"
}
