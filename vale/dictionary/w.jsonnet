local word = import './word.jsonnet';
[
  word.new('WAL', 'S', 'noun') { abbreviation: true, elaboration: 'write-ahead log' },
  word.new('walkthrough', 'S', 'noun'),
  word.new('web', '', 'noun') { swaps: { '[Ww]orld [Ww]ide [Ww]eb': 'web' } },
  word.new('Webpack', '', 'noun'),
  word.new('webserver', 'S', 'noun'),
  word.new('Webex', '', 'noun') { description: 'https://www.webex.com/', product: true, swaps: { webex: 'Webex' } },
  word.new('Wi-Fi', '', 'noun') { description: 'wireless fidelity', swaps: { '(?:WiFi|wifi)': 'Wi-Fi' } },
  word.new('WildFly', '', 'noun') { description: 'WildFly, formerly known as JBoss AS, or simply JBoss, is an application server written by JBoss, now developed by Red Hat (https://en.wikipedia.org/wiki/WildFly)', product: true },
  word.new('windows_exporter', 'S', 'noun') { description: 'The Prometheus exporter for Windows machines (https://github.com/prometheus-community/windows_exporter)', product: true },
  word.new('worktree', '', 'noun'),
]
