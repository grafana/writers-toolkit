local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'existence',
  level: 'warning',
  link: 'https://developers.google.com/style/possessives#product,-feature,-and-company-names',
  message: |||
    Don't form a possessive from a feature name, product name, or trademark, regardless of who owns it.
    Instead, use the name as a modifier or rewrite to use a word like of to indicate the relationship.
  |||,
  tokens: [
    (if def.word[std.length(def.word) - 1] == 's'
     then def.word + "'"
     else def.word + "'s")
    for def in defs
    if 'product' in def && def.product
  ],
})
