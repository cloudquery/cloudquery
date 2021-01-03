package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/cloudquery/cloudquery/providers/common"
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
	AllocatedStorage                 *int64
	AssociatedRoles                  []*ClusterRole `gorm:"constraint:OnDelete:CASCADE;"`
	AvailabilityZones                *string
	BacktrackConsumedChangeRecords   *int64
	BacktrackWindow                  *int64
	BackupRetentionPeriod            *int64
	Capacity                         *int64
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
	Port                             *int64
	PreferredBackupWindow            *string
	PreferredMaintenanceWindow       *string
	ReadReplicaIdentifiers           *string
	ReaderEndpoint                   *string
	ReplicationSourceIdentifier      *string

	ScalingConfigAutoPause             *bool
	ScalingConfigMaxCapacity           *int64
	ScalingConfigMinCapacity           *int64
	ScalingConfigSecondsUntilAutoPause *int64
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
	PromotionTier               *int64
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

func (c *Client) transformClusterRole(value *rds.DBClusterRole) *ClusterRole {
	return &ClusterRole{
		AccountID:   c.accountID,
		Region:      c.region,
		FeatureName: value.FeatureName,
		RoleArn:     value.RoleArn,
		Status:      value.Status,
	}
}

func (c *Client) transformClusterRoles(values []*rds.DBClusterRole) []*ClusterRole {
	var tValues []*ClusterRole
	for _, v := range values {
		tValues = append(tValues, c.transformClusterRole(v))
	}
	return tValues
}

func (c *Client) transformClusterMember(value *rds.DBClusterMember) *ClusterMember {
	return &ClusterMember{
		AccountID:                   c.accountID,
		Region:                      c.region,
		ClusterParameterGroupStatus: value.DBClusterParameterGroupStatus,
		InstanceIdentifier:          value.DBInstanceIdentifier,
		IsClusterWriter:             value.IsClusterWriter,
		PromotionTier:               value.PromotionTier,
	}
}

func (c *Client) transformClusterMembers(values []*rds.DBClusterMember) []*ClusterMember {
	var tValues []*ClusterMember
	for _, v := range values {
		tValues = append(tValues, c.transformClusterMember(v))
	}
	return tValues
}

func (c *Client) transformClusterOptionGroupStatus(value *rds.DBClusterOptionGroupStatus) *ClusterOptionGroupStatus {
	return &ClusterOptionGroupStatus{
		AccountID:              c.accountID,
		Region:                 c.region,
		ClusterOptionGroupName: value.DBClusterOptionGroupName,
		Status:                 value.Status,
	}
}

func (c *Client) transformClusterOptionGroupStatuss(values []*rds.DBClusterOptionGroupStatus) []*ClusterOptionGroupStatus {
	var tValues []*ClusterOptionGroupStatus
	for _, v := range values {
		tValues = append(tValues, c.transformClusterOptionGroupStatus(v))
	}
	return tValues
}

func (c *Client) transformClusterDomainMembership(value *rds.DomainMembership) *ClusterDomainMembership {
	return &ClusterDomainMembership{
		AccountID:   c.accountID,
		Region:      c.region,
		Domain:      value.Domain,
		FQDN:        value.FQDN,
		IAMRoleName: value.IAMRoleName,
		Status:      value.Status,
	}
}

func (c *Client) transformClusterDomainMemberships(values []*rds.DomainMembership) []*ClusterDomainMembership {
	var tValues []*ClusterDomainMembership
	for _, v := range values {
		tValues = append(tValues, c.transformClusterDomainMembership(v))
	}
	return tValues
}

func (c *Client) transformClusterVpcSecurityGroupMembership(value *rds.VpcSecurityGroupMembership) *ClusterVpcSecurityGroupMembership {
	return &ClusterVpcSecurityGroupMembership{
		AccountID:          c.accountID,
		Region:             c.region,
		Status:             value.Status,
		VpcSecurityGroupId: value.VpcSecurityGroupId,
	}
}

func (c *Client) transformClusterVpcSecurityGroupMemberships(values []*rds.VpcSecurityGroupMembership) []*ClusterVpcSecurityGroupMembership {
	var tValues []*ClusterVpcSecurityGroupMembership
	for _, v := range values {
		tValues = append(tValues, c.transformClusterVpcSecurityGroupMembership(v))
	}
	return tValues
}

