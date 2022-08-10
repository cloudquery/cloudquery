package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger           hclog.Logger
	DoClient         *godo.Client
	Regions          []string
	SpacesRegion     string
	CredentialStatus DoCredentialStruct
	S3               *s3.Client
}

type DoCredentialStruct struct {
	Api    bool
	Spaces bool
}

type SpacesCredentialsProvider struct {
	SpacesAccessKey   string
	SpacesAccessKeyId string
}

type AwsLogger struct {
	l hclog.Logger
}

type SpacesEndpointResolver struct{}

const MaxItemsPerPage = 200

var defaultSpacesRegions = []string{"nyc3", "sfo3", "ams3", "sgp1", "fra1"}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func (c *Client) WithSpacesRegion(region string) *Client {
	return &Client{
		logger:       c.Logger().With("spaces_region", region),
		DoClient:     c.DoClient,
		SpacesRegion: region,
		S3:           c.S3,
	}
}

func (s SpacesCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     s.SpacesAccessKeyId,
		SecretAccessKey: s.SpacesAccessKey,
		Source:          "digitalocean",
	}, nil
}

func (SpacesEndpointResolver) ResolveEndpoint(_, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:    fmt.Sprintf("https://%s.digitaloceanspaces.com", region),
		Source: aws.EndpointSourceCustom,
	}, nil
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	if providerConfig.Token == "" {
		providerConfig.Token = getTokenFromEnv()
	}
	if providerConfig.Token == "" {
		return nil, diag.FromError(errors.New("missing API token"), diag.USER)
	}

	credStatus := DoCredentialStruct{
		Api:    true,
		Spaces: true,
	}

	if providerConfig.SpacesAccessKey == "" || providerConfig.SpacesAccessKeyId == "" {
		providerConfig.SpacesAccessKeyId, providerConfig.SpacesAccessKey = getSpacesTokenFromEnv()
	}
	if providerConfig.SpacesAccessKey == "" || providerConfig.SpacesAccessKeyId == "" {
		credStatus.Spaces = false
	}

	awsCfg, err := awscfg.LoadDefaultConfig(context.Background(),
		awscfg.WithCredentialsProvider(SpacesCredentialsProvider{providerConfig.SpacesAccessKey, providerConfig.SpacesAccessKeyId}),
		awscfg.WithEndpointResolver(SpacesEndpointResolver{}),
	)

	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}

	if providerConfig.SpacesDebugLogging {
		awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		awsCfg.Logger = AwsLogger{logger}
	}

	spacesRegions := defaultSpacesRegions
	if len(providerConfig.SpacesRegions) > 0 {
		spacesRegions = providerConfig.SpacesRegions
	}

	client := Client{
		logger:           logger,
		DoClient:         godo.NewFromToken(providerConfig.Token),
		Regions:          spacesRegions,
		SpacesRegion:     "nyc3",
		S3:               s3.NewFromConfig(awsCfg),
		CredentialStatus: credStatus,
	}
	return &client, nil
}

func getTokenFromEnv() string {
	doToken := os.Getenv("DIGITALOCEAN_TOKEN")
	doAccessToken := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if doToken != "" {
		return doToken
	}
	if doAccessToken != "" {
		return doAccessToken
	}
	return ""
}

func getSpacesTokenFromEnv() (string, string) {
	spacesAccessKey := os.Getenv("SPACES_ACCESS_KEY_ID")
	spacesSecretKey := os.Getenv("SPACES_SECRET_ACCESS_KEY")
	if spacesAccessKey == "" {
		return "", ""
	}
	if spacesSecretKey == "" {
		return "", ""
	}
	return spacesAccessKey, spacesSecretKey
}

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	if classification == logging.Warn {
		a.l.Warn(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug(fmt.Sprintf(format, v...))
	}
}
