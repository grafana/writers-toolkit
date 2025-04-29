local word = import './word.jsonnet';
[
  word.new('FAQ', 'S', 'noun') { abbreviation: true, elaboration: 'frequently asked question', established_abbreviation: true },
  word.new('Fargate', 'M', 'noun'),
  word.new('FCP', '', 'noun') { abbreviation: true, elaboration: 'First Contentful Paint' },
  word.new('FID', '', 'noun') { abbreviation: true, elaboration: 'First Input Delay' },
  word.new('Figma', '', 'noun') { description: 'Figma design tool (https://www.figma.com/login)', product: true, swaps: { figma: 'Figma' } },
  word.new('filename', 'S', 'noun') { swaps: { 'file name': 'filename', 'file names': 'filenames' } },
  word.new('Firehose', 'M', 'noun') { Amazon: true, product: true },
  word.new('firewall rules', '', 'noun') { swaps: { firewalls: 'firewall rules' } },
  word.new('FreeBSD', '', 'noun') { description: 'FreeBSD operating system', product: true },
  word.new('frontend', 'S', 'noun') { swaps: { 'front[ -]end': 'frontend', 'front[ -]ends': 'frontends' } },
]
