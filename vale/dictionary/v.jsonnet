local word = import './word.jsonnet';
[
  word.new('Velero', '', 'noun') { description: 'Velero is an open source tool to safely backup and restore, perform disaster recovery, and migrate Kubernetes cluster resources and persistent volumes.', product: true },
  word.new('versus', '', 'preposition') { description: 'against (https://en.wiktionary.org/wiki/versus)', swaps: { 'vs\\.': 'versus' } },
  word.new('viewport', 'S', 'noun') { description: 'A viewport is a polygon viewing region in computer graphics (https://en.wikipedia.org/wiki/Viewport)' },
  word.new('Vite', '', 'noun') { description: 'Next Generation Frontend Tooling (https://vitejs.dev/)', product: true },
  word.new('VPC', 'S', 'noun') { abbreviation: true, elaboration: 'virtual private cloud', established_abbreviation: true },
  word.new('VU', 'S', 'noun') { abbreviation: true, elaboration: 'virtual user' },
  word.new('VUH', 'S', 'noun') { abbreviation: true, elaboration: 'virtual user hour' },
]
