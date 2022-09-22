// Code generated by codegen; DO NOT EDIT.

package emr

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:      "aws_emr_clusters",
		Resolver:  fetchEmrClusters,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
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
				Resolver: schema.PathResolver("ClusterArn"),
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
				Name:     "applications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Applications"),
			},
			{
				Name:     "auto_scaling_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingRole"),
			},
			{
				Name:     "auto_terminate",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoTerminate"),
			},
			{
				Name:     "configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configurations"),
			},
			{
				Name:     "custom_ami_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomAmiId"),
			},
			{
				Name:     "ebs_root_volume_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EbsRootVolumeSize"),
			},
			{
				Name:     "ec2_instance_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ec2InstanceAttributes"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "instance_collection_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceCollectionType"),
			},
			{
				Name:     "kerberos_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KerberosAttributes"),
			},
			{
				Name:     "log_encryption_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogEncryptionKmsKeyId"),
			},
			{
				Name:     "log_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogUri"),
			},
			{
				Name:     "master_public_dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterPublicDnsName"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "normalized_instance_hours",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NormalizedInstanceHours"),
			},
			{
				Name:     "os_release_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OSReleaseLabel"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "placement_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementGroups"),
			},
			{
				Name:     "release_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReleaseLabel"),
			},
			{
				Name:     "repo_upgrade_on_boot",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepoUpgradeOnBoot"),
			},
			{
				Name:     "requested_ami_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequestedAmiVersion"),
			},
			{
				Name:     "running_ami_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RunningAmiVersion"),
			},
			{
				Name:     "scale_down_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScaleDownBehavior"),
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityConfiguration"),
			},
			{
				Name:     "service_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceRole"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "step_concurrency_level",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StepConcurrencyLevel"),
			},
			{
				Name:     "termination_protected",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TerminationProtected"),
			},
			{
				Name:     "visible_to_all_users",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VisibleToAllUsers"),
			},
		},
	}
}
