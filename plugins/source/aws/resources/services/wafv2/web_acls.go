package wafv2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafv2/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func WebAcls() *schema.Table {
	tableName := "aws_wafv2_web_acls"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_WebACL.html`,
		Resolver:            fetchWafv2WebAcls,
		PreResourceResolver: getWebAcl,
		Transform:           transformers.TransformWithStruct(&models.WebACLWrapper{}, transformers.WithUnwrapStructFields("WebACL")),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveWebACLTags,
			},
			{
				Name:     "resources_for_web_acl",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: resolveWafv2webACLResourcesForWebACL,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	config := wafv2.ListWebACLsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100),
	}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- output.WebACLs

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func getWebAcl(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	webAcl := resource.Item.(types.WebACLSummary)

	webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: cl.WAFScope}
	webAclOutput, err := svc.GetWebACL(ctx, &webAclConfig, func(o *wafv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	cfg := wafv2.GetLoggingConfigurationInput{
		ResourceArn: webAclOutput.WebACL.ARN,
	}

	loggingConfigurationOutput, err := svc.GetLoggingConfiguration(ctx, &cfg, func(o *wafv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if client.IsAWSError(err, "WAFNonexistentItemException") {
			cl.Logger().Debug().Err(err).Msg("Logging configuration not found for")
		} else {
			cl.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed with error")
		}
	}

	var webAclLoggingConfiguration *types.LoggingConfiguration
	if loggingConfigurationOutput != nil {
		webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
	}

	resource.Item = &models.WebACLWrapper{
		WebACL:               webAclOutput.WebACL,
		LoggingConfiguration: webAclLoggingConfiguration,
	}
	return nil
}

func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*models.WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	resourceArns := []string{}
	if cl.WAFScope == types.ScopeCloudfront {
		cloudfrontService := cl.Services().Cloudfront
		params := &cloudfront.ListDistributionsByWebACLIdInput{
			WebACLId: webACL.Id,
			MaxItems: aws.Int32(100),
		}
		for {
			output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(o *cloudfront.Options) {
				o.Region = cl.Region
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
		var resourceType types.ResourceType
		for _, resType := range resourceType.Values() {
			output, err := service.ListResourcesForWebACL(ctx,
				&wafv2.ListResourcesForWebACLInput{
					WebACLArn:    webACL.ARN,
					ResourceType: resType,
				}, func(o *wafv2.Options) {
					o.Region = cl.Region
				})
			if err != nil {
				return err
			}
			resourceArns = append(resourceArns, output.ResourceArns...)
		}
	}
	return resource.Set(c.Name, resourceArns)
}
func resolveWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*models.WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: webACL.ARN}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
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
