local newWord(word, affixes, po) = {
  // word is the word to be defined.
  word: word,
  // affixes is a string of Hunspell affixes that can be applied to the word.
  affixes: affixes,
  // po is the part of speech of the word.
  // TODO: Determine the full list of acceptable values.
  // Known acceptable values:
  //   - 'noun'
  //   - 'verb'
  //   - 'adjective'
  po: po,

  // description is a string that defines the word.
  // It's mostly used for jargon or products.
  description: null,

  // acronym is a boolean that indicates whether the word is an acronym.
  acronym: false,
  // established_acronym is a boolean that indicates whether the word is an established acronym that doesn't need explanation in general.
  established_acronym: false,

  // product is a boolean that indicates whether the word is a product name.
  product: false,

  // Amazon is a boolean that indicates whether the word is an Amazon product.
  Amazon: false,
};

{
  words: [
    newWord('ACL', 'S', 'noun') { acronym: true },
    newWord('Aerospike', '', 'noun'),
    newWord('Agent', '', 'noun') { product: true },
    newWord('Alertmanager', 'MS', 'noun'),
    newWord('allowlist', 'DGS', 'verb'),
    newWord('allowlist', 'S', 'noun'),
    newWord('Alloy', '', 'noun') { product: true },
    newWord('Ansible', '', 'adjective'),
    newWord('Ansible', '', 'noun'),
    newWord('Apdex', '', 'noun'),
    newWord('API', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('autoscale', 'DGRS', 'verb'),
    newWord('AWS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('backport', 'DGS', 'verb'),
    newWord('Beyla', '', 'noun') { product: true },
    newWord('blackbox', '', 'noun'),
    newWord('blockquote', 'S', 'noun'),
    newWord('boolean', '', 'noun'),
    newWord('BPF', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Brotli', '', 'noun') { description: 'Brotli is a lossless data compression algorithm developed by Google' },
    newWord('burndown', '', 'adjective'),
    newWord('bundler', 'S', 'noun'),
    newWord('cAdvisor', 'M', 'noun'),
    newWord('CDN', '', 'noun') { acronym: true, description: 'Content Delivery Network' },
    newWord('CLA', '', 'noun') { acronym: true, description: 'Contributor License Agreement' },
    newWord('CLI', '', 'noun') { acronym: true, established_acronym: true },
    newWord('CloudWatch', '', 'noun') { Amazon: true, product: true },
    newWord('codespace', 'S', 'noun') { description: "A codespace is a development environment that's hosted in the cloud. https://docs.github.com/en/codespaces/overview" },
    newWord('Codespaces', '', 'noun') { description: 'GitHub Codespaces https://docs.github.com/en/codespaces/overview', product: true },
    newWord('composable', '', 'adjective'),
    newWord('Couchbase', 'M', 'noun'),
    newWord('Data Firehose', '', 'noun') { Amazon: true, product: true },
    newWord('CPU', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('CRD', 'S', 'noun'),
    newWord('CSS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('CSV', '', 'noun') { acronym: true, established_acronym: true },
    newWord('CVE', 'S', 'noun'),
    newWord('Databricks', '', 'noun'),
    newWord('Datadog', '', 'adjective'),
    newWord('Datadog', 'M', 'noun'),
    newWord('deliverable', 'S', 'noun'),
    newWord('disaggregate', 'DS', 'verb'),
    newWord('distroless', '', 'adjective'),
    newWord('DOM', '', 'noun') { acronym: true, established_acronym: true },
    newWord("don'ts", '', 'noun'),
    newWord('downsample', 'DG', 'verb'),
    newWord('duplicate', 'dDSN', 'noun'),
    newWord('Dynatrace', 'M', 'noun'),
    newWord('eBPF', '', 'noun') { acronym: true, established_acronym: true },
    newWord('enablement', '', 'noun'),
    newWord('enqueue', 'DS', 'verb'),
    newWord('ESLint', 'M', 'noun'),
    newWord('etcd', '', 'noun'),
    newWord('FAQ', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('Fargate', 'M', 'noun'),
    newWord('Firehose', 'M', 'noun'),
    newWord('GNU', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Goldmark', 'M', 'noun'),
    newWord('goroutine', 'S', 'noun'),
    newWord('GPU', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('Grafana', '', 'adjective'),
    newWord('Grafana', '', 'noun') { product: true },
    newWord('Gravatar', '', 'adjective'),
    newWord('Gravatar', 'M', 'noun'),
    newWord('Graylog', 'M', 'noun'),
    newWord('GUI', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('Gzip', '', 'noun'),
    newWord('heatmap', 'S', 'noun'),
    newWord('hostname', 'S', 'noun'),
    newWord('HPA', 'S', 'noun'),
    newWord('HTML', '', 'noun') { acronym: true, established_acronym: true },
    newWord('HTTP', '', 'noun') { acronym: true, established_acronym: true },
    newWord('HTTPS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('IDE', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('inclusivity', '', 'noun'),
    newWord('ingester', 'MS', 'noun'),
    newWord('Istio', '', 'noun'),
    newWord('Jaeger', '', 'adjective'),
    newWord('Jaeger', 'M', 'noun'),
    newWord('JAR', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Jira', 'M', 'noun'),
    newWord('JMeter', 'M', 'noun'),
    newWord('JPG', '', 'noun') { acronym: true, established_acronym: true },
    newWord('JSON', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Jsonnet', 'M', 'noun'),
    newWord('JSX', '', 'noun') { acronym: true, established_acronym: true },
    newWord('JUnit', 'M', 'noun'),
    newWord('KEDA', 'M', 'noun'),
    newWord('Kibana', 'M', 'noun'),
    newWord('Kinesis', 'M', 'noun'),
    newWord('KPI', 'S', 'noun'),
    newWord('Kprobe', 'S', 'noun'),
    newWord('kubelet', '', 'noun'),
    newWord('LESS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Linode', 'M', 'noun'),
    newWord('Loki', '', 'noun') { product: true },
    newWord('lookup', 'S', 'noun'),
    newWord('marshal', 'u', 'verb'),
    newWord('matcher', 'S', 'noun'),
    newWord('memberlist', '', 'noun'),
    newWord('Mimir', 'M', 'noun') { product: true },
    newWord('misconfiguration', 'S', 'noun'),
    newWord('mixin', 'S', 'noun'),
    newWord('Moodle', 'M', 'noun'),
    newWord('namespace', 'S', 'noun'),
    newWord('nginx', '', 'noun'),
    newWord('OAuth', '', 'noun'),
    newWord('Okta', 'M', 'noun'),
    newWord('OnCall', 'M', 'noun'),
    newWord('OpenTelemetry', '', 'adjective'),
    newWord('OpenTelemetry', 'M', 'noun'),
    newWord('OSS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('OTel', '', 'adjective'),
    newWord('OTel', 'M', 'noun'),
    newWord('OTLP', '', 'noun') { acronym: true, established_acronym: true },
    newWord('overbill', 'DG', 'verb'),
    newWord('overutilization', 'S', 'noun'),
    newWord('Parca', 'M', 'noun'),
    newWord('PDF', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('performant', '', 'adjective'),
    newWord('Phlare', 'M', 'noun') { product: true },
    newWord('PHP', '', 'noun') { acronym: true, established_acronym: true },
    newWord('PNG', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Podman', 'M', 'noun'),
    newWord('preconfigure', 'D', 'verb'),
    newWord('profile', 'DGRS', 'verb'),
    newWord('Promtail', 'M', 'noun'),
    newWord('provision', 'dD', 'verb'),
    newWord('proxy', 'DG', 'verb'),
    newWord('Puppetfile', 'S', 'noun'),
    newWord('PVC', 'S', 'noun'),
    newWord('Pyroscope', '', 'noun') { product: true },
    newWord('quantile', 'S', 'noun'),
    newWord('querier', 'MS', 'noun'),
    newWord('query', '', 'noun'),
    newWord('query', 'DGS', 'verb'),
    newWord('queryable', '', 'adjective'),
    newWord('RAM', '', 'noun') { acronym: true, established_acronym: true },
    newWord('RDS', '', 'noun') { acronym: true, description: 'Amazon Relational Database Service', product: true },
    newWord('Relational Database Service', '', 'noun') { Amazon: true, description: 'Amazon Relational Database Service', product: true },
    newWord('reachability', '', 'noun') { description: 'A product metric in Synthetic Monitoring' },
    newWord('React', 'M', 'noun'),
    newWord('redirection', 'S', 'noun'),
    newWord('register', 'dG', 'verb'),
    newWord('register', 'uD', 'adjective'),
    newWord('retry', 'DGS', 'verb'),
    newWord('retryable', '', 'adjective'),
    newWord('REPL', '', 'noun') { acronym: true, established_acronym: true },
    newWord('rollout', 'S', 'noun'),
    newWord('Rollup', '', 'noun') { description: 'The JavaScript module bundler (https://rollupjs.org/)', product: true },
    newWord('RSA', '', 'noun') { acronym: true, established_acronym: true },
    newWord('runbook', 'S', 'noun'),
    newWord('sandbox', 'DG', 'verb'),
    newWord('scheduler', 'MS', 'noun') { description: 'A Kubernetes component that schedules workloads' },
    newWord('SCM', '', 'noun') { acronym: true, established_acronym: true },
    newWord('SCSS', '', 'noun') { acronym: true, established_acronym: true },
    newWord('SDK', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('SEO', '', 'noun') { acronym: true, established_acronym: true },
    newWord('serverless', '', 'adjective'),
    newWord('shard', 'DG', 'verb'),
    newWord('shortcode', 'S', 'noun'),
    newWord('showback', 'S', 'noun'),
    newWord('siloed', '', 'adjective'),
    newWord('SLA', 'S', 'noun') { acronym: true },
    newWord('SLI', 'S', 'noun') { acronym: true },
    newWord('SLO', 'S', 'noun') { acronym: true },
    newWord('spanset', 'S', 'noun'),
    newWord('Splunk', '', 'adjective'),
    newWord('Splunk', 'M', 'noun'),
    newWord('SQL', '', 'noun') { acronym: true, established_acronym: true },
    newWord('SRE', 'S', 'noun') { acronym: true },
    newWord('SSD', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('SSH', '', 'noun') { acronym: true, established_acronym: true },
    newWord('SSL', '', 'noun') { acronym: true, established_acronym: true },
    newWord('SSO', '', 'noun') { acronym: true, established_acronym: true },
    newWord('submenu', 'S', 'noun'),
    newWord('subnet', 'S', 'noun'),
    newWord('subquery', 'S', 'noun'),
    newWord('subtask', 'S', 'noun'),
    newWord('SUT', '', 'noun') { acronym: true, description: 'System Under Test' },
    newWord('SVG', '', 'noun') { acronym: true, established_acronym: true },
    newWord('systemd', '', 'noun'),
    newWord('TCP', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Tempo', '', 'noun') { product: true },
    newWord('Thanos', 'M', 'noun'),
    newWord('toolset', 'S', 'noun'),
    newWord('tooltip', 'S', 'noun'),
    newWord('traceroute', 'S', 'noun'),
    newWord('tracepoint', 'S', 'noun'),
    newWord('triage', 'D', 'verb'),
    newWord('TSDB', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('TTL', 'S', 'noun'),
    newWord('UI', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('UX', '', 'noun') { acronym: true, established_acronym: true },
    newWord('uprobe', 'S', 'noun'),
    newWord('URI', 'S', 'noun') { acronym: true, established_acronym: true },
    newWord('URL', '', 'noun') { acronym: true, established_acronym: true },
    newWord('USB', '', 'noun') { acronym: true, established_acronym: true },
    newWord('UTC', '', 'noun') { acronym: true, established_acronym: true },
    newWord('UTF', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Velero', '', 'noun') { description: 'Velero is an open source tool to safely backup and restore, perform disaster recovery, and migrate Kubernetes cluster resources and persistent volumes.', product: true },
    newWord('Vite', '', 'noun') { description: 'Next Generation Frontend Tooling (https://vitejs.dev/)', product: true },
    newWord('viewport', 'S', 'noun') { description: 'A viewport is a polygon viewing region in computer graphics (https://en.wikipedia.org/wiki/Viewport)' },
    newWord('VU', 'S', 'noun') { acronym: true },
    newWord('VUH', 'S', 'noun') { acronym: true },
    newWord('WAL', 'S', 'noun') { acronym: true },
    newWord('walkthrough', 'S', 'noun'),
    newWord('Webpack', '', 'noun'),
    newWord('webserver', 'S', 'noun'),
    newWord('windows_exporter', 'S', 'noun') { description: 'The Prometheus exporter for Windows machines (https://github.com/prometheus-community/windows_exporter)', product: true },
    newWord('worktree', '', 'noun'),
    newWord('XML', '', 'noun') { acronym: true, established_acronym: true },
    newWord('XSS', '', 'noun') { acronym: true },
    newWord('YAML', '', 'noun') { acronym: true, established_acronym: true },
    newWord('ZIP', '', 'noun') { acronym: true, established_acronym: true },
    newWord('Zipkin', '', 'adjective'),
    newWord('Zipkin', 'M', 'noun'),
  ],
}
