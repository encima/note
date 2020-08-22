# {{.Title}} 
## {{.Status}}

### Context
 * Source: {{.Source}}
 * Person: {{.Person}}
 * Type: {{.Type}}
 * Links:
   [i](./index.md) - {{range $i, $l := .Links}}{{- if $i}} - {{end}}{{.}}{{end}}

