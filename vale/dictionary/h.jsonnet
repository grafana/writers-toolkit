local word = import './word.jsonnet';
[
  word.new('Hashicorp', '', 'noun') { elaboration: 'A company that produces a suite of infrastructure automation products.' },
  word.new('hashmod', '', 'noun') { description: 'A program that distributes an enabled / disabled flag across a cluster without coordination, using consistent hashing.', product: true },
  word.new('heatmap', 'S', 'noun'),
  word.new('hostname', 'S', 'noun'),
  word.new('HPA', 'S', 'noun'),
  word.new('HTML', '', 'noun') { abbreviation: true, elaboration: 'HyperText Markup Language', established_abbreviation: true },
  word.new('HTTP', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol', established_abbreviation: true },
  word.new('HTTPS', '', 'noun') { abbreviation: true, elaboration: 'HyperText Transfer Protocol Secure', established_abbreviation: true, swaps: { HTTPs: 'HTTPS' } },
]
