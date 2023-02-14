package client

import (
	"crypto/x509"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	CACert           string `json:"ca_cert,omitempty"`
}

func (s *Spec) Options() (*clickhouse.Options, error) {
	options, err := clickhouse.ParseDSN(s.ConnectionString)
	if err != nil {
		return nil, err
	}

	// set database name to "default", if empty
	if len(options.Auth.Database) == 0 {
		options.Auth.Database = "default"
	}

	if tlsConfig := options.TLS; tlsConfig != nil && len(s.CACert) > 0 {
		// read file
		caCert, err := os.ReadFile(s.CACert)
		if err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
			// no such file. treat as plain input
			caCert = []byte(s.CACert)
		}

		if tlsConfig.RootCAs == nil {
			tlsConfig.RootCAs, err = x509.SystemCertPool()
			if err != nil {
				return nil, err
			}
		}

		tlsConfig.RootCAs.AppendCertsFromPEM(caCert)
	}

	return options, nil
}
