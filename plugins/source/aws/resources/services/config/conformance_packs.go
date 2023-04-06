package config

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConformancePacks() *schema.Table {
	tableName := "aws_config_conformance_packs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_ConformancePackDetail.html`,
		Resolver:    fetchConfigConformancePacks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&types.ConformancePackDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConformancePackArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			conformancePackRuleCompliances(),
		},
	}
}

func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	var ae smithy.APIError
	for {
		resp, err := c.Services().Configservice.DescribeConformancePacks(ctx, &config)

		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		if (c.Region == "af-south-1" || c.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}

		if err != nil {
			return err
		}
		res <- resp.ConformancePackDetails
		if aws.ToString(resp.NextToken) == "" {
			break
		}
		config.NextToken = resp.NextToken
	}
	return nil
}
