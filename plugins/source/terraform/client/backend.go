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
)

type BackendType string

// BackendConfigBlock - abstract backend config
type BackendConfigBlock struct {
	BackendName string           `json:"name"`
	Type        BackendType      `json:"type"`
	Config      *json.RawMessage `json:"config"`
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

func NewS3TerraformBackend(config *BackendConfigBlock) (*TerraformBackend, error) {
	var b S3BackendConfig
	if err := json.Unmarshal(*config.Config, &b); err != nil {
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
		// if it has RoleArn use it instead
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
		BackendType: config.Type,
		BackendName: config.BackendName,
		Data:        terraformData,
	}, nil
}

func NewLocalTerraformBackend(config *BackendConfigBlock) (*TerraformBackend, error) {
	var b LocalBackendConfig
	if err := json.Unmarshal(*config.Config, &b); err != nil {
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
		BackendType: config.Type,
		BackendName: config.BackendName,
		Data:        terraformData,
	}, nil
}

// NewBackend initialize function
func NewBackend(cfg *BackendConfigBlock) (*TerraformBackend, error) {
	if cfg.Config == nil {
		return nil, errors.New("missing `config` in spec")
	}

	switch cfg.Type {
	case LOCAL:
		localBackend, err := NewLocalTerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return localBackend, nil
	case S3:
		s3Backend, err := NewS3TerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return s3Backend, nil
	default:
		return nil, fmt.Errorf("unsupported backend: %q", cfg.Type)
	}
}
