package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func documentVersions() *schema.Table {
	tableName := "aws_ssm_document_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentVersionInfo.html`,
		Resolver:    fetchSsmDocumentVersions,
		Transform:   transformers.TransformWithStruct(&types.DocumentVersionInfo{}, transformers.WithPrimaryKeys("DocumentVersion")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "document_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchSsmDocumentVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm
	item := parent.Item.(*types.DocumentDescription)

	params := ssm.ListDocumentVersionsInput{
		Name: item.Name,
	}
	for {
		output, err := svc.ListDocumentVersions(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.DocumentVersions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
