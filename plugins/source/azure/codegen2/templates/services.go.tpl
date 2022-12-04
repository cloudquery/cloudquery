package client

import (
  "github.com/Azure/azure-sdk-for-go/sdk/azcore"
  {{- range .}}
  "{{.ImportPath}}"
  {{- end}}
)

type Services struct {
  {{- range .}}
  {{.Service | ToCamel}}{{.ClientName}} *{{.Service}}.{{.ClientName}}
  {{- end}}
}

func InitServices(subscriptionId string, azCred azcore.TokenCredential) (Services, error) {
  var services Services
  {{range .}}
  {{if .NewFuncHasSubscriptionId}}
  {{.Service | ToCamel}}{{.ClientName}}, err := {{.Service}}.{{.NewFuncName}}(subscriptionId, azCred, nil)
  if err != nil {
    return services, err
  }
  services.{{.Service | ToCamel}}{{.ClientName}} = {{.Service | ToCamel}}{{.ClientName}}
  {{else}}
  {{.Service | ToCamel}}{{.ClientName}}, err := {{.Service}}.{{.NewFuncName}}(azCred, nil)
  if err != nil {
    return services, err
  }
  services.{{.Service | ToCamel}}{{.ClientName}} = {{.Service | ToCamel}}{{.ClientName}}
  {{end}}
  {{end}}
  return services, nil
}