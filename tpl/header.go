package tpl

var GlobalHeaderTpl = `header {
	{{- range $key, $value := .Headers }}
	{{ $key }} "{{ $value }}"
	{{- end }}
}
`

type GlobalHeader struct {
	Headers map[string]string
}

var SubHeaderTpl = `@{{ .SubID }} path{{if .PathSegment}} /{{ .PathSegment }}/{{ .SubID }} /{{ .PathSegment }}/{{ .SubID }}/{{else}} /{{ .SubID }} /{{ .SubID }}/{{end}}
header @{{ .SubID }} {
	defer
	{{- range $key, $value := .Headers }}
	{{ $key }} "{{ $value }}"
	{{- end }}
}
`

type SubHeader struct {
	SubID       string
	PathSegment string
	Headers     map[string]string
}
