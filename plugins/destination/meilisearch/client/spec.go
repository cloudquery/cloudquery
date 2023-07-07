package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/meilisearch/meilisearch-go"
	"github.com/valyala/fasthttp"
)

type Spec struct {
	// required
	Host   string `json:"host,omitempty"`
	APIKey string `json:"api_key,omitempty"`

	// optional
	Timeout time.Duration `json:"timeout,omitempty"`
	CACert  string        `json:"ca_cert,omitempty"`

	BatchSize      int                  `json:"batch_size,omitempty"`
	BatchSizeBytes int                  `json:"batch_size_bytes,omitempty"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) validate() error {
	switch {
	case len(s.Host) == 0:
		return fmt.Errorf("empty \"host\" value")
	case len(s.APIKey) == 0:
		return fmt.Errorf("empty \"api_key\" value")
	default:
		return nil
	}
}

func (s *Spec) setDefaults() {
	if s.Timeout == 0 {
		s.Timeout = 5 * time.Minute
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
		Timeout: s.Timeout,
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
		return nil, fmt.Errorf("failed to append \"ca_cert\" value")
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
