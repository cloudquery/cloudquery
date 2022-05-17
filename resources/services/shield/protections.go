package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource protections --config gen.hcl --output .
func Protections() *schema.Table {
	return &schema.Table{
		Name:          "aws_shield_protections",
		Description:   "An object that represents a resource that is under DDoS protection.",
		Resolver:      fetchShieldProtections,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
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
				Description: "The AWS tags of the resource.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveShieldProtectionTags,
			},
			{
				Name:        "application_automatic_response_configuration_status",
				Description: "Indicates whether automatic application layer DDoS mitigation is enabled for the protection",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationLayerAutomaticResponseConfiguration.Status"),
			},
			{
				Name:        "health_check_ids",
				Description: "The unique identifier (ID) for the Route 53 health check that's associated with the protection",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the protection",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the protection",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) of the protection",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ProtectionArn"),
			},
			{
				Name:        "resource_arn",
				Description: "The ARN (Amazon Resource Name) of the Amazon Web Services resource that is protected",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchShieldProtections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.ListProtectionsInput{}
	for {
		output, err := svc.ListProtections(ctx, &config, func(o *shield.Options) {
			o.Region = c.Region
		})
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		res <- output.Protections

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func ResolveShieldProtectionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Protection)
	cli := meta.(*client.Client)
	svc := cli.Services().Shield
	config := shield.ListTagsForResourceInput{ResourceARN: r.ProtectionArn}

	output, err := svc.ListTagsForResource(ctx, &config, func(o *shield.Options) {
		o.Region = cli.Region
	})
	if err != nil {
		if cli.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}

	tags := map[string]string{}
	client.TagsIntoMap(output.Tags, tags)

	return resource.Set(c.Name, tags)
}
