package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2

	config := wafv2.ListWebACLsInput{
		Scope: c.WAFScope,
		Limit: aws.Int32(100),
	}
	for {
		output, err := service.ListWebACLs(ctx, &config)
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: c.WAFScope}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			cfg := wafv2.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.ARN,
			}

			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if client.IsAWSError(err, "WAFNonexistentItemException") {
					c.Logger().Debug().Err(err).Msg("Logging configuration not found for")
				} else {
					c.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed with error")
				}
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
func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	resourceArns := []string{}
	if cl.WAFScope == types.ScopeCloudfront {
		cloudfrontService := cl.Services().Cloudfront
		params := &cloudfront.ListDistributionsByWebACLIdInput{
			WebACLId: webACL.Id,
			MaxItems: aws.Int32(100),
		}
		for {
			output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(options *cloudfront.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for _, item := range output.DistributionList.Items {
				resourceArns = append(resourceArns, *item.ARN)
			}
			if aws.ToString(output.DistributionList.NextMarker) == "" {
				break
			}
			params.Marker = output.DistributionList.NextMarker
		}
	} else {
		output, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN})
		if err != nil {
			return err
		}
		resourceArns = output.ResourceArns
	}
	return resource.Set(c.Name, resourceArns)
}
func resolveWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: webACL.ARN}
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
	return resource.Set(c.Name, outputTags)
}
