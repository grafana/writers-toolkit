local word = import './word.jsonnet';
[
  word.new('IBM', '', 'noun') { abbreviation: true, established_abbreviation: true, product: true },
  word.new('ICU', '', 'noun') { abbreviation: true, elaboration: 'International Components for Unicode' },
  word.new('IDE', 'S', 'noun') { abbreviation: true, elaboration: 'integrated development environment', established_abbreviation: true },
  word.new('inclusivity', '', 'noun'),
  word.new('ingester', 'MS', 'noun'),
  word.new('Istio', '', 'noun'),
]
