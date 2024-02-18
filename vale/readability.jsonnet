std.prune({
  'Grafana/ReadabilityAutomatedReadability.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/AutomatedReadability.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 8.',
  }),
  'Grafana/ReadabilityColemanLiau.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/ColemanLiau.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 9.',
  }),
  'Grafana/ReadabilityFleschKincaid.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/FleschKincaid.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 8.',
  }),
  'Grafana/ReadabilityFleschReadingEase.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/FleschReadingEase.yml') {
    condition: '< 100',
    level: 'suggestion',
    message: '%s aim for above 70.',
  }),
  'Grafana/ReadabilityGunningFog.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/GunningFog.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 10.',
  }),
  'Grafana/ReadabilityLIX.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/LIX.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 35.',
  }),
  'Grafana/ReadabilitySMOG.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/SMOG.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: '%s aim for below 10.',
  }),
})
