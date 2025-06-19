local word = import './word.jsonnet';
[
  word.new('YAML', '', 'noun') { abbreviation: true, elaboration: "YAML Ain't Markup Language", established_abbreviation: true },
  word.new('YugabyteDB', '', 'noun') { product: true, description: "https://docs.yugabyte.com/", swaps: {yugabyte: 'YugabyteDB'} },
]
