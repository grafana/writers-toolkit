local word = import './word.jsonnet';
[
  word.new('Data Firehose', '', 'noun') { Amazon: true, product: true },
  word.new('Databricks', '', 'noun') { product: true },
  word.new('Datadog', '', 'adjective'),
  word.new('Datadog', '', 'noun') { product: true },
  word.new('deliverable', 'S', 'noun'),
  word.new('disaggregate', 'DS', 'verb'),
  word.new('distroless', '', 'adjective'),
  word.new('DOM', '', 'noun') { abbreviation: true, elaboration: 'Document Object Model', established_abbreviation: true },
  word.new("don'ts", '', 'noun'),
  word.new('downsample', 'DG', 'verb'),
  word.new('duplicate', 'dDSN', 'noun'),
  word.new('Dynatrace', '', 'noun') { product: true },
]
