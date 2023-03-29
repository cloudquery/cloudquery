package client

import (
	"crypto/x509"
	"fmt"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
)

type Spec struct {
	Cluster          string `json:"cluster,omitempty"`
	ConnectionString string `json:"connection_string,omitempty"`
	CACert           string `json:"ca_cert,omitempty"`

	Engine *queries.Engine `json:"engine,omitempty"`
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

		if ok := tlsConfig.RootCAs.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("failed to append \"ca_cert\" value")
		}
	}

	return options, nil
}

func (s *Spec) SetDefaults() {
	if s.Engine == nil {
		s.Engine = queries.DefaultEngine()
	}
}

func (s *Spec) Validate() error {
	return s.Engine.Validate()
}
