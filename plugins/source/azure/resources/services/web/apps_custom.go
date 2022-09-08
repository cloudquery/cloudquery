package web

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

type publishProfile struct {
	PublishUrl string `xml:"publishUrl,attr"`
	UserName   string `xml:"userName,attr"`
	UserPWD    string `xml:"userPWD,attr"`
}

type publishData struct {
	XMLName     xml.Name         `xml:"publishData"`
	PublishData []publishProfile `xml:"publishProfile"`
}

func PublishingProfiles() *schema.Table {
	return &schema.Table{
		Name:     "azure_web_app_publishing_profiles",
		Resolver: fetchPublishingProfiles,
		Columns: []schema.Column{
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name: "publish_url",
				Type: schema.TypeString,
			},
			{
				Name: "user_name",
				Type: schema.TypeString,
			},
			{
				Name:     "user_pwd",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserPWD"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchPublishingProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(web.Site)
	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.ListPublishingProfileXMLWithSecrets(ctx, *p.ResourceGroup, *p.Name, web.CsmPublishingProfileOptions{})
	if err != nil {
		return diag.WrapError(err)
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(response.Body); err != nil {
		return diag.WrapError(err)
	}
	var profileData publishData
	if err = xml.Unmarshal(buf.Bytes(), &profileData); err != nil {
		return diag.WrapError(err)
	}

	res <- profileData.PublishData
	return nil
}

func fetchVnetConnections(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	site := resource.Item.(web.Site)
	svc := meta.(*client.Client).Services().Web.Apps

	if site.SiteConfig == nil || site.SiteConfig.VnetName == nil {
		return nil
	}

	response, err := svc.GetVnetConnection(ctx, *site.ResourceGroup, *site.Name, *site.SiteConfig.VnetName)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.VnetInfoProperties != nil {
		vnetConnection := make(map[string]interface{})
		if response.Name != nil {
			vnetConnection["name"] = response.Name
		}
		if response.ID != nil {
			vnetConnection["id"] = response.ID
		}
		if response.Type != nil {
			vnetConnection["type"] = response.Type
		}
		vnetConnection["properties"] = response.VnetInfoProperties
		b, err := json.Marshal(vnetConnection)
		if err != nil {
			return errors.WithStack(err)
		}
		return errors.WithStack(resource.Set(c.Name, b))
	}
	return nil
}
