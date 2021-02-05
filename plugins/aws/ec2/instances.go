package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Instance struct {
	_                     interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                    uint        `gorm:"primarykey"`
	AccountID             string      `neo:"unique"`
	Region                string      `neo:"unique"`
	AmiLaunchIndex        *int64
	Architecture          *string
	BlockDeviceMappings   []*InstanceBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`
	CapacityReservationId *string

	CapacityReservationPreference     *string
	CapacityReservationTargetId       *string
	CapacityReservationTargetGroupArn *string

	ClientToken *string

	CpuOptionsCoreCount      *int64
	CpuOptionsThreadsPerCore *int64

	EbsOptimized                            *bool
	ElasticGpuAssociations                  []*InstanceElasticGpuAssociation                  `gorm:"constraint:OnDelete:CASCADE;"`
	ElasticInferenceAcceleratorAssociations []*InstanceElasticInferenceAcceleratorAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	EnaSupport                              *bool
	HibernationOptionsConfigured            *bool
	Hypervisor                              *string

	IamInstanceProfileArn *string
	IamInstanceProfileId  *string

	ImageId           *string
	InstanceId        *string `neo:"unique"`
	InstanceLifecycle *string
	InstanceType      *string
	KernelId          *string
	KeyName           *string
	LaunchTime        *time.Time
	Licenses          []*InstanceLicenseConfiguration `gorm:"constraint:OnDelete:CASCADE;"`

	MetadataOptionsHttpEndpoint            *string
	MetadataOptionsHttpPutResponseHopLimit *int64
	MetadataOptionsHttpTokens              *string
	MetadataOptionsState                   *string

	MonitoringState *string

	NetworkInterfaces []*InstanceNetworkInterface `gorm:"constraint:OnDelete:CASCADE;"`
	OutpostArn        *string

	PlacementAffinity             *string
	PlacementAvailabilityZone     *string
	PlacementGroupName            *string
	PlacementHostId               *string
	PlacementHostResourceGroupArn *string
	PlacementPartitionNumber      *int64
	PlacementSpreadDomain         *string
	PlacementTenancy              *string

	Platform              *string
	PrivateDnsName        *string
	PrivateIpAddress      *string
	ProductCodes          []*InstanceProductCode `gorm:"constraint:OnDelete:CASCADE;"`
	PublicDnsName         *string
	PublicIpAddress       *string
	RamdiskId             *string
	RootDeviceName        *string
	RootDeviceType        *string
	SecurityGroups        []*InstanceGroupIdentifier `gorm:"constraint:OnDelete:CASCADE;"`
	SourceDestCheck       *bool
	SpotInstanceRequestId *string
	SriovNetSupport       *string

	StateCode *int64
	StateName *string

	StateReasonCode    *string
	StateReasonMessage *string

	StateTransitionReason *string
	SubnetId              *string
	Tags                  []*InstanceTag `gorm:"constraint:OnDelete:CASCADE;"`
	VirtualizationType    *string
	VpcId                 *string
}

func (Instance) TableName() string {
	return "aws_ec2_instances"
}

type InstanceBlockDeviceMapping struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	DeviceName *string

	AttachTime          *time.Time
	DeleteOnTermination *bool
	Status              *string
	VolumeId            *string
}

func (InstanceBlockDeviceMapping) TableName() string {
	return "aws_ec2_instance_block_device_mappings"
}

type InstanceElasticGpuAssociation struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ElasticGpuAssociationId    *string
	ElasticGpuAssociationState *string
	ElasticGpuAssociationTime  *string
	ElasticGpuId               *string
}

func (InstanceElasticGpuAssociation) TableName() string {
	return "aws_ec2_instance_elastic_gpu_associations"
}

type InstanceElasticInferenceAcceleratorAssociation struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ElasticInferenceAcceleratorArn              *string
	ElasticInferenceAcceleratorAssociationId    *string
	ElasticInferenceAcceleratorAssociationState *string
	ElasticInferenceAcceleratorAssociationTime  *time.Time
}

func (InstanceElasticInferenceAcceleratorAssociation) TableName() string {
	return "aws_ec2_instance_elastic_inference_accelerator_associations"
}

type InstanceLicenseConfiguration struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	LicenseConfigurationArn *string
}

func (InstanceLicenseConfiguration) TableName() string {
	return "aws_ec2_instance_license_configurations"
}

type InstanceNetworkInterface struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AssociationCarrierIp     *string
	AssociationIpOwnerId     *string
	AssociationPublicDnsName *string
	AssociationPublicIp      *string

	AttachmentTime                *time.Time
	AttachmentId                  *string
	AttachmentDeleteOnTermination *bool
	AttachmentDeviceIndex         *int64
	AttachmentStatus              *string

	Description        *string
	InterfaceType      *string
	Ipv6Addresses      []*InstanceIpv6Address `gorm:"constraint:OnDelete:CASCADE;"`
	MacAddress         *string
	NetworkInterfaceId *string
	OwnerId            *string
	PrivateDnsName     *string
	PrivateIpAddress   *string
	PrivateIpAddresses []*InstancePrivateIpAddress `gorm:"constraint:OnDelete:CASCADE;"`
	SourceDestCheck    *bool
	Status             *string
	SubnetId           *string
	VpcId              *string
}

func (InstanceNetworkInterface) TableName() string {
	return "aws_ec2_instance_network_interfaces"
}

type InstanceGroupIdentifier struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	GroupId   *string
	GroupName *string
}

func (InstanceGroupIdentifier) TableName() string {
	return "aws_ec2_instance_group_identifiers"
}

type InstanceIpv6Address struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Ipv6Address *string
}

func (InstanceIpv6Address) TableName() string {
	return "aws_ec2_instance_ipv6_addresses"
}

type InstancePrivateIpAddress struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	AssociationCarrierIp     *string
	AssociationIpOwnerId     *string
	AssociationPublicDnsName *string
	AssociationPublicIp      *string

	Primary          *bool
	PrivateDnsName   *string
	PrivateIpAddress *string
}

func (InstancePrivateIpAddress) TableName() string {
	return "aws_ec2_instance_private_ip_addresses"
}

type InstanceProductCode struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ProductCodeId   *string
	ProductCodeType *string
}

func (InstanceProductCode) TableName() string {
	return "aws_ec2_instance_product_codes"
}

type InstanceTag struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint `neo:"ignore"`

	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (InstanceTag) TableName() string {
	return "aws_ec2_instance_tags"
}

func (c *Client) transformInstanceBlockDeviceMapping(value *ec2.InstanceBlockDeviceMapping) *InstanceBlockDeviceMapping {
	res := InstanceBlockDeviceMapping{
		DeviceName: value.DeviceName,
		AccountID:  c.accountID,
		Region:     c.region,
	}

	if value.Ebs != nil {
		res.AttachTime = value.Ebs.AttachTime
		res.DeleteOnTermination = value.Ebs.DeleteOnTermination
		res.Status = value.Ebs.Status
		res.VolumeId = value.Ebs.VolumeId
	}

	return &res
}

func (c *Client) transformInstanceBlockDeviceMappings(values []*ec2.InstanceBlockDeviceMapping) []*InstanceBlockDeviceMapping {
	var tValues []*InstanceBlockDeviceMapping
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceBlockDeviceMapping(v))
	}
	return tValues
}

func (c *Client) transformInstanceElasticGpuAssociation(value *ec2.ElasticGpuAssociation) *InstanceElasticGpuAssociation {
	return &InstanceElasticGpuAssociation{
		AccountID:                  c.accountID,
		Region:                     c.region,
		ElasticGpuAssociationId:    value.ElasticGpuAssociationId,
		ElasticGpuAssociationState: value.ElasticGpuAssociationState,
		ElasticGpuAssociationTime:  value.ElasticGpuAssociationTime,
		ElasticGpuId:               value.ElasticGpuId,
	}
}

func (c *Client) transformInstanceElasticGpuAssociations(values []*ec2.ElasticGpuAssociation) []*InstanceElasticGpuAssociation {
	var tValues []*InstanceElasticGpuAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceElasticGpuAssociation(v))
	}
	return tValues
}

func (c *Client) transformInstanceElasticInferenceAcceleratorAssociation(value *ec2.ElasticInferenceAcceleratorAssociation) *InstanceElasticInferenceAcceleratorAssociation {
	return &InstanceElasticInferenceAcceleratorAssociation{
		AccountID:                                c.accountID,
		Region:                                   c.region,
		ElasticInferenceAcceleratorArn:           value.ElasticInferenceAcceleratorArn,
		ElasticInferenceAcceleratorAssociationId: value.ElasticInferenceAcceleratorAssociationId,
		ElasticInferenceAcceleratorAssociationState: value.ElasticInferenceAcceleratorAssociationState,
		ElasticInferenceAcceleratorAssociationTime:  value.ElasticInferenceAcceleratorAssociationTime,
	}
}

func (c *Client) transformInstanceElasticInferenceAcceleratorAssociations(values []*ec2.ElasticInferenceAcceleratorAssociation) []*InstanceElasticInferenceAcceleratorAssociation {
	var tValues []*InstanceElasticInferenceAcceleratorAssociation
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceElasticInferenceAcceleratorAssociation(v))
	}
	return tValues
}

func (c *Client) transformInstanceLicenseConfiguration(value *ec2.LicenseConfiguration) *InstanceLicenseConfiguration {
	return &InstanceLicenseConfiguration{
		AccountID:               c.accountID,
		Region:                  c.region,
		LicenseConfigurationArn: value.LicenseConfigurationArn,
	}
}

func (c *Client) transformInstanceLicenseConfigurations(values []*ec2.LicenseConfiguration) []*InstanceLicenseConfiguration {
	var tValues []*InstanceLicenseConfiguration
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceLicenseConfiguration(v))
	}
	return tValues
}

func (c *Client) transformInstanceGroupIdentifier(value *ec2.GroupIdentifier) *InstanceGroupIdentifier {
	return &InstanceGroupIdentifier{
		AccountID: c.accountID,
		Region:    c.region,
		GroupId:   value.GroupId,
		GroupName: value.GroupName,
	}
}

func (c *Client) transformInstanceGroupIdentifiers(values []*ec2.GroupIdentifier) []*InstanceGroupIdentifier {
	var tValues []*InstanceGroupIdentifier
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceGroupIdentifier(v))
	}
	return tValues
}

func (c *Client) transformInstanceIpv6Address(value *ec2.InstanceIpv6Address) *InstanceIpv6Address {
	return &InstanceIpv6Address{
		AccountID:   c.accountID,
		Region:      c.region,
		Ipv6Address: value.Ipv6Address,
	}
}

func (c *Client) transformInstanceIpv6Addresss(values []*ec2.InstanceIpv6Address) []*InstanceIpv6Address {
	var tValues []*InstanceIpv6Address
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceIpv6Address(v))
	}
	return tValues
}

func (c *Client) transformInstancePrivateIpAddress(value *ec2.InstancePrivateIpAddress) *InstancePrivateIpAddress {
	res := InstancePrivateIpAddress{
		AccountID:        c.accountID,
		Region:           c.region,
		Primary:          value.Primary,
		PrivateDnsName:   value.PrivateDnsName,
		PrivateIpAddress: value.PrivateIpAddress,
	}

	if value.Association != nil {
		res.AssociationCarrierIp = value.Association.CarrierIp
		res.AssociationIpOwnerId = value.Association.IpOwnerId
		res.AssociationPublicDnsName = value.Association.PublicDnsName
		res.AssociationPublicIp = value.Association.PublicIp
	}

	return &res
}

func (c *Client) transformInstancePrivateIpAddresss(values []*ec2.InstancePrivateIpAddress) []*InstancePrivateIpAddress {
	var tValues []*InstancePrivateIpAddress
	for _, v := range values {
		tValues = append(tValues, c.transformInstancePrivateIpAddress(v))
	}
	return tValues
}

func (c *Client) transformInstanceNetworkInterface(value *ec2.InstanceNetworkInterface) *InstanceNetworkInterface {
	res := InstanceNetworkInterface{
		AccountID:          c.accountID,
		Region:             c.region,
		Description:        value.Description,
		InterfaceType:      value.InterfaceType,
		Ipv6Addresses:      c.transformInstanceIpv6Addresss(value.Ipv6Addresses),
		MacAddress:         value.MacAddress,
		NetworkInterfaceId: value.NetworkInterfaceId,
		OwnerId:            value.OwnerId,
		PrivateDnsName:     value.PrivateDnsName,
		PrivateIpAddress:   value.PrivateIpAddress,
		PrivateIpAddresses: c.transformInstancePrivateIpAddresss(value.PrivateIpAddresses),
		SourceDestCheck:    value.SourceDestCheck,
		Status:             value.Status,
		SubnetId:           value.SubnetId,
		VpcId:              value.VpcId,
	}

	if value.Attachment != nil {
		res.AttachmentTime = value.Attachment.AttachTime
		res.AttachmentId = value.Attachment.AttachmentId
		res.AttachmentDeleteOnTermination = value.Attachment.DeleteOnTermination
		res.AttachmentDeviceIndex = value.Attachment.DeviceIndex
		res.AttachmentStatus = value.Attachment.Status
	}

	if value.Association != nil {
		res.AssociationCarrierIp = value.Association.CarrierIp
		res.AssociationIpOwnerId = value.Association.IpOwnerId
		res.AssociationPublicDnsName = value.Association.PublicDnsName
		res.AssociationPublicIp = value.Association.PublicIp
	}

	return &res
}

func (c *Client) transformInstanceNetworkInterfaces(values []*ec2.InstanceNetworkInterface) []*InstanceNetworkInterface {
	var tValues []*InstanceNetworkInterface
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceNetworkInterface(v))
	}
	return tValues
}

func (c *Client) transformInstanceProductCode(value *ec2.ProductCode) *InstanceProductCode {
	return &InstanceProductCode{
		AccountID:       c.accountID,
		Region:          c.region,
		ProductCodeId:   value.ProductCodeId,
		ProductCodeType: value.ProductCodeType,
	}
}

func (c *Client) transformInstanceProductCodes(values []*ec2.ProductCode) []*InstanceProductCode {
	var tValues []*InstanceProductCode
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceProductCode(v))
	}
	return tValues
}

func (c *Client) transformInstanceTag(value *ec2.Tag) *InstanceTag {
	return &InstanceTag{
		AccountID: c.accountID,
		Region:    c.region,
		Key:       value.Key,
		Value:     value.Value,
	}
}

func (c *Client) transformInstanceTags(values []*ec2.Tag) []*InstanceTag {
	var tValues []*InstanceTag
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceTag(v))
	}
	return tValues
}

func (c *Client) transformInstance(value *ec2.Instance) *Instance {
	res := Instance{
		Region:                                  c.region,
		AccountID:                               c.accountID,
		AmiLaunchIndex:                          value.AmiLaunchIndex,
		Architecture:                            value.Architecture,
		BlockDeviceMappings:                     c.transformInstanceBlockDeviceMappings(value.BlockDeviceMappings),
		CapacityReservationId:                   value.CapacityReservationId,
		ClientToken:                             value.ClientToken,
		EbsOptimized:                            value.EbsOptimized,
		ElasticGpuAssociations:                  c.transformInstanceElasticGpuAssociations(value.ElasticGpuAssociations),
		ElasticInferenceAcceleratorAssociations: c.transformInstanceElasticInferenceAcceleratorAssociations(value.ElasticInferenceAcceleratorAssociations),
		EnaSupport:                              value.EnaSupport,
		Hypervisor:                              value.Hypervisor,
		ImageId:                                 value.ImageId,
		InstanceId:                              value.InstanceId,
		InstanceLifecycle:                       value.InstanceLifecycle,
		InstanceType:                            value.InstanceType,
		KernelId:                                value.KernelId,
		KeyName:                                 value.KeyName,
		LaunchTime:                              value.LaunchTime,
		Licenses:                                c.transformInstanceLicenseConfigurations(value.Licenses),
		NetworkInterfaces:                       c.transformInstanceNetworkInterfaces(value.NetworkInterfaces),
		OutpostArn:                              value.OutpostArn,
		Platform:                                value.Platform,
		PrivateDnsName:                          value.PrivateDnsName,
		PrivateIpAddress:                        value.PrivateIpAddress,
		ProductCodes:                            c.transformInstanceProductCodes(value.ProductCodes),
		PublicDnsName:                           value.PublicDnsName,
		PublicIpAddress:                         value.PublicIpAddress,
		RamdiskId:                               value.RamdiskId,
		RootDeviceName:                          value.RootDeviceName,
		RootDeviceType:                          value.RootDeviceType,
		SecurityGroups:                          c.transformInstanceGroupIdentifiers(value.SecurityGroups),
		SourceDestCheck:                         value.SourceDestCheck,
		SpotInstanceRequestId:                   value.SpotInstanceRequestId,
		SriovNetSupport:                         value.SriovNetSupport,
		StateTransitionReason:                   value.StateTransitionReason,
		SubnetId:                                value.SubnetId,
		Tags:                                    c.transformInstanceTags(value.Tags),
		VirtualizationType:                      value.VirtualizationType,
		VpcId:                                   value.VpcId,
	}

	if value.CpuOptions != nil {
		res.CpuOptionsCoreCount = value.CpuOptions.CoreCount
		res.CpuOptionsThreadsPerCore = value.CpuOptions.ThreadsPerCore
	}

	if value.CapacityReservationSpecification != nil {
		res.CapacityReservationPreference = value.CapacityReservationSpecification.CapacityReservationPreference
		if value.CapacityReservationSpecification.CapacityReservationTarget != nil {
			res.CapacityReservationTargetId = value.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationId
			res.CapacityReservationTargetGroupArn = value.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupArn
		}
	}

	if value.HibernationOptions != nil {
		res.HibernationOptionsConfigured = value.HibernationOptions.Configured
	}

	if value.IamInstanceProfile != nil {
		res.IamInstanceProfileArn = value.IamInstanceProfile.Arn
		res.IamInstanceProfileId = value.IamInstanceProfile.Id
	}

	if value.MetadataOptions != nil {
		res.MetadataOptionsHttpEndpoint = value.MetadataOptions.HttpEndpoint
		res.MetadataOptionsHttpTokens = value.MetadataOptions.HttpTokens
		res.MetadataOptionsState = value.MetadataOptions.State
		res.MetadataOptionsHttpPutResponseHopLimit = value.MetadataOptions.HttpPutResponseHopLimit
	}

	if value.Monitoring != nil {
		res.MonitoringState = value.Monitoring.State
	}

	if value.Placement != nil {
		res.PlacementAffinity = value.Placement.Affinity
		res.PlacementAvailabilityZone = value.Placement.AvailabilityZone
		res.PlacementGroupName = value.Placement.GroupName
		res.PlacementHostId = value.Placement.HostId
		res.PlacementHostResourceGroupArn = value.Placement.HostResourceGroupArn
		res.PlacementPartitionNumber = value.Placement.PartitionNumber
		res.PlacementSpreadDomain = value.Placement.SpreadDomain
		res.PlacementTenancy = value.Placement.Tenancy
	}

	if value.State != nil {
		res.StateCode = value.State.Code
		res.StateName = value.State.Name
	}

	if value.StateReason != nil {
		res.StateReasonCode = value.StateReason.Code
		res.StateReasonMessage = value.StateReason.Message
	}

	return &res
}

func (c *Client) transformInstances(values []*ec2.Instance) []*Instance {
	var tValues []*Instance
	for _, v := range values {
		tValues = append(tValues, c.transformInstance(v))
	}
	return tValues
}

var InstanceTables = []interface{}{
	&Instance{},
	&InstanceBlockDeviceMapping{},
	&InstanceElasticGpuAssociation{},
	&InstanceElasticInferenceAcceleratorAssociation{},
	&InstanceLicenseConfiguration{},
	&InstanceNetworkInterface{},
	&InstanceGroupIdentifier{},
	&InstanceIpv6Address{},
	&InstancePrivateIpAddress{},
	&InstanceProductCode{},
	&InstanceTag{},
}

func (c *Client) instances(gConfig interface{}) error {
	var config ec2.DescribeInstancesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(InstanceTables...)
	for {
		output, err := c.svc.DescribeInstances(&config)
		if err != nil {
			return err
		}
		for _, reservation := range output.Reservations {
			c.log.Info("Fetched resources", zap.String("resource", "ec2.instances"), zap.Int("count", len(reservation.Instances)))
			c.db.ChunkedCreate(c.transformInstances(reservation.Instances))
		}
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
