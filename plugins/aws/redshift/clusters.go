package redshift

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Cluster struct {
	_                                interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                               uint        `gorm:"primarykey"`
	AccountID                        string      `neo:"unique"`
	Region                           string      `neo:"unique"`
	AllowVersionUpgrade              *bool
	AutomatedSnapshotRetentionPeriod *int32
	AvailabilityZone                 *string
	ClusterAvailabilityStatus        *string
	ClusterCreateTime                *time.Time
	ClusterIdentifier                *string                        `neo:"unique"`
	ClusterNodes                     []*ClusterNode                 `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterParameterGroups           []*ClusterParameterGroupStatus `gorm:"constraint:OnDelete:CASCADE;"`
	ClusterPublicKey                 *string
	ClusterRevisionNumber            *string
	ClusterSecurityGroups            []*ClusterSecurityGroupMembership `gorm:"constraint:OnDelete:CASCADE;"`

	ClusterSnapshotCopyStatusDestinationRegion             *string
	ClusterSnapshotCopyStatusManualSnapshotRetentionPeriod *int32
	ClusterSnapshotCopyStatusRetentionPeriod               *int64
	ClusterSnapshotCopyStatusSnapshotCopyGrantName         *string

	ClusterStatus          *string
	ClusterSubnetGroupName *string
	ClusterVersion         *string
	DBName                 *string

	DataTransferProgressCurrentRateInMegaBytesPerSecond    *float64
	DataTransferProgressDataTransferredInMegaBytes         *int64
	DataTransferProgressElapsedTimeInSeconds               *int64
	DataTransferProgressEstimatedTimeToCompletionInSeconds *int64
	DataTransferProgressStatus                             *string
	DataTransferProgressTotalDataInMegaBytes               *int64

	DeferredMaintenanceWindows []*ClusterDeferredMaintenanceWindow `gorm:"constraint:OnDelete:CASCADE;"`

	ElasticIpStatusElasticIp *string
	ElasticIpStatusStatus    *string

	ElasticResizeNumberOfNodeOptions *string
	Encrypted                        *bool

	EndpointAddress *string
	EndpointPort    *int32

	EnhancedVpcRouting                     *bool
	ExpectedNextSnapshotScheduleTime       *time.Time
	ExpectedNextSnapshotScheduleTimeStatus *string

	HsmStatusHsmClientCertificateIdentifier *string
	HsmStatusHsmConfigurationIdentifier     *string
	HsmStatusStatus                         *string

	IamRoles                       []*ClusterIamRole `gorm:"constraint:OnDelete:CASCADE;"`
	KmsKeyId                       *string
	MaintenanceTrackName           *string
	ManualSnapshotRetentionPeriod  *int32
	MasterUsername                 *string
	ModifyStatus                   *string
	NextMaintenanceWindowStartTime *time.Time
	NodeType                       *string
	NumberOfNodes                  *int32
	PendingActions                 []*ClusterPendingActions `gorm:"constraint:OnDelete:CASCADE;"`

	PendingModifiedValuesAutomatedSnapshotRetentionPeriod *int32
	PendingModifiedValuesClusterIdentifier                *string
	PendingModifiedValuesClusterType                      *string
	PendingModifiedValuesClusterVersion                   *string
	PendingModifiedValuesEncryptionType                   *string
	PendingModifiedValuesEnhancedVpcRouting               *bool
	PendingModifiedValuesMaintenanceTrackName             *string
	PendingModifiedValuesMasterUserPassword               *string
	PendingModifiedValuesNodeType                         *string
	PendingModifiedValuesNumberOfNodes                    *int32
	PendingModifiedValuesPubliclyAccessible               *bool

	PreferredMaintenanceWindow *string
	PubliclyAccessible         *bool

	ResizeInfoAllowCancelResize *bool
	ResizeInfoResizeType        *string

	RestoreStatusCurrentRestoreRateInMegaBytesPerSecond *float64
	RestoreStatusElapsedTimeInSeconds                   *int64
	RestoreStatusEstimatedTimeToCompletionInSeconds     *int64
	RestoreStatusProgressInMegaBytes                    *int64
	RestoreStatusSnapshotSizeInMegaBytes                *int64
	RestoreStatusStatus                                 *string

	SnapshotScheduleIdentifier *string
	SnapshotScheduleState      *string
	Tags                       []*ClusterTag `gorm:"constraint:OnDelete:CASCADE;"`
	VpcId                      *string
	VpcSecurityGroups          []*ClusterVpcSecurityGroupMembership `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_redshift_clusters"
}

