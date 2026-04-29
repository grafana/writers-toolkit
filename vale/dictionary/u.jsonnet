local word = import './word.jsonnet';
[
  word.new('UI', 'S', 'noun') { abbreviation: true, elaboration: 'user interface', established_abbreviation: true },
  word.new('UID', 'S', 'noun') { abbreviation: true, elaboration: 'unique identifier', established_abbreviation: true },
  word.new('uprobe', 'S', 'noun'),
  word.new('unary', '', 'adjective'),
  word.new('underprovision', 'DG', 'verb'),
  word.new('undock', 'DGS', 'verb') { description: 'To change the behavior of a Grafana dashboard sidebar so that it sits next to other dashboard elements instead of floating over them.' },
  word.new('ungroup', 'DGS', 'verb') { description: 'To remove Grafana dashboard panels from a row or tab.' },
  word.new('URI', 'S', 'noun') { abbreviation: true, elaboration: 'Uniform Resource Identifier', established_abbreviation: true },
  word.new('URL', 'S', 'noun') { abbreviation: true, elaboration: 'Uniform Resource Locator', established_abbreviation: true, swaps: { url: 'URL', urls: 'URLs' } },
  word.new('USB', '', 'noun') { abbreviation: true, elaboration: 'Universal Serial Bus', established_abbreviation: true },
  word.new('UTC', '', 'noun') { abbreviation: true, elaboration: 'Coordinated Universal Time', established_abbreviation: true },
  word.new('UTF', '', 'noun') { abbreviation: true, elaboration: 'Unicode Transformation Format', established_abbreviation: true },
  word.new('UUID', 'S', 'noun') { abbreviation: true, elaboration: 'universally unique identifier', established_abbreviation: true },
  word.new('UX', '', 'noun') { abbreviation: true, elaboration: 'user experience', established_abbreviation: true },
]
