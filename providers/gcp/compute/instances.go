package compute

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Instance struct {
	ID                                                  uint `gorm:"primarykey"`
	ProjectID                                           string
	CanIpForward                                        bool
	ConfidentialInstanceConfigEnableConfidentialCompute bool
	CpuPlatform                                         string
	CreationTimestamp                                   string
	DeletionProtection                                  bool
	Description                                         string
	Disks                                               []*InstanceAttachedDisk `gorm:"constraint:OnDelete:CASCADE;"`
	DisplayDeviceEnableDisplay                          bool
	Fingerprint                                         string
	GuestAccelerators                                   []*InstanceAcceleratorConfig `gorm:"constraint:OnDelete:CASCADE;"`
	Hostname                                            string
	ResourceID                                          uint64
	Kind                                                string
	LabelFingerprint                                    string
	//Labels                            []*InstanceLabels `gorm:"constraint:OnDelete:CASCADE;"`
	LastStartTimestamp      string
	LastStopTimestamp       string
	LastSuspendedTimestamp  string
	MachineType             string
	Metadata                []*InstanceMetadataItem `gorm:"constraint:OnDelete:CASCADE;"`
	MinCpuPlatform          string
	Name                    string
	NetworkInterfaces       []*InstanceNetworkInterface `gorm:"constraint:OnDelete:CASCADE;"`
	PrivateIpv6GoogleAccess string

	ReservationAffinityConsumeReservationType string
	ReservationAffinityKey                    string
	ReservationAffinityValues                 []*InstanceReservationAffinityValue `gorm:"constraint:OnDelete:CASCADE;"`
	ResourcePolicies                          []*InstanceResourcePolicy           `gorm:"constraint:OnDelete:CASCADE;"`

	SchedulingAutomaticRestart  *bool
	SchedulingMinNodeCpus       int64
	SchedulingNodeAffinities    []*InstanceSchedulingNodeAffinity `gorm:"constraint:OnDelete:CASCADE;"`
	SchedulingOnHostMaintenance string
	SchedulingPreemptible       bool

	SelfLink        string
	ServiceAccounts []*InstanceServiceAccount `gorm:"constraint:OnDelete:CASCADE;"`

	ShieldedInstanceConfigEnableIntegrityMonitoring      bool
	ShieldedInstanceConfigEnableSecureBoot               bool
	ShieldedInstanceConfigEnableVtpm                     bool
	ShieldedInstanceIntegrityPolicyUpdateAutoLearnPolicy bool

	StartRestricted bool
	Status          string
	StatusMessage   string
	Tags            []*InstanceTag `gorm:"constraint:OnDelete:CASCADE;"`
	Zone            string
}

type InstanceAttachedDisk struct {
	ID                    uint `gorm:"primarykey"`
	InstanceID            uint
	AutoDelete            bool
	Boot                  bool
	DeviceName            string
	DiskSizeGb            int64
	GuestOsFeatures       []*InstanceGuestOsFeature `gorm:"constraint:OnDelete:CASCADE;"`
	Index                 int64
	InitializeDescription string
	InitializeDiskName    string
	InitializeDiskSizeGb  int64
	InitializeDiskType    string
	//InitializeLabels           []*InstanceAttachedDiskInitializeLabel `gorm:"constraint:OnDelete:CASCADE;"`
	InitializeOnUpdateAction   string
	InitializeResourcePolicies []*InstanceAttachedDiskInitializeResourcePolicy `gorm:"constraint:OnDelete:CASCADE;"`
	InitializeSourceImage      string
	InitializeSourceSnapshot   string
	Interface                  string
	Kind                       string
	Licenses                   []*InstanceAttachedDiskLicense `gorm:"constraint:OnDelete:CASCADE;"`
	Mode                       string
	Source                     string
	Type                       string
}

type InstanceGuestOsFeature struct {
	ID                     uint `gorm:"primarykey"`
	InstanceAttachedDiskID uint
	Type                   string
}

type InstanceAttachedDiskInitializeResourcePolicy struct {
	ID                     uint `gorm:"primarykey"`
	InstanceAttachedDiskID uint
	Value                  string
}

type InstanceAttachedDiskLicense struct {
	ID                     uint `gorm:"primarykey"`
	InstanceAttachedDiskID uint
	Value                  string
}

