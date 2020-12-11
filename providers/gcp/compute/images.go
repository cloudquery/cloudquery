package compute

import (
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Image struct {
	ID                    uint `gorm:"primarykey"`
	ProjectID             string
	Region                string
	ArchiveSizeBytes      int64
	CreationTimestamp     string
	DeprecatedDeleted     string
	DeprecatedDeprecated  string
	DeprecatedObsolete    string
	DeprecatedReplacement string
	DeprecatedState       string
	Description           string
	DiskSizeGb            int64
	Family                string
	GuestOsFeatures       []*ImageGuestOsFeature `gorm:"constraint:OnDelete:CASCADE;"`
	ResourceID            uint64
	Id                    uint64
	Kind                  string
	LabelFingerprint      string
	//Labels                               []*ImageLabel `gorm:"constraint:OnDelete:CASCADE;"`
	LicenseCodes         []*ImageLicenseCode `gorm:"constraint:OnDelete:CASCADE;"`
	Licenses             []*ImageLicense     `gorm:"constraint:OnDelete:CASCADE;"`
	Name                 string
	RawDiskContainerType string
	RawDiskSha1Checksum  string
	RawDiskSource        string
	SelfLink             string
	SourceDisk           string
	SourceDiskId         string
	SourceImage          string
	SourceImageId        string
	SourceSnapshot       string
	SourceSnapshotId     string
	SourceType           string
	Status               string
	StorageLocations     []*ImageStorageLocation `gorm:"constraint:OnDelete:CASCADE;"`
}

type ImageGuestOsFeature struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint
	Type    string
}

type ImageLicenseCode struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint
	Value   int64
}
type ImageLicense struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint
	Value   string
}

type ImageStorageLocation struct {
	ID      uint `gorm:"primarykey"`
	ImageID uint
	Value   string
}

func (c *Client) transformImageGuestOsFeature(value *compute.GuestOsFeature) *ImageGuestOsFeature {
	return &ImageGuestOsFeature{
		Type: value.Type,
	}
}

func (c *Client) transformImageGuestOsFeatures(values []*compute.GuestOsFeature) []*ImageGuestOsFeature {
	var tValues []*ImageGuestOsFeature
	for _, v := range values {
		tValues = append(tValues, c.transformImageGuestOsFeature(v))
	}
	return tValues
}

func (c *Client) transformImageLicenseCodes(values []int64) []*ImageLicenseCode {
	var tValues []*ImageLicenseCode
	for _, v := range values {
		tValues = append(tValues, &ImageLicenseCode{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformImageLicenses(values []string) []*ImageLicense {
	var tValues []*ImageLicense
	for _, v := range values {
		tValues = append(tValues, &ImageLicense{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformImageStorageLocations(values []string) []*ImageStorageLocation {
	var tValues []*ImageStorageLocation
	for _, v := range values {
		tValues = append(tValues, &ImageStorageLocation{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformImage(value *compute.Image) *Image {
	return &Image{
		Region:                c.region,
		ProjectID:             c.projectID,
		ArchiveSizeBytes:      value.ArchiveSizeBytes,
		CreationTimestamp:     value.CreationTimestamp,
		DeprecatedDeleted:     value.Deprecated.Deleted,
		DeprecatedDeprecated:  value.Deprecated.Deprecated,
		DeprecatedObsolete:    value.Deprecated.Obsolete,
		DeprecatedReplacement: value.Deprecated.Replacement,
		DeprecatedState:       value.Deprecated.State,
		Description:           value.Description,
		DiskSizeGb:            value.DiskSizeGb,
		Family:                value.Family,
		GuestOsFeatures:       c.transformImageGuestOsFeatures(value.GuestOsFeatures),
		ResourceID:            value.Id,
		Id:                    value.Id,
		Kind:                  value.Kind,
		LabelFingerprint:      value.LabelFingerprint,
		LicenseCodes:          c.transformImageLicenseCodes(value.LicenseCodes),
		Licenses:              c.transformImageLicenses(value.Licenses),
		Name:                  value.Name,
		RawDiskContainerType:  value.RawDisk.ContainerType,
		RawDiskSha1Checksum:   value.RawDisk.Sha1Checksum,
		RawDiskSource:         value.RawDisk.Source,
		SelfLink:              value.SelfLink,
		SourceDisk:            value.SourceDisk,
		SourceDiskId:          value.SourceDiskId,
		SourceImage:           value.SourceImage,
		SourceImageId:         value.SourceImageId,
		SourceSnapshot:        value.SourceSnapshot,
		SourceSnapshotId:      value.SourceSnapshotId,
		SourceType:            value.SourceType,
		Status:                value.Status,
		StorageLocations:      c.transformImageStorageLocations(value.StorageLocations),
	}
}

func (c *Client) transformImages(values []*compute.Image) []*Image {
	var tValues []*Image
	for _, v := range values {
		tValues = append(tValues, c.transformImage(v))
	}
	return tValues
}

func (c *Client) images(gConfig interface{}) error {
	var config elb.DescribeLoadBalancersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeImage"] {
		err := c.db.AutoMigrate(
			&Image{},
			&ImageGuestOsFeature{},
			&ImageLicenseCode{},
			&ImageLicense{},
			&ImageStorageLocation{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeImage"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.Images.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&Image{})
		common.ChunkedCreate(c.db, c.transformImages(output.Items))
		c.log.Info("populating images", zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
