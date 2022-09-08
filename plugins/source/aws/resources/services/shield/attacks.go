package shield

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)


func Attacks() *schema.Table {
	return &schema.Table{
		Name:          "aws_shield_attacks",
		Description:   "The details of a DDoS attack",
		Resolver:      fetchShieldAttacks,
		Multiplex:     client.AccountMultiplex,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "attack_counters",
				Description: "List of counters that describe the attack for the specified time period",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AttackCounters"),
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the attack",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AttackId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "end_time",
				Description: "The time the attack ended, in Unix time in seconds",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "mitigations",
				Description: "List of mitigation actions taken for the attack",
				Type:        schema.TypeStringArray,
				Resolver:    resolveAttacksMitigations,
			},
			{
				Name:        "resource_arn",
				Description: "The ARN (Amazon Resource Name) of the resource that was attacked",
				Type:        schema.TypeString,
			},
			{
				Name:        "start_time",
				Description: "The time the attack started, in Unix time in seconds",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "attack_properties",
				Description: "Details of a Shield event",
				Type:        schema.TypeJSON,
			},
			{
				Name: "sub_resources",
				Description: "The attack information for the specified SubResource",
				Type: 			schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchShieldAttacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	end := time.Now()
	start := end.Add(-time.Hour * 24)
	config := shield.ListAttacksInput{
		EndTime:   &types.TimeRange{ToExclusive: &end},
		StartTime: &types.TimeRange{FromInclusive: &start},
	}
	for {
		output, err := svc.ListAttacks(ctx, &config)
		if err != nil {
			return err
		}
		for _, a := range output.AttackSummaries {
			config := shield.DescribeAttackInput{AttackId: a.AttackId}
			attack, err := svc.DescribeAttack(ctx, &config, func(o *shield.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- attack.Attack
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveAttacksMitigations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.AttackDetail)
	mitigations := make([]string, 0, len(r.Mitigations))
	for _, m := range r.Mitigations {
		mitigations = append(mitigations, *m.MitigationName)
	}
	return diag.WrapError(resource.Set(c.Name, mitigations))
}

func resolveAttackPropertiesTopContributors(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.AttackProperty)
	marshalledJson := make(map[string]interface{})
	for _, c := range r.TopContributors {
		marshalledJson[*c.Name] = c.Value
	}
	return diag.WrapError(resource.Set(c.Name, marshalledJson))
}
