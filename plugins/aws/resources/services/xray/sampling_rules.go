package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource sampling_rules --config gen.hcl --output .
func SamplingRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_xray_sampling_rules",
		Description:  "A SamplingRule.",
		Resolver:     fetchXraySamplingRules,
		Multiplex:    client.ServiceAccountRegionMultiplexer("xray"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "tags",
				Description: "A list of Tags that specify information about the sampling rule.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveXraySamplingRuleTags,
			},
			{
				Name:        "created_at",
				Description: "When the rule was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified_at",
				Description: "When the rule was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "fixed_rate",
				Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("SamplingRule.FixedRate"),
			},
			{
				Name:        "http_method",
				Description: "Matches the HTTP method of a request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.HTTPMethod"),
			},
			{
				Name:        "host",
				Description: "Matches the hostname from a request URL.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.Host"),
			},
			{
				Name:        "priority",
				Description: "The priority of the sampling rule.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SamplingRule.Priority"),
			},
			{
				Name:        "reservoir_size",
				Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SamplingRule.ReservoirSize"),
			},
			{
				Name:        "resource_arn",
				Description: "Matches the ARN of the Amazon Web Services resource on which the service runs.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ResourceARN"),
			},
			{
				Name:        "service_name",
				Description: "Matches the name that the service uses to identify itself in segments.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ServiceName"),
			},
			{
				Name:        "service_type",
				Description: "Matches the origin that the service uses to identify its type in segments.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.ServiceType"),
			},
			{
				Name:        "url_path",
				Description: "Matches the path from a request URL.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamplingRule.URLPath"),
			},
			{
				Name:        "version",
				Description: "The version of the sampling rule format (1).",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SamplingRule.Version"),
			},
			{
				Name:        "attributes",
				Description: "Matches attributes derived from the request.",
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
	c := meta.(*client.Client)
	svc := c.Services().Xray
	input := xray.GetSamplingRulesInput{}
	for {
		output, err := svc.GetSamplingRules(ctx, &input, func(o *xray.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.SamplingRuleRecords

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func ResolveXraySamplingRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sr := resource.Item.(types.SamplingRuleRecord)
	cl := meta.(*client.Client)
	svc := cl.Services().Xray
	params := xray.ListTagsForResourceInput{ResourceARN: sr.SamplingRule.RuleARN}

	output, err := svc.ListTagsForResource(ctx, &params, func(o *xray.Options) {
		o.Region = cl.Region
	})
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
