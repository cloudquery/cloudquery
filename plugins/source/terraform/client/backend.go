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
	BackendName string              `json:"name"`
	Local       *LocalBackendConfig `json:"local,omitempty"`
	S3          *S3BackendConfig    `json:"s3,omitempty"`
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
	if config.S3 == nil {
		return nil, errors.New("cannot parse s3 backend config")
	}

	if config.S3.Region == "" {
		if region, err := s3manager.GetBucketRegion(
			context.Background(),
			session.Must(session.NewSession()),
			config.S3.Bucket,
			"us-east-1",
		); err != nil {
			return nil, err
		} else { //nolint:revive
			config.S3.Region = region
		}
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(config.S3.Region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	awsCfg := &aws.Config{}
	if config.S3.RoleArn != "" {
		// if it has RoleArn use it instead
		parsedArn, err := arn.Parse(config.S3.RoleArn)
		if err != nil {
			return nil, err
		}
		creds := stscreds.NewCredentials(sess, parsedArn.String())
		awsCfg.Credentials = creds
	}
	svc := s3.New(sess, awsCfg)

	// get the tf state file
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.S3.Bucket),
		Key:    aws.String(config.S3.Key),
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
	if config.Local == nil {
		return nil, errors.New("cannot parse local backend config")
	}

	f, err := os.Open(config.Local.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read tfstate from %s", config.Local.Path)
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
	if cfg.Local != nil && cfg.S3 != nil {
		return nil, errors.New("cannot have both local and s3 backends")
	}

	if cfg.Local != nil {
		localBackend, err := NewLocalTerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return localBackend, nil
	}
	if cfg.S3 != nil {
		s3Backend, err := NewS3TerraformBackend(cfg)
		if err != nil {
			return nil, err
		}
		return s3Backend, nil
	}

	return nil, errors.New("unsupported backend: specify local or s3")
}
