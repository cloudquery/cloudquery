package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Operations() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_operations",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html`,
		Resolver:    fetchApprunnerOperations,
		Transform:   transformers.TransformWithStruct(&types.OperationSummary{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
		},
	}
}
