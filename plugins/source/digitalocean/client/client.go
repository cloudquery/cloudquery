package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	MaxItemsPerPage = 200

	firstSpacesRegion = "nyc3"
)

var defaultSpacesRegions = []string{"nyc3", "sfo3", "ams3", "sgp1", "fra1", "syd1"}

type Client struct {
	logger           zerolog.Logger
	DoClient         *godo.Client
	Regions          []string
	SpacesRegion     string
	CredentialStatus DoCredentialStruct
	Services         *Services
}

type DoCredentialStruct struct {
	Api    bool
	Spaces bool
}

type SpacesCredentialsProvider struct {
	SpacesAccessKey   string
	SpacesAccessKeyId string
}

type SpacesEndpointResolver struct{}

type DoLogger struct {
	l zerolog.Logger
}

func (s SpacesCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     s.SpacesAccessKeyId,
		SecretAccessKey: s.SpacesAccessKey,
		Source:          "digitalocean",
	}, nil
}

func (c *Client) WithSpacesRegion(region string) *Client {
	return &Client{
		logger:       c.Logger().With().Str("spaces_region", region).Logger(),
		DoClient:     c.DoClient,
		SpacesRegion: region,
		Services:     initServices(c.DoClient, c.Services.Spaces),
	}
}

func (SpacesEndpointResolver) ResolveEndpoint(_, region string, options ...any) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:    fmt.Sprintf("https://%s.digitaloceanspaces.com", region),
		Source: aws.EndpointSourceCustom,
	}, nil
}

type Services struct {
	Account        AccountService
	Cdn            CdnService
	BillingHistory BillingHistoryService
	Monitoring     MonitoringService
	Balance        BalanceService
	Certificates   CertificatesService
	Databases      DatabasesService
	Domains        DomainsService
	Droplets       DropletsService
	Firewalls      FirewallsService
	FloatingIps    FloatingIpsService
	Images         ImagesService
	Keys           KeysService
	LoadBalancers  LoadBalancersService
	Projects       ProjectsService
	Regions        RegionsService
	Registry       RegistryService
	Sizes          SizesService
	Snapshots      SnapshotsService
	Storage        StorageService
	Vpcs           VpcsService
	Spaces         SpacesService
}

type ServicesRegionMap map[string]*Services

// ServicesManager will hold the entire map of region services
type ServicesManager struct {
	services ServicesRegionMap
}

func (s *ServicesManager) ServicesByRegion(region string) *Services {
	return s.services[region]
}

func New(logger zerolog.Logger, doSpec Spec) (schema.ClientMeta, error) {
	credStatus := DoCredentialStruct{
		Api:    true,
		Spaces: true,
	}

	if doSpec.SpacesAccessKey == "" || doSpec.SpacesAccessKeyId == "" {
		doSpec.SpacesAccessKeyId, doSpec.SpacesAccessKey = getSpacesTokenFromEnv()
	}
	if doSpec.SpacesAccessKey == "" || doSpec.SpacesAccessKeyId == "" {
		credStatus.Spaces = false
	}

	awsCfg, err := awscfg.LoadDefaultConfig(context.Background(),
		awscfg.WithCredentialsProvider(SpacesCredentialsProvider{doSpec.SpacesAccessKey, doSpec.SpacesAccessKeyId}),
		awscfg.WithEndpointResolverWithOptions(SpacesEndpointResolver{}),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if doSpec.SpacesDebugLogging {
		awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		awsCfg.Logger = DoLogger{logger}
	}

	spacesRegions := defaultSpacesRegions
	if len(doSpec.SpacesRegions) > 0 {
		spacesRegions = doSpec.SpacesRegions
	}

	doClient := godo.NewFromToken(doSpec.Token)

	return &Client{
		logger:           logger,
		DoClient:         doClient,
		Regions:          spacesRegions,
		SpacesRegion:     firstSpacesRegion,
		CredentialStatus: credStatus,
		Services:         initServices(doClient, s3.NewFromConfig(awsCfg)),
	}, nil
}

func (a DoLogger) Logf(classification logging.Classification, format string, v ...any) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c Client) ID() string {
	return c.SpacesRegion
}

func initServices(doClient *godo.Client, spacesService SpacesService) *Services {
	return &Services{
		Account:        doClient.Account,
		Cdn:            doClient.CDNs,
		BillingHistory: doClient.BillingHistory,
		Monitoring:     doClient.Monitoring,
		Balance:        doClient.Balance,
		Certificates:   doClient.Certificates,
		Databases:      doClient.Databases,
		Domains:        doClient.Domains,
		Droplets:       doClient.Droplets,
		Firewalls:      doClient.Firewalls,
		FloatingIps:    doClient.FloatingIPs,
		Images:         doClient.Images,
		Keys:           doClient.Keys,
		LoadBalancers:  doClient.LoadBalancers,
		Projects:       doClient.Projects,
		Regions:        doClient.Regions,
		Registry:       doClient.Registry,
		Sizes:          doClient.Sizes,
		Snapshots:      doClient.Snapshots,
		Storage:        doClient.Storage,
		Vpcs:           doClient.VPCs,
		Spaces:         spacesService,
	}
}
