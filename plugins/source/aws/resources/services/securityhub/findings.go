package securityhub

import (
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	tableName := "aws_securityhub_findings"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html.
` + "The `request_account_id` and `request_region` columns are added to show the account and region of where the request was made from." + `
This is useful when multi region and account aggregation is enabled.`,
		Resolver: fetchFindings,
		Transform: client.TransformWithStruct(&types.AwsSecurityFinding{},
			transformers.WithPrimaryKeys("AwsAccountId", "Region", "CreatedAt", "Description", "GeneratorId", "Id", "ProductArn", "SchemaVersion", "Title"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: []schema.Column{
			{
				Name:            "request_account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "request_region",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
