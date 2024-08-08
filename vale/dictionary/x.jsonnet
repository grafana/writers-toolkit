local word = import './word.jsonnet';
[
  word.new('XML', '', 'noun') { abbreviation: true, elaboration: 'Extensible Markup Language', established_abbreviation: true },
  word.new('XSS', '', 'noun') { abbreviation: true, elaboration: 'cross-site scripting' },
]
