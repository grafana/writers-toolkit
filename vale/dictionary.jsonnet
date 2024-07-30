// The structure of the word object is documented in https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/#word-metadata.
local newWord(word, affixes, po) = {
  abbreviation: false,
  affixes: affixes,
  Amazon: false,
  Apache: false,
  Google: false,
  description: null,
  established_abbreviation: false,
  po: po,
  product: false,
  word: word,
};

{
  words: [
    newWord('ACL', 'S', 'noun') { abbreviation: true },
    newWord('Aerospike', '', 'noun'),
    newWord('Agent', '', 'noun') { product: true },
    newWord('Alertmanager', 'MS', 'noun'),
    newWord('allowlist', 'DGS', 'verb'),
    newWord('allowlist', 'S', 'noun'),
    newWord('Alloy', '', 'noun') { product: true },
    newWord('Ansible', '', 'adjective'),
    newWord('Ansible', '', 'noun'),
    newWord('Apdex', '', 'noun'),
    newWord('API', 'S', 'noun') { abbreviation: true, elaboration: 'Application Programming Interface', established_abbreviation: true },
    newWord('APT', '', 'noun') { abbreviation: true, description: 'https://en.wikipedia.org/wiki/APT_(software)', elaboration: 'Advanced package tool', established_abbreviation: true },
    newWord('Asserts', '', 'noun') { description: 'https://grafana.com/products/cloud/asserts/', product: true },
    newWord('autoscale', 'DGS', 'verb'),
    newWord('autoscaler', 'S', 'noun'),
    newWord('AWS', '', 'noun') { abbreviation: true, elaboration: 'Amazon Web Services', established_abbreviation: true, product: true },
    newWord('backport', 'DGS', 'verb'),
    newWord('Beyla', '', 'noun') { product: true },
    newWord('blackbox', '', 'noun'),
    newWord('blockquote', 'S', 'noun'),
    newWord('boolean', '', 'noun'),
    newWord('BPF', '', 'noun') { abbreviation: true, elaboration: 'Berkeley Packet Filter', established_abbreviation: true },
    newWord('Brotli', '', 'noun') { description: 'Brotli is a lossless data compression algorithm developed by Google' },
    newWord('burndown', '', 'adjective'),
    newWord('bundler', 'S', 'noun'),
    newWord('cAdvisor', 'M', 'noun'),
    newWord('callee', 'S', 'noun') { description: 'A function that is called by another function' },
    newWord('CDN', '', 'noun') { abbreviation: true, elaboration: 'Content Delivery Network' },
    newWord('CLA', '', 'noun') { abbreviation: true, elaboration: 'Contributor License Agreement' },
    newWord('CLI', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('CloudWatch', '', 'noun') { Amazon: true, product: true },
    newWord('codespace', 'S', 'noun') { description: "A codespace is a development environment that's hosted in the cloud. https://docs.github.com/en/codespaces/overview" },
    newWord('Codespaces', '', 'noun') { description: 'GitHub Codespaces https://docs.github.com/en/codespaces/overview', product: true },
    newWord('composable', '', 'adjective'),
    newWord('configure', 'mpDS', 'verb'),
    newWord('contentful', '', 'adjective') { description: 'Having content. Used in Web Vital metrics, such as Largest Contentful Paint: https://web.dev/articles/lcp' },
    newWord('CORS', '', 'noun') { abbreviation: true, established_abbreviation: true, description: 'Cross-origin resource sharing. Allows a web page to access restricted resources from a server on a domain different than the domain that served the web page.' },
    newWord('Couchbase', 'M', 'noun'),
    newWord('CLS', '', 'noun') { abbreviation: true, elaboration: 'Cumulative Layout Shift' },
    newWord('Data Firehose', '', 'noun') { Amazon: true, product: true },
    newWord('CPU', 'S', 'noun') { abbreviation: true, elaboration: 'central processing unit', established_abbreviation: true },
    newWord('CRD', 'S', 'noun') { abbreviation: true, elaboration: 'Custom Resource Definition' },
    newWord('CSS', '', 'noun') { abbreviation: true, elaboration: 'Cascading Style Sheets', established_abbreviation: true },
    newWord('CSV', '', 'noun') { abbreviation: true, elaboration: 'Comma-separated values', established_abbreviation: true },
    newWord('CVE', 'S', 'noun'),
    newWord('Databricks', '', 'noun') { product: true },
    newWord('Datadog', '', 'adjective'),
    newWord('Datadog', '', 'noun') { product: true },
    newWord('deliverable', 'S', 'noun'),
    newWord('disaggregate', 'DS', 'verb'),
    newWord('distroless', '', 'adjective'),
    newWord('DOM', '', 'noun') { abbreviation: true, elaboration: 'Document Object Model', established_abbreviation: true },
    newWord("don'ts", '', 'noun'),
    newWord('downsample', 'DG', 'verb'),
    newWord('duplicate', 'dDSN', 'noun'),
    newWord('Dynatrace', '', 'noun') { product: true },
    newWord('eBPF', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('EKS', '', 'noun') { abbreviation: true, elaboration: 'Elastic Kubernetes Service', product: true },
    newWord('Elastic Kubernetes Service', '', 'noun') { Amazon: true, product: true },
    newWord('enablement', '', 'noun'),
    newWord('enqueue', 'DS', 'verb'),
    newWord('ESLint', 'M', 'noun'),
    newWord('etcd', '', 'noun'),
    newWord('FAQ', 'S', 'noun') { abbreviation: true, elaboration: 'frequently asked question', established_abbreviation: true },
    newWord('Fargate', 'M', 'noun'),
    newWord('FCP', '', 'noun') { abbreviation: true, elaboration: 'First Contentful Paint' },
    newWord('FID', '', 'noun') { abbreviation: true, elaboration: 'First Input Delay' },
    newWord('Firehose', 'M', 'noun') { Amazon: true, product: true },
    newWord('glob', 'G', 'noun') { description: 'In computer programming, glob patterns specify sets of filenames with wildcard characters.' },
    newWord('globbing', '', 'verb') { description: 'In computer programming, glob patterns specify sets of filenames with wildcard characters.' },
    newWord('GNU', '', 'noun') { abbreviation: true, established_abbreviation: true, product: true },
    newWord('Goldmark', '', 'noun') { description: 'Goldmark is a markdown parser written in Go (https://github.com/yuin/goldmark)' },
    newWord('Kubernetes Engine', '', 'noun') { Google: true, product: true },
    newWord('GKE', '', 'noun') { abbreviation: true, elaboration: 'Google Kubernetes Engine', product: true },  // Doesn't need Google: true because Google is in the name.
    newWord('goroutine', 'S', 'noun'),
    newWord('GPU', 'S', 'noun') { abbreviation: true, elaboration: 'graphics processing unit', established_abbreviation: true },
    newWord('Grafana', '', 'adjective'),
    newWord('Grafana', '', 'noun') { product: true },
    newWord('Gravatar', '', 'adjective'),
    newWord('Gravatar', '', 'noun') { product: true },
    newWord('Graylog', '', 'noun') { product: true },
    newWord('GUI', 'S', 'noun') { abbreviation: true, elaboration: 'graphical user interface', established_abbreviation: true },
    newWord('Gzip', '', 'noun'),
    newWord('heatmap', 'S', 'noun'),
    newWord('hostname', 'S', 'noun'),
    newWord('HPA', 'S', 'noun'),
    newWord('HTML', '', 'noun') { abbreviation: true, elaboration: 'HyperText Markup Language', established_abbreviation: true },
    newWord('HTTP', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol', established_abbreviation: true },
    newWord('HTTPS', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol Secure', established_abbreviation: true },
    newWord('IBM', '', 'noun') { abbreviation: true, established_abbreviation: true, product: true },
    newWord('IDE', 'S', 'noun') { abbreviation: true, elaboration: 'integrated development environment', established_abbreviation: true },
    newWord('inclusivity', '', 'noun'),
    newWord('ingester', 'MS', 'noun'),
    newWord('Istio', '', 'noun'),
    newWord('Jaeger', '', 'adjective'),
    newWord('Jaeger', '', 'noun') { product: true },
    newWord('JAR', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('Jira', '', 'noun') { product: true },
    newWord('JMeter', 'M', 'noun'),
    newWord('JPG', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('JSON', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('Jsonnet', 'M', 'noun'),
    newWord('JSX', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('JUnit', 'M', 'noun'),
    newWord('KEDA', 'M', 'noun') { abbreviation: true, elaboration: 'Kubernetes-based Event-Driven Autoscaling' },
    newWord('keepalive', '', 'noun') { description: 'A message sent by one device to another to check that the link between the two is operating (https://en.wikipedia.org/wiki/Keepalive).' },
    newWord('Kibana', '', 'noun') { product: true },
    newWord('Killercoda', '', 'noun') { product: true },
    newWord('Kinesis', '', 'noun') { Amazon: true, product: true },
    newWord('KPI', 'S', 'noun') { abbreviation: true, elaboration: 'key performance indicator' },
    newWord('Kprobe', 'S', 'noun'),
    newWord('kubelet', '', 'noun'),
    newWord('LCP', '', 'noun') { abbreviation: true, elaboration: 'Largest Contentful Paint' },
    newWord('LESS', '', 'noun') { abbreviation: true, elaboration: 'Leaner Style Sheets', established_abbreviation: true },
    newWord('Linode', 'M', 'noun'),
    newWord('Loki', '', 'noun') { product: true },
    newWord('lookup', 'S', 'noun'),
    newWord('marshal', 'u', 'verb'),
    newWord('manage', 'uD', 'verb'),
    newWord('matcher', 'S', 'noun'),
    newWord('memberlist', '', 'noun'),
    newWord('Mesos', '', 'noun') { Apache: true, product: true, description: 'Apache Mesos' },
    newWord('Mimir', 'M', 'noun') { product: true },
    newWord('misconfiguration', 'S', 'noun'),
    newWord('mixin', 'S', 'noun'),
    newWord('Moodle', '', 'noun') { product: true },
    newWord('namespace', 'S', 'noun'),
    newWord('nginx', '', 'noun'),
    newWord('OAuth', '', 'noun'),
    newWord('Okta', '', 'noun') { product: true },
    newWord('OnCall', '', 'noun') { product: true },
    newWord('OpenTelemetry', '', 'adjective'),
    newWord('OpenTelemetry', '', 'noun') { product: true },
    newWord('OSS', '', 'noun') { abbreviation: true, elaboration: 'open source software', established_abbreviation: true },
    newWord('OTel', '', 'adjective'),
    newWord('OTel', '', 'noun') { product: true },
    newWord('OTLP', '', 'noun') { abbreviation: true, elaboration: 'OpenTelemetry Protocol', established_abbreviation: true },
    newWord('overbill', 'DG', 'verb'),
    newWord('overutilization', 'S', 'noun'),
    newWord('Parca', '', 'noun') { product: true },
    newWord('PDF', 'S', 'noun') { abbreviation: true, elaboration: 'Portable Document Format', established_abbreviation: true },
    newWord('performant', '', 'adjective'),
    newWord('Phlare', 'M', 'noun') { product: true },
    newWord('PHP', '', 'noun') { abbreviation: true, elaboration: 'PHP: Hypertext Preprocessor', established_abbreviation: true },
    newWord('PNG', '', 'noun') { abbreviation: true, elaboration: 'Portable Network Graphics', established_abbreviation: true },
    newWord('Podman', '', 'noun') { product: true },
    newWord('profile', 'DGRS', 'verb'),
    newWord('Promtail', '', 'noun') { product: true },
    newWord('provision', 'dD', 'verb'),
    newWord('proxy', 'DG', 'verb'),
    newWord('Puppetfile', 'S', 'noun'),
    newWord('PVC', 'S', 'noun') { abbreviation: true, elaboration: 'Persistent Volume Claim' },
    newWord('Pyroscope', '', 'noun') { product: true },
    newWord('quantile', 'S', 'noun'),
    newWord('querier', 'MS', 'noun'),
    newWord('query', '', 'noun'),
    newWord('query', 'DGS', 'verb'),
    newWord('queryless', '', 'adjective') { description: 'In contrast to using a query language like PromQL or SQL' },
    newWord('queryable', '', 'adjective'),
    newWord('RAM', '', 'noun') { abbreviation: true, elaboration: 'random access memory', established_abbreviation: true },
    newWord('RBAC', '', 'noun') { abbreviation: true, elaboration: 'Role-based Access Control', established_abbreviation: true },
    newWord('RDS', '', 'noun') { abbreviation: true, elaboration: 'Relational Database Service', product: true },
    newWord('Relational Database Service', '', 'noun') { Amazon: true, description: 'Amazon Relational Database Service', product: true },
    newWord('reachability', '', 'noun') { description: 'A product metric in Synthetic Monitoring' },
    newWord('React', '', 'noun') { product: true },
    newWord('redirection', 'S', 'noun'),
    newWord('register', 'dG', 'verb'),
    newWord('register', 'uD', 'adjective'),
    newWord('retry', 'DGS', 'verb'),
    newWord('retryable', '', 'adjective'),
    newWord('REPL', '', 'noun') { abbreviation: true, elaboration: 'read-eval-print loop', established_abbreviation: true },
    newWord('rollout', 'S', 'noun'),
    newWord('Rollup', '', 'noun') { description: 'The JavaScript module bundler (https://rollupjs.org/)', product: true },
    newWord('RPM', '', 'noun') { abbreviation: true, description: 'https://en.wikipedia.org/wiki/RPM_Package_Manager', elaboration: 'RPM Package Manager', established_abbreviation: true },
    newWord('RSA', '', 'noun') { abbreviation: true, elaboration: 'Rivest–Shamir–Adleman', established_abbreviation: true },
    newWord('runbook', 'S', 'noun'),
    newWord('sandbox', 'DG', 'verb'),
    newWord('scheduler', 'MS', 'noun') { description: 'A Kubernetes component that schedules workloads' },
    newWord('SCM', '', 'noun') { abbreviation: true, elaboration: 'source code management', established_abbreviation: true },
    newWord('SCSS', '', 'noun') { abbreviation: true, elaboration: 'Sassy CSS', established_abbreviation: true },
    newWord('SDK', 'S', 'noun') { abbreviation: true, elaboration: 'software development kit', established_abbreviation: true },
    newWord('SEO', '', 'noun') { abbreviation: true, elaboration: 'search engine optimization', established_abbreviation: true },
    newWord('serverless', '', 'adjective'),
    newWord('shard', 'DG', 'verb'),
    newWord('shortcode', 'S', 'noun'),
    newWord('showback', 'S', 'noun'),
    newWord('siloed', '', 'adjective'),
    newWord('SLA', 'S', 'noun') { abbreviation: true, elaboration: 'service level agreement' },
    newWord('SLI', 'S', 'noun') { abbreviation: true, elaboration: 'service level indicator' },
    newWord('SLO', 'S', 'noun') { abbreviation: true, elaboration: 'service level objective' },
    newWord('spanset', 'S', 'noun'),
    newWord('Splunk', '', 'adjective'),
    newWord('Splunk', '', 'noun') { product: true },
    newWord('SQL', '', 'noun') { abbreviation: true, elaboration: 'Structured Query Language', established_abbreviation: true },
    newWord('SRE', 'S', 'noun') { abbreviation: true, elaboration: 'site reliability engineering' },
    newWord('SSD', 'S', 'noun') { abbreviation: true, elaboration: 'solid-state drive', established_abbreviation: true },
    newWord('SSH', '', 'noun') { abbreviation: true, elaboration: 'Secure Shell', established_abbreviation: true },
    newWord('SSL', '', 'noun') { abbreviation: true, elaboration: 'Secure Sockets Layer', established_abbreviation: true },
    newWord('SSO', '', 'noun') { abbreviation: true, elaboration: 'single sign-on', established_abbreviation: true },
    newWord('submenu', 'S', 'noun'),
    newWord('subnet', 'S', 'noun'),
    newWord('subquery', 'S', 'noun'),
    newWord('subtask', 'S', 'noun'),
    newWord('SUT', '', 'noun') { abbreviation: true, elaboration: 'System Under Test' },
    newWord('SVG', '', 'noun') { abbreviation: true, elaboration: 'Scalable Vector Graphics', established_abbreviation: true },
    newWord('systemd', '', 'noun'),
    newWord('TCP', '', 'noun') { abbreviation: true, elaboration: 'Transmission Control Protocol', established_abbreviation: true },
    newWord('Tempo', '', 'noun') { product: true },
    newWord('Thanos', '', 'noun') { product: true },
    newWord('TLS', '', 'noun') { abbreviation: true, description: 'A cryptographic protocol designed to provide secure communications over network.', elaboration: 'Transport Layer Security', established_abbreviation: true },
    newWord('toolset', 'S', 'noun'),
    newWord('tooltip', 'S', 'noun'),
    newWord('traceroute', 'S', 'noun'),
    newWord('tracepoint', 'S', 'noun'),
    newWord('triage', 'D', 'verb'),
    newWord('TSDB', 'S', 'noun') { abbreviation: true, elaboration: 'time-series database', established_abbreviation: true },
    newWord('TTL', 'S', 'noun') { abbreviation: true, elaboration: 'time to live' },
    newWord('UI', 'S', 'noun') { abbreviation: true, elaboration: 'user interface', established_abbreviation: true },
    newWord('UX', '', 'noun') { abbreviation: true, elaboration: 'user experience', established_abbreviation: true },
    newWord('uprobe', 'S', 'noun'),
    newWord('URI', 'S', 'noun') { abbreviation: true, elaboration: 'Uniform Resource Identifier', established_abbreviation: true },
    newWord('URL', '', 'noun') { abbreviation: true, elaboration: 'Uniform Resource Locator', established_abbreviation: true },
    newWord('USB', '', 'noun') { abbreviation: true, elaboration: 'Universal Serial Bus', established_abbreviation: true },
    newWord('UTC', '', 'noun') { abbreviation: true, elaboration: 'Coordinated Universal Time', established_abbreviation: true },
    newWord('UTF', '', 'noun') { abbreviation: true, elaboration: 'Unicode Transformation Format', established_abbreviation: true },
    newWord('Velero', '', 'noun') { description: 'Velero is an open source tool to safely backup and restore, perform disaster recovery, and migrate Kubernetes cluster resources and persistent volumes.', product: true },
    newWord('Vite', '', 'noun') { description: 'Next Generation Frontend Tooling (https://vitejs.dev/)', product: true },
    newWord('viewport', 'S', 'noun') { description: 'A viewport is a polygon viewing region in computer graphics (https://en.wikipedia.org/wiki/Viewport)' },
    newWord('VPC', 'S', 'noun') { abbreviation: true, elaboration: 'virtual private cloud', established_abbreviation: true },
    newWord('VU', 'S', 'noun') { abbreviation: true, elaboration: 'virtual user' },
    newWord('VUH', 'S', 'noun') { abbreviation: true, elaboration: 'virtual user hour' },
    newWord('WAL', 'S', 'noun') { abbreviation: true, elaboration: 'write-ahead log' },
    newWord('walkthrough', 'S', 'noun'),
    newWord('Webpack', '', 'noun'),
    newWord('webserver', 'S', 'noun'),
    newWord('WildFly', '', 'noun') { description: 'WildFly, formerly known as JBoss AS, or simply JBoss, is an application server written by JBoss, now developed by Red Hat (https://en.wikipedia.org/wiki/WildFly)', product: true },
    newWord('windows_exporter', 'S', 'noun') { description: 'The Prometheus exporter for Windows machines (https://github.com/prometheus-community/windows_exporter)', product: true },
    newWord('worktree', '', 'noun'),
    newWord('XML', '', 'noun') { abbreviation: true, elaboration: 'Extensible Markup Language', established_abbreviation: true },
    newWord('XSS', '', 'noun') { abbreviation: true, elaboration: 'cross-site scripting' },
    newWord('YAML', '', 'noun') { abbreviation: true, elaboration: "YAML Ain't Markup Language", established_abbreviation: true },
    newWord('ZIP', '', 'noun') { abbreviation: true, established_abbreviation: true },
    newWord('Zipkin', '', 'adjective'),
    newWord('Zipkin', '', 'noun') { product: true },
    newWord('zlib', '', 'noun') { description: 'zlib is a general-purpose lossless data-compression library.' },
  ],
}
