local word = import './word.jsonnet';
[

  word.new('backend', 'S', 'noun') { swaps: { 'back[ -]end': 'backend' } },
  word.new('backoff', '', 'noun') { description: 'Refers to exponential backoff (retry logic)' },
  word.new('backport', 'DGS', 'verb'),
  word.new('basemap', '', 'noun') { description: 'A type of map layer available in geomap visualizations' },
  word.new('Beyla', '', 'noun') { product: true },
  word.new('blackbox', '', 'noun'),
  word.new('blocklist', 'DGS', 'verb') { swaps: { blacklisted: 'blocklisted', blacklisting: 'blocklisting', blacklists: 'blocklists' } },
  word.new('blocklist', 'S', 'noun') { swaps: { blacklist: 'blocklist' } },
  word.new('blockquote', 'S', 'noun'),
  word.new('Bollinger', '', 'noun') { description: 'Bollinger Bands are a tool that help traders assess market volatility' },
  word.new('boolean', '', 'noun'),
  word.new('BoringCrypto', '', 'noun') { product: true, description: 'An open-source cryptographic library used by BoringSSL and other user-space applications' },
  word.new('BPF', '', 'noun') { abbreviation: true, elaboration: 'Berkeley Packet Filter', established_abbreviation: true },
  word.new('Brotli', '', 'noun') { description: 'Brotli is a lossless data compression algorithm developed by Google' },
  word.new('burndown', '', 'adjective'),
  word.new('bundler', 'S', 'noun'),
]
