package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbCertificates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DocDB

	input := &docdb.DescribeCertificatesInput{}

	for {
		output, err := svc.DescribeCertificates(ctx, input)
		if err != nil {
			return err
		}
		if len(output.Certificates) == 0 {
			return nil
		}
		res <- output.Certificates
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}
