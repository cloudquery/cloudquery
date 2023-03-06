package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

type LastAccessed struct {
	Arn   string `json:"arn"`
	JobId string `json:"job_id"`
	types.ServiceLastAccessed
}

func userLastAccessedDetails() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_user_last_accessed_details",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Arn", "ServiceNamespace")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func roleLastAccessedDetails() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_role_last_accessed_details",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Arn", "ServiceNamespace")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func groupLastAccessedDetails() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_group_last_accessed_details",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Arn", "ServiceNamespace")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func policyLastAccessedDetails() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_policy_last_accessed_details",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Arn", "ServiceNamespace")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam
	job := parent.Item.(Job)
	config := iam.GetServiceLastAccessedDetailsInput{
		JobId:    &job.JobId,
		MaxItems: aws.Int32(1000),
	}

	for {
		details, err := svc.GetServiceLastAccessedDetails(ctx, &config)
		if err != nil {
			return err
		}

		switch details.JobStatus {
		case types.JobStatusTypeInProgress:
			time.Sleep(time.Second)
			continue
		case types.JobStatusTypeFailed:
			return fmt.Errorf("failed to get last accessed details with error: %s - %s", *details.Error.Code, *details.Error.Message)
		case types.JobStatusTypeCompleted:
			for _, detail := range details.ServicesLastAccessed {
				res <- LastAccessed{
					Arn:                 job.Arn,
					JobId:               job.JobId,
					ServiceLastAccessed: detail,
				}
			}
			if details.Marker == nil {
				return nil
			}
			if details.Marker != nil {
				config.Marker = details.Marker
			}
		}
	}
}
