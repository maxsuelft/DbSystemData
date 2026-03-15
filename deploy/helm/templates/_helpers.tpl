{{/*
Expand the name of the chart.
*/}}
{{- define "dbsystemdata.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
*/}}
{{- define "dbsystemdata.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "dbsystemdata.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "dbsystemdata.labels" -}}
helm.sh/chart: {{ include "dbsystemdata.chart" . }}
{{ include "dbsystemdata.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "dbsystemdata.selectorLabels" -}}
app.kubernetes.io/name: {{ include "dbsystemdata.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app: dbsystemdata
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "dbsystemdata.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "dbsystemdata.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Namespace - uses the release namespace from helm install -n <namespace>
*/}}
{{- define "dbsystemdata.namespace" -}}
{{- .Release.Namespace }}
{{- end }}
