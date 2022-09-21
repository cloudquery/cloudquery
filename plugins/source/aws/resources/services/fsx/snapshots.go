// Code generated by codegen; DO NOT EDIT.

package fsx

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:      "aws_fsx_snapshots",
		Resolver:  fetchFsxSnapshots,
		Multiplex: client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:          "administrative_actions",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("AdministrativeActions"),
				IgnoreInTests: true,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "lifecycle_transition_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LifecycleTransitionReason"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "snapshot_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotId"),
			},
			{
				Name:     "volume_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VolumeId"),
			},
		},
	}
}
