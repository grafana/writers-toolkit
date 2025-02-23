local word = import './word.jsonnet';
[
  word.new('manage', 'uD', 'verb'),
  word.new('Markdown', '', 'noun') { product: true, swaps: { markdown: 'Markdown' } },
  word.new('marshal', 'u', 'verb'),
  word.new('matcher', 'S', 'noun'),
  word.new('media type', 'S', 'noun') { swaps: { '(?:content|media)-?type': 'media type', 'content type': 'media type' } },
  word.new('memberlist', '', 'noun'),
  word.new('Memcached', '', 'noun') { product: true, swaps: { memcached: 'Memcached' } },
  word.new('metadata', '', 'noun') { swaps: { 'meta[- ]data': 'metadata' } },
  word.new('Metrics Drilldown', '', 'noun') { product: true },
  word.new('Mesos', '', 'noun') { Apache: true, product: true, description: 'Apache Mesos' },
  word.new('middleware', 'S', 'noun'),
  word.new('Mimir', 'M', 'noun') { product: true },
  word.new('misconfiguration', 'S', 'noun'),
  word.new('mixin', 'S', 'noun') { swaps: { 'mix[- ]in': 'mixin' } },
  word.new('Moodle', '', 'noun') { product: true },
  word.new('MySQL', '', 'noun') { product: true, swaps: { mysql: 'MySQL' } },
]
