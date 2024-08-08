local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'conditional',
  message: 'Use the full Amazon product name in the first instance.',
  link: 'https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#amazon-products',
  level: 'warning',

  local regexp = std.join('|', [
    '%s' % def.word
    for def in defs
    if 'product' in def && def.product && 'Amazon' in def && def.Amazon
  ]),
  first: '\\b(%s)\\b' % regexp,
  second: 'Amazon (%s)' % regexp,
  scope: 'text',
})
