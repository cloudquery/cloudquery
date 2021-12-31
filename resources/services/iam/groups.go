package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_groups",
		Description:  "Contains information about an IAM group entity.",
		Resolver:     fetchIamGroups,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "policies",
				Description: "List of policies attached to group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveIamGroupPolicies,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the group. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the group was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "id",
				Description: "The stable and unique string identifying the group. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupId"),
			},
			{
				Name:        "name",
				Description: "The friendly name that identifies the group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupName"),
			},
			{
				Name:        "path",
				Description: "The path to the group. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			IamGroupPolicies(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Groups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Group)
	svc := meta.(*client.Client).Services().IAM
	config := iam.ListAttachedGroupPoliciesInput{
		GroupName: r.GroupName,
	}
	response, err := svc.ListAttachedGroupPolicies(ctx, &config)
	if err != nil {
		return err
	}
	policyMap := map[string]*string{}
	for _, p := range response.AttachedPolicies {
		policyMap[*p.PolicyArn] = p.PolicyName
	}
	return resource.Set(c.Name, policyMap)
}
