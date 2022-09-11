    mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ or .MockDefinitionType .AzureStructName }}{}
	require.Nil(t, faker.FakeObject(&data))

	{{if .GetFunction}}
	id :=  "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	getData := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeObject(&getData))
	{{end}}