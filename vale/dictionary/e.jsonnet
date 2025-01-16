local word = import './word.jsonnet';
[
  word.new('eBPF', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('EKS', '', 'noun') { abbreviation: true, elaboration: 'Elastic Kubernetes Service', product: true },
  word.new('Elastic Kubernetes Service', '', 'noun') { Amazon: true, product: true },
  word.new('email', '', 'noun') { swaps: { '(?:[eE]-mail)': 'email' } },
  word.new('enablement', '', 'noun'),
  word.new('enqueue', 'DS', 'verb'),
  word.new('Entra', '', 'noun') { Azure: true, product: true },
  word.new('ESLint', 'M', 'noun'),
  word.new('etcd', '', 'noun'),
]
