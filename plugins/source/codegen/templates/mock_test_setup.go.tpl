	mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,{{ range .MockRelations }}
			{{ . }},{{ end  }}
		},
	}

	data := {{ .AzurePackageName }}.{{ or .MockDefinitionType .AzureStructName }}{}
	require.Nil(t, faker.FakeObject(&data))
	{{ if (or .GetFunction .Table.Relations) }}
	// Ensure name and ID are consistent so we can reference it in other mock
	name :=  "test"
	data.Name = &name

	// Use correct Azure ID format
	id :=  "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id
	{{ end}}


	{{if .GetFunction}}

	getData := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeObject(&getData))
	{{end}}