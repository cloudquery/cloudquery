package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource sampling_rules --config gen.hcl --output .
func SamplingRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_xray_sampling_rules",
		Description:  "A SamplingRule (https://docsawsamazoncom/xray/latest/api/API_SamplingRulehtml) and its metadata",
		Resolver:     fetchXraySamplingRules,
		Multiplex:    client.ServiceAccountRegionMultiplexer("xray"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXraySamplingRuleTags,
			},
			{
				Name:        "created_at",
				Description: "When the rule was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_at",
				Description: "When the rule was last modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "fixed_rate",
				Description: "The percentage of matching requests to instrument, after the reservoir is exhausted",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("SamplingRule.FixedRate"),
			},
			{
				Name:        "http_method",
				Description: "Matches the HTTP method of a request",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.HTTPMethod"),
			},
			{
				Name:        "host",
				Description: "Matches the hostname from a request URL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.Host"),
			},
			{
				Name:        "priority",
				Description: "The priority of the sampling rule",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SamplingRule.Priority"),
			},
			{
				Name:        "reservoir_size",
				Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SamplingRule.ReservoirSize"),
			},
			{
				Name:        "resource_arn",
				Description: "Matches the ARN of the Amazon Web Services resource on which the service runs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ResourceARN"),
			},
			{
				Name:        "service_name",
				Description: "Matches the name that the service uses to identify itself in segments",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ServiceName"),
			},
			{
				Name:        "service_type",
				Description: "Matches the origin that the service uses to identify its type in segments",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ServiceType"),
			},
			{
				Name:        "url_path",
				Description: "Matches the path from a request URL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.URLPath"),
			},
			{
				Name:        "version",
				Description: "The version of the sampling rule format (1)",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SamplingRule.Version"),
			},
			{
				Name:        "attributes",
				Description: "Matches attributes derived from the request",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SamplingRule.Attributes"),
			},
			{
				Name:        "arn",
				Description: "The ARN of the sampling rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.RuleARN"),
			},
			{
				Name:        "rule_name",
				Description: "The name of the sampling rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.RuleName"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchXraySamplingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := xray.NewGetSamplingRulesPaginator(meta.(*client.Client).Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.SamplingRuleRecords
	}
	return nil
}
func resolveXraySamplingRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sr := resource.Item.(types.SamplingRuleRecord)
	cl := meta.(*client.Client)
	svc := cl.Services().Xray
	params := xray.ListTagsForResourceInput{ResourceARN: sr.SamplingRule.RuleARN}

	output, err := svc.ListTagsForResource(ctx, &params)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}

	tags := map[string]string{}
	client.TagsIntoMap(output.Tags, tags)

	return diag.WrapError(resource.Set(c.Name, tags))
}
