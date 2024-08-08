local word = import './word.jsonnet';
[
  word.new('Jaeger', '', 'adjective'),
  word.new('Jaeger', '', 'noun') { product: true },
  word.new('JAR', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('Jira', '', 'noun') { product: true },
  word.new('JMeter', 'M', 'noun'),
  word.new('JPG', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('JSON', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('Jsonnet', 'M', 'noun'),
  word.new('JSX', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('JUnit', 'M', 'noun'),
]
