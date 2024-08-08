local word = import './word.jsonnet';
[
  word.new('heatmap', 'S', 'noun'),
  word.new('hostname', 'S', 'noun'),
  word.new('HPA', 'S', 'noun'),
  word.new('HTML', '', 'noun') { abbreviation: true, elaboration: 'HyperText Markup Language', established_abbreviation: true },
  word.new('HTTP', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol', established_abbreviation: true },
  word.new('HTTPS', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol Secure', established_abbreviation: true, swaps: { HTTPs: 'HTTPS' } },
]
