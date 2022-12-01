package client

import (
  "github.com/Azure/azure-sdk-for-go/sdk/azcore"
  {{- range .}}
  "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/{{.ImportPath}}"
  {{- end}}
)

type Services struct {
  {{- range .}}
  {{.Service | ToCamel}}{{.SubService | ToCamel}} *{{.Service}}.{{.ClientName}}
  {{- end}}
}

func InitServices(subscriptionId string, azCred azcore.TokenCredential) (Services, error) {
  var services Services
  {{range .}}
  {{.Service | ToCamel}}{{.SubService | ToCamel}}, err := {{.Service}}.{{.NewFuncName}}(subscriptionId, azCred, nil)
  if err != nil {
    return services, err
  }
  services.{{.Service | ToCamel}}{{.SubService | ToCamel}} = {{.Service | ToCamel}}{{.SubService | ToCamel}}
  {{end}}
  return services, nil
}