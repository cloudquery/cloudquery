package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EmrClusters() *schema.Table {
	return &schema.Table{
		Name:          "aws_emr_clusters",
		Description:   "The detailed description of the cluster.",
		Resolver:      fetchEmrClusters,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
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
				Name:        "applications",
				Description: "The applications installed on this cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Applications"),
			},
			{
				Name:        "auto_scaling_role",
				Description: "An IAM role for automatic scaling policies",
				Type:        schema.TypeString,
			},
			{
				Name:        "auto_terminate",
				Description: "Specifies whether the cluster should terminate after completing all steps.",
				Type:        schema.TypeBool,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name of the cluster.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "configurations",
				Description: "The list of Configurations supplied to the EMR cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Configurations"),
			},
			{
				Name:          "custom_ami_id",
				Description:   "The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "ebs_root_volume_size",
				Description: "The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance",
				Type:        schema.TypeInt,
			},
			{
				Name:     "ec2_instance_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ec2InstanceAttributes"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_collection_type",
				Description: "The instance group configuration of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:          "kerberos_attributes",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("KerberosAttributes"),
				IgnoreInTests: true,
			},
			{
				Name:          "log_encryption_kms_key_id",
				Description:   "The AWS KMS customer master key (CMK) used for encrypting log files",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "log_uri",
				Description: "The path to the Amazon S3 location where logs for this cluster are stored.",
				Type:        schema.TypeString,
			},
			{
				Name:        "master_public_dns_name",
				Description: "The DNS name of the master node",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "normalized_instance_hours",
				Description: "An approximation of the cost of the cluster, represented in m1.small/hours",
				Type:        schema.TypeInt,
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "placement_groups",
				Description: "Placement group configured for an Amazon EMR cluster.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("PlacementGroups"),
			},
			{
				Name:        "release_label",
				Description: "The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster",
				Type:        schema.TypeString,
			},
			{
				Name:        "repo_upgrade_on_boot",
				Description: "Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:          "requested_ami_version",
				Description:   "The AMI version requested for this cluster.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "running_ami_version",
				Description:   "The AMI version running on this cluster.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "scale_down_behavior",
				Description: "The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.",
				Type:        schema.TypeString,
			},
			{
				Name:          "security_configuration",
				Description:   "The name of the security configuration applied to the cluster.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "service_role",
				Description: "The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.",
				Type:        schema.TypeString,
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:        "step_concurrency_level",
				Description: "Specifies the number of steps that can be executed concurrently.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "tags",
				Description: "A list of tags associated with a cluster.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "termination_protected",
				Description: "Indicates whether Amazon EMR will lock the cluster to prevent the EC2 instances from being terminated by an API call or user intervention, or in the event of a cluster error.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "visible_to_all_users",
				Description: "Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config emr.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().EMR
	for {
		response, err := svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		for _, c := range response.Clusters {
			out, err := svc.DescribeCluster(ctx, &emr.DescribeClusterInput{ClusterId: c.Id})
			if err != nil {
				return err
			}
			res <- out.Cluster
		}
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
