std.prune({
  'Grafana/ReadabilityAutomatedReadability.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/AutomatedReadability.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'Automated Readability Index: %s (aim for below 8).',
  }),
  'Grafana/ReadabilityColemanLiau.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/ColemanLiau.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'Coleman–Liau Index grade: %s (aim for below 9).',
  }),
  'Grafana/ReadabilityFleschKincaid.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/FleschKincaid.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'Flesch–Kincaid grade level: %s (aim for below 8).',
  }),
  'Grafana/ReadabilityFleschReadingEase.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/FleschReadingEase.yml') {
    condition: '< 100',
    level: 'suggestion',
    message: 'Flesch reading ease: %s (aim for above 70).',
  }),
  'Grafana/ReadabilityGunningFog.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/GunningFog.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'Gunning-Fog index: %s (aim for below 10).',
  }),
  'Grafana/ReadabilityLIX.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/LIX.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'LIX score: %s (aim for below 35).',
  }),
  'Grafana/ReadabilitySMOG.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Readability/SMOG.yml') {
    condition: '> 0',
    level: 'suggestion',
    message: 'SMOG score: %s (aim for below 10).',
  }),
})
