local word = import './word.jsonnet';
[
  word.new('cAdvisor', 'M', 'noun') { description: 'cAdvisor (Container Advisor) provides container users an understanding of the resource usage and performance characteristics of their running containers. (https://github.com/google/cadvisor)', swaps: { cadvisor: 'cAdvisor' } },
  word.new('callee', 'S', 'noun') { description: 'A function that is called by another function' },
  word.new('CDN', '', 'noun') { abbreviation: true, elaboration: 'Content Delivery Network' },
  word.new('CentOS', '', 'noun') { description: 'A Linux distribution that provides a free, community-supported computing platform functionally compatible with its upstream source, Red Hat Enterprise Linux (RHEL).', product: true, swaps: { '[cC]entos': 'CentOS' } },
  word.new('checkbox', 'S', 'noun') { description: 'A small box that can be checked or unchecked', swaps: { 'check[- ]box': 'checkbox' } },
  word.new('CLA', '', 'noun') { abbreviation: true, elaboration: 'Contributor License Agreement' },
  word.new('clear', 'S', 'verb') { swaps: { 'un(?:check|select)': 'clear' } },
  word.new('CLI', '', 'noun') { abbreviation: true, established_abbreviation: true },
  word.new('clickjack', 'DSG', 'noun') { description: 'A type of attack where a malicious website tricks a user into clicking on a hidden element on another website' },
  word.new('CloudWatch', '', 'noun') { Amazon: true, product: true },
  word.new('CLS', '', 'noun') { abbreviation: true, elaboration: 'Cumulative Layout Shift' },
  word.new('CMS', '', 'noun') { abbreviation: true, elaboration: 'content management system' },
  word.new('codespace', 'S', 'noun') { description: "A codespace is a development environment that's hosted in the cloud. https://docs.github.com/en/codespaces/overview" },
  word.new('Codespaces', '', 'noun') { description: 'GitHub Codespaces https://docs.github.com/en/codespaces/overview', product: true },
  word.new('comment', 'S', 'noun'),
  word.new('comment', 'uDG', 'verb'),
  word.new('composable', '', 'adjective'),
  word.new('configure', 'mpDS', 'verb'),
  word.new('contentful', '', 'adjective') { description: 'Having content. Used in Web Vital metrics, such as Largest Contentful Paint: https://web.dev/articles/lcp' },
  word.new('CORS', '', 'noun') { abbreviation: true, established_abbreviation: true, description: 'Cross-origin resource sharing. Allows a web page to access restricted resources from a server on a domain different than the domain that served the web page.' },
  word.new('Couchbase', 'M', 'noun'),
  word.new('CPU', 'S', 'noun') { abbreviation: true, elaboration: 'central processing unit', established_abbreviation: true },
  word.new('CRD', 'S', 'noun') { abbreviation: true, elaboration: 'Custom Resource Definition', established_abbreviation: true },
  word.new('CSS', '', 'noun') { abbreviation: true, elaboration: 'Cascading Style Sheets', established_abbreviation: true },
  word.new('CSV', 'S', 'noun') { abbreviation: true, elaboration: 'Comma-separated values', established_abbreviation: true },
  word.new('CVE', 'S', 'noun'),
]
