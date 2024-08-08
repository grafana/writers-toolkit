local word = import './word.jsonnet';
[
  word.new('quantile', 'S', 'noun'),
  word.new('querier', 'MS', 'noun'),
  word.new('query', '', 'noun'),
  word.new('query', 'DGS', 'verb'),
  word.new('queryable', '', 'adjective'),
  word.new('queryless', '', 'adjective') { description: 'In contrast to using a query language like PromQL or SQL' },
]
