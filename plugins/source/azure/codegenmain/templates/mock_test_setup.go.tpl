    mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeData(&data, options.WithIgnoreInterface(true), options.WithFieldsToIgnore("Response"), options.WithRandomMapAndSliceMinSize(1), options.WithRandomMapAndSliceMaxSize(1)))