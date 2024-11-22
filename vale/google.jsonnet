std.prune({
  'Grafana/styles/Grafana/GrafanaGoogleAMPM.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/AMPM.yml')),
  // Replaced by Grafana/styles/GrafanaAcronyms.yml.
  'Grafana/styles/Grafana/GrafanaGoogleAcronyms.yml': null,
  // Not sure that this is that useful.
  'Grafana/styles/Grafana/GrafanaGoogleColons.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleContractions.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Contractions.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleDateFormat.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/DateFormat.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleEllipses.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Ellipses.yml')),
  // Replaced by Grafana/styles/GrafanaExclamation.yml.
  'Grafana/styles/Grafana/GrafanaGoogleExclamation.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleFirstPerson.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/FirstPerson.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleGender.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Gender.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleGenderBias.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/GenderBias.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleHeadingPunctuation.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/HeadingPunctuation.yml')),
  // Replaced by Grafana/styles/GrafanaHeadings.yml.
  'Grafana/styles/Grafana/GrafanaGoogleHeadings.yml': null,
  // Replaced by Grafana/styles/GrafanaLatin.yml.
  'Grafana/styles/Grafana/GrafanaGoogleLatin.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleLyHyphens.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/LyHyphens.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleOptionalPlurals.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/OptionalPlurals.yml')),
  // Replaced by Grafana/styles/GrafanaOrdinal.yml.
  'Grafana/styles/Grafana/GrafanaGoogleOrdinal.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleOxfordComma.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/OxfordComma.yml') { level: 'suggestion' }),
  // Replaced by Grafana/styles/GrafanaParentheses.yml.
  'Grafana/styles/Grafana/GrafanaGoogleParens.yml': null,
  'Grafana/styles/Grafana/GrafanaGooglePassive.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Passive.yml')),
  'Grafana/styles/Grafana/GrafanaGooglePeriods.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Periods.yml')),
  // Google.Quotes more often a false positive for usage.
  'Grafana/styles/Grafana/GrafanaGoogleQuotes.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleRanges.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Ranges.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleSemicolons.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Semicolons.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleSlang.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Slang.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleSpacing.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Spacing.yml')),
  'Grafana/styles/Grafana/GrafanaGoogleSpelling.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Spelling.yml')),
  // More often than not, we need to be consistent with Prometheus units or units used Grafana UI which are not represented in the form encouraged by SI.
  'Grafana/styles/Grafana/GrafanaGoogleUnits.yml': null,
  // Replaced by Grafana/styles/GrafanaWe.yml.
  'Grafana/styles/Grafana/GrafanaGoogleWe.yml': null,
  'Grafana/styles/Grafana/GrafanaGoogleWill.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Will.yml') {
    message: super.message + '\n\n' + |||
      Use present tense for statements that describe general behavior that's not associated with a particular time.
    |||,
  }),
  // Replaced by Grafana/styles/GrafanaWordList.yml.
  'Grafana/styles/Grafana/GrafanaGoogleWordList.yml': null,
})
