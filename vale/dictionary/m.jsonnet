local word = import './word.jsonnet';
[
  word.new('manage', 'uD', 'verb'),
  word.new('marshal', 'u', 'verb'),
  word.new('matcher', 'S', 'noun'),
  word.new('memberlist', '', 'noun'),
  word.new('Mesos', '', 'noun') { Apache: true, product: true, description: 'Apache Mesos' },
  word.new('middleware', 'S', 'noun'),
  word.new('Mimir', 'M', 'noun') { product: true },
  word.new('misconfiguration', 'S', 'noun'),
  word.new('mixin', 'S', 'noun'),
  word.new('Moodle', '', 'noun') { product: true },
]
