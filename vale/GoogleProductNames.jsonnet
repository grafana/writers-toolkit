local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'conditional',
  message: 'Use the full Google product name in the first instance.',
  link: 'https://grafana.com/docs/writers-toolkit/write/style-guide/capitalization-punctuation/#google-products',
  level: 'warning',

  local regexp = std.join('|', [
    '%s' % def.word
    for def in defs
    if 'product' in def && def.product && 'Google' in def && def.Google
  ]),
  first: '\\b(%s)\\b' % regexp,
  second: 'Google (%s)' % regexp,
  scope: 'text',
})
