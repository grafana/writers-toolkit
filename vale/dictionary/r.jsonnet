local word = import './word.jsonnet';
[
  word.new('RAM', '', 'noun') { abbreviation: true, elaboration: 'random access memory', established_abbreviation: true },
  word.new('RBAC', '', 'noun') { abbreviation: true, elaboration: 'Role-based Access Control', established_abbreviation: true },
  word.new('RCA workbench', '', 'noun') { description: 'RCA workbench provides a comprehensive exploration of potential causes for specific issues over time and dependencies, offering access to relevant metrics, logs, and traces for efficient troubleshooting and analysis.', product: true, swaps: { 'RCA Workbench': 'RCA workbench', 'the RCA [Ww]orkbench': 'RCA workbench' } },
  word.new('RDS', '', 'noun') { abbreviation: true, elaboration: 'Relational Database Service', product: true },
  word.new('Relational Database Service', '', 'noun') { Amazon: true, description: 'Amazon Relational Database Service', product: true },
  word.new('reachability', '', 'noun') { description: 'A product metric in Synthetic Monitoring' },
  word.new('React', '', 'noun') { product: true },
  word.new('redirection', 'S', 'noun'),
  word.new('Redis', '', 'noun') { product: true, swaps: { redis: 'Redis' } },
  word.new('register', 'dG', 'verb'),
  word.new('register', 'uD', 'adjective'),
  word.new('regular expression', 's', 'noun') { swaps: { 'regexp?': 'regular expression', 'regex[ep]?s': 'regular expression' } },
  word.new('repository', 'S', 'noun') { swaps: { repo: 'repository', repos: 'repositories' } },
  word.new('retry', 'DGS', 'verb'),
  word.new('retryable', '', 'adjective'),
  word.new('REPL', '', 'noun') { abbreviation: true, elaboration: 'read-eval-print loop', established_abbreviation: true },
  word.new('RHEL', '', 'noun') { abbreviation: true, description: 'RedHat Enterprise Linux', established_abbreviation: true, product: true },
  word.new('rollout', 'S', 'noun'),
  word.new('Rollup', '', 'noun') { description: 'The JavaScript module bundler (https://rollupjs.org/)', product: true },
  word.new('RPM', '', 'noun') { abbreviation: true, description: 'https://en.wikipedia.org/wiki/RPM_Package_Manager', elaboration: 'RPM Package Manager', established_abbreviation: true },
  word.new('RSA', '', 'noun') { abbreviation: true, elaboration: 'Rivest–Shamir–Adleman', established_abbreviation: true },
  word.new('runbook', 'S', 'noun'),
]
