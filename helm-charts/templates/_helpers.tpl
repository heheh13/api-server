{{- define "labels" -}}
chart-name: {{.Chart.Name | quote}}
release-Name : {{.Release.Name | quote}}
{{- range $key, $val := .Values.Labels}}
{{$key}}: {{$val | quote}}
{{- end}}
{{- end}}