    pager := {{ template "fetch_call.go.tpl" . }}
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return err
        }
        res <- page.Value
    }