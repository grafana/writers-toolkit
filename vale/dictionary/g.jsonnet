local word = import './word.jsonnet';
[
  word.new('GEM', '', 'noun') { abbreviation: true, elaboration: 'Grafana Enterprise Metrics', product: true, swaps: { GME: 'GEM' } },
  word.new('GKE', '', 'noun') { abbreviation: true, elaboration: 'Google Kubernetes Engine', product: true },  // Doesn't need Google: true because Google is in the name.
  word.new('Git', '', 'noun') { product: true, swaps: { git: 'Git' } },
  word.new('GitHub', '', 'noun') { product: true, swaps: { '(?:github|gitHub|Github)': 'GitHub' } },
  word.new('GitLab', '', 'noun') { product: true, swaps: { '(?:gitlab|gitLab|Gitlab)': 'GitLab' } },
  word.new('glob', 'G', 'noun') { description: 'In computer programming, glob patterns specify sets of filenames with wildcard characters.' },
  word.new('globbing', '', 'verb') { description: 'In computer programming, glob patterns specify sets of filenames with wildcard characters.' },
  word.new('GNU', '', 'noun') { abbreviation: true, established_abbreviation: true, product: true },
  word.new('Goldmark', '', 'noun') { description: 'Goldmark is a markdown parser written in Go (https://github.com/yuin/goldmark)' },
  word.new('goroutine', 'S', 'noun'),
  word.new('GPG', '', 'noun') { abbreviation: true, description: 'GNU Privacy Guard', established_abbreviation: true },
  word.new('GPU', 'S', 'noun') { abbreviation: true, elaboration: 'graphics processing unit', established_abbreviation: true },
  word.new('Grafana', '', 'adjective'),
  word.new('Grafana', '', 'noun') { product: true, swaps: { grafana: 'Grafana' } },
  word.new('Gravatar', '', 'adjective'),
  word.new('Gravatar', '', 'noun') { product: true },
  word.new('Graylog', '', 'noun') { product: true },
  word.new('GUI', 'S', 'noun') { abbreviation: true, elaboration: 'graphical user interface', established_abbreviation: true },
  word.new('Gunicorn', '', 'noun') { product: true },
  word.new('Gzip', '', 'noun'),
]
