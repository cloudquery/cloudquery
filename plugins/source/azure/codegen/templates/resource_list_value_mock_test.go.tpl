{{template "mock_test_base.go.tpl" .}}

func create{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	{{template "mock_test_setup.go.tpl" .}}

    {{if .MockListResult }}
	{{if eq .MockListResult "CQ_CODEGEN_DIRECT_RESPONSE" }}
	result := data
	{{else}}
    result := {{ .AzurePackageName }}.{{ .MockListResult }}{Value: &[]{{ .AzurePackageName }}.{{ or .MockValueType .AzureStructName }}{data}}
	{{end}}
	{{else}}
	result := {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{Value: &[]{{ .AzurePackageName }}.{{ or .MockValueType .AzureStructName }}{data}}
	{{end}}

	{{template "mock_test_assert.go.tpl" .}}
}