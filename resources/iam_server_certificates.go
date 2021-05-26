package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamServerCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_server_certificates",
		Resolver:     fetchIamServerCertificates,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "path",
				Type: schema.TypeString,
			},
			{
				Name: "server_certificate_id",
				Type: schema.TypeString,
			},
			{
				Name: "server_certificate_name",
				Type: schema.TypeString,
			},
			{
				Name: "expiration",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "upload_date",
				Type: schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamServerCertificates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config iam.ListServerCertificatesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListServerCertificates(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ServerCertificateMetadataList
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
