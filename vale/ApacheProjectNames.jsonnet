local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'conditional',
  message: 'Use the full Apache project name in the first instance.',
  link: 'https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#apache-projects',
  level: 'warning',

  local regexp = std.join('|', [
    '%s' % def.word
    for def in defs
    if 'product' in def && def.product && 'Apache' in def && def.Apache
  ]),
  first: '\\b(%s)\\b' % regexp,
  second: 'Apache (%s)' % regexp,
  scope: 'text',
})
