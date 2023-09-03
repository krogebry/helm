{/*
Expand the name of the chart.
*/}}
{{- define "default_name" -}}
{{ .Values.app.name }}-{{ .Values.shared_clusteri.app_name }}
{{- end }}
{{- define "app.name" -}}
{{- .Values.name | default (include "default_name" .) }}
{{- end }}

{{- define "mtf-namespace-name" -}}
{{ .Values.shared_clusteri.app_name }}-{{ .Values.app.env }}
{{- end }}

{{- define "mtf-fqdn" -}}
{{ .Values.shared_clusteri.app_name }}-{{ include "mtf-namespace-name" . }}.{{ .Values.shared_clusteri.cluster_family_name }}-{{.Values.shared_clusteri.cluster_family_id }}-{{.Values.shared_clusteri.env_name }}-{{ .Values.shared_clusteri.region}}.{{ .Values.shared_clusteri.cloud }}.{{ .Values.shared_clusteri.domain }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "app.fullname" -}}
{{- default ( include "app.name" . ) .Values.fullname }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "app.chart" -}}
{{- .Chart.Name }}-{{- .Chart.Version }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "app.labels" -}}
helm.sh/chart: {{ include "app.chart" . }}
{{ include "app.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "app.selectorLabels" -}}
app.kubernetes.io/name: {{ include "app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "app.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "app.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Format the ARN of the service account to use
*/}}
{{- define "app.serviceAccountEnvName" -}}
{{- if eq .Values.infra.env_name "stage" }}
{{- printf "stg" }}
{{- else }}
{{- printf "%s" .Values.infra.env_name }}
{{- end }}
{{- end }}

{{- define "oauth2-proxy.version" -}}
{{ trimPrefix "v" (lower (.Values.oauth2_proxy.image.tag )) }}
{{- end -}}

{{- define "oauth2-proxy.configMapName" -}}
{{ printf "%s-%s-config" (include "app.name" .) .Values.oauth2_proxy.containerName}}
{{- end -}}

{{- define "telegraf.configMapName" -}}
{{ printf "%s-%s-config" (include "app.name" .) .Values.telegraf.containerName}}
{{- end -}}

{{- define "virtual_service.rest.port" -}}
{{- if .Values.oauth2_proxy.enabled -}}
{{ .Values.oauth2_proxy.service.port }}
{{- else -}}
{{ .Values.service.port }}
{{- end -}}
{{- end -}}
