package ecr

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"go.uber.org/zap"
)


type Repository struct {
	_   interface{} `neo:"raw:MERGE (a:AWSAccount {account_id: $account_id}) MERGE (a) - [:Resource] -> (n)"`
	ID uint `gorm:"primarykey"`
	AccountID string
	Region string
	CreatedAt *time.Time

	EncryptionType   *string
	EncryptionKmsKey *string

	ImageScanningConfigurationScanOnPush *bool
	ImageTagMutability                   *string
	RegistryId                           *string
	ARN                                  *string
	Name                                 *string
	URI                                  *string
	Images []*Image `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Repository) TableName() string {
	return "aws_ecr_repositories"
}

type Image struct {
	ID        uint        `gorm:"primarykey"`
	RepositoryID uint `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	ArtifactMediaType      *string
	ImageDigest            *string
	ImageManifestMediaType *string
	ImagePushedAt          *time.Time

	ImageScanFindingsSeverityCounts               []*ImageSeverityCount `gorm:"constraint:OnDelete:CASCADE;"`
	ImageScanFindingsImageScanCompletedAt         *time.Time
	ImageScanFindingsVulnerabilitySourceUpdatedAt *time.Time

	ImageScanStatusDescription *string
	ImageScanStatusStatus      *string
	ImageSizeInBytes           *int64
	ImageTags                  []*ImageTags `gorm:"constraint:OnDelete:CASCADE;"`
	RegistryId                 *string
	RepositoryName             *string
}

func (Image) TableName() string {
	return "aws_ecr_images"
}

type ImageSeverityCount struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Severity string
	Count    int32
}

func (ImageSeverityCount) TableName() string {
	return "aws_ecr_image_severity_counts"
}

type ImageTags struct {
	ID        uint   `gorm:"primarykey"`
	ImageID   uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Value string
}

func (ImageTags) TableName() string {
	return "aws_ecr_image_tags"
}

func (c *Client) transformRepository(value *types.Repository) *Repository {
		tValue := Repository {
			AccountID:          c.accountID,
			Region:             c.region,
			CreatedAt:          value.CreatedAt,
			ImageTagMutability: aws.String(string(value.ImageTagMutability)),
			RegistryId:         value.RegistryId,
			ARN:                value.RepositoryArn,
			Name:               value.RepositoryName,
			URI:                value.RepositoryUri,
		}

		if value.EncryptionConfiguration != nil {
			tValue.EncryptionType = aws.String(string(value.EncryptionConfiguration.EncryptionType))
			tValue.EncryptionKmsKey = value.EncryptionConfiguration.KmsKey
		}

		if value.ImageScanningConfiguration != nil {
			tValue.ImageScanningConfigurationScanOnPush = &value.ImageScanningConfiguration.ScanOnPush
		}

		return &tValue
}

func (c *Client) transformImages(values *[]types.ImageDetail) []*Image {
	var tValues []*Image
	for _, value := range *values {
		tValue := Image{
			AccountID:              c.accountID,
			Region:                 c.region,
			ArtifactMediaType:      value.ArtifactMediaType,
			ImageDigest:            value.ImageDigest,
			ImageManifestMediaType: value.ImageManifestMediaType,
			ImagePushedAt:          value.ImagePushedAt,
			ImageSizeInBytes:       value.ImageSizeInBytes,
			ImageTags:              c.transformImageTags(&value.ImageTags),
			RegistryId:             value.RegistryId,
			RepositoryName:         value.RepositoryName,
		}

		if value.ImageScanFindingsSummary != nil {
			tValue.ImageScanFindingsImageScanCompletedAt = value.ImageScanFindingsSummary.ImageScanCompletedAt
			tValue.ImageScanFindingsVulnerabilitySourceUpdatedAt = value.ImageScanFindingsSummary.VulnerabilitySourceUpdatedAt
			for severity, count := range value.ImageScanFindingsSummary.FindingSeverityCounts {
				tValue.ImageScanFindingsSeverityCounts = append(tValue.ImageScanFindingsSeverityCounts,
					&ImageSeverityCount{
						AccountID: c.accountID,
						Region:    c.region,
						Severity:  severity,
						Count:     count,
					})
			}
		}

		if value.ImageScanStatus != nil {
			tValue.ImageScanStatusDescription = value.ImageScanStatus.Description
			tValue.ImageScanStatusStatus = aws.String(string(value.ImageScanStatus.Status))
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformImageTags(values *[]string) []*ImageTags {
	var tValues []*ImageTags
	for _, v := range *values {
		tValues = append(tValues, &ImageTags{
			AccountID: c.accountID,
			Region:    c.region,
			Value:     v,
		})
	}
	return tValues
}

type ImageConfig struct {
	Filter string
}

var ImageTables = []interface{}{
	&Repository{},
	&Image{},
	&ImageSeverityCount{},
	&ImageTags{},
}

func (c *Client) images(_ interface{}) error {
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ImageTables...)
	ctx := context.Background()
	var describeRepositoriesInput ecr.DescribeRepositoriesInput
	var describeImagesInput ecr.DescribeImagesInput
	var maxResults int32
	maxResults = 1000
	describeRepositoriesInput.MaxResults = &maxResults
	describeImagesInput.MaxResults = &maxResults
	totalImages := 0
	for {
		outputRepos, err := c.svc.DescribeRepositories(ctx, &describeRepositoriesInput)
		if err != nil {
			return err
		}
		for _, repo := range outputRepos.Repositories {
			describeImagesInput.RepositoryName = repo.RepositoryName
			describeImagesInput.NextToken = nil
			tRepo := c.transformRepository(&repo)
			for {
				outputDescribeImages, err := c.svc.DescribeImages(ctx, &describeImagesInput)
				if err != nil {
					return err
				}
				tRepo.Images = append(tRepo.Images, c.transformImages(&outputDescribeImages.ImageDetails)...)
				totalImages += len(outputDescribeImages.ImageDetails)
				c.log.Info("Fetched resources", zap.String("resource", "ecr.images"), zap.Int("count", len(outputDescribeImages.ImageDetails)))
				
				if aws.ToString(outputDescribeImages.NextToken) == "" {
					break
				}
				describeImagesInput.NextToken = outputDescribeImages.NextToken
			}
			c.db.InsertOne(tRepo)
		}
		if aws.ToString(outputRepos.NextToken) == "" {
			break
		}
		describeRepositoriesInput.NextToken = outputRepos.NextToken
	}

	if totalImages == 0 {
		c.log.Info("Fetched resources", zap.String("resource", "ecr.images"), zap.Int("count", 0))
	}
	return nil
}
