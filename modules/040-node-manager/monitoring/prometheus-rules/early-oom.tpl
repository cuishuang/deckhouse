{{- if .Values.nodeManager.earlyOomEnabled }}
- name: d8.early-oom.availability
  rules:
  - alert: D8EarlyOOMPodIsNotReady
    expr: min by (pod) (early_oom_psi_unavailable{namespace="d8-cloud-instance-manager", pod=~"early-oom-.*"}) == 1
    for: 3m
    labels:
      severity_level: "8"
      tier: cluster
      d8_module: node-manager
      d8_component: early-oom
    annotations:
      plk_protocol_version: "1"
      plk_markup_format: "markdown"
      plk_create_group_if_not_exists__d8_early_oom_malfunctioning: "D8EarlyOOMPodIsNotReady,tier=cluster,prometheus=deckhouse,kubernetes=~kubernetes"
      plk_grouped_by__d8_early_oom_malfunctioning: "D8EarlyOOMPodIsNotReady,tier=cluster,prometheus=deckhouse,kubernetes=~kubernetes"
      plk_labels_as_annotations: "pod"
      summary: >
        The {{`{{$labels.pod}}`}} Pod has detected unavailable PSI subsystem.
        Check logs for additional information:
        ```
        kubectl -n d8-cloud-instance-manager logs {{`{{$labels.pod}}`}} -c psi-monitor
        ```
{{- end }}
