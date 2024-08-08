local word = import './word.jsonnet';
[
  word.new('KEDA', 'M', 'noun') { abbreviation: true, elaboration: 'Kubernetes-based Event-Driven Autoscaling' },
  word.new('keepalive', '', 'noun') { description: 'A message sent by one device to another to check that the link between the two is operating (https://en.wikipedia.org/wiki/Keepalive).' },
  word.new('Kibana', '', 'noun') { product: true },
  word.new('Killercoda', '', 'noun') { product: true },
  word.new('Kinesis', '', 'noun') { Amazon: true, product: true },
  word.new('KPI', 'S', 'noun') { abbreviation: true, elaboration: 'key performance indicator' },
  word.new('Kprobe', 'S', 'noun'),
  word.new('kubelet', '', 'noun'),
  word.new('Kubernetes Engine', '', 'noun') { Google: true, product: true },
]
