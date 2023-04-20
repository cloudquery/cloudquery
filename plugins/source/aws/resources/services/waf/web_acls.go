package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func WebAcls() *schema.Table {
	tableName := "aws_waf_web_acls"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html`,
		Resolver:    fetchWafWebAcls,
		Transform:   transformers.TransformWithStruct(&WebACLWrapper{}, transformers.WithUnwrapStructFields("WebACL")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafWebACLTags,
			},
		},
	}
}

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func fetchWafWebAcls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config)
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := waf.GetWebACLInput{WebACLId: webAcl.WebACLId}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			cfg := waf.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.WebACLArn,
			}
			// TODO: Look into refactoring this as a column resolver
			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				c.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed")
			}

			var webAclLoggingConfiguration *types.LoggingConfiguration
			if loggingConfigurationOutput != nil {
				webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
			}

			res <- &WebACLWrapper{
				webAclOutput.WebACL,
				webAclLoggingConfiguration,
			}
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	// Resolve tags for resource
	awsClient := meta.(*client.Client)
	service := awsClient.Services().Waf
	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: webACL.WebACLArn}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig)
		if err != nil {
			return err
		}
		for _, t := range tags.TagInfoForResource.TagList {
			outputTags[*t.Key] = t.Value
		}
		if aws.ToString(tags.NextMarker) == "" {
			break
		}
		tagsConfig.NextMarker = tags.NextMarker
	}
	return resource.Set("tags", outputTags)
}
