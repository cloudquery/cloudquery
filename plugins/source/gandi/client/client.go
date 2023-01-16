package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/go-gandi/go-gandi"
	"github.com/go-gandi/go-gandi/config"
	"github.com/rs/zerolog"
)

type Services struct {
	CertificateClient   CertificateClient
	DomainClient        DomainClient
	LiveDNSClient       LiveDNSClient
	SimpleHostingClient SimpleHostingClient
}

type Client struct {
	logger    zerolog.Logger
	sharingID string

	Services Services
}

const MaxItemsPerPage = 200

func New(logger zerolog.Logger, services Services, sharingId string) Client {
	return Client{
		logger:    logger,
		sharingID: sharingId,

		Services: services,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sharingID
}

func Configure(_ context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	gaSpec := &Spec{}
	if err := s.UnmarshalSpec(gaSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gandi spec: %w", err)
	}

	services, err := getGandiServices(gaSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, *services, gaSpec.SharingID)
	return &cl, nil
}

func getGandiServices(spec *Spec) (*Services, error) {
	if spec.APIKey == "" {
		return nil, errors.New("no API key provided")
	}

	gCfg := config.Config{
		APIKey:    spec.APIKey,
		SharingID: spec.SharingID,
		Debug:     spec.GandiDebug,
		APIURL:    spec.EndpointURL,
	}
	if spec.Timeout > 0 {
		gCfg.Timeout = time.Duration(spec.Timeout) * time.Second
	}
	return &Services{
		CertificateClient:   gandi.NewCertificateClient(gCfg),
		DomainClient:        gandi.NewDomainClient(gCfg),
		LiveDNSClient:       gandi.NewLiveDNSClient(gCfg),
		SimpleHostingClient: gandi.NewSimpleHostingClient(gCfg),
	}, nil
}
