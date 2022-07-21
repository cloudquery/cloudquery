package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gopkg.in/yaml.v3"
)

type BackendType string

// currently supported backends type
// full list - https://www.terraform.io/docs/language/settings/backends/index.html
const (
	LOCAL BackendType = "local"
	S3    BackendType = "s3"
)

// BackendConfigBlock - abstract backend config
type BackendConfigBlock struct {
	BackendName string                 `yaml:"name"`
	BackendType string                 `yaml:"backend"`
	ConfigAttrs map[string]interface{} `yaml:",inline"`
}

type TerraformBackend struct {
	BackendType BackendType
	BackendName string
	Data        *TerraformData
}

type LocalBackendConfig struct {
	Path string `yaml:"path"`
}

type S3BackendConfig struct {
	Bucket  string `yaml:"bucket"`
	Key     string `yaml:"key"`
	Region  string `yaml:"region"`
	RoleArn string `yaml:"role_arn,omitempty"`
}

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

func NewS3TerraformBackend(config *BackendConfigBlock) (*TerraformBackend, error) {
	var b S3BackendConfig

	cfgBytes, _ := yaml.Marshal(config.ConfigAttrs)
	if err := yaml.Unmarshal(cfgBytes, &b); err != nil {
		return nil, fmt.Errorf("cannot parse s3 backend config: %w", err)
	}

	if b.Region == "" {
		if region, err := s3manager.GetBucketRegion(
			context.Background(),
			session.Must(session.NewSession()),
			b.Bucket,
			"us-east-1",
		); err != nil {
			return nil, err
		} else { //nolint:revive
			b.Region = region
		}
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(b.Region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	awsCfg := &aws.Config{}
	if b.RoleArn != "" {
		// if has RoleArn use it instead
		parsedArn, err := arn.Parse(b.RoleArn)
		if err != nil {
			return nil, err
		}
		creds := stscreds.NewCredentials(sess, parsedArn.String())
		awsCfg.Credentials = creds
	}
	svc := s3.New(sess, awsCfg)

	// get the tf state file
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.Bucket),
		Key:    aws.String(b.Key),
	})
	if err != nil {
		return nil, err
	}

	terraformData, err := parseAndValidate(result.Body)
	if err != nil {
		return nil, err
	}

	return &TerraformBackend{
		BackendType: S3,
		BackendName: config.BackendName,
		Data:        terraformData,
	}, nil
}

func NewLocalTerraformBackend(config *BackendConfigBlock) (*TerraformBackend, error) {
	var b LocalBackendConfig

	cfgBytes, _ := yaml.Marshal(config.ConfigAttrs)
	if err := yaml.Unmarshal(cfgBytes, &b); err != nil {
		return nil, fmt.Errorf("cannot parse local backend config: %w", err)
	}

	f, err := os.Open(b.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read tfstate from %s", b.Path)
	}
	defer f.Close()

	terraformData, err := parseAndValidate(f)
	if err != nil {
		return nil, err
	}

	return &TerraformBackend{
		BackendType: LOCAL,
		BackendName: config.BackendName,
		Data:        terraformData,
	}, nil
}

// NewBackend initialize function
func NewBackend(cfg *BackendConfigBlock) (*TerraformBackend, error) {
	switch cfg.BackendType {
	case "local":
		localBackend, err := NewLocalTerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return localBackend, nil
	case "s3":
		s3Backend, err := NewS3TerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return s3Backend, nil
	default:
		return nil, errors.New("unsupported backend")
	}
}
