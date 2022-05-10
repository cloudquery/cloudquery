package shield

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource attacks --config gen.hcl --output .
func Attacks() *schema.Table {
	return &schema.Table{
		Name:          "aws_shield_attacks",
		Description:   "The details of a DDoS attack",
		Resolver:      fetchShieldAttacks,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
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
				Resolver:    resolveAttacksAttackCounters,
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the attack",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AttackId"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_shield_attack_properties",
				Description: "Details of a Shield event",
				Resolver:    fetchShieldAttackProperties,
				Columns: []schema.Column{
					{
						Name:        "attack_cq_id",
						Description: "Unique CloudQuery ID of aws_shield_attacks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "attack_layer",
						Description: "The type of Shield event that was observed",
						Type:        schema.TypeString,
					},
					{
						Name:        "attack_property_identifier",
						Description: "Defines the Shield event property information that is provided",
						Type:        schema.TypeString,
					},
					{
						Name:        "top_contributors",
						Description: "Contributor objects for the top five contributors to a Shield event",
						Type:        schema.TypeJSON,
						Resolver:    resolveAttackPropertiesTopContributors,
					},
					{
						Name:        "total",
						Description: "The total contributions made to this Shield event by all contributors",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "unit",
						Description: "The unit used for the ContributorValue property",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_shield_attack_sub_resources",
				Description: "The attack information for the specified SubResource",
				Resolver:    fetchShieldAttackSubResources,
				Columns: []schema.Column{
					{
						Name:        "attack_cq_id",
						Description: "Unique CloudQuery ID of aws_shield_attacks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "attack_vectors",
						Description: "The list of attack types and associated counters",
						Type:        schema.TypeJSON,
						Resolver:    resolveAttackSubResourcesAttackVectors,
					},
					{
						Name:        "counters",
						Description: "The counters that describe the details of the attack",
						Type:        schema.TypeJSON,
						Resolver:    resolveAttackSubResourcesCounters,
					},
					{
						Name:        "id",
						Description: "The unique identifier (ID) of the SubResource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The SubResource type",
						Type:        schema.TypeString,
					},
				},
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
		output, err := svc.ListAttacks(ctx, &config, func(o *shield.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, a := range output.AttackSummaries {
			config := shield.DescribeAttackInput{AttackId: a.AttackId}
			attack, err := svc.DescribeAttack(ctx, &config, func(o *shield.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
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
func resolveAttacksAttackCounters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.AttackDetail)
	json, err := json.Marshal(r.AttackCounters)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, json)
}
func resolveAttacksMitigations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.AttackDetail)
	mitigations := make([]string, 0, len(r.Mitigations))
	for _, m := range r.Mitigations {
		mitigations = append(mitigations, *m.MitigationName)
	}
	return resource.Set(c.Name, mitigations)
}
func fetchShieldAttackProperties(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*types.AttackDetail)
	res <- r.AttackProperties
	return nil
}
func resolveAttackPropertiesTopContributors(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.AttackProperty)
	json := make(map[string]interface{})
	for _, c := range r.TopContributors {
		json[*c.Name] = c.Value
	}
	return resource.Set(c.Name, json)
}
func fetchShieldAttackSubResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*types.AttackDetail)
	res <- r.SubResources
	return nil
}
func resolveAttackSubResourcesAttackVectors(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.SubResourceSummary)
	json, err := json.Marshal(r.AttackVectors)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, json)
}

func resolveAttackSubResourcesCounters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.SubResourceSummary)
	json, err := json.Marshal(r.Counters)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, json)
}
