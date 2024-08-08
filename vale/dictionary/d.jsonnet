local word = import './word.jsonnet';
[
  word.new('data center', 'S', 'noun') { swaps: { datacenter: 'data center', datacenters: 'data centers' } },
  word.new('data source', 'S', 'noun') { swaps: { 'data-?source': 'data source', 'data-?sources': 'data sources' } },
  word.new('Data Firehose', '', 'noun') { Amazon: true, product: true, swaps: { '(?:(?<!Data )Firehose|Kinesis Data Firehose|Kinesis Firehose)': 'Data Firehose' } },
  word.new('Databricks', '', 'noun') { product: true },
  word.new('Datadog', '', 'adjective'),
  word.new('Datadog', '', 'noun') { product: true },
  word.new('deliverable', 'S', 'noun'),
  word.new('disaggregate', 'DS', 'verb'),
  word.new('distroless', '', 'adjective'),
  word.new('DOM', '', 'noun') { abbreviation: true, elaboration: 'Document Object Model', established_abbreviation: true },
  word.new("don'ts", '', 'noun'),
  word.new('downsample', 'DG', 'verb'),
  word.new('duplicate', 'dDSN', 'noun') { swaps: { 'de-duplicate': 'deduplicate', 'de-duplicated': 'deduplicated', 'de-duplicates': 'deduplicates', 'de-duplication': 'deduplication' } },
  word.new('Dynatrace', '', 'noun') { product: true },
]
