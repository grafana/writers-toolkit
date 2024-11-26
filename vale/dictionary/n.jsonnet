local word = import './word.jsonnet';
[
  word.new('namespace', 'S', 'noun'),
  word.new('NAT', '', 'noun') { abbreviation: true, description: 'Network Address Translation', established_abbreviation: true },
  word.new('nginx', '', 'noun'),
]
