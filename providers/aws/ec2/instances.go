package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"time"
)

type Instance struct {
	ID                                      uint `gorm:"primarykey"`
	AccountID                               string
	Region                                  string
	AmiLaunchIndex                          *int64
	Architecture                            *string
	BlockDeviceMappings                     []*InstanceBlockDeviceMapping `gorm:"constraint:OnDelete:CASCADE;"`
	CapacityReservationId                   *string
	CapacityReservationSpecification        *InstanceCapacityReservationSpecificationResponse `gorm:"constraint:OnDelete:CASCADE;"`
	ClientToken                             *string
	CpuOptions                              *ec2.CpuOptions `gorm:"embedded;embeddedPrefix:cpu_options_"`
	EbsOptimized                            *bool
	ElasticGpuAssociations                  []*InstanceElasticGpuAssociation                  `gorm:"constraint:OnDelete:CASCADE;"`
	ElasticInferenceAcceleratorAssociations []*InstanceElasticInferenceAcceleratorAssociation `gorm:"constraint:OnDelete:CASCADE;"`
	EnaSupport                              *bool
	HibernationOptions                      *ec2.HibernationOptions `gorm:"embedded;embeddedPrefix:hibernation_options_"`
	Hypervisor                              *string
	IamInstanceProfile                      *ec2.IamInstanceProfile `gorm:"embedded;embeddedPrefix:iam_instance_profile_"`
	ImageId                                 *string
	InstanceId                              *string
	InstanceLifecycle                       *string
	InstanceType                            *string
	KernelId                                *string
	KeyName                                 *string
	LaunchTime                              *time.Time
	Licenses                                []*InstanceLicenseConfiguration      `gorm:"constraint:OnDelete:CASCADE;"`
	MetadataOptions                         *ec2.InstanceMetadataOptionsResponse `gorm:"embedded;embeddedPrefix:metadata_options_"`
	Monitoring                              *ec2.Monitoring                      `gorm:"embedded;embeddedPrefix:monitoring_"`
	NetworkInterfaces                       []*InstanceNetworkInterface          `gorm:"constraint:OnDelete:CASCADE;"`
	OutpostArn                              *string
	Placement                               *ec2.Placement `gorm:"embedded;embeddedPrefix:placement_"`
	Platform                                *string
	PrivateDnsName                          *string
	PrivateIpAddress                        *string
	ProductCodes                            []*InstanceProductCode `gorm:"constraint:OnDelete:CASCADE;"`
	PublicDnsName                           *string
	PublicIpAddress                         *string
	RamdiskId                               *string
	RootDeviceName                          *string
	RootDeviceType                          *string
	SecurityGroups                          []*InstanceGroupIdentifier `gorm:"constraint:OnDelete:CASCADE;"`
	SourceDestCheck                         *bool
	SpotInstanceRequestId                   *string
	SriovNetSupport                         *string
	State                                   *ec2.InstanceState `gorm:"embedded;embeddedPrefix:state_"`
	StateReason                             *ec2.StateReason   `gorm:"embedded;embeddedPrefix:state_reason_"`
	StateTransitionReason                   *string
	SubnetId                                *string
	Tags                                    []*InstanceTag `gorm:"constraint:OnDelete:CASCADE;"`
	VirtualizationType                      *string
	VpcId                                   *string
}

type InstanceBlockDeviceMapping struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	DeviceName *string
	Ebs        *ec2.EbsInstanceBlockDevice `gorm:"embedded;embeddedPrefix:ebs_"`
}

type InstanceCapacityReservationSpecificationResponse struct {
	ID                            uint `gorm:"primarykey"`
	InstanceID                    uint
	CapacityReservationPreference *string
	CapacityReservationTarget     *ec2.CapacityReservationTargetResponse `gorm:"embedded"`
}

type InstanceElasticGpuAssociation struct {
	ID                         uint `gorm:"primarykey"`
	InstanceID                 uint
	ElasticGpuAssociationId    *string
	ElasticGpuAssociationState *string
	ElasticGpuAssociationTime  *string
	ElasticGpuId               *string
}

type InstanceElasticInferenceAcceleratorAssociation struct {
	ID                                          uint `gorm:"primarykey"`
	InstanceID                                  uint
	ElasticInferenceAcceleratorArn              *string
	ElasticInferenceAcceleratorAssociationId    *string
	ElasticInferenceAcceleratorAssociationState *string
	ElasticInferenceAcceleratorAssociationTime  *time.Time
}

type InstanceLicenseConfiguration struct {
	ID                      uint `gorm:"primarykey"`
	InstanceID              uint
	LicenseConfigurationArn *string
}

type InstanceNetworkInterface struct {
	ID                 uint `gorm:"primarykey"`
	InstanceID         uint
	Association        *ec2.InstanceNetworkInterfaceAssociation `gorm:"embedded;embeddedPrefix:association_"`
	Attachment         *ec2.InstanceNetworkInterfaceAttachment  `gorm:"embedded;embeddedPrefix:attachment_"`
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

type InstanceGroupIdentifier struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	GroupId    *string
	GroupName  *string
}

type InstanceIpv6Address struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint
	Ipv6Address                *string
}

type InstancePrivateIpAddress struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint
	Association                *ec2.InstanceNetworkInterfaceAssociation `gorm:"embedded;embeddedPrefix:association_"`
	Primary                    *bool
	PrivateDnsName             *string
	PrivateIpAddress           *string
}

type InstanceProductCode struct {
	ID              uint `gorm:"primarykey"`
	InstanceID      uint
	ProductCodeId   *string
	ProductCodeType *string
}

type InstanceTag struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	Key        *string
	Value      *string
}

