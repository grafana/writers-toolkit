local word = import './word.jsonnet';
[
  word.new('Parca', '', 'noun') { product: true },
  word.new('PDF', 'S', 'noun') { abbreviation: true, elaboration: 'Portable Document Format', established_abbreviation: true },
  word.new('performant', '', 'adjective'),
  word.new('Phlare', 'M', 'noun') { product: true },
  word.new('PHP', '', 'noun') { abbreviation: true, elaboration: 'PHP: Hypertext Preprocessor', established_abbreviation: true },
  word.new('PNG', '', 'noun') { abbreviation: true, elaboration: 'Portable Network Graphics', established_abbreviation: true },
  word.new('Podman', '', 'noun') { product: true },
  word.new('profile', 'DGRS', 'verb'),
  word.new('Promtail', '', 'noun') { product: true },
  word.new('provision', 'dD', 'verb'),
  word.new('proxy', 'DG', 'verb'),
  word.new('Puppetfile', 'S', 'noun'),
  word.new('PVC', 'S', 'noun') { abbreviation: true, elaboration: 'Persistent Volume Claim' },
  word.new('Pyroscope', '', 'noun') { product: true },
]
