package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type BackendType string

// BackendConfigBlock - abstract backend config
type BackendConfigBlock struct {
	BackendName string           `json:"name"`
	Type        BackendType      `json:"type"`
	Config      *json.RawMessage `json:"config"`
}

func (bc *BackendConfigBlock) Validate() error {
	if bc.Config == nil {
		return errors.New("missing `config` in spec")
	}
	return nil
}

type TerraformBackend struct {
	BackendType BackendType
	BackendName string
	Data        *TerraformData
}

type LocalBackendConfig struct {
	Path string `json:"path"`
}

type S3BackendConfig struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	Region  string `json:"region"`
	RoleArn string `json:"role_arn,omitempty"`
}

// currently supported backends type
// full list - https://www.terraform.io/docs/language/settings/backends/index.html
const (
	LOCAL BackendType = "local"
	S3    BackendType = "s3"
)

// parseAndValidate received reader turn in into TerraformData state and validate the state version
func parseAndValidate(reader io.Reader) (*TerraformData, error) {
	var s TerraformData
	if err := json.NewDecoder(reader).Decode(&s.State); err != nil {
		return nil, fmt.Errorf("invalid tf state file")
	}
	if s.State.Version != StateVersion {
		return nil, fmt.Errorf("unsupported state version %d", s.State.Version)
	}
	return &s, nil
}

func NewS3TerraformBackend(ctx context.Context, cfg *BackendConfigBlock) (*TerraformBackend, error) {
	var b S3BackendConfig
	if err := json.Unmarshal(*cfg.Config, &b); err != nil {
		return nil, fmt.Errorf("cannot parse s3 backend config: %w", err)
	}

	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	if b.Region == "" {
		region, err := manager.GetBucketRegion(ctx, s3.NewFromConfig(awsConfig), b.Bucket,
			func(options *s3.Options) { options.Region = "us-east-1" })
		if err != nil {
			return nil, err
		}
		b.Region = region
	}
	awsConfig.Region = b.Region

	if err != nil {
		return nil, err
	}

	if b.RoleArn != "" {
		// if it has RoleArn use it instead
		parsedArn, err := arn.Parse(b.RoleArn)
		if err != nil {
			return nil, err
		}
		creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awsConfig), parsedArn.String())
		awsConfig.Credentials = creds
	}
	svc := s3.NewFromConfig(awsConfig)

	// get the tf state file
	result, err := svc.GetObject(ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(b.Bucket),
			Key:    aws.String(b.Key),
		},
	)
	if err != nil {
		return nil, err
	}

	terraformData, err := parseAndValidate(result.Body)
	if err != nil {
		return nil, err
	}

	return &TerraformBackend{
		BackendType: cfg.Type,
		BackendName: cfg.BackendName,
		Data:        terraformData,
	}, nil
}

func NewLocalTerraformBackend(cfg *BackendConfigBlock) (*TerraformBackend, error) {
	var b LocalBackendConfig
	if err := json.Unmarshal(*cfg.Config, &b); err != nil {
		return nil, fmt.Errorf("cannot parse local backend config: %w", err)
	}

	f, err := os.Open(b.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read tfstate from %s: %w", b.Path, err)
	}
	defer f.Close()

	terraformData, err := parseAndValidate(f)
	if err != nil {
		return nil, err
	}

	return &TerraformBackend{
		BackendType: cfg.Type,
		BackendName: cfg.BackendName,
		Data:        terraformData,
	}, nil
}

// NewBackend initialize function
func NewBackend(ctx context.Context, cfg *BackendConfigBlock) (*TerraformBackend, error) {
	switch cfg.Type {
	case LOCAL:
		return NewLocalTerraformBackend(cfg)
	case S3:
		return NewS3TerraformBackend(ctx, cfg)
	default:
		return nil, fmt.Errorf("unsupported backend: %q", cfg.Type)
	}
}