func (c *Client) transformInstanceBlockDeviceMapping(value *ec2.InstanceBlockDeviceMapping) *InstanceBlockDeviceMapping {
	return &InstanceBlockDeviceMapping{
		DeviceName: value.DeviceName,
		Ebs:        value.Ebs,
	}
}

func (c *Client) transformInstanceBlockDeviceMappings(values []*ec2.InstanceBlockDeviceMapping) []*InstanceBlockDeviceMapping {
	var tValues []*InstanceBlockDeviceMapping
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceBlockDeviceMapping(v))
	}
	return tValues
}

func (c *Client) transformInstanceCapacityReservationSpecificationResponse(value *ec2.CapacityReservationSpecificationResponse) *InstanceCapacityReservationSpecificationResponse {
	return &InstanceCapacityReservationSpecificationResponse{
		CapacityReservationPreference: value.CapacityReservationPreference,
		CapacityReservationTarget:     value.CapacityReservationTarget,
	}
}

func (c *Client) transformInstanceElasticGpuAssociation(value *ec2.ElasticGpuAssociation) *InstanceElasticGpuAssociation {
	return &InstanceElasticGpuAssociation{
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
		ElasticInferenceAcceleratorArn:              value.ElasticInferenceAcceleratorArn,
		ElasticInferenceAcceleratorAssociationId:    value.ElasticInferenceAcceleratorAssociationId,
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
	return &InstancePrivateIpAddress{
		Association:      value.Association,
		Primary:          value.Primary,
		PrivateDnsName:   value.PrivateDnsName,
		PrivateIpAddress: value.PrivateIpAddress,
	}
}

func (c *Client) transformInstancePrivateIpAddresss(values []*ec2.InstancePrivateIpAddress) []*InstancePrivateIpAddress {
	var tValues []*InstancePrivateIpAddress
	for _, v := range values {
		tValues = append(tValues, c.transformInstancePrivateIpAddress(v))
	}
	return tValues
}

func (c *Client) transformInstanceNetworkInterface(value *ec2.InstanceNetworkInterface) *InstanceNetworkInterface {
	return &InstanceNetworkInterface{
		Association:        value.Association,
		Attachment:         value.Attachment,
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
		Key:   value.Key,
		Value: value.Value,
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
	return &Instance{
		Region:                                  c.region,
		AccountID:                               c.accountID,
		AmiLaunchIndex:                          value.AmiLaunchIndex,
		Architecture:                            value.Architecture,
		BlockDeviceMappings:                     c.transformInstanceBlockDeviceMappings(value.BlockDeviceMappings),
		CapacityReservationId:                   value.CapacityReservationId,
		CapacityReservationSpecification:        c.transformInstanceCapacityReservationSpecificationResponse(value.CapacityReservationSpecification),
		ClientToken:                             value.ClientToken,
		CpuOptions:                              value.CpuOptions,
		EbsOptimized:                            value.EbsOptimized,
		ElasticGpuAssociations:                  c.transformInstanceElasticGpuAssociations(value.ElasticGpuAssociations),
		ElasticInferenceAcceleratorAssociations: c.transformInstanceElasticInferenceAcceleratorAssociations(value.ElasticInferenceAcceleratorAssociations),
		EnaSupport:                              value.EnaSupport,
		HibernationOptions:                      value.HibernationOptions,
		Hypervisor:                              value.Hypervisor,
		IamInstanceProfile:                      value.IamInstanceProfile,
		ImageId:                                 value.ImageId,
		InstanceId:                              value.InstanceId,
		InstanceLifecycle:                       value.InstanceLifecycle,
		InstanceType:                            value.InstanceType,
		KernelId:                                value.KernelId,
		KeyName:                                 value.KeyName,
		LaunchTime:                              value.LaunchTime,
		Licenses:                                c.transformInstanceLicenseConfigurations(value.Licenses),
		MetadataOptions:                         value.MetadataOptions,
		Monitoring:                              value.Monitoring,
		NetworkInterfaces:                       c.transformInstanceNetworkInterfaces(value.NetworkInterfaces),
		OutpostArn:                              value.OutpostArn,
		Placement:                               value.Placement,
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
		State:                                   value.State,
		StateReason:                             value.StateReason,
		StateTransitionReason:                   value.StateTransitionReason,
		SubnetId:                                value.SubnetId,
		Tags:                                    c.transformInstanceTags(value.Tags),
		VirtualizationType:                      value.VirtualizationType,
		VpcId:                                   value.VpcId,
	}
}

func (c *Client) transformInstances(values []*ec2.Instance) []*Instance {
	var tValues []*Instance
	for _, v := range values {
		tValues = append(tValues, c.transformInstance(v))
	}
	return tValues
}

func (c *Client) instances(gConfig interface{}) error {
	var config ec2.DescribeInstancesInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["ec2Instance"] {
		err := c.db.AutoMigrate(
			&Instance{},
			&InstanceBlockDeviceMapping{},
			&InstanceCapacityReservationSpecificationResponse{},
			&InstanceElasticGpuAssociation{},
			&InstanceElasticInferenceAcceleratorAssociation{},
			&InstanceLicenseConfiguration{},
			&InstanceNetworkInterface{},
			&InstanceGroupIdentifier{},
			&InstanceIpv6Address{},
			&InstancePrivateIpAddress{},
			&InstanceProductCode{},
			&InstanceTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["ec2Instance"] = true
	}
	for {
		output, err := c.svc.DescribeInstances(&config)
		if err != nil {
			return err
		}
		c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Instance{})
		for _, reservation := range output.Reservations {
			c.log.Info("Fetched resources", zap.Int("count", len(reservation.Instances)))
			common.ChunkedCreate(c.db, c.transformInstances(reservation.Instances))
		}
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