type ClusterNode struct {
	ID               uint   `gorm:"primarykey"`
	AccountID        string `gorm:"-"`
	Region           string `gorm:"-"`
	ClusterID        uint   `neo:"ignore"`
	NodeRole         *string
	PrivateIPAddress *string
	PublicIPAddress  *string
}

func (ClusterNode) TableName() string {
	return "aws_redshift_cluster_nodes"
}

type ClusterParameterStatus struct {
	ID                             uint   `gorm:"primarykey"`
	AccountID                      string `gorm:"-"`
	Region                         string `gorm:"-"`
	ClusterParameterGroupStatusID  uint   `neo:"ignore"`
	ParameterApplyErrorDescription *string
	ParameterApplyStatus           *string
	ParameterName                  *string
}

func (ClusterParameterStatus) TableName() string {
	return "aws_redshift_cluster_parameter_statuses"
}

type ClusterParameterGroupStatus struct {
	ID                   uint                      `gorm:"primarykey"`
	AccountID            string                    `gorm:"-"`
	Region               string                    `gorm:"-"`
	ClusterID            uint                      `neo:"ignore"`
	List                 []*ClusterParameterStatus `gorm:"constraint:OnDelete:CASCADE;"`
	ParameterApplyStatus *string
	ParameterGroupName   *string
}

func (ClusterParameterGroupStatus) TableName() string {
	return "aws_redshift_cluster_parameter_group_statuses"
}

type ClusterSecurityGroupMembership struct {
	ID                       uint   `gorm:"primarykey"`
	AccountID                string `gorm:"-"`
	Region                   string `gorm:"-"`
	ClusterID                uint   `neo:"ignore"`
	ClusterSecurityGroupName *string
	Status                   *string
}

func (ClusterSecurityGroupMembership) TableName() string {
	return "aws_redshift_cluster_security_group_memberships"
}

type ClusterDeferredMaintenanceWindow struct {
	ID                         uint   `gorm:"primarykey"`
	AccountID                  string `gorm:"-"`
	Region                     string `gorm:"-"`
	ClusterID                  uint   `neo:"ignore"`
	DeferMaintenanceEndTime    *time.Time
	DeferMaintenanceIdentifier *string
	DeferMaintenanceStartTime  *time.Time
}

func (ClusterDeferredMaintenanceWindow) TableName() string {
	return "aws_redshift_cluster_deferred_maintenance_windows"
}

type ClusterIamRole struct {
	ID          uint   `gorm:"primarykey"`
	AccountID   string `gorm:"-"`
	Region      string `gorm:"-"`
	ClusterID   uint   `neo:"ignore"`
	ApplyStatus *string
	IamRoleArn  *string
}

func (ClusterIamRole) TableName() string {
	return "aws_redshift_cluster_iam_roles"
}

type ClusterPendingActions struct {
	ID        uint `gorm:"primarykey"`
	ClusterID uint
	Value     string
}

func (ClusterPendingActions) TableName() string {
	return "aws_redshift_cluster_pending_actions"
}

type ClusterTag struct {
	ID        uint   `gorm:"primarykey"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`
	ClusterID uint   `neo:"ignore"`
	Key       *string
	Value     *string
}

func (ClusterTag) TableName() string {
	return "aws_redshift_cluster_tags"
}

type ClusterVpcSecurityGroupMembership struct {
	ID                 uint   `gorm:"primarykey"`
	AccountID          string `gorm:"-"`
	Region             string `gorm:"-"`
	ClusterID          uint   `neo:"ignore"`
	Status             *string
	VpcSecurityGroupId *string
}

func (ClusterVpcSecurityGroupMembership) TableName() string {
	return "aws_redshift_cluster_vpc_security_group_memberships"
}

