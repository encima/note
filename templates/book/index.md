# {{.Title}}

{{if .Subtitle}}## {{.Subtitle}}{{end}}

**By**: {{range $i, $a := .Authors}}{{- if $i}}, {{end}}{{- $a}}{{end}}

**Published**: {{.Published -}}
{{range .Chapters}}

[Chapter {{.Number}} - {{.Name}}]({{.Link}}){{end}}