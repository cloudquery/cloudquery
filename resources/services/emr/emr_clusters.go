package emr

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EmrClusters() *schema.Table {
	return &schema.Table{
		Name:         "aws_emr_clusters",
		Description:  "The detailed description of the cluster.",
		Resolver:     fetchEmrClusters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
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
				Name:        "applications",
				Description: "The applications installed on this cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEMRClusterJSONField(func(c *types.Cluster) interface{} { return c.Applications }),
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
				Name:        "arn",
				Description: "The Amazon Resource Name of the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClusterArn"),
			},
			{
				Name:        "configurations",
				Description: "The list of Configurations supplied to the EMR cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEMRClusterJSONField(func(c *types.Cluster) interface{} { return c.Configurations }),
			},
			{
				Name:        "custom_ami_id",
				Description: "The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ebs_root_volume_size",
				Description: "The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance",
				Type:        schema.TypeInt,
			},
			{
				Name:        "ec2_instance_attribute_additional_master_security_groups",
				Description: "A list of additional Amazon EC2 security group IDs for the master node.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.AdditionalMasterSecurityGroups"),
			},
			{
				Name:        "ec2_instance_attribute_additional_slave_security_groups",
				Description: "A list of additional Amazon EC2 security group IDs for the core and task nodes.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.AdditionalSlaveSecurityGroups"),
			},
			{
				Name:        "ec2_instance_attribute_availability_zone",
				Description: "The Availability Zone in which the cluster will run.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.Ec2AvailabilityZone"),
			},
			{
				Name:        "ec2_instance_attribute_key_name",
				Description: "The name of the Amazon EC2 key pair to use when connecting with SSH into the master node as a user named \"hadoop\".",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.Ec2KeyName"),
			},
			{
				Name:        "ec2_instance_attribute_subnet_id",
				Description: "Set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.Ec2SubnetId"),
			},
			{
				Name:        "ec2_instance_attribute_emr_managed_master_security_group",
				Description: "The identifier of the Amazon EC2 security group for the master node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.EmrManagedMasterSecurityGroup"),
			},
			{
				Name:        "ec2_instance_attribute_emr_managed_slave_security_group",
				Description: "The identifier of the Amazon EC2 security group for the core and task nodes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.EmrManagedSlaveSecurityGroup"),
			},
			{
				Name:        "ec2_instance_attribute_iam_instance_profile",
				Description: "The IAM role that was specified when the cluster was launched",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.IamInstanceProfile"),
			},
			{
				Name:        "ec2_instance_attribute_requested_availability_zones",
				Description: "Specifies one or more Availability Zones in which to launch EC2 cluster instances when the EC2-Classic network configuration is supported.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.RequestedEc2AvailabilityZones"),
			},
			{
				Name:        "ec2_instance_attribute_requested_subnet_ids",
				Description: "Specifies the unique identifier of one or more Amazon EC2 subnets in which to launch EC2 cluster instances.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.RequestedEc2SubnetIds"),
			},
			{
				Name:        "ec2_instance_attribute_service_access_security_group",
				Description: "The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Ec2InstanceAttributes.ServiceAccessSecurityGroup"),
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
				Name:        "kerberos_kdc_admin_password",
				Description: "The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KerberosAttributes.KdcAdminPassword"),
			},
			{
				Name:        "kerberos_realm",
				Description: "The name of the Kerberos realm to which all nodes in a cluster belong",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KerberosAttributes.Realm"),
			},
			{
				Name:        "kerberos_ad_domain_join_password",
				Description: "The Active Directory password for ADDomainJoinUser.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KerberosAttributes.ADDomainJoinPassword"),
			},
			{
				Name:        "kerberos_ad_domain_join_user",
				Description: "Required only when establishing a cross-realm trust with an Active Directory domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KerberosAttributes.ADDomainJoinUser"),
			},
			{
				Name:        "kerberos_cross_realm_trust_principal_password",
				Description: "Required only when establishing a cross-realm trust with a KDC in a different realm",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("KerberosAttributes.CrossRealmTrustPrincipalPassword"),
			},
			{
				Name:        "log_encryption_kms_key_id",
				Description: "The AWS KMS customer master key (CMK) used for encrypting log files",
				Type:        schema.TypeString,
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
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.",
				Type:        schema.TypeString,
			},
			{
				Name:        "placement_groups",
				Description: "Placement group configured for an Amazon EMR cluster.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEMRClusterJSONField(func(c *types.Cluster) interface{} { return c.PlacementGroups }),
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
				Name:        "requested_ami_version",
				Description: "The AMI version requested for this cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "running_ami_version",
				Description: "The AMI version running on this cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "scale_down_behavior",
				Description: "The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_configuration",
				Description: "The name of the security configuration applied to the cluster.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_role",
				Description: "The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the cluster.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.State"),
			},
			{
				Name:        "state_change_reason_code",
				Description: "The programmatic code for the state change reason.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.StateChangeReason.Code"),
			},
			{
				Name:        "state_change_reason_message",
				Description: "The descriptive message for the state change reason.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.StateChangeReason.Message"),
			},
			{
				Name:        "creation_date_time",
				Description: "The creation date and time of the cluster.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.CreationDateTime"),
			},
			{
				Name:        "end_date_time",
				Description: "The date and time when the cluster was terminated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.EndDateTime"),
			},
			{
				Name:        "ready_date_time",
				Description: "The date and time when the cluster was ready to run steps.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.Timeline.ReadyDateTime"),
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
				Resolver: resolveEMRClusterJSONField(func(c *types.Cluster) interface{} {
					tags := make(map[string]string, len(c.Tags))
					for _, t := range c.Tags {
						tags[aws.ToString(t.Key)] = aws.ToString(t.Value)
					}
					return tags
				}),
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
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEmrClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config emr.ListClustersInput
	c := meta.(*client.Client)
	svc := c.Services().EMR
	for {
		response, err := svc.ListClusters(ctx, &config, func(options *emr.Options) {
			options.Region = c.Region
		})
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

func resolveEMRClusterJSONField(getter func(c *types.Cluster) interface{}) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		cl, ok := resource.Item.(*types.Cluster)
		if !ok {
			return fmt.Errorf("not a %T instance: %T", c, resource.Item)
		}
		b, err := json.Marshal(getter(cl))
		if err != nil {
			return err
		}
		return resource.Set(c.Name, b)
	}
}
