local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'conditional',
  message: "Spell out '%s', if it's unfamiliar to the audience.",
  link: 'https://developers.google.com/style/abbreviations',
  level: 'suggestion',
  ignorecase: false,
  // Ensures that the existence of 'first' implies the existence of 'second'.
  first: '\\b([A-Z]{3,5})\\b',
  second: '(?:\\b[A-Z][a-z]+ )+\\(([A-Z]{3,5})\\)',
  // ... with the exception of these:
  exceptions: [
    '%s' % def.word
    for def in defs
    if 'abbreviation' in def && def.abbreviation && 'established_abbreviation' in def && def.established_abbreviation
  ],
})
