    var response api.{{ .StructName }}
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.ID = to.Ptr(id)

    mock{{ .Client }}.EXPECT().{{ .Method }}({{- range .Params }}gomock.Any(), {{ end }}gomock.Any()).
        Return(response, nil).MinTimes(1)