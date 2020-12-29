package redshift

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Cluster struct {
	ID                                     uint `gorm:"primarykey"`
	AccountID                              string
	Region                                 string
	AllowVersionUpgrade                    *bool
	AutomatedSnapshotRetentionPeriod       *int64
	AvailabilityZone                       *string
	ClusterAvailabilityStatus              *string
	ClusterCreateTime                      *time.Time
	ClusterIdentifier                      *string
	ClusterNodes                           []*ClusterNode                 `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterParameterGroups                 []*ClusterParameterGroupStatus `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterPublicKey                       *string
	ClusterRevisionNumber                  *string
	ClusterSecurityGroups                  []*ClusterSecurityGroupMembership   `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterSnapshotCopyStatus              *redshift.ClusterSnapshotCopyStatus `gorm:"embedded;embeddedPrefix:cluster_snapshot_copy_status_"`
	ClusterStatus                          *string
	ClusterSubnetGroupName                 *string
	ClusterVersion                         *string
	DBName                                 *string
	DataTransferProgress                   *redshift.DataTransferProgress      `gorm:"embedded;embeddedPrefix:data_transfer_progress_"`
	DeferredMaintenanceWindows             []*ClusterDeferredMaintenanceWindow `gorm:"constraint:OnDelete:CASCADE;"`
	ElasticIpStatus                        *redshift.ElasticIpStatus           `gorm:"embedded;embeddedPrefix:elastic_ip_status_"`
	ElasticResizeNumberOfNodeOptions       *string
	Encrypted                              *bool
	Endpoint                               *redshift.Endpoint `gorm:"embedded;embeddedPrefix:endpoint_"`
	EnhancedVpcRouting                     *bool
	ExpectedNextSnapshotScheduleTime       *time.Time
	ExpectedNextSnapshotScheduleTimeStatus *string
	HsmStatus                              *redshift.HsmStatus `gorm:"embedded;embeddedPrefix:hsm_status_"`
	IamRoles                               []*ClusterIamRole   `gorm:"constraint:OnDelete:CASCADE;"`
	KmsKeyId                               *string
	MaintenanceTrackName                   *string
	ManualSnapshotRetentionPeriod          *int64
	MasterUsername                         *string
	ModifyStatus                           *string
	NextMaintenanceWindowStartTime         *time.Time
	NodeType                               *string
	NumberOfNodes                          *int64
	PendingActions                         *string
	PendingModifiedValues                  *redshift.PendingModifiedValues `gorm:"embedded;embeddedPrefix:pending_modified_values_"`
	PreferredMaintenanceWindow             *string
	PubliclyAccessible                     *bool
	ResizeInfo                             *redshift.ResizeInfo    `gorm:"embedded;embeddedPrefix:resize_info_"`
	RestoreStatus                          *redshift.RestoreStatus `gorm:"embedded;embeddedPrefix:restore_status_"`
	SnapshotScheduleIdentifier             *string
	SnapshotScheduleState                  *string
	Tags                                   []*ClusterTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                                  *string
	VpcSecurityGroups                      []*ClusterVpcSecurityGroupMembership `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_redshift_clusters"
}

type ClusterNode struct {
	ID               uint `gorm:"primarykey"`
	ClusterID        uint
	NodeRole         *string
	PrivateIPAddress *string
	PublicIPAddress  *string
}

func (ClusterNode) TableName() string {
	return "aws_redshift_cluster_nodes"
}

type ClusterParameterGroupStatus struct {
	ID                   uint `gorm:"primarykey"`
	ClusterID            uint
	List                 []*ClusterParameterStatus `gorm:"constraint:OnDelete:CASCADE;"`
	ParameterApplyStatus *string
	ParameterGroupName   *string
}

func (ClusterParameterGroupStatus) TableName() string {
	return "aws_redshift_cluster_parameter_group_statuses"
}

