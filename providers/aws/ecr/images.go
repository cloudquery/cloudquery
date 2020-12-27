package ecr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/cloudquery/cloudquery/providers/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)


type Image struct {
	ID uint `gorm:"primarykey"`
	AccountID string
	Region string

	ArtifactMediaType *string
	ImageDigest *string
	ImageManifestMediaType *string
	ImagePushedAt *time.Time

	ImageScanFindingsSeverityCounts []*ImageSeverityCount `gorm:"constraint:OnDelete:CASCADE;"`
	ImageScanFindingsImageScanCompletedAt *time.Time
	ImageScanFindingsVulnerabilitySourceUpdatedAt *time.Time

	ImageScanStatusDescription *string
	ImageScanStatusStatus *string
	ImageSizeInBytes *int64
	ImageTags []*ImageTags `gorm:"constraint:OnDelete:CASCADE;"`
	RegistryId *string
	RepositoryName *string
}

func (Image) TableName() string {
	return "aws_ecr_images"
}

type ImageSeverityCount struct {
	ID uint
	ImageID uint
	Severity string
	Count *int64
}

func (ImageSeverityCount) TableName() string {
	return "aws_ecr_image_severity_counts"
}

type ImageTags struct {
	ID uint `gorm:"primarykey"`
	ImageID uint
	Value *string
}

func (ImageTags) TableName() string {
	return "aws_ecr_image_tags"
}


func (c *Client) transformImages(values []*ecr.ImageDetail) []*Image {
	var tValues []*Image
	for _, value := range values {
		tValue := Image{
			AccountID: c.accountID,
			Region: c.region,
			ArtifactMediaType: value.ArtifactMediaType,
			ImageDigest: value.ImageDigest,
			ImageManifestMediaType: value.ImageManifestMediaType,
			ImagePushedAt: value.ImagePushedAt,
			ImageSizeInBytes: value.ImageSizeInBytes,
			ImageTags: c.transformImageTags(value.ImageTags),
			RegistryId: value.RegistryId,
			RepositoryName: value.RepositoryName,
		}

		if value.ImageScanFindingsSummary != nil {
			tValue.ImageScanFindingsImageScanCompletedAt = value.ImageScanFindingsSummary.ImageScanCompletedAt
			tValue.ImageScanFindingsVulnerabilitySourceUpdatedAt = value.ImageScanFindingsSummary.VulnerabilitySourceUpdatedAt
			for severity, count := range value.ImageScanFindingsSummary.FindingSeverityCounts {
				tValue.ImageScanFindingsSeverityCounts = append(tValue.ImageScanFindingsSeverityCounts,
					&ImageSeverityCount{
						Severity: severity,
						Count:    count,
					})
			}
		}

		if value.ImageScanStatus != nil {
			tValue.ImageScanStatusDescription = value.ImageScanStatus.Description
			tValue.ImageScanStatusStatus = value.ImageScanStatus.Status
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformImageTags(values []*string) []*ImageTags {
	var tValues []*ImageTags
	for _, v := range values {
		tValues = append(tValues, &ImageTags{
			Value: v,
		})
	}
	return tValues
}

type ImageConfig struct {
	Filter string
}

func MigrateImage(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Image{},
		&ImageSeverityCount{},
		&ImageTags{},
	)
	if err != nil {
		return err
	}

	return nil
}


func (c *Client) images(_ interface{}) error {
	c.db.Where("region = ?", c.region).Where("account_id = ?", c.accountID).Delete(&Image{})

	var describeRepositoriesInput ecr.DescribeRepositoriesInput
	var describeImagesInput ecr.DescribeImagesInput
	var maxResults int64
	maxResults = 1000
	describeRepositoriesInput.MaxResults = &maxResults
	describeImagesInput.MaxResults = &maxResults
	totalImages := 0
	for {
		outputRepos, err := c.svc.DescribeRepositories(&describeRepositoriesInput)
		if err != nil {
			return err
		}
		for _, repo := range outputRepos.Repositories {
			for {
				describeImagesInput.RepositoryName = repo.RepositoryName
				outputDescribeImages, err := c.svc.DescribeImages(&describeImagesInput)
				if err != nil {
					return err
				}
				common.ChunkedCreate(c.db, c.transformImages(outputDescribeImages.ImageDetails))
				totalImages += len(outputDescribeImages.ImageDetails)
				c.log.Info("Fetched resources", zap.String("resource", "ecr.images"), zap.Int("count", len(outputDescribeImages.ImageDetails)))
				if aws.StringValue(outputDescribeImages.NextToken) == "" {
					break
				}
				describeImagesInput.NextToken = outputDescribeImages.NextToken
			}
		}
		if aws.StringValue(outputRepos.NextToken) == "" {
			break
		}
		describeRepositoriesInput.NextToken = outputRepos.NextToken
	}

	if totalImages == 0 {
		c.log.Info("Fetched resources", zap.String("resource", "ecr.images"), zap.Int("count", 0))
	}
	return nil
}
