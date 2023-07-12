package config

import (
	"context"
	"errors"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ConformancePackArn"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			conformancePackRuleCompliances(),
		},
	}
}

func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	var ae smithy.APIError
	configService := cl.Services().Configservice
	paginator := configservice.NewDescribeConformancePacksPaginator(configService, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *configservice.Options) {
			options.Region = cl.Region
		})
		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		if (cl.Region == "af-south-1" || cl.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}
		if err != nil {
			return err
		}
		res <- page.ConformancePackDetails
	}
	return nil
}
