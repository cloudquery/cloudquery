package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53ReusableDelegationSets() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_reusable_delegation_sets",
		Resolver:     fetchRoute53DelegationSets,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "name_servers",
				Type: schema.TypeStringArray,
			},
			{
				Name: "caller_reference",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53DelegationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config route53.ListReusableDelegationSetsInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListReusableDelegationSets(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DelegationSets
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
