local word = import './word.jsonnet';
[
  word.new('ZIP', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('Zipkin', '', 'adjective'),
  word.new('Zipkin', '', 'noun') { product: true },
  word.new('zlib', '', 'noun') { description: 'zlib is a general-purpose lossless data-compression library.' },
]