type ClusterParameterStatus struct {
	ID                             uint `gorm:"primarykey"`
	ClusterParameterGroupStatusID  uint
	ParameterApplyErrorDescription *string
	ParameterApplyStatus           *string
	ParameterName                  *string
}

func (ClusterParameterStatus) TableName() string {
	return "aws_redshift_cluster_parameter_statuses"
}

type ClusterSecurityGroupMembership struct {
	ID                       uint `gorm:"primarykey"`
	ClusterID                uint
	ClusterSecurityGroupName *string
	Status                   *string
}

func (ClusterSecurityGroupMembership) TableName() string {
	return "aws_redshift_cluster_security_group_memberships"
}

type ClusterDeferredMaintenanceWindow struct {
	ID                         uint `gorm:"primarykey"`
	ClusterID                  uint
	DeferMaintenanceEndTime    *time.Time
	DeferMaintenanceIdentifier *string
	DeferMaintenanceStartTime  *time.Time
}

func (ClusterDeferredMaintenanceWindow) TableName() string {
	return "aws_redshift_cluster_deferred_maintenance_windows"
}

type ClusterIamRole struct {
	ID          uint `gorm:"primarykey"`
	ClusterID   uint
	ApplyStatus *string
	IamRoleArn  *string
}

func (ClusterIamRole) TableName() string {
	return "aws_redshift_cluster_iam_roles"
}

type ClusterTag struct {
	ID        uint `gorm:"primarykey"`
	ClusterID uint
	Key       *string
	Value     *string
}

func (ClusterTag) TableName() string {
	return "aws_redshift_cluster_tags"
}

type ClusterVpcSecurityGroupMembership struct {
	ID                 uint `gorm:"primarykey"`
	ClusterID          uint
	Status             *string
	VpcSecurityGroupId *string
}

func (ClusterVpcSecurityGroupMembership) TableName() string {
	return "aws_redshift_cluster_vpc_security_group_memberships"
}

func (c *Client) transformClusterNode(value *redshift.ClusterNode) *ClusterNode {
	return &ClusterNode{
		NodeRole:         value.NodeRole,
		PrivateIPAddress: value.PrivateIPAddress,
		PublicIPAddress:  value.PublicIPAddress,
	}
}

func (c *Client) transformClusterNodes(values []*redshift.ClusterNode) []*ClusterNode {
	var tValues []*ClusterNode
	for _, v := range values {
		tValues = append(tValues, c.transformClusterNode(v))
	}
	return tValues
}

func (c *Client) transformClusterParameterStatus(value *redshift.ClusterParameterStatus) *ClusterParameterStatus {
	return &ClusterParameterStatus{
		ParameterApplyErrorDescription: value.ParameterApplyErrorDescription,
		ParameterApplyStatus:           value.ParameterApplyStatus,
		ParameterName:                  value.ParameterName,
	}
}

func (c *Client) transformClusterParameterStatuss(values []*redshift.ClusterParameterStatus) []*ClusterParameterStatus {
	var tValues []*ClusterParameterStatus
	for _, v := range values {
		tValues = append(tValues, c.transformClusterParameterStatus(v))
	}
	return tValues
}

func (c *Client) transformClusterParameterGroupStatus(value *redshift.ClusterParameterGroupStatus) *ClusterParameterGroupStatus {
	return &ClusterParameterGroupStatus{
		List:                 c.transformClusterParameterStatuss(value.ClusterParameterStatusList),
		ParameterApplyStatus: value.ParameterApplyStatus,
		ParameterGroupName:   value.ParameterGroupName,
	}
}

func (c *Client) transformClusterParameterGroupStatuss(values []*redshift.ClusterParameterGroupStatus) []*ClusterParameterGroupStatus {
	var tValues []*ClusterParameterGroupStatus
	for _, v := range values {
		tValues = append(tValues, c.transformClusterParameterGroupStatus(v))
	}
	return tValues
}

