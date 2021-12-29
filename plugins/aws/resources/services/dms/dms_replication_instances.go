package dms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cq-provider-aws/client"
)

func DmsReplicationInstances() *schema.Table {
	return &schema.Table{
		Name:         "aws_dms_replication_instances",
		Description:  "Provides information that defines a replication instance.",
		Resolver:     fetchDmsReplicationInstances,
		Multiplex:    client.ServiceAccountRegionMultiplexer("dms"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "arn"}},
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
				Description: "Any tags assigned to the resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "allocated_storage",
				Description: "The amount of storage (in gigabytes) that is allocated for the replication instance.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "auto_minor_version_upgrade",
				Description: "Boolean value indicating if minor version upgrades will be automatically applied to the instance.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone for the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_name_servers",
				Description: "The DNS name servers supported for the replication instance to access your on-premise source or target database.",
				Type:        schema.TypeString,
			},
			{
				Name:        "engine_version",
				Description: "The engine version number of the replication instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "free_until",
				Description: "The expiration date of the free replication instance that is part of the Free DMS program.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "instance_create_time",
				Description: "The time the replication instance was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "kms_key_id",
				Description: "An KMS key identifier that is used to encrypt the data on the replication instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "multi_az",
				Description: "Specifies whether the replication instance is a Multi-AZ deployment",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("MultiAZ"),
			},
			{
				Name:        "pending_modified_values_allocated_storage",
				Description: "The amount of storage (in gigabytes) that is allocated for the replication instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("PendingModifiedValues.AllocatedStorage"),
			},
			{
				Name:        "pending_modified_values_engine_version",
				Description: "The engine version number of the replication instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.EngineVersion"),
			},
			{
				Name:        "pending_modified_values_multi_az",
				Description: "Specifies whether the replication instance is a Multi-AZ deployment",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("PendingModifiedValues.MultiAZ"),
			},
			{
				Name:        "pending_modified_values_class",
				Description: "The compute and memory capacity of the replication instance as defined for the specified replication instance class",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PendingModifiedValues.ReplicationInstanceClass"),
			},
			{
				Name:        "preferred_maintenance_window",
				Description: "The maintenance window times for the replication instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "publicly_accessible",
				Description: "Specifies the accessibility options for the replication instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the replication instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationInstanceArn"),
			},
			{
				Name:        "class",
				Description: "The compute and memory capacity of the replication instance as defined for the specified replication instance class",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationInstanceClass"),
			},
			{
				Name:        "identifier",
				Description: "The replication instance identifier is a required parameter",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationInstanceIdentifier"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the replication instance.  Deprecated: This member has been deprecated.",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("ReplicationInstancePrivateIpAddress"),
			},
			{
				Name:        "private_ip_addresses",
				Description: "One or more private IP addresses for the replication instance.",
				Type:        schema.TypeInetArray,
				Resolver:    schema.IPAddressesResolver("ReplicationInstancePrivateIpAddresses"),
			},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the replication instance.  Deprecated: This member has been deprecated.",
				Type:        schema.TypeInet,
				Resolver:    schema.IPAddressResolver("ReplicationInstancePublicIpAddress"),
			},
			{
				Name:        "public_ip_addresses",
				Description: "One or more public IP addresses for the replication instance.",
				Type:        schema.TypeInetArray,
				Resolver:    schema.IPAddressesResolver("ReplicationInstancePublicIpAddresses"),
			},
			{
				Name:        "status",
				Description: "The status of the replication instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationInstanceStatus"),
			},
			{
				Name:        "replication_subnet_group_description",
				Description: "A description for the replication subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationSubnetGroup.ReplicationSubnetGroupDescription"),
			},
			{
				Name:        "replication_subnet_group_identifier",
				Description: "The identifier of the replication instance subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationSubnetGroup.ReplicationSubnetGroupIdentifier"),
			},
			{
				Name:        "replication_subnet_group_subnet_group_status",
				Description: "The status of the subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationSubnetGroup.SubnetGroupStatus"),
			},
			{
				Name:        "replication_subnet_group_vpc_id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReplicationSubnetGroup.VpcId"),
			},
			{
				Name:        "secondary_availability_zone",
				Description: "The Availability Zone of the standby replication instance in a Multi-AZ deployment.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_dms_replication_instance_replication_subnet_group_subnets",
				Description: "In response to a request by the DescribeReplicationSubnetGroups operation, this object identifies a subnet by its given Availability Zone, subnet identifier, and status.",
				Resolver:    fetchDmsReplicationInstanceReplicationSubnetGroupSubnets,
				Columns: []schema.Column{
					{
						Name:        "replication_instance_cq_id",
						Description: "Unique CloudQuery ID of aws_dms_replication_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "subnet_availability_zone_name",
						Description: "The name of the Availability Zone.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:        "subnet_identifier",
						Description: "The subnet identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_status",
						Description: "The status of the subnet.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_dms_replication_instance_vpc_security_groups",
				Description: "Describes the status of a security group associated with the virtual private cloud (VPC) hosting your replication and DB instances.",
				Resolver:    fetchDmsReplicationInstanceVpcSecurityGroups,
				Columns: []schema.Column{
					{
						Name:        "replication_instance_cq_id",
						Description: "Unique CloudQuery ID of aws_dms_replication_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "status",
						Description: "The status of the VPC security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_security_group_id",
						Description: "The VPC security group ID.",
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
func fetchDmsReplicationInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DMS

	var describeReplicationInstancesInput *databasemigrationservice.DescribeReplicationInstancesInput
	describeReplicationInstancesOutput, err := svc.DescribeReplicationInstances(ctx, describeReplicationInstancesInput, func(options *databasemigrationservice.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	if len(describeReplicationInstancesOutput.ReplicationInstances) == 0 {
		return nil
	}

	listTagsForResourceInput := databasemigrationservice.ListTagsForResourceInput{}
	for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
		listTagsForResourceInput.ResourceArnList = append(listTagsForResourceInput.ResourceArnList, *replicationInstance.ReplicationInstanceArn)
	}
	var listTagsForResourceOutput *databasemigrationservice.ListTagsForResourceOutput
	listTagsForResourceOutput, err = svc.ListTagsForResource(ctx, &listTagsForResourceInput, func(options *databasemigrationservice.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	replicationInstanceTags := make(map[string]map[string]interface{})
	for _, tag := range listTagsForResourceOutput.TagList {
		if replicationInstanceTags[*tag.ResourceArn] == nil {
			replicationInstanceTags[*tag.ResourceArn] = make(map[string]interface{})
		}
		replicationInstanceTags[*tag.ResourceArn][*tag.Key] = *tag.Value
	}

	for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
		wrapper := DmsReplicationInstanceWrapper{
			ReplicationInstance: replicationInstance,
			Tags:                replicationInstanceTags[*replicationInstance.ReplicationInstanceArn],
		}
		res <- wrapper
	}
	return nil
}

func fetchDmsReplicationInstanceReplicationSubnetGroupSubnets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	replicationInstance, ok := parent.Item.(DmsReplicationInstanceWrapper)
	if !ok {
		return fmt.Errorf("not dms replication instance")
	}
	res <- replicationInstance.ReplicationSubnetGroup.Subnets
	return nil
}

func fetchDmsReplicationInstanceVpcSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	replicationInstance, ok := parent.Item.(DmsReplicationInstanceWrapper)
	if !ok {
		return fmt.Errorf("not dms replication instance")
	}
	res <- replicationInstance.VpcSecurityGroups
	return nil
}

type DmsReplicationInstanceWrapper struct {
	types.ReplicationInstance
	Tags map[string]interface{}
}
