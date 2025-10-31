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

var SubHeaderTpl = `@{{ .SubID }} path /{{ .SubID }} /{{ .SubID }}/
header @{{ .SubID }} {
	defer
	{{- range $key, $value := .Headers }}
	{{ $key }} "{{ $value }}"
	{{- end }}
}
`

type SubHeader struct {
	SubID   string
	Headers map[string]string
}
