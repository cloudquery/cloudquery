package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pkg/errors"
)

func Accounts() *schema.Table {
	tableName := "aws_s3_accounts"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchS3Accounts,
		Transform: transformers.TransformWithStruct(&models.PublicAccessBlockConfigurationWrapper{}, transformers.WithUnwrapStructFields("PublicAccessBlockConfiguration")),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchS3Accounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().S3control
	var accountConfig s3control.GetPublicAccessBlockInput
	accountConfig.AccountId = aws.String(c.AccountID)
	resp, err := svc.GetPublicAccessBlock(ctx, &accountConfig)

	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		var nspabc *s3controlTypes.NoSuchPublicAccessBlockConfiguration
		if !errors.As(err, &nspabc) {
			return err
		}
		res <- models.PublicAccessBlockConfigurationWrapper{ConfigExists: false}
	} else {
		res <- models.PublicAccessBlockConfigurationWrapper{PublicAccessBlockConfiguration: *resp.PublicAccessBlockConfiguration, ConfigExists: true}
	}

	return nil
}
