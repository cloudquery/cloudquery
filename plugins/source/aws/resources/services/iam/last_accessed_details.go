package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type LastAccessed struct {
	Arn   string `json:"arn"`
	JobId string `json:"job_id"`
	types.ServiceLastAccessed
}

func userLastAccessedDetails() *schema.Table {
	tableName := "aws_iam_user_last_accessed_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchUserLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("ServiceNamespace"), transformers.WithSkipFields("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "user_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func roleLastAccessedDetails() *schema.Table {
	tableName := "aws_iam_role_last_accessed_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchRoleLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("ServiceNamespace"), transformers.WithSkipFields("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "role_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func groupLastAccessedDetails() *schema.Table {
	tableName := "aws_iam_group_last_accessed_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchGroupLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("ServiceNamespace"), transformers.WithSkipFields("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "group_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func policyLastAccessedDetails() *schema.Table {
	tableName := "aws_iam_policy_last_accessed_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html`,
		Resolver:    fetchPolicyLastAccessedDetails,
		Transform:   transformers.TransformWithStruct(&LastAccessed{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithPrimaryKeys("Arn", "ServiceNamespace")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchUserLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	user := parent.Item.(*types.User)
	return fetchLastAccessedDetails(ctx, meta, user.Arn, res)
}

func fetchRoleLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	role := parent.Item.(*types.Role)
	return fetchLastAccessedDetails(ctx, meta, role.Arn, res)
}

func fetchGroupLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	group := parent.Item.(types.Group)
	return fetchLastAccessedDetails(ctx, meta, group.Arn, res)
}

func fetchPolicyLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	policy := parent.Item.(types.ManagedPolicyDetail)
	return fetchLastAccessedDetails(ctx, meta, policy.Arn, res)
}

func fetchLastAccessedDetails(ctx context.Context, meta schema.ClientMeta, arn *string, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	generateConfig := iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         arn,
		Granularity: types.AccessAdvisorUsageGranularityTypeActionLevel,
	}
	output, err := svc.GenerateServiceLastAccessedDetails(ctx, &generateConfig, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	jobId := output.JobId

	config := iam.GetServiceLastAccessedDetailsInput{
		JobId:    jobId,
		MaxItems: aws.Int32(1000),
	}

	for {
		details, err := svc.GetServiceLastAccessedDetails(ctx, &config, func(options *iam.Options) {
			options.Region = cl.Region
		})
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
					Arn:                 *arn,
					JobId:               *jobId,
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
