package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	{{range .Packages}}"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/{{.}}"
    {{end}}"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            "azure",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
            {{range .Resources}}"{{.ServiceName}}": {{.Function}}(),
            {{end}}
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
