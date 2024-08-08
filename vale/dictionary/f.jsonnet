local word = import './word.jsonnet';
[
  word.new('FAQ', 'S', 'noun') { abbreviation: true, elaboration: 'frequently asked question', established_abbreviation: true },
  word.new('Fargate', 'M', 'noun'),
  word.new('FCP', '', 'noun') { abbreviation: true, elaboration: 'First Contentful Paint' },
  word.new('FID', '', 'noun') { abbreviation: true, elaboration: 'First Input Delay' },
  word.new('Firehose', 'M', 'noun') { Amazon: true, product: true },
]