func (c *Client) transformClusterSecurityGroupMembership(value *redshift.ClusterSecurityGroupMembership) *ClusterSecurityGroupMembership {
	return &ClusterSecurityGroupMembership{
		ClusterSecurityGroupName: value.ClusterSecurityGroupName,
		Status:                   value.Status,
	}
}

func (c *Client) transformClusterSecurityGroupMemberships(values []*redshift.ClusterSecurityGroupMembership) []*ClusterSecurityGroupMembership {
	var tValues []*ClusterSecurityGroupMembership
	for _, v := range values {
		tValues = append(tValues, c.transformClusterSecurityGroupMembership(v))
	}
	return tValues
}

func (c *Client) transformClusterDeferredMaintenanceWindow(value *redshift.DeferredMaintenanceWindow) *ClusterDeferredMaintenanceWindow {
	return &ClusterDeferredMaintenanceWindow{
		DeferMaintenanceEndTime:    value.DeferMaintenanceEndTime,
		DeferMaintenanceIdentifier: value.DeferMaintenanceIdentifier,
		DeferMaintenanceStartTime:  value.DeferMaintenanceStartTime,
	}
}

func (c *Client) transformClusterDeferredMaintenanceWindows(values []*redshift.DeferredMaintenanceWindow) []*ClusterDeferredMaintenanceWindow {
	var tValues []*ClusterDeferredMaintenanceWindow
	for _, v := range values {
		tValues = append(tValues, c.transformClusterDeferredMaintenanceWindow(v))
	}
	return tValues
}

func (c *Client) transformClusterIamRole(value *redshift.ClusterIamRole) *ClusterIamRole {
	return &ClusterIamRole{
		ApplyStatus: value.ApplyStatus,
		IamRoleArn:  value.IamRoleArn,
	}
}

func (c *Client) transformClusterIamRoles(values []*redshift.ClusterIamRole) []*ClusterIamRole {
	var tValues []*ClusterIamRole
	for _, v := range values {
		tValues = append(tValues, c.transformClusterIamRole(v))
	}
	return tValues
}

func (c *Client) transformClusterTag(value *redshift.Tag) *ClusterTag {
	return &ClusterTag{
		Key:   value.Key,
		Value: value.Value,
	}
}

func (c *Client) transformClusterTags(values []*redshift.Tag) []*ClusterTag {
	var tValues []*ClusterTag
	for _, v := range values {
		tValues = append(tValues, c.transformClusterTag(v))
	}
	return tValues
}

func (c *Client) transformClusterVpcSecurityGroupMembership(value *redshift.VpcSecurityGroupMembership) *ClusterVpcSecurityGroupMembership {
	return &ClusterVpcSecurityGroupMembership{
		Status:             value.Status,
		VpcSecurityGroupId: value.VpcSecurityGroupId,
	}
}

func (c *Client) transformClusterVpcSecurityGroupMemberships(values []*redshift.VpcSecurityGroupMembership) []*ClusterVpcSecurityGroupMembership {
	var tValues []*ClusterVpcSecurityGroupMembership
	for _, v := range values {
		tValues = append(tValues, c.transformClusterVpcSecurityGroupMembership(v))
	}
	return tValues
}

