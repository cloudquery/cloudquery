package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Documents() *schema.Table {
	tableName := "aws_ssm_documents"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentDescription.html`,
		Resolver:            fetchSsmDocuments,
		PreResourceResolver: getDocument,
		Transform:           transformers.TransformWithStruct(&types.DocumentDescription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveDocumentARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: resolveDocumentPermission,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			documentVersions(),
		},
	}
}

func fetchSsmDocuments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm

	params := ssm.ListDocumentsInput{
		Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}},
	}
	for {
		output, err := svc.ListDocuments(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.DocumentIdentifiers

		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}

func getDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm
	d := resource.Item.(types.DocumentIdentifier)

	dd, err := svc.DescribeDocument(ctx, &ssm.DescribeDocumentInput{Name: d.Name})
	if err != nil {
		return err
	}

	resource.Item = dd.Document
	return nil
}

func resolveDocumentPermission(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) (exitErr error) {
	d := resource.Item.(*types.DocumentDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	input := ssm.DescribeDocumentPermissionInput{
		Name:           d.Name,
		PermissionType: types.DocumentPermissionTypeShare,
	}
	var permissions []*ssm.DescribeDocumentPermissionOutput
	for {
		output, err := svc.DescribeDocumentPermission(ctx, &input)
		if err != nil {
			return err
		}
		permissions = append(permissions, output)
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return resource.Set(col.Name, permissions)
}

func resolveDocumentARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	d := resource.Item.(*types.DocumentDescription)
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "ssm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("document/%s", aws.ToString(d.Name)),
	}.String())
}
