package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLightsailCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := lightsail.GetCertificatesInput{
		IncludeCertificateDetails: true,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetCertificates(ctx, &input)
	if err != nil {
		return err
	}
	for _, cer := range response.Certificates {
		res <- cer.CertificateDetail
	}
	return nil
}
