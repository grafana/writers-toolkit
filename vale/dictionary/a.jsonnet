local word = import './word.jsonnet';
[
  word.new('ACL', 'S', 'noun') { abbreviation: true },
  word.new('Aerospike', '', 'noun'),
  word.new('Agent', '', 'noun') { product: true },
  word.new('Alertmanager', 'MS', 'noun'),
  word.new('allowlist', 'DGS', 'verb'),
  word.new('allowlist', 'S', 'noun'),
  word.new('Alloy', '', 'noun') { product: true },
  word.new('Ansible', '', 'adjective'),
  word.new('Ansible', '', 'noun'),
  word.new('Apdex', '', 'noun'),
  word.new('API', 'S', 'noun') { abbreviation: true, elaboration: 'Application Programming Interface', established_abbreviation: true },
  word.new('APT', '', 'noun') { abbreviation: true, description: 'https://en.wikipedia.org/wiki/APT_(software)', elaboration: 'Advanced package tool', established_abbreviation: true },
  word.new('Asserts', '', 'noun') { description: 'https://grafana.com/products/cloud/asserts/', product: true },
  word.new('autoscale', 'DGS', 'verb'),
  word.new('autoscaler', 'S', 'noun'),
  word.new('AWS', '', 'noun') { abbreviation: true, elaboration: 'Amazon Web Services', established_abbreviation: true, product: true },
]
