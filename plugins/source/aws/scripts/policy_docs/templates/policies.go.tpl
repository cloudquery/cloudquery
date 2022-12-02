## Policies

CloudQuery SQL Policies for AWS. See [Github Readme](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies) for installation instructions.

### Policies and Compliance Frameworks Available

{{- range .}}
#### {{ .Name }}

##### Queries
{{ .Name }} performs the following checks:
{{- _, $query := range .Queries }}
  - {{ $query }}
{{- end }}

##### Requirements
{{ .Name }} requires the following tables to be synced before the policy is run:

```
{{- _, $table := range .Tables }}
  - {{ $table }}
{{- end }}
```
{{- end }}