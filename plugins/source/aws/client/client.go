package client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/plugin-sdk/v2/backend"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	ServicesManager ServicesManager
	logger          zerolog.Logger
	// this is set by table clientList
	AccountID            string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
	Partition            string
	LanguageCode         string
	Backend              backend.Backend
	specificRegions      bool
}

type AwsLogger struct {
	l zerolog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type ServicesPartitionAccountRegionMap map[string]map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services         ServicesPartitionAccountRegionMap
	wafScopeServices map[string]map[string]*Services
}

const (
	defaultRegion         = "us-east-1"
	defaultVar            = "default"
	cloudfrontScopeRegion = defaultRegion
)

var errInvalidRegion = errors.New("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}
var errRetrievingCredentials = errors.New("error retrieving AWS credentials (see logs for details). Please verify your credentials and try again")

func (s *ServicesManager) ServicesByPartitionAccountAndRegion(partition, accountId, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[partition][accountId][region]
}

func (s *ServicesManager) ServicesByAccountForWAFScope(partition, accountId string) *Services {
	return s.wafScopeServices[partition][accountId]
}

func (s *ServicesManager) InitServices(details svcsDetail) {
	if details.region != "" {
		s.InitServicesForPartitionAccountAndRegion(details.partition, details.accountId, details.region, details.svcs)
	} else {
		s.InitServicesForPartitionAccountAndScope(details.partition, details.accountId, details.svcs)
	}
}

func (s *ServicesManager) InitServicesForPartitionAccountAndRegion(partition, accountId, region string, svcs Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = make(map[string]*Services)
	}
	s.services[partition][accountId][region] = &svcs
}

func (s *ServicesManager) InitServicesForPartitionAccountAndScope(partition, accountId string, svcs Services) {
	if s.wafScopeServices == nil {
		s.wafScopeServices = make(map[string]map[string]*Services)
	}
	if s.wafScopeServices[partition] == nil {
		s.wafScopeServices[partition] = make(map[string]*Services)
	}
	s.wafScopeServices[partition][accountId] = &svcs
}

func NewAwsClient(logger zerolog.Logger, b backend.Backend) Client {
	return Client{
		Backend: b,
		ServicesManager: ServicesManager{
			services: ServicesPartitionAccountRegionMap{},
		},
		logger: logger,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	idStrings := []string{
		c.AccountID,
		c.Region,
		c.AutoscalingNamespace,
		string(c.WAFScope),
		c.LanguageCode,
	}

	return strings.TrimRight(strings.Join(idStrings, ":"), ":")
}

func (c *Client) Services() *Services {
	s := c.ServicesManager.ServicesByPartitionAccountAndRegion(c.Partition, c.AccountID, c.Region)
	if s == nil && c.WAFScope == wafv2types.ScopeCloudfront {
		return c.ServicesManager.ServicesByAccountForWAFScope(c.Partition, c.AccountID)
	}
	return s
}

func (c *Client) withPartitionAccountIDAndRegion(partition, accountID, region string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             c.WAFScope,
		Backend:              c.Backend,
	}
}

func (c *Client) withPartitionAccountIDRegionAndNamespace(partition, accountID, region, namespace string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("autoscaling_namespace", namespace).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: namespace,
		WAFScope:             c.WAFScope,
		Backend:              c.Backend,
	}
}

func (c *Client) withPartitionAccountIDRegionAndScope(partition, accountID, region string, scope wafv2types.Scope) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("waf_scope", string(scope)).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             scope,
		Backend:              c.Backend,
	}
}

func (c *Client) withLanguageCode(code string) *Client {
	newC := *c
	newC.LanguageCode = code
	newC.logger = newC.logger.With().Str("language_code", code).Logger()
	return &newC
}

// This is the entrypoint into configuring the AWS plugin. It is called by the plugin initialization in resources/plugin/aws.go
func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var awsPluginSpec Spec
	err := spec.UnmarshalSpec(&awsPluginSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	err = awsPluginSpec.Validate()
	if err != nil {
		return nil, fmt.Errorf("spec validation failed: %w", err)
	}

	awsPluginSpec.SetDefaults()

	client := NewAwsClient(logger, opts.Backend)

	var adminAccountSts AssumeRoleAPIClient

	if awsPluginSpec.Organization != nil {
		var err error
		awsPluginSpec.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, &awsPluginSpec)
		if err != nil {
			logger.Error().Err(err).Msg("error getting child accounts")
			return nil, err
		}
	}
	if len(awsPluginSpec.Accounts) == 0 {
		awsPluginSpec.Accounts = append(awsPluginSpec.Accounts, Account{
			ID: defaultVar,
		})
	}

	initLock := sync.Mutex{}
	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(awsPluginSpec.InitializationConcurrency)
	for _, account := range awsPluginSpec.Accounts {
		account := account
		errorGroup.Go(func() error {
			svcsDetail, err := client.setupAWSAccount(gtx, logger, &awsPluginSpec, adminAccountSts, account)
			if err != nil {
				return err
			}
			initLock.Lock()
			for _, details := range svcsDetail {
				client.ServicesManager.InitServices(details)
			}
			defer initLock.Unlock()
			return nil
		})
	}
	err = errorGroup.Wait()
	if err != nil {
		return nil, err
	}

	if len(client.ServicesManager.services) == 0 {
		return nil, fmt.Errorf("no enabled accounts instantiated")
	}
	return &client, nil
}

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...any) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}
