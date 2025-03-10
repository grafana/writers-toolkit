local word = import './word.jsonnet';
[
  word.new('TCP', '', 'noun') { abbreviation: true, elaboration: 'Transmission Control Protocol', established_abbreviation: true },
  word.new('templated', '', 'adjective'),
  word.new('Tempo', '', 'noun') { product: true, swaps: { tempo: 'Tempo' } },
  word.new('Thanos', '', 'noun') { product: true },
  word.new('Threema', '', 'noun') { description: 'https://threema.ch/en', product: true, swaps: { threema: 'Threema' } },
  word.new('TLS', '', 'noun') { abbreviation: true, description: 'A cryptographic protocol designed to provide secure communications over network.', elaboration: 'Transport Layer Security', established_abbreviation: true },
  word.new('toolset', 'S', 'noun'),
  word.new('tooltip', 'S', 'noun'),
  word.new('tracepoint', 'S', 'noun'),
  word.new('TraceQL', '', 'noun') { swaps: { 'trace(?:ql|QL)': 'TraceQL' } },
  word.new('traceroute', 'S', 'noun'),
  word.new('Traces Drilldown', '', 'noun') { product: true },
  word.new('triage', 'D', 'verb'),
  word.new('TSDB', 'S', 'noun') { abbreviation: true, elaboration: 'time-series database', established_abbreviation: true },
  word.new('TTL', 'S', 'noun') { abbreviation: true, elaboration: 'time to live' },
]
