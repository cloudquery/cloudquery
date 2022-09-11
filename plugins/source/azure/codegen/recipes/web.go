package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

type publishProfile struct {
	PublishUrl string
	UserName   string
	UserPWD    string
}

func Web() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &web.Site{},
					listFunction:       "List",
					subServiceOverride: "Apps",
					mockListResult:     "AppCollection",
					relations:          []string{"siteAuthSettings()", "publishingProfiles()", "vnetConnections()", "publishingProfiles()"},
				},
				{
					azureStruct:          &web.SiteAuthSettings{},
					listFunction:         "GetAuthSettings",
					listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
					listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name"},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation: true,
				},
				{
					azureStruct:          &web.VnetInfo{},
					listFunction:         "GetVnetConnection",
					listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
					listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name", "*site.SiteConfig.VnetName"},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					subServiceOverride: "VnetConnections",
					isRelation:         true,
				},
				{
					azureStruct: &publishProfile{},
					helpers: []string{`type PublishProfile struct {
						PublishUrl string ` + "`" + `xml:"publishUrl,attr"` + "`" + `
						UserName   string ` + "`" + `xml:"userName,attr"` + "`" + `
						UserPWD    string ` + "`" + `xml:"userPWD,attr"` + "`" + `
					}`, `type publishData struct {
						XMLName     xml.Name ` + "`" + `xml:"publishUrl,attr"` + "`" + `
						PublishData []PublishProfile ` + "`" + `xml:"PublishProfile"` + "`" + `
					}`,
					},
					listFunction:         `ListPublishingProfileXMLWithSecrets`,
					listFunctionArgsInit: []string{"site := parent.Item.(web.Site)"},
					listFunctionArgs:     []string{"*site.ResourceGroup", "*site.Name", "web.CsmPublishingProfileOptions{}"},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
				
					buf := new(bytes.Buffer)
					if _, err = buf.ReadFrom(response.Body); err != nil {
						return errors.WithStack(err)
					}
					var profileData publishData
					if err = xml.Unmarshal(buf.Bytes(), &profileData); err != nil {
						return errors.WithStack(err)
					}
				
					res <- profileData.PublishData`,
					subServiceOverride: "PublishingProfiles",
					isRelation:         true,
				},
			},
			serviceNameOverride: "Web",
		},
	}

	return generateResources(resourcesByTemplates)
}
