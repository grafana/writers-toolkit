local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'conditional',
  message: 'Use the full Palantir product name in the first instance.',
  link: 'https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#palantir-products',
  level: 'warning',

  local regexp = std.join('|', [
    '%s' % def.word
    for def in defs
    if 'product' in def && def.product && 'Palantir' in def && def.Palantir
  ]),
  first: '\\b(%s)\\b' % regexp,
  second: 'Palantir (%s)' % regexp,
  scope: 'text',
})