func (c *Client) transformClusters(values *[]types.Cluster) []*Cluster {
	var tValues []*Cluster
	for _, value := range *values {
		tValue := Cluster{
			AccountID:                              c.accountID,
			Region:                                 c.region,
			AllowVersionUpgrade:                    &value.AllowVersionUpgrade,
			AutomatedSnapshotRetentionPeriod:       &value.AutomatedSnapshotRetentionPeriod,
			AvailabilityZone:                       value.AvailabilityZone,
			ClusterAvailabilityStatus:              value.ClusterAvailabilityStatus,
			ClusterCreateTime:                      value.ClusterCreateTime,
			ClusterIdentifier:                      value.ClusterIdentifier,
			ClusterNodes:                           c.transformClusterNodes(&value.ClusterNodes),
			ClusterParameterGroups:                 c.transformClusterParameterGroupStatuss(&value.ClusterParameterGroups),
			ClusterPublicKey:                       value.ClusterPublicKey,
			ClusterRevisionNumber:                  value.ClusterRevisionNumber,
			ClusterSecurityGroups:                  c.transformClusterSecurityGroupMemberships(&value.ClusterSecurityGroups),
			ClusterStatus:                          value.ClusterStatus,
			ClusterSubnetGroupName:                 value.ClusterSubnetGroupName,
			ClusterVersion:                         value.ClusterVersion,
			DBName:                                 value.DBName,
			DeferredMaintenanceWindows:             c.transformClusterDeferredMaintenanceWindows(&value.DeferredMaintenanceWindows),
			ElasticResizeNumberOfNodeOptions:       value.ElasticResizeNumberOfNodeOptions,
			Encrypted:                              &value.Encrypted,
			EnhancedVpcRouting:                     &value.EnhancedVpcRouting,
			ExpectedNextSnapshotScheduleTime:       value.ExpectedNextSnapshotScheduleTime,
			ExpectedNextSnapshotScheduleTimeStatus: value.ExpectedNextSnapshotScheduleTimeStatus,
			IamRoles:                               c.transformClusterIamRoles(&value.IamRoles),
			KmsKeyId:                               value.KmsKeyId,
			MaintenanceTrackName:                   value.MaintenanceTrackName,
			ManualSnapshotRetentionPeriod:          &value.ManualSnapshotRetentionPeriod,
			MasterUsername:                         value.MasterUsername,
			ModifyStatus:                           value.ModifyStatus,
			NextMaintenanceWindowStartTime:         value.NextMaintenanceWindowStartTime,
			NodeType:                               value.NodeType,
			NumberOfNodes:                          &value.NumberOfNodes,
			PendingActions:                         c.transformClusterPendingActions(&value.PendingActions),
			PreferredMaintenanceWindow:             value.PreferredMaintenanceWindow,
			PubliclyAccessible:                     &value.PubliclyAccessible,
			SnapshotScheduleIdentifier:             value.SnapshotScheduleIdentifier,
			SnapshotScheduleState:                  aws.String(string(value.SnapshotScheduleState)),
			Tags:                                   c.transformClusterTags(&value.Tags),
			VpcId:                                  value.VpcId,
			VpcSecurityGroups:                      c.transformClusterVpcSecurityGroupMemberships(&value.VpcSecurityGroups),
		}
		if value.ClusterSnapshotCopyStatus != nil {
			tValue.ClusterSnapshotCopyStatusDestinationRegion = value.ClusterSnapshotCopyStatus.DestinationRegion
			tValue.ClusterSnapshotCopyStatusManualSnapshotRetentionPeriod = &value.ClusterSnapshotCopyStatus.ManualSnapshotRetentionPeriod
			tValue.ClusterSnapshotCopyStatusRetentionPeriod = &value.ClusterSnapshotCopyStatus.RetentionPeriod
			tValue.ClusterSnapshotCopyStatusSnapshotCopyGrantName = value.ClusterSnapshotCopyStatus.SnapshotCopyGrantName

		}
		if value.DataTransferProgress != nil {
			tValue.DataTransferProgressCurrentRateInMegaBytesPerSecond = value.DataTransferProgress.CurrentRateInMegaBytesPerSecond
			tValue.DataTransferProgressDataTransferredInMegaBytes = &value.DataTransferProgress.DataTransferredInMegaBytes
			tValue.DataTransferProgressElapsedTimeInSeconds = value.DataTransferProgress.ElapsedTimeInSeconds
			tValue.DataTransferProgressEstimatedTimeToCompletionInSeconds = value.DataTransferProgress.EstimatedTimeToCompletionInSeconds
			tValue.DataTransferProgressStatus = value.DataTransferProgress.Status
			tValue.DataTransferProgressTotalDataInMegaBytes = &value.DataTransferProgress.TotalDataInMegaBytes

		}
		if value.ElasticIpStatus != nil {

			tValue.ElasticIpStatusElasticIp = value.ElasticIpStatus.ElasticIp
			tValue.ElasticIpStatusStatus = value.ElasticIpStatus.Status

		}
		if value.Endpoint != nil {

			tValue.EndpointAddress = value.Endpoint.Address
			tValue.EndpointPort = &value.Endpoint.Port

		}
		if value.HsmStatus != nil {

			tValue.HsmStatusHsmClientCertificateIdentifier = value.HsmStatus.HsmClientCertificateIdentifier
			tValue.HsmStatusHsmConfigurationIdentifier = value.HsmStatus.HsmConfigurationIdentifier
			tValue.HsmStatusStatus = value.HsmStatus.Status

		}
		if value.PendingModifiedValues != nil {

			tValue.PendingModifiedValuesAutomatedSnapshotRetentionPeriod = value.PendingModifiedValues.AutomatedSnapshotRetentionPeriod
			tValue.PendingModifiedValuesClusterIdentifier = value.PendingModifiedValues.ClusterIdentifier
			tValue.PendingModifiedValuesClusterType = value.PendingModifiedValues.ClusterType
			tValue.PendingModifiedValuesClusterVersion = value.PendingModifiedValues.ClusterVersion
			tValue.PendingModifiedValuesEncryptionType = value.PendingModifiedValues.EncryptionType
			tValue.PendingModifiedValuesEnhancedVpcRouting = value.PendingModifiedValues.EnhancedVpcRouting
			tValue.PendingModifiedValuesMaintenanceTrackName = value.PendingModifiedValues.MaintenanceTrackName
			tValue.PendingModifiedValuesMasterUserPassword = value.PendingModifiedValues.MasterUserPassword
			tValue.PendingModifiedValuesNodeType = value.PendingModifiedValues.NodeType
			tValue.PendingModifiedValuesNumberOfNodes = value.PendingModifiedValues.NumberOfNodes
			tValue.PendingModifiedValuesPubliclyAccessible = value.PendingModifiedValues.PubliclyAccessible

		}
		if value.ResizeInfo != nil {
			tValue.ResizeInfoAllowCancelResize = &value.ResizeInfo.AllowCancelResize
			tValue.ResizeInfoResizeType = value.ResizeInfo.ResizeType
		}
		if value.RestoreStatus != nil {

			tValue.RestoreStatusCurrentRestoreRateInMegaBytesPerSecond = &value.RestoreStatus.CurrentRestoreRateInMegaBytesPerSecond
			tValue.RestoreStatusElapsedTimeInSeconds = &value.RestoreStatus.ElapsedTimeInSeconds
			tValue.RestoreStatusEstimatedTimeToCompletionInSeconds = &value.RestoreStatus.EstimatedTimeToCompletionInSeconds
			tValue.RestoreStatusProgressInMegaBytes = &value.RestoreStatus.ProgressInMegaBytes
			tValue.RestoreStatusSnapshotSizeInMegaBytes = &value.RestoreStatus.SnapshotSizeInMegaBytes
			tValue.RestoreStatusStatus = value.RestoreStatus.Status

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterNodes(values *[]types.ClusterNode) []*ClusterNode {
	var tValues []*ClusterNode
	for _, value := range *values {
		tValue := ClusterNode{
			AccountID:        c.accountID,
			Region:           c.region,
			NodeRole:         value.NodeRole,
			PrivateIPAddress: value.PrivateIPAddress,
			PublicIPAddress:  value.PublicIPAddress,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterParameterStatuss(values *[]types.ClusterParameterStatus) []*ClusterParameterStatus {
	var tValues []*ClusterParameterStatus
	for _, value := range *values {
		tValue := ClusterParameterStatus{
			AccountID:                      c.accountID,
			Region:                         c.region,
			ParameterApplyErrorDescription: value.ParameterApplyErrorDescription,
			ParameterApplyStatus:           value.ParameterApplyStatus,
			ParameterName:                  value.ParameterName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterParameterGroupStatuss(values *[]types.ClusterParameterGroupStatus) []*ClusterParameterGroupStatus {
	var tValues []*ClusterParameterGroupStatus
	for _, value := range *values {
		tValue := ClusterParameterGroupStatus{
			AccountID:            c.accountID,
			Region:               c.region,
			List:                 c.transformClusterParameterStatuss(&value.ClusterParameterStatusList),
			ParameterApplyStatus: value.ParameterApplyStatus,
			ParameterGroupName:   value.ParameterGroupName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterSecurityGroupMemberships(values *[]types.ClusterSecurityGroupMembership) []*ClusterSecurityGroupMembership {
	var tValues []*ClusterSecurityGroupMembership
	for _, value := range *values {
		tValue := ClusterSecurityGroupMembership{
			AccountID:                c.accountID,
			Region:                   c.region,
			ClusterSecurityGroupName: value.ClusterSecurityGroupName,
			Status:                   value.Status,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterDeferredMaintenanceWindows(values *[]types.DeferredMaintenanceWindow) []*ClusterDeferredMaintenanceWindow {
	var tValues []*ClusterDeferredMaintenanceWindow
	for _, value := range *values {
		tValue := ClusterDeferredMaintenanceWindow{
			AccountID:                  c.accountID,
			Region:                     c.region,
			DeferMaintenanceEndTime:    value.DeferMaintenanceEndTime,
			DeferMaintenanceIdentifier: value.DeferMaintenanceIdentifier,
			DeferMaintenanceStartTime:  value.DeferMaintenanceStartTime,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterIamRoles(values *[]types.ClusterIamRole) []*ClusterIamRole {
	var tValues []*ClusterIamRole
	for _, value := range *values {
		tValue := ClusterIamRole{
			AccountID:   c.accountID,
			Region:      c.region,
			ApplyStatus: value.ApplyStatus,
			IamRoleArn:  value.IamRoleArn,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformClusterPendingActions(values *[]string) []*ClusterPendingActions {
	var tValues []*ClusterPendingActions
	for _, v := range *values {
		tValues = append(tValues, &ClusterPendingActions{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformClusterTags(values *[]types.Tag) []*ClusterTag {
	var tValues []*ClusterTag
	for _, value := range *values {
		tValue := ClusterTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       value.Key,
			Value:     value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformClusterVpcSecurityGroupMemberships(values *[]types.VpcSecurityGroupMembership) []*ClusterVpcSecurityGroupMembership {
	var tValues []*ClusterVpcSecurityGroupMembership
	for _, value := range *values {
		tValue := ClusterVpcSecurityGroupMembership{
			AccountID:          c.accountID,
			Region:             c.region,
			Status:             value.Status,
			VpcSecurityGroupId: value.VpcSecurityGroupId,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type ClusterConfig struct {
	Filter string
}

var ClusterTables = []interface{}{
	&Cluster{},
	&ClusterNode{},
	&ClusterParameterGroupStatus{},
	&ClusterParameterStatus{},
	&ClusterSecurityGroupMembership{},
	&ClusterDeferredMaintenanceWindow{},
	&ClusterIamRole{},
	&ClusterTag{},
	&ClusterVpcSecurityGroupMembership{},
}

func (c *Client) clusters(gConfig interface{}) error {
	var config redshift.DescribeClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)

	for {
		output, err := c.svc.DescribeClusters(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformClusters(&output.Clusters))
		c.log.Info("Fetched resources", zap.String("resource", "redshift.clusters"), zap.Int("count", len(output.Clusters)))
		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
