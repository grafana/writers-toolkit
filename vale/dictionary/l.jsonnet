local word = import './word.jsonnet';
[
  word.new('LangChain', '', 'noun') { description: 'https://www.langchain.com/', product: true, swaps: { langchain: 'LangChain' } },
  word.new('launchd', '', 'noun') { description: 'An open-source service management framework used with macOS.', product: true },
  word.new('LCP', '', 'noun') { abbreviation: true, elaboration: 'Largest Contentful Paint' },
  word.new('LESS', '', 'noun') { abbreviation: true, elaboration: 'Leaner Style Sheets', established_abbreviation: true },
  word.new('LLM', 'S', 'noun') { abbreviation: true, elaboration: 'large language model' },
  word.new('Linode', 'M', 'noun'),
  word.new('LogQL', '', 'noun') { description: 'The Grafana Loki log query language.', swaps: { 'log(?:ql|QL)': 'LogQL' } },
  word.new('Logs Drilldown', '', 'noun') { product: true },
  word.new('Loki', '', 'noun') { product: true, swaps: { loki: 'Loki' } },
  word.new('lookup', 'S', 'noun'),
  word.new('loopback', '', 'noun') { description: 'A network interface that is used to send data back to the same device.' },
  word.new('Lucene', '', 'noun') { product: true, swaps: { lucene: 'Lucene' } },
]
