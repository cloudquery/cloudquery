// Code generated by codegen; DO NOT EDIT.
package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
    "github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
    {{- range . }}
    "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/{{ . }}"
    {{- end }}
)

func initServices(subscriptionID string, credentials azcore.TokenCredential, options *arm.ClientOptions) (*Services, error) {
	var services Services
	var err error

	{{ range . }}
	services.{{ . | ToCamel }}, err = {{ . }}.New{{ . | ToCamel }}Client(subscriptionID, credentials, options)
	if err != nil {
	    return nil, err
	}
    {{ end }}

	return &services, nil
}

type Services struct {
    {{- range . }}
		{{ . | ToCamel }} *{{ . }}.{{ . | ToCamel }}Client
    {{- end }}
}