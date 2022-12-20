package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config rds.DescribeCertificatesInput
	c := meta.(*client.Client)
	svc := c.Services().Rds
	for {
		response, err := svc.DescribeCertificates(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