func (c *Client) transformCluster(value *rds.DBCluster) *Cluster {
	res := Cluster{
		Region:                           c.region,
		AccountID:                        c.accountID,
		ActivityStreamKinesisStreamName:  value.ActivityStreamKinesisStreamName,
		ActivityStreamKmsKeyId:           value.ActivityStreamKmsKeyId,
		ActivityStreamMode:               value.ActivityStreamMode,
		ActivityStreamStatus:             value.ActivityStreamStatus,
		AllocatedStorage:                 value.AllocatedStorage,
		AssociatedRoles:                  c.transformClusterRoles(value.AssociatedRoles),
		AvailabilityZones:                common.StringListToString(value.AvailabilityZones),
		BacktrackConsumedChangeRecords:   value.BacktrackConsumedChangeRecords,
		BacktrackWindow:                  value.BacktrackWindow,
		BackupRetentionPeriod:            value.BackupRetentionPeriod,
		Capacity:                         value.Capacity,
		CharacterSetName:                 value.CharacterSetName,
		CloneGroupId:                     value.CloneGroupId,
		ClusterCreateTime:                value.ClusterCreateTime,
		CopyTagsToSnapshot:               value.CopyTagsToSnapshot,
		CrossAccountClone:                value.CrossAccountClone,
		CustomEndpoints:                  common.StringListToString(value.CustomEndpoints),
		ClusterArn:                       value.DBClusterArn,
		ClusterIdentifier:                value.DBClusterIdentifier,
		ClusterMembers:                   c.transformClusterMembers(value.DBClusterMembers),
		ClusterOptionGroupMemberships:    c.transformClusterOptionGroupStatuss(value.DBClusterOptionGroupMemberships),
		ClusterParameterGroup:            value.DBClusterParameterGroup,
		SubnetGroup:                      value.DBSubnetGroup,
		DatabaseName:                     value.DatabaseName,
		DbClusterResourceId:              value.DbClusterResourceId,
		DeletionProtection:               value.DeletionProtection,
		DomainMemberships:                c.transformClusterDomainMemberships(value.DomainMemberships),
		EarliestBacktrackTime:            value.EarliestBacktrackTime,
		EarliestRestorableTime:           value.EarliestRestorableTime,
		EnabledCloudwatchLogsExports:     common.StringListToString(value.EnabledCloudwatchLogsExports),
		Endpoint:                         value.Endpoint,
		Engine:                           value.Engine,
		EngineMode:                       value.EngineMode,
		EngineVersion:                    value.EngineVersion,
		GlobalWriteForwardingRequested:   value.GlobalWriteForwardingRequested,
		GlobalWriteForwardingStatus:      value.GlobalWriteForwardingStatus,
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
		ReadReplicaIdentifiers:           common.StringListToString(value.ReadReplicaIdentifiers),
		ReaderEndpoint:                   value.ReaderEndpoint,
		ReplicationSourceIdentifier:      value.ReplicationSourceIdentifier,
		Status:                           value.Status,
		StorageEncrypted:                 value.StorageEncrypted,
		VpcSecurityGroups:                c.transformClusterVpcSecurityGroupMemberships(value.VpcSecurityGroups),
	}

	if value.ScalingConfigurationInfo != nil {
		res.ScalingConfigAutoPause = value.ScalingConfigurationInfo.AutoPause
		res.ScalingConfigMaxCapacity = value.ScalingConfigurationInfo.MaxCapacity
		res.ScalingConfigMinCapacity = value.ScalingConfigurationInfo.MinCapacity
		res.ScalingConfigSecondsUntilAutoPause = value.ScalingConfigurationInfo.SecondsUntilAutoPause
		res.ScalingConfigTimeoutAction = value.ScalingConfigurationInfo.TimeoutAction
	}

	return &res
}

func (c *Client) transformClusters(values []*rds.DBCluster) []*Cluster {
	var tValues []*Cluster
	for _, v := range values {
		tValues = append(tValues, c.transformCluster(v))
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
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)

	for {
		output, err := c.svc.DescribeDBClusters(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusters(output.DBClusters))
		c.log.Info("Fetched resources", zap.String("resource", "rds.clusters"), zap.Int("count", len(output.DBClusters)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
