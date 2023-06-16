# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for AWS. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies) for installation instructions.

{{- range .}}
## {{ .Name }}

### Requirements
{{ .Name }} requires the following tables to be synced before the policy is executed:

```yaml
tables:
{{- range $table := .Tables }}
  - {{ $table }}
{{- end }}
```

### Queries

{{ .Name }} performs the following checks:
{{- range $query := .Queries }}{{ if not $query.View }}
  - {{ $query.Title }}
{{- end }}{{- end }}

{{- if .DependentViews }}

### Dependent Views

{{ .Name }} depends on the following views:
{{$createdViews := .CreatedViews }}
{{- $num_created := 0}}
{{- range $v := .DependentViews }}
  - {{ $v }}{{ if index $createdViews $v }}<sup>*</sup>{{ $num_created = add $num_created 1 }}{{end}}
{{- end }}

{{- if gt $num_created 0}}

  <sup>*</sup> {{if eq $num_created 1}}This view is{{else}}These views are{{end}} automatically created or updated by this policy.
{{- end}}
{{- end }}
{{- if .UnusedViews }}

### Unused Views

{{ .Name }} creates {{if eq (len .UnusedViews) 1}}this view but does not use it:{{else}}these views but does not use them:{{end}}
{{range $v := .UnusedViews }}
  - {{ $v }}
{{- end }}
{{end}}

{{- end }}
