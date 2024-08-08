local word = import './word.jsonnet';
[
  word.new('OAuth', '', 'noun'),
  word.new('Okta', '', 'noun') { product: true },
  word.new('OnCall', '', 'noun') { product: true },
  word.new('OpenTelemetry', '', 'adjective'),
  word.new('OpenTelemetry', '', 'noun') { product: true },
  word.new('OSS', '', 'noun') { abbreviation: true, elaboration: 'open source software', established_abbreviation: true },
  word.new('OTel', '', 'adjective'),
  word.new('OTel', '', 'noun') { product: true },
  word.new('OTLP', '', 'noun') { abbreviation: true, elaboration: 'OpenTelemetry Protocol', established_abbreviation: true },
  word.new('overbill', 'DG', 'verb'),
  word.new('overutilization', 'S', 'noun'),
]
