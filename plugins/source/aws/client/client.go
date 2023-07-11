package client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
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
	Backend              state.Client
	specificRegions      bool
	Spec                 *Spec
	// Do not rely on this field, it will be removed once https://github.com/aws/aws-sdk-go-v2/issues/2163 is resolved
	AWSConfig *aws.Config
}

type AwsLogger struct {
	l zerolog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type ServicesPartitionAccountMap map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services ServicesPartitionAccountMap
}

const (
	defaultRegion              = "us-east-1"
	defaultVar                 = "default"
	awsCloudfrontScopeRegion   = defaultRegion
	awsCnCloudfrontScopeRegion = "cn-north-1"
)

var errInvalidRegion = errors.New("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}
var errRetrievingCredentials = errors.New("error retrieving AWS credentials (see logs for details). Please verify your credentials and try again")

var ErrPaidAPIsNotEnabled = errors.New("not fetching resource because `use_paid_apis` is set to false")

func (s *ServicesManager) ServicesByPartitionAccount(partition, accountId string) *Services {
	return s.services[partition][accountId]
}

func (s *ServicesManager) InitServices(details svcsDetail) {
	s.InitServicesForPartitionAccount(details.partition, details.accountId, details.svcs)
}

func (s *ServicesManager) InitServicesForPartitionAccount(partition, accountId string, svcs Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = &svcs
	}

	s.services[partition][accountId].Regions = funk.UniqString(append(s.services[partition][accountId].Regions, svcs.Regions...))
}

func NewAwsClient(logger zerolog.Logger, spec *Spec) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesPartitionAccountMap{},
		},
		logger: logger,
		Spec:   spec,
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
	return c.ServicesManager.ServicesByPartitionAccount(c.Partition, c.AccountID)
}

func (c *Client) Duplicate() *Client {
	duplicateClient := *c
	return &duplicateClient
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
		Spec:                 c.Spec,
		AWSConfig:            c.AWSConfig,
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
		Spec:                 c.Spec,
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
		Spec:                 c.Spec,
	}
}

func (c *Client) withLanguageCode(code string) *Client {
	newC := *c
	newC.LanguageCode = code
	newC.logger = newC.logger.With().Str("language_code", code).Logger()
	return &newC
}

// Configure is the entrypoint into configuring the AWS plugin. It is called by the plugin initialization in resources/plugin/aws.go
func Configure(ctx context.Context, logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("spec validation failed: %w", err)
	}
	spec.SetDefaults()

	client := NewAwsClient(logger, &spec)

	var adminAccountSts AssumeRoleAPIClient

	if client.Spec.Organization != nil {
		var err error
		client.Spec.Accounts, adminAccountSts, err = loadOrgAccounts(ctx, logger, client.Spec)
		if err != nil {
			logger.Error().Err(err).Msg("error getting child accounts")
			return nil, err
		}
	}
	if len(client.Spec.Accounts) == 0 {
		client.Spec.Accounts = []Account{
			{
				ID: defaultVar,
			},
		}
	}

	initLock := sync.Mutex{}
	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(client.Spec.InitializationConcurrency)
	for _, account := range client.Spec.Accounts {
		account := account
		errorGroup.Go(func() error {
			svcsDetail, err := client.setupAWSAccount(gtx, logger, client.Spec, adminAccountSts, account)
			if err != nil {
				return err
			}
			if svcsDetail == nil {
				return nil
			}
			initLock.Lock()
			defer initLock.Unlock()
			client.ServicesManager.InitServices(*svcsDetail)
			return nil
		})
	}
	if err := errorGroup.Wait(); err != nil {
		return nil, err
	}

	if len(client.ServicesManager.services) == 0 {
		// This is a special error case where we found active accounts, but just weren't able to assume a role in any of them
		if client.Spec.Organization != nil && len(client.Spec.Accounts) > 0 && client.Spec.Organization.MemberCredentials == nil {
			return nil, fmt.Errorf("discovered %d accounts in the AWS Organization, but the credentials specified in 'admin_account' were unable to assume a role in the member accounts. Verify that the role you are trying to assume (arn:aws:iam::<account_id>:role/%s) exists. If you need to use a different set of credentials to do the role assumption use 'member_trusted_principal'", len(client.Spec.Accounts), client.Spec.Organization.ChildAccountRoleName)
		}
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