func (c *Client) transformCluster(value *redshift.Cluster) *Cluster {
	return &Cluster{
		Region:                                 c.region,
		AccountID:                              c.accountID,
		AllowVersionUpgrade:                    value.AllowVersionUpgrade,
		AutomatedSnapshotRetentionPeriod:       value.AutomatedSnapshotRetentionPeriod,
		AvailabilityZone:                       value.AvailabilityZone,
		ClusterAvailabilityStatus:              value.ClusterAvailabilityStatus,
		ClusterCreateTime:                      value.ClusterCreateTime,
		ClusterIdentifier:                      value.ClusterIdentifier,
		ClusterNodes:                           c.transformClusterNodes(value.ClusterNodes),
		ClusterParameterGroups:                 c.transformClusterParameterGroupStatuss(value.ClusterParameterGroups),
		ClusterPublicKey:                       value.ClusterPublicKey,
		ClusterRevisionNumber:                  value.ClusterRevisionNumber,
		ClusterSecurityGroups:                  c.transformClusterSecurityGroupMemberships(value.ClusterSecurityGroups),
		ClusterSnapshotCopyStatus:              value.ClusterSnapshotCopyStatus,
		ClusterStatus:                          value.ClusterStatus,
		ClusterSubnetGroupName:                 value.ClusterSubnetGroupName,
		ClusterVersion:                         value.ClusterVersion,
		DBName:                                 value.DBName,
		DataTransferProgress:                   value.DataTransferProgress,
		DeferredMaintenanceWindows:             c.transformClusterDeferredMaintenanceWindows(value.DeferredMaintenanceWindows),
		ElasticIpStatus:                        value.ElasticIpStatus,
		ElasticResizeNumberOfNodeOptions:       value.ElasticResizeNumberOfNodeOptions,
		Encrypted:                              value.Encrypted,
		Endpoint:                               value.Endpoint,
		EnhancedVpcRouting:                     value.EnhancedVpcRouting,
		ExpectedNextSnapshotScheduleTime:       value.ExpectedNextSnapshotScheduleTime,
		ExpectedNextSnapshotScheduleTimeStatus: value.ExpectedNextSnapshotScheduleTimeStatus,
		HsmStatus:                              value.HsmStatus,
		IamRoles:                               c.transformClusterIamRoles(value.IamRoles),
		KmsKeyId:                               value.KmsKeyId,
		MaintenanceTrackName:                   value.MaintenanceTrackName,
		ManualSnapshotRetentionPeriod:          value.ManualSnapshotRetentionPeriod,
		MasterUsername:                         value.MasterUsername,
		ModifyStatus:                           value.ModifyStatus,
		NextMaintenanceWindowStartTime:         value.NextMaintenanceWindowStartTime,
		NodeType:                               value.NodeType,
		NumberOfNodes:                          value.NumberOfNodes,
		PendingActions:                         common.StringListToString(value.PendingActions),
		PendingModifiedValues:                  value.PendingModifiedValues,
		PreferredMaintenanceWindow:             value.PreferredMaintenanceWindow,
		PubliclyAccessible:                     value.PubliclyAccessible,
		ResizeInfo:                             value.ResizeInfo,
		RestoreStatus:                          value.RestoreStatus,
		SnapshotScheduleIdentifier:             value.SnapshotScheduleIdentifier,
		SnapshotScheduleState:                  value.SnapshotScheduleState,
		Tags:                                   c.transformClusterTags(value.Tags),
		VpcId:                                  value.VpcId,
		VpcSecurityGroups:                      c.transformClusterVpcSecurityGroupMemberships(value.VpcSecurityGroups),
	}
}

func (c *Client) transformClusters(values []*redshift.Cluster) []*Cluster {
	var tValues []*Cluster
	for _, v := range values {
		tValues = append(tValues, c.transformCluster(v))
	}
	return tValues
}

func MigrateClusters(db *gorm.DB) error {
	return db.AutoMigrate(
		&Cluster{},
		&ClusterNode{},
		&ClusterParameterGroupStatus{},
		&ClusterParameterStatus{},
		&ClusterSecurityGroupMembership{},
		&ClusterDeferredMaintenanceWindow{},
		&ClusterIamRole{},
		&ClusterTag{},
		&ClusterVpcSecurityGroupMembership{},
	)
}

func (c *Client) clusters(gConfig interface{}) error {
	var config redshift.DescribeClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	for {
		output, err := c.svc.DescribeClusters(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Cluster{})
		common.ChunkedCreate(c.db, c.transformClusters(output.Clusters))
		c.log.Info("Fetched resources", zap.String("resource", "redshift.clusters"), zap.Int("count", len(output.Clusters)))
		if aws.StringValue(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
