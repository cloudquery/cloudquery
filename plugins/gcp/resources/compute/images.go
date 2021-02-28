package compute

import (
	"google.golang.org/api/compute/v1"
)

type Image struct {
	_                     interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                    uint        `gorm:"primarykey"`
	ProjectID             string      `neo:"unique"`
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
	ResourceID            uint64                 `neo:"unique"`
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

func (Image) TableName() string {
	return "gcp_compute_images"
}

type ImageGuestOsFeature struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Type      string
}

func (ImageGuestOsFeature) TableName() string {
	return "gcp_compute_image_guest_os_features"
}

type ImageLicenseCode struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     int64
}

func (ImageLicenseCode) TableName() string {
	return "gcp_compute_image_license_codes"
}

type ImageLicense struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     string
}

func (ImageLicense) TableName() string {
	return "gcp_compute_image_licenses"
}

type ImageStorageLocation struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     string
}

func (ImageStorageLocation) TableName() string {
	return "gcp_compute_image_storage_locations"
}

func (c *Client) transformImageGuestOsFeature(value *compute.GuestOsFeature) *ImageGuestOsFeature {
	return &ImageGuestOsFeature{
		ProjectID: c.projectID,
		Type:      value.Type,
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
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformImageLicenses(values []string) []*ImageLicense {
	var tValues []*ImageLicense
	for _, v := range values {
		tValues = append(tValues, &ImageLicense{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformImageStorageLocations(values []string) []*ImageStorageLocation {
	var tValues []*ImageStorageLocation
	for _, v := range values {
		tValues = append(tValues, &ImageStorageLocation{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformImage(value *compute.Image) *Image {
	res := Image{
		Region:            c.region,
		ProjectID:         c.projectID,
		ArchiveSizeBytes:  value.ArchiveSizeBytes,
		CreationTimestamp: value.CreationTimestamp,
		Description:       value.Description,
		DiskSizeGb:        value.DiskSizeGb,
		Family:            value.Family,
		GuestOsFeatures:   c.transformImageGuestOsFeatures(value.GuestOsFeatures),
		ResourceID:        value.Id,
		Kind:              value.Kind,
		LabelFingerprint:  value.LabelFingerprint,
		LicenseCodes:      c.transformImageLicenseCodes(value.LicenseCodes),
		Licenses:          c.transformImageLicenses(value.Licenses),
		Name:              value.Name,
		SelfLink:          value.SelfLink,
		SourceDisk:        value.SourceDisk,
		SourceDiskId:      value.SourceDiskId,
		SourceImage:       value.SourceImage,
		SourceImageId:     value.SourceImageId,
		SourceSnapshot:    value.SourceSnapshot,
		SourceSnapshotId:  value.SourceSnapshotId,
		SourceType:        value.SourceType,
		Status:            value.Status,
		StorageLocations:  c.transformImageStorageLocations(value.StorageLocations),
	}

	if value.Deprecated != nil {
		res.DeprecatedDeleted = value.Deprecated.Deleted
		res.DeprecatedDeprecated = value.Deprecated.Deprecated
		res.DeprecatedObsolete = value.Deprecated.Obsolete
		res.DeprecatedReplacement = value.Deprecated.Replacement
		res.DeprecatedState = value.Deprecated.State
	}

	if value.RawDisk != nil {
		res.RawDiskContainerType = value.RawDisk.ContainerType
		res.RawDiskSha1Checksum = value.RawDisk.Sha1Checksum
		res.RawDiskSource = value.RawDisk.Source
	}

	return &res
}

func (c *Client) transformImages(values []*compute.Image) []*Image {
	var tValues []*Image
	for _, v := range values {
		tValues = append(tValues, c.transformImage(v))
	}
	return tValues
}

var ImageTables = []interface{}{
	&Image{},
	&ImageGuestOsFeature{},
	&ImageLicenseCode{},
	&ImageLicense{},
	&ImageStorageLocation{},
}

func (c *Client) images(_ interface{}) error {

	c.db.Where("project_id", c.projectID).Delete(ImageTables...)
	nextPageToken := ""
	for {
		call := c.svc.Images.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformImages(output.Items))
		c.log.Info("Fetched resources", "resource", "compute.images", "count", len(output.Items))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
