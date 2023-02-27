package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ContactLists() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_contact_lists",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html`,
		Resolver:            fetchSesContactLists,
		PreResourceResolver: getContactList,
		Transform:           transformers.TransformWithStruct(&sesv2.GetContactListOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactListName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
