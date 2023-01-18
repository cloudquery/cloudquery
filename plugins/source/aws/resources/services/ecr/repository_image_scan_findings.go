package ecr

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecr/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RepositoryImageScanFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_image_scan_findings",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html`,
		Resolver:    fetchEcrRepositoryImageScanFindings,
		Transform:   transformers.TransformWithStruct(&models.ImageScanWrapper{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
		},
	}
}
