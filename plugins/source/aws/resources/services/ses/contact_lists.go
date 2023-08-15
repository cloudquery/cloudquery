package ses

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ContactLists() *schema.Table {
	tableName := "aws_ses_contact_lists"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html`,
		Resolver:            fetchSesContactLists,
		PreResourceResolver: getContactList,
		Transform:           transformers.TransformWithStruct(&sesv2.GetContactListOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ContactListName"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchSesContactLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2

	p := sesv2.NewListContactListsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *sesv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ContactLists
	}

	return nil
}

func getContactList(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2
	item := resource.Item.(types.ContactList)

	getOutput, err := svc.GetContactList(ctx,
		&sesv2.GetContactListInput{ContactListName: item.ContactListName},
		func(o *sesv2.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	resource.SetItem(getOutput)

	return nil
}
