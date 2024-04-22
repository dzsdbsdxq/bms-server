package global

{{- if .HasGlobal }}

import "bms-server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}