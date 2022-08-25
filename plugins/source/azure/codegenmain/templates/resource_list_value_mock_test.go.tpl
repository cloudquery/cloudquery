{{template "base_mock_test.go.tpl" .}}

func create{{ .AzureService }}{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	faker.SetIgnoreInterface(true)
	require.Nil(t, faker.FakeData(&data))

    {{if .MockListResult}}
    list := {{ .AzurePackageName }}.{{ .MockListResult }}{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}
	{{else}}
	list := {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}
	{{end}}

	{{ range or .MockListFunctionArgsInit .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(list, nil)
	return s
}