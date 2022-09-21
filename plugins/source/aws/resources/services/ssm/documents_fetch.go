package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsmDocuments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM

	params := ssm.ListDocumentsInput{
		Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}},
	}
	for {
		output, err := svc.ListDocuments(ctx, &params)
		if err != nil {
			return err
		}

		for _, d := range output.DocumentIdentifiers {
			dd, err := svc.DescribeDocument(ctx, &ssm.DescribeDocumentInput{Name: d.Name})
			if err != nil {
				return err
			}
			res <- dd.Document
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}

func ssmDocumentPostResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) (exitErr error) {
	d := resource.Item.(*types.DocumentDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().SSM

	input := ssm.DescribeDocumentPermissionInput{
		Name:           d.Name,
		PermissionType: types.DocumentPermissionTypeShare,
	}
	var accountIDs []string
	var infoList []types.AccountSharingInfo
	for {
		output, err := svc.DescribeDocumentPermission(ctx, &input)
		if err != nil {
			return err
		}
		accountIDs = append(accountIDs, output.AccountIds...)
		infoList = append(infoList, output.AccountSharingInfoList...)
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	if err := resource.Set("account_ids", accountIDs); err != nil {
		return err
	}
	return resource.Set("account_sharing_info", infoList)
}

func resolveDocumentARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	d := resource.Item.(*types.DocumentDescription)
	cl := meta.(*client.Client)
	return resource.Set(c.Name, cl.ARN("ssm", "document", *d.Name))
}
