{{template "mock_test_base.go.tpl" .}}

func create{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	{{template "mock_test_setup.go.tpl" .}}

    {{if .MockListResult}}
    result := {{ .AzurePackageName }}.{{ .MockListResult }}{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}
	{{else}}
	result := {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}
	{{end}}

	{{template "mock_test_assert.go.tpl" .}}
}