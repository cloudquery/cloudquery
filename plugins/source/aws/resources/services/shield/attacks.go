// Code generated by codegen; DO NOT EDIT.

package shield

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Attacks() *schema.Table {
	return &schema.Table{
		Name:      "aws_shield_attacks",
		Resolver:  fetchShieldAttacks,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AttackId"),
				Description: `The unique identifier (ID) of the attack`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attack_counters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttackCounters"),
			},
			{
				Name:     "attack_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AttackProperties"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndTime"),
			},
			{
				Name:     "mitigations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Mitigations"),
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceArn"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "sub_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SubResources"),
			},
		},
	}
}
