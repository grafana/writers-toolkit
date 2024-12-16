local static_exceptions = [
  'Adaptive Metrics',
  'Ansible',
  'Application Observability',
  'AWS Observability',
  'AzureAD',
  'Beyla',
  'BoltDB',
  'Cloud Logs',
  'Cloud Metrics',
  'Cloud Traces',
  'Cluster',
  'Frontend Observability',
  'Generic OAuth',
  'GitHub',
  'GitLab',
  'Google',
  'GKE Autopilot',
  'Grafana',
  'Grafana Agent',
  'Grafana Agent Flow',
  'Grafana Cloud',
  'Grafana Cloud Traces',
  'Grafana Enterprise Logs',
  'Grafana Enterprise Metrics',
  'Grafana Enterprise Traces',
  'Grafana Labs',
  'Graylog',
  'Helm',
  'Hugo',
  'IBM Cloud',
  'Kibana',
  'Kubernetes Monitoring',
  'LogQL',
  'Loki',
  'macOS',
  'Memcached',
  'Microsoft',
  'Mimir',
  'Node Exporter',
  'OAuth',
  'Okta',
  'OpenCost',
  'OpenTelemetry',
  'OpenTelemetry Collector',
  'Personally Identifiable Information',
  'Prometheus',
  'PromQL',
  'Promtail',
  'Pyroscope',
  'Redis',
  'Service Level Objective',
  'Tempo',
  'Terraform',
  'Thanos',
  'TraceQL',
  'WebSockets',
  "What's new",
  "Writers' Toolkit",
  'YouTube',
];
local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'capitalization',
  message: |||
    Use sentence-style capitalization for '%s'.

    Vale considers multi-word exceptions such as _Grafana Enterprise Metrics_ as a single correctly cased word.

    If your heading contains capitalized words that represent product names, you need to add those words to the Grafana dictionary or the list of static exceptions in https://github.com/grafana/writers-toolkit/blob/main/vale/Headings.jsonnet for them to be considered correctly cased.
  |||,
  link: 'https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings',
  level: 'warning',
  scope: 'heading',
  match: '$sentence',
  threshold: 0.3,
  exceptions:
    [
      'Amazon ' + def.word
      for def in defs
      if 'product' in def && def.product && 'Amazon' in def && def.Amazon
    ] +
    [
      def.word
      for def in defs
      if 'product' in def && def.product
    ] +
    [
      def.word
      for def in defs
      if 'abbreviation' in def && def.abbreviation
    ] +
    [
      def.elaboration
      for def in defs
      if ('abbreviation' in def && def.abbreviation) && ('elaboration' in def && std.any(std.map(function(c) local cp = std.codepoint(c); cp >= 65 && cp < 97, std.stringChars(def.elaboration))))
    ] + static_exceptions,
})
