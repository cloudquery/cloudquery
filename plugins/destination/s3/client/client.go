package client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/filetypes/v2"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec

	s3Client   *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	*filetypes.Client
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "s3").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal s3 spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(c.pluginSpec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	cfg.Region = c.pluginSpec.Region
	c.s3Client = s3.NewFromConfig(cfg)
	c.uploader = manager.NewUploader(c.s3Client)
	c.downloader = manager.NewDownloader(c.s3Client)

	// we want to run this test because we want it to fail early if the bucket is not accessible
	if err := c.testWriteAccess(ctx, cfg); err != nil {
		return nil, err
	}

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (c *Client) testWriteAccess(ctx context.Context, cfg aws.Config) error {
	stsClient := sts.NewFromConfig(cfg)
	callerIdentity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	userArn := aws.ToString(callerIdentity.Arn)
	iamClient := iam.NewFromConfig(cfg)

	userArnParsed, err := arn.Parse(userArn)
	if err != nil {
		return err
	}

	timeNow := time.Now().UTC()
	bucketArn := arn.ARN{
		Partition: userArnParsed.Partition,
		Service:   "s3",
		Resource: strings.Join([]string{
			c.pluginSpec.Bucket,
			replacePathVariables(c.pluginSpec.Path, "TEST_TABLE", "TEST_UUID", timeNow),
		}, "/"),
	}.String()

	input := &iam.SimulatePrincipalPolicyInput{
		PolicySourceArn: &userArn,
		ActionNames: []string{
			"s3:PutObject",
		},
		ResourceArns: []string{
			bucketArn,
		},
	}

	output, err := iamClient.SimulatePrincipalPolicy(ctx, input)
	if err != nil {
		return err
	}

	anyAllowed := false
	for _, o := range output.EvaluationResults {
		if o.EvalDecision == types.PolicyEvaluationDecisionTypeAllowed {
			anyAllowed = true
		}
	}
	if !anyAllowed {
		return errors.New("Write access to S3 bucket is not allowed")
	}

	return nil
}
