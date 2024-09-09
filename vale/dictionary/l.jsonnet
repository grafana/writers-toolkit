local word = import './word.jsonnet';
[
  word.new('LangChain', '', 'noun') { description: 'https://www.langchain.com/', product: true, swaps: { langchain: 'LangChain' } },
  word.new('LCP', '', 'noun') { abbreviation: true, elaboration: 'Largest Contentful Paint' },
  word.new('LESS', '', 'noun') { abbreviation: true, elaboration: 'Leaner Style Sheets', established_abbreviation: true },
  word.new('LLM', 'S', 'noun') { abbreviation: true, elaboration: 'large language model' },
  word.new('Linode', 'M', 'noun'),
  word.new('LogQL', '', 'noun') { description: 'The Grafana Loki log query language.', swaps: { 'log(?:ql|QL)': 'LogQL' } },
  word.new('Loki', '', 'noun') { product: true, swaps: { loki: 'Loki' } },
  word.new('lookup', 'S', 'noun'),
]
