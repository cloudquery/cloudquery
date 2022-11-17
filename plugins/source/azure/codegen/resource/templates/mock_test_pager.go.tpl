    var response api.{{ .StructName }}
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

    mock{{ .Client }}.EXPECT().{{ .Method }}({{- range .Params }}gomock.Any(), {{ end }}gomock.Any()).
        Return(client.CreatePager(response)).MinTimes(1)