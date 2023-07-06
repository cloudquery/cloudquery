package ssm

import (
	"context"
	"fmt"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveDocumentARN,
				PrimaryKey: true,
			},
			{
				Name:     "permissions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDocumentPermission,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			documentVersions(),
		},
	}
}

func fetchSsmDocuments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	params := ssm.ListDocumentsInput{
		Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}},
	}
	paginator := ssm.NewListDocumentsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DocumentIdentifiers
	}
	return nil
}

func getDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm
	d := resource.Item.(types.DocumentIdentifier)

	dd, err := svc.DescribeDocument(ctx, &ssm.DescribeDocumentInput{Name: d.Name}, func(o *ssm.Options) {
		o.Region = cl.Region
	})
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
	// No paginator
	for {
		output, err := svc.DescribeDocumentPermission(ctx, &input, func(o *ssm.Options) {
			o.Region = cl.Region
		})
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