type InstanceAcceleratorConfig struct {
	ID               uint `gorm:"primarykey"`
	InstanceID       uint
	AcceleratorCount int64
	AcceleratorType  string
}

type InstanceMetadataItem struct {
	ID          uint `gorm:"primarykey"`
	InstanceID  uint
	Fingerprint string
	Key         string
	Value       *string
	Kind        string
}

type InstanceNetworkInterface struct {
	ID            uint `gorm:"primarykey"`
	InstanceID    uint
	AccessConfigs []*InstanceAccessConfig `gorm:"constraint:OnDelete:CASCADE;"`
	AliasIpRanges []*InstanceAliasIpRange `gorm:"constraint:OnDelete:CASCADE;"`
	Fingerprint   string
	Ipv6Address   string
	Kind          string
	Name          string
	Network       string
	NetworkIP     string
	Subnetwork    string
}

type InstanceAccessConfig struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint
	Kind                       string
	Name                       string
	NatIP                      string
	NetworkTier                string
	PublicPtrDomainName        string
	SetPublicPtr               bool
	Type                       string
}

type InstanceAliasIpRange struct {
	ID                         uint `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint
	IpCidrRange                string
	SubnetworkRangeName        string
}

type InstanceReservationAffinityValue struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	Value      string
}
type InstanceResourcePolicy struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	Value      string
}

type InstanceSchedulingNodeAffinity struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	Key        string
	Operator   string
	Value      string
}

type InstanceServiceAccount struct {
	ID         uint `gorm:"primarykey"`
	InstanceID uint
	Email      string
	Scope      string
}

type InstanceTag struct {
	ID          uint `gorm:"primarykey"`
	InstanceID  uint
	Fingerprint string
	Value       string `gorm:"constraint:OnDelete:CASCADE;"`
}

func (c *Client) transformInstanceGuestOsFeatures(values []*compute.GuestOsFeature) []*InstanceGuestOsFeature {
	var tValues []*InstanceGuestOsFeature
	for _, v := range values {
		tValues = append(tValues, &InstanceGuestOsFeature{
			Type: v.Type,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAttachedDiskLicenses(values []string) []*InstanceAttachedDiskLicense {
	var tValues []*InstanceAttachedDiskLicense
	for _, v := range values {
		tValues = append(tValues, &InstanceAttachedDiskLicense{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAttachedDisk(value *compute.AttachedDisk) *InstanceAttachedDisk {
	return &InstanceAttachedDisk{
		AutoDelete:      value.AutoDelete,
		Boot:            value.Boot,
		DeviceName:      value.DeviceName,
		DiskSizeGb:      value.DiskSizeGb,
		GuestOsFeatures: c.transformInstanceGuestOsFeatures(value.GuestOsFeatures),
		Index:           value.Index,
		Interface:       value.Interface,
		Kind:            value.Kind,
		Licenses:        c.transformInstanceAttachedDiskLicenses(value.Licenses),
		Mode:            value.Mode,
		Source:          value.Source,
		Type:            value.Type,
	}
}

func (c *Client) transformInstanceAttachedDisks(values []*compute.AttachedDisk) []*InstanceAttachedDisk {
	var tValues []*InstanceAttachedDisk
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAttachedDisk(v))
	}
	return tValues
}

func (c *Client) transformInstanceAcceleratorConfig(value *compute.AcceleratorConfig) *InstanceAcceleratorConfig {
	return &InstanceAcceleratorConfig{
		AcceleratorCount: value.AcceleratorCount,
		AcceleratorType:  value.AcceleratorType,
	}
}

func (c *Client) transformInstanceAcceleratorConfigs(values []*compute.AcceleratorConfig) []*InstanceAcceleratorConfig {
	var tValues []*InstanceAcceleratorConfig
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAcceleratorConfig(v))
	}
	return tValues
}

func (c *Client) transformInstanceMetadataItems(value *compute.Metadata) []*InstanceMetadataItem {
	var tValues []*InstanceMetadataItem
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceMetadataItem{
			Fingerprint: value.Fingerprint,
			Key:         v.Key,
			Value:       v.Value,
			Kind:        value.Kind,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAccessConfig(value *compute.AccessConfig) *InstanceAccessConfig {
	return &InstanceAccessConfig{
		Kind:                value.Kind,
		Name:                value.Name,
		NatIP:               value.NatIP,
		NetworkTier:         value.NetworkTier,
		PublicPtrDomainName: value.PublicPtrDomainName,
		SetPublicPtr:        value.SetPublicPtr,
		Type:                value.Type,
	}
}

func (c *Client) transformInstanceAccessConfigs(values []*compute.AccessConfig) []*InstanceAccessConfig {
	var tValues []*InstanceAccessConfig
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAccessConfig(v))
	}
	return tValues
}

func (c *Client) transformInstanceAliasIpRange(value *compute.AliasIpRange) *InstanceAliasIpRange {
	return &InstanceAliasIpRange{
		IpCidrRange:         value.IpCidrRange,
		SubnetworkRangeName: value.SubnetworkRangeName,
	}
}

func (c *Client) transformInstanceAliasIpRanges(values []*compute.AliasIpRange) []*InstanceAliasIpRange {
	var tValues []*InstanceAliasIpRange
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAliasIpRange(v))
	}
	return tValues
}

func (c *Client) transformInstanceNetworkInterface(value *compute.NetworkInterface) *InstanceNetworkInterface {
	return &InstanceNetworkInterface{
		AccessConfigs: c.transformInstanceAccessConfigs(value.AccessConfigs),
		AliasIpRanges: c.transformInstanceAliasIpRanges(value.AliasIpRanges),
		Fingerprint:   value.Fingerprint,
		Ipv6Address:   value.Ipv6Address,
		Kind:          value.Kind,
		Name:          value.Name,
		Network:       value.Network,
		NetworkIP:     value.NetworkIP,
		Subnetwork:    value.Subnetwork,
	}
}

func (c *Client) transformInstanceNetworkInterfaces(values []*compute.NetworkInterface) []*InstanceNetworkInterface {
	var tValues []*InstanceNetworkInterface
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceNetworkInterface(v))
	}
	return tValues
}

func (c *Client) transformInstanceReservationAffinityValues(values []string) []*InstanceReservationAffinityValue {
	var tValues []*InstanceReservationAffinityValue
	for _, v := range values {
		tValues = append(tValues, &InstanceReservationAffinityValue{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceResourcePolicies(values []string) []*InstanceResourcePolicy {
	var tValues []*InstanceResourcePolicy
	for _, v := range values {
		tValues = append(tValues, &InstanceResourcePolicy{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceSchedulingNodeAffinities(values []*compute.SchedulingNodeAffinity) []*InstanceSchedulingNodeAffinity {
	var tValues []*InstanceSchedulingNodeAffinity
	for _, v := range values {
		for _, v1 := range v.Values {
			tValues = append(tValues, &InstanceSchedulingNodeAffinity{
				Key:      v.Key,
				Operator: v.Operator,
				Value:    v1,
			})
		}
	}
	return tValues
}

func (c *Client) transformInstanceServiceAccounts(values []*compute.ServiceAccount) []*InstanceServiceAccount {
	var tValues []*InstanceServiceAccount
	for _, v := range values {
		for _, scope := range v.Scopes {
			tValues = append(tValues, &InstanceServiceAccount{
				Email: v.Email,
				Scope: scope,
			})
		}
	}
	return tValues
}

func (c *Client) transformTags(value *compute.Tags) []*InstanceTag {
	var tValues []*InstanceTag
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceTag{
			Fingerprint: value.Fingerprint,
			Value:       v,
		})
	}
	return tValues
}

func (c *Client) transformMetadataItems(value *compute.Metadata) []*InstanceMetadataItem {
	var tValues []*InstanceMetadataItem
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceMetadataItem{
			Fingerprint: value.Fingerprint,
			Key:         v.Key,
			Value:       v.Value,
			Kind:        value.Kind,
		})
	}
	return tValues
}

func (c *Client) transformInstance(value *compute.Instance) *Instance {
	res := Instance{
		ProjectID:    c.projectID,
		CanIpForward: value.CanIpForward,

		CpuPlatform:                value.CpuPlatform,
		CreationTimestamp:          value.CreationTimestamp,
		DeletionProtection:         value.DeletionProtection,
		Description:                value.Description,
		Disks:                      c.transformInstanceAttachedDisks(value.Disks),
		Fingerprint:                value.Fingerprint,
		GuestAccelerators:          c.transformInstanceAcceleratorConfigs(value.GuestAccelerators),
		Hostname:                   value.Hostname,
		ResourceID:                 value.Id,
		Kind:                       value.Kind,
		LabelFingerprint:           value.LabelFingerprint,
		LastStartTimestamp:         value.LastStartTimestamp,
		LastStopTimestamp:          value.LastStopTimestamp,
		LastSuspendedTimestamp:     value.LastSuspendedTimestamp,
		MachineType:                value.MachineType,
		Metadata:                   c.transformMetadataItems(value.Metadata),
		MinCpuPlatform:             value.MinCpuPlatform,
		Name:                       value.Name,
		NetworkInterfaces:          c.transformInstanceNetworkInterfaces(value.NetworkInterfaces),
		PrivateIpv6GoogleAccess:    value.PrivateIpv6GoogleAccess,

		ResourcePolicies: c.transformInstanceResourcePolicies(value.ResourcePolicies),

		SelfLink:        value.SelfLink,
		ServiceAccounts: c.transformInstanceServiceAccounts(value.ServiceAccounts),

		StartRestricted: value.StartRestricted,
		Status:          value.Status,
		StatusMessage:   value.StatusMessage,
		Tags:            c.transformTags(value.Tags),
		Zone:            value.Zone,
	}

	if value.DisplayDevice != nil {
		res.DisplayDeviceEnableDisplay = value.DisplayDevice.EnableDisplay
	}

	if value.ConfidentialInstanceConfig != nil {
		res.ConfidentialInstanceConfigEnableConfidentialCompute = value.ConfidentialInstanceConfig.EnableConfidentialCompute
	}

	if value.ReservationAffinity != nil {
		res.ReservationAffinityKey = value.ReservationAffinity.Key
		res.ReservationAffinityConsumeReservationType = value.ReservationAffinity.ConsumeReservationType
		res.ReservationAffinityValues = c.transformInstanceReservationAffinityValues(value.ReservationAffinity.Values)
	}
	if value.Scheduling != nil {
		res.SchedulingAutomaticRestart = value.Scheduling.AutomaticRestart
		res.SchedulingMinNodeCpus = value.Scheduling.MinNodeCpus
		res.SchedulingOnHostMaintenance = value.Scheduling.OnHostMaintenance
		res.SchedulingNodeAffinities = c.transformInstanceSchedulingNodeAffinities(value.Scheduling.NodeAffinities)
	}

	if value.ShieldedInstanceConfig != nil {
		res.ShieldedInstanceConfigEnableIntegrityMonitoring = value.ShieldedInstanceConfig.EnableIntegrityMonitoring
		res.ShieldedInstanceConfigEnableSecureBoot = value.ShieldedInstanceConfig.EnableSecureBoot
		res.ShieldedInstanceConfigEnableVtpm = value.ShieldedInstanceConfig.EnableVtpm
		res.ShieldedInstanceIntegrityPolicyUpdateAutoLearnPolicy = value.ShieldedInstanceIntegrityPolicy.UpdateAutoLearnPolicy
	}

	return &res
}

func (c *Client) transformInstances(values []*compute.Instance) []*Instance {
	var tValues []*Instance
	for _, v := range values {
		tValues = append(tValues, c.transformInstance(v))
	}
	return tValues
}

type InstanceConfig struct {
	Filter string
}

func (c *Client) instances(gConfig interface{}) error {
	var config InstanceConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeInstance"] {
		err := c.db.AutoMigrate(
			&Instance{},
			&InstanceAttachedDisk{},
			&InstanceAttachedDiskLicense{},
			&InstanceGuestOsFeature{},
			&InstanceAcceleratorConfig{},
			&InstanceMetadataItem{},
			&InstanceNetworkInterface{},
			&InstanceAccessConfig{},
			&InstanceAliasIpRange{},
			&InstanceSchedulingNodeAffinity{},
			&InstanceServiceAccount{},
			&InstanceTag{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeInstance"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.Instances.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&Instance{})
		var tValues []*Instance
		for _, items := range output.Items {
			tValues = append(tValues, c.transformInstances(items.Instances)...)
		}
		common.ChunkedCreate(c.db, tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.instances"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
