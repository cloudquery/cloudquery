func build{{ .SubService | ToCamel }}PreResolver(t *testing.T, ctrl *gomock.Controller, c *client.Services) {
	{{- with .Fetcher.PreResolver }}
	if c.{{ .Service }} == nil {
		c.{{ .Service }} = new(service.{{ .Service }}Client)
	}
	{{ .Service | ToLower }}Client := c.{{ .Service }}
	if {{ .Service | ToLower }}Client.{{ .Client }} == nil {
		{{ .Service | ToLower }}Client.{{ .Client }} = mocks.NewMock{{ .Client }}(ctrl)
	}

	mock{{ .Client }} := {{ .Service | ToLower }}Client.{{ .Client }}.(*mocks.Mock{{ .Client }})

	{{ template "mock_fetch_call.go.tpl" . }}
	{{- end -}}
}