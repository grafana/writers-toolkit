local word = import './word.jsonnet';
[
  word.new('IBM', '', 'noun') { abbreviation: true, established_abbreviation: true, product: true },
  word.new('ICU', '', 'noun') { abbreviation: true, elaboration: 'International Components for Unicode' },
  word.new('IDE', 'S', 'noun') { abbreviation: true, elaboration: 'integrated development environment', established_abbreviation: true },
  word.new('inclusivity', '', 'noun'),
  word.new('InfluxDB', '', 'noun') { product: true, swaps: { 'influx[Dd][Bb]': 'InfluxDB', 'Influx[Dd]b': 'InfluxDB', 'Influxd[Bb]': 'InfluxDB' } },
  word.new('ingester', 'MS', 'noun'),
  word.new('inode', 'S', 'noun') { description: 'A data structure in a Unix-style file system that describes a file-system object such as a file or a directory (https://en.wikipedia.org/wiki/Inode)' },
  word.new('instrumentation', 'S', 'noun'),
  word.new('intrinsic', 'S', 'noun'),
  word.new('Istio', '', 'noun'),
]
