local word = import './word.jsonnet';
[
  word.new('KEDA', 'M', 'noun') { abbreviation: true, elaboration: 'Kubernetes-based Event-Driven Autoscaling' },
  word.new('keepalive', '', 'noun') { description: 'A message sent by one device to another to check that the link between the two is operating (https://en.wikipedia.org/wiki/Keepalive).' },
  word.new('Keycloak', '', 'noun') { description: 'Open source identity and access management solution.', product: true },
  word.new('Kibana', '', 'noun') { product: true },
  word.new('Killercoda', '', 'noun') { product: true },
  word.new('Kinesis', '', 'noun') { Amazon: true, product: true },
  word.new('Kotlin', '', 'noun') { product: true, swaps: { 'kotlin': 'Kotlin' } },
  word.new('KPI', 'S', 'noun') { abbreviation: true, elaboration: 'key performance indicator' },
  word.new('KQL', '', 'noun') { Azure: true, product: true, elaboration: 'Kusto Query Language' },
  word.new('Kprobe', 'S', 'noun'),
  word.new('kubelet', '', 'noun'),
  word.new('Kubernetes', '', 'noun') { product: true, swap: '(?:[kK]8s|kubernetes)' },
  word.new('Kubernetes Engine', '', 'noun') { Google: true, product: true },
  word.new('Kusto', '', 'noun') { Azure: true, description: 'An Azure query language.', product: true },
  word.new('Kustomize', '', 'noun') { description: 'A Kubernetes native configuration management tool.', product: true },
]
