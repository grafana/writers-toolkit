local word = import './word.jsonnet';
[
  word.new('backend', 'S', 'noun') { swaps: { 'back[ -]end': 'backend' } },
  word.new('backport', 'DGS', 'verb'),
  word.new('Beyla', '', 'noun') { product: true },
  word.new('blackbox', '', 'noun'),
  word.new('blocklist', 'DGS', 'verb') { swaps: { blacklisted: 'blocklisted', blacklisting: 'blocklisting', blacklists: 'blocklists' } },
  word.new('blocklist', 'S', 'noun') { swaps: { blacklist: 'blocklist' } },
  word.new('blockquote', 'S', 'noun'),
  word.new('boolean', '', 'noun'),
  word.new('BPF', '', 'noun') { abbreviation: true, elaboration: 'Berkeley Packet Filter', established_abbreviation: true },
  word.new('Brotli', '', 'noun') { description: 'Brotli is a lossless data compression algorithm developed by Google' },
  word.new('burndown', '', 'adjective'),
  word.new('bundler', 'S', 'noun'),
]
