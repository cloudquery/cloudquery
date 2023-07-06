package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Profiles() *schema.Table {
	tableName := "aws_signer_signing_profiles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/signer/latest/api/API_GetSigningProfile.html`,
		Resolver:            fetchProfiles,
		PreResourceResolver: getProfile,
		Transform:           transformers.TransformWithStruct(&signer.GetSigningProfileOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithPrimaryKeys("ProfileVersionArn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "signer"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Signer
	config := signer.ListSigningProfilesInput{}

	paginator := signer.NewListSigningProfilesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *signer.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Profiles
	}
	return nil
}

func getProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Signer
	a := resource.Item.(types.SigningProfile)

	profile, err := svc.GetSigningProfile(ctx, &signer.GetSigningProfileInput{ProfileName: a.ProfileName}, func(o *signer.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = profile
	return nil
}
