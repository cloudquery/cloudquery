package rds

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Cluster struct {
	_                                interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                               uint        `gorm:"primarykey"`
	AccountID                        string
	Region                           string
	ActivityStreamKinesisStreamName  *string
	ActivityStreamKmsKeyId           *string
	ActivityStreamMode               *string
	ActivityStreamStatus             *string
	AllocatedStorage                 *int32
	AssociatedRoles                  []*ClusterRole `gorm:"constraint:OnDelete:CASCADE;"`
	AvailabilityZones                *string
	BacktrackConsumedChangeRecords   *int64
	BacktrackWindow                  *int64
	BackupRetentionPeriod            *int32
	Capacity                         *int32
	CharacterSetName                 *string
	CloneGroupId                     *string
	ClusterCreateTime                *time.Time
	CopyTagsToSnapshot               *bool
	CrossAccountClone                *bool
	CustomEndpoints                  *string
	ClusterArn                       *string `neo:"unique"`
	ClusterIdentifier                *string
	ClusterMembers                   []*ClusterMember            `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterOptionGroupMemberships    []*ClusterOptionGroupStatus `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterParameterGroup            *string
	SubnetGroup                      *string
	DatabaseName                     *string
	DbClusterResourceId              *string
	DeletionProtection               *bool
	DomainMemberships                []*ClusterDomainMembership `gorm:"constraint:OnDelete:CASCADE;"`
	EarliestBacktrackTime            *time.Time
	EarliestRestorableTime           *time.Time
	EnabledCloudwatchLogsExports     *string
	Endpoint                         *string
	Engine                           *string
	EngineMode                       *string
	EngineVersion                    *string
	GlobalWriteForwardingRequested   *bool
	GlobalWriteForwardingStatus      *string
	HostedZoneId                     *string
	HttpEndpointEnabled              *bool
	IAMDatabaseAuthenticationEnabled *bool
	KmsKeyId                         *string
	LatestRestorableTime             *time.Time
	MasterUsername                   *string
	MultiAZ                          *bool
	PercentProgress                  *string
	Port                             *int32
	PreferredBackupWindow            *string
	PreferredMaintenanceWindow       *string
	ReadReplicaIdentifiers           *string
	ReaderEndpoint                   *string
	ReplicationSourceIdentifier      *string

	ScalingConfigAutoPause             *bool
	ScalingConfigMaxCapacity           *int32
	ScalingConfigMinCapacity           *int32
	ScalingConfigSecondsUntilAutoPause *int32
	ScalingConfigTimeoutAction         *string

	Status            *string
	StorageEncrypted  *bool
	VpcSecurityGroups []*ClusterVpcSecurityGroupMembership `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_rds_clusters"
}

type ClusterRole struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	FeatureName *string
	RoleArn     *string
	Status      *string
}

func (ClusterRole) TableName() string {
	return "aws_rds_cluster_roles"
}

type ClusterMember struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ClusterParameterGroupStatus *string
	InstanceIdentifier          *string
	IsClusterWriter             *bool
	PromotionTier               *int32
}

func (ClusterMember) TableName() string {
	return "aws_rds_cluster_members"
}

type ClusterOptionGroupStatus struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ClusterOptionGroupName *string
	Status                 *string
}

func (ClusterOptionGroupStatus) TableName() string {
	return "aws_rds_cluster_option_group_statuses"
}

type ClusterDomainMembership struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Domain      *string
	FQDN        *string
	IAMRoleName *string
	Status      *string
}

func (ClusterDomainMembership) TableName() string {
	return "aws_rds_cluster_domain_membership"
}

type ClusterVpcSecurityGroupMembership struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Status             *string
	VpcSecurityGroupId *string
}

func (ClusterVpcSecurityGroupMembership) TableName() string {
	return "aws_rds_cluster_vpc_security_group_memberships"
}


func (c *Client) transformClusterRoles(values *[]types.DBClusterRole) []*ClusterRole {
	var tValues []*ClusterRole
	for _, value := range *values {
		tValues = append(tValues, &ClusterRole{
			AccountID:   c.accountID,
			Region:      c.region,
			FeatureName: value.FeatureName,
			RoleArn:     value.RoleArn,
			Status:      value.Status,
		})
	}
	return tValues
}

func (c *Client) transformClusterMembers(values *[]types.DBClusterMember) []*ClusterMember {
	var tValues []*ClusterMember
	for _, value := range *values {
		tValues = append(tValues, &ClusterMember{
			AccountID:                   c.accountID,
			Region:                      c.region,
			ClusterParameterGroupStatus: value.DBClusterParameterGroupStatus,
			InstanceIdentifier:          value.DBInstanceIdentifier,
			IsClusterWriter:             &value.IsClusterWriter,
			PromotionTier:               value.PromotionTier,
		})
	}
	return tValues
}

func (c *Client) transformClusterOptionGroupStatuss(values *[]types.DBClusterOptionGroupStatus) []*ClusterOptionGroupStatus {
	var tValues []*ClusterOptionGroupStatus
	for _, value := range *values {
		tValues = append(tValues, &ClusterOptionGroupStatus{
			AccountID:              c.accountID,
			Region:                 c.region,
			ClusterOptionGroupName: value.DBClusterOptionGroupName,
			Status:                 value.Status,
		})
	}
	return tValues
}

func (c *Client) transformClusterDomainMemberships(values *[]types.DomainMembership) []*ClusterDomainMembership {
	var tValues []*ClusterDomainMembership
	for _, value := range *values {
		tValues = append(tValues, &ClusterDomainMembership{
			AccountID:   c.accountID,
			Region:      c.region,
			Domain:      value.Domain,
			FQDN:        value.FQDN,
			IAMRoleName: value.IAMRoleName,
			Status:      value.Status,
		})
	}
	return tValues
}

func (c *Client) transformClusterVpcSecurityGroupMemberships(values *[]types.VpcSecurityGroupMembership) []*ClusterVpcSecurityGroupMembership {
	var tValues []*ClusterVpcSecurityGroupMembership
	for _, value := range *values {
		tValues = append(tValues, &ClusterVpcSecurityGroupMembership{
			AccountID:          c.accountID,
			Region:             c.region,
			Status:             value.Status,
			VpcSecurityGroupId: value.VpcSecurityGroupId,
		})
	}
	return tValues
}


func (c *Client) transformClusters(values *[]types.DBCluster) []*Cluster {
	var tValues []*Cluster
	for _, value := range *values {
		res := Cluster{
			Region:                           c.region,
			AccountID:                        c.accountID,
			ActivityStreamKinesisStreamName:  value.ActivityStreamKinesisStreamName,
			ActivityStreamKmsKeyId:           value.ActivityStreamKmsKeyId,
			ActivityStreamMode:               aws.String(string(value.ActivityStreamMode)),
			ActivityStreamStatus:             aws.String(string(value.ActivityStreamStatus)),
			AllocatedStorage:                 value.AllocatedStorage,
			AssociatedRoles:                  c.transformClusterRoles(&value.AssociatedRoles),
			AvailabilityZones:                common.StringListToString(&value.AvailabilityZones),
			BacktrackConsumedChangeRecords:   value.BacktrackConsumedChangeRecords,
			BacktrackWindow:                  value.BacktrackWindow,
			BackupRetentionPeriod:            value.BackupRetentionPeriod,
			Capacity:                         value.Capacity,
			CharacterSetName:                 value.CharacterSetName,
			CloneGroupId:                     value.CloneGroupId,
			ClusterCreateTime:                value.ClusterCreateTime,
			CopyTagsToSnapshot:               value.CopyTagsToSnapshot,
			CrossAccountClone:                value.CrossAccountClone,
			CustomEndpoints:                  common.StringListToString(&value.CustomEndpoints),
			ClusterArn:                       value.DBClusterArn,
			ClusterIdentifier:                value.DBClusterIdentifier,
			ClusterMembers:                   c.transformClusterMembers(&value.DBClusterMembers),
			ClusterOptionGroupMemberships:    c.transformClusterOptionGroupStatuss(&value.DBClusterOptionGroupMemberships),
			ClusterParameterGroup:            value.DBClusterParameterGroup,
			SubnetGroup:                      value.DBSubnetGroup,
			DatabaseName:                     value.DatabaseName,
			DbClusterResourceId:              value.DbClusterResourceId,
			DeletionProtection:               value.DeletionProtection,
			DomainMemberships:                c.transformClusterDomainMemberships(&value.DomainMemberships),
			EarliestBacktrackTime:            value.EarliestBacktrackTime,
			EarliestRestorableTime:           value.EarliestRestorableTime,
			EnabledCloudwatchLogsExports:     common.StringListToString(&value.EnabledCloudwatchLogsExports),
			Endpoint:                         value.Endpoint,
			Engine:                           value.Engine,
			EngineMode:                       value.EngineMode,
			EngineVersion:                    value.EngineVersion,
			GlobalWriteForwardingRequested:   value.GlobalWriteForwardingRequested,
			GlobalWriteForwardingStatus:      aws.String(string(value.GlobalWriteForwardingStatus)),
			HostedZoneId:                     value.HostedZoneId,
			HttpEndpointEnabled:              value.HttpEndpointEnabled,
			IAMDatabaseAuthenticationEnabled: value.IAMDatabaseAuthenticationEnabled,
			KmsKeyId:                         value.KmsKeyId,
			LatestRestorableTime:             value.LatestRestorableTime,
			MasterUsername:                   value.MasterUsername,
			MultiAZ:                          value.MultiAZ,
			PercentProgress:                  value.PercentProgress,
			Port:                             value.Port,
			PreferredBackupWindow:            value.PreferredBackupWindow,
			PreferredMaintenanceWindow:       value.PreferredMaintenanceWindow,
			ReadReplicaIdentifiers:           common.StringListToString(&value.ReadReplicaIdentifiers),
			ReaderEndpoint:                   value.ReaderEndpoint,
			ReplicationSourceIdentifier:      value.ReplicationSourceIdentifier,
			Status:                           value.Status,
			StorageEncrypted:                 &value.StorageEncrypted,
			VpcSecurityGroups:                c.transformClusterVpcSecurityGroupMemberships(&value.VpcSecurityGroups),
		}

		if value.ScalingConfigurationInfo != nil {
			res.ScalingConfigAutoPause = value.ScalingConfigurationInfo.AutoPause
			res.ScalingConfigMaxCapacity = value.ScalingConfigurationInfo.MaxCapacity
			res.ScalingConfigMinCapacity = value.ScalingConfigurationInfo.MinCapacity
			res.ScalingConfigSecondsUntilAutoPause = value.ScalingConfigurationInfo.SecondsUntilAutoPause
			res.ScalingConfigTimeoutAction = value.ScalingConfigurationInfo.TimeoutAction
		}
		tValues = append(tValues, &res)
	}
	return tValues
}

var ClusterTables = []interface{}{
	&Cluster{},
	&ClusterRole{},
	&ClusterMember{},
	&ClusterOptionGroupStatus{},
	&ClusterDomainMembership{},
	&ClusterVpcSecurityGroupMembership{},
}

func (c *Client) clusters(gConfig interface{}) error {
	var config rds.DescribeDBClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)

	for {
		output, err := c.svc.DescribeDBClusters(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusters(&output.DBClusters))
		c.log.Info("Fetched resources", zap.String("resource", "rds.clusters"), zap.Int("count", len(output.DBClusters)))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
