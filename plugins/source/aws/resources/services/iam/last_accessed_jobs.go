package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

type Job struct {
	Arn   string `json:"arn"`
	JobId string `json:"job_id"`
}

func userLastAccessedJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_user_last_accessed_jobs",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateServiceLastAccessedDetails.html`,
		Resolver:    fetchUserLastAccessedJobs,
		Transform:   transformers.TransformWithStruct(&Job{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
		Relations: []*schema.Table{userLastAccessedDetails()},
	}
}

func roleLastAccessedJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_role_last_accessed_jobs",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateServiceLastAccessedDetails.html`,
		Resolver:    fetchRoleLastAccessedJobs,
		Transform:   transformers.TransformWithStruct(&Job{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
		Relations: []*schema.Table{roleLastAccessedDetails()},
	}
}

func groupLastAccessedJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_group_last_accessed_jobs",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateServiceLastAccessedDetails.html`,
		Resolver:    fetchGroupLastAccessedJobs,
		Transform:   transformers.TransformWithStruct(&Job{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
		Relations: []*schema.Table{groupLastAccessedDetails()},
	}
}

func policyLastAccessedJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_policy_last_accessed_jobs",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateServiceLastAccessedDetails.html`,
		Resolver:    fetchPolicyLastAccessedJobs,
		Transform:   transformers.TransformWithStruct(&Job{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
		Relations: []*schema.Table{policyLastAccessedDetails()},
	}
}

func fetchUserLastAccessedJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	user := parent.Item.(*types.User)
	return fetchLastAccessedJobs(ctx, meta, user.Arn, res)
}

func fetchRoleLastAccessedJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	role := parent.Item.(*types.Role)
	return fetchLastAccessedJobs(ctx, meta, role.Arn, res)
}

func fetchGroupLastAccessedJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	group := parent.Item.(types.Group)
	return fetchLastAccessedJobs(ctx, meta, group.Arn, res)
}

func fetchPolicyLastAccessedJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	policy := parent.Item.(types.ManagedPolicyDetail)
	return fetchLastAccessedJobs(ctx, meta, policy.Arn, res)
}

func fetchLastAccessedJobs(ctx context.Context, meta schema.ClientMeta, arn *string, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam
	config := iam.GenerateServiceLastAccessedDetailsInput{
		Arn:         arn,
		Granularity: types.AccessAdvisorUsageGranularityTypeActionLevel,
	}
	output, err := svc.GenerateServiceLastAccessedDetails(ctx, &config)
	if err != nil {
		return err
	}

	res <- Job{Arn: *arn, JobId: *output.JobId}

	return nil
}
