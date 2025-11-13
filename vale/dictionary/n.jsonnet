local word = import './word.jsonnet';
[
  word.new('namespace', 'S', 'noun'),
  word.new('navigation', 'S', 'noun'),
  word.new('NAT', '', 'noun') { abbreviation: true, description: 'Network Address Translation', established_abbreviation: true },
  word.new('nginx', '', 'noun'),
  word.new('Netlink', '', 'noun') { product: true, description: 'Netlink is a socket family used for inter-process communication (IPC) between both the kernel and userspace processes: https://en.wikipedia.org/wiki/Netlink' },
]
