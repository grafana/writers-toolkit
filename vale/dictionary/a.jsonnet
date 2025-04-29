local word = import './word.jsonnet';
[
  word.new('ACL', 'S', 'noun') { abbreviation: true },
  word.new('ADOT', '', 'noun') { abbreviation: true, description: 'AWS Distro for OpenTelemetry', established_abbreviation: true, product: true },
  word.new('AI Observability', '', 'noun') { product: true, swaps: { 'Grafana AI observability': 'Grafana AI Observability' } },
  word.new('Adaptive Metrics', '', 'noun') { swaps: { '\\b(?:[aA]daptive metrics|adaptive Metrics)\\b': 'Adaptive Metrics' } },
  word.new('Aerospike', '', 'noun'),
  word.new('after', '', 'preposition') { swaps: { Once: 'After' } },
  word.new('Agent', '', 'noun') { product: true },
  word.new('Alertmanager', 'MS', 'noun') { swaps: { '(?:alert[Mm]anager|[Aa]lert [Mm]anager|AlertManager)': 'Alertmanager' } },
  word.new('allowlist', 'DGS', 'verb') { swaps: { whitelisted: 'allowlisted', whitelisting: 'allowlisting', whitelists: 'allowlists' } },
  word.new('allowlist', 'S', 'noun') { swaps: { whitelist: 'allowlist' } },
  word.new('Alloy', '', 'noun') { product: true },
  word.new('anonymize', 'DGS', 'verb') { description: 'https://dictionary.cambridge.org/dictionary/english/anonymize' },
  word.new('Ansible', '', 'adjective'),
  word.new('Ansible', '', 'noun'),
  word.new('Apdex', '', 'noun'),
  word.new('API', 'S', 'noun') { abbreviation: true, elaboration: 'Application Programming Interface', established_abbreviation: true },
  word.new('APT', '', 'noun') { abbreviation: true, description: 'https://en.wikipedia.org/wiki/APT_(software)', elaboration: 'Advanced package tool', established_abbreviation: true },
  word.new('ARN', '', 'noun') { Amazon: true, abbreviation: true, description: 'Amazon Resource Name', established_abbreviation: true, product: true },
  word.new('Asserts', '', 'noun') { description: 'https://grafana.com/products/cloud/asserts/', product: true },
  word.new('autoscale', 'DGS', 'verb'),
  word.new('autoscaler', 'S', 'noun'),
  word.new('AWS', '', 'noun') { abbreviation: true, elaboration: 'Amazon Web Services', established_abbreviation: true, product: true },
  // Note that AWS Distro for OpenTelemetry Collector is an Amazon product but should not be called Amazon AWS Distro for OpenTelemetry Collector.
  // https://aws.amazon.com/otel/
  word.new('AWS Distro for OpenTelemetry Collector', '', 'noun') { description: 'An AWS-supported distribution of the OpenTelemetry project.', product: true },
  // Note that AWS X-Ray is an Amazon product but should not be called Amazon AWS X-Ray.
  // https://docs.aws.amazon.com/xray/
  word.new('AWS X-Ray', '', 'noun') { description: 'A service that collects data about requests that your application serves.', product: true },
]
