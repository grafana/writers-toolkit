local word = import './word.jsonnet';
[
  word.new('YAML', '', 'noun') { abbreviation: true, elaboration: "YAML Ain't Markup Language", established_abbreviation: true },
  word.new('Yugabyte', '', 'noun') { product: true },
]
