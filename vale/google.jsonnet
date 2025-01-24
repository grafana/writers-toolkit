std.prune({
  'Grafana/styles/Grafana/GoogleAMPM.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/AMPM.yml')),
  // Replaced by Grafana/Acronyms.yml.
  'Grafana/styles/Grafana/GoogleAcronyms.yml': null,
  // Not sure that this is that useful.
  'Grafana/styles/Grafana/GoogleColons.yml': null,
  'Grafana/styles/Grafana/GoogleContractions.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Contractions.yml')),
  'Grafana/styles/Grafana/GoogleDateFormat.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/DateFormat.yml')),
  'Grafana/styles/Grafana/GoogleEllipses.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Ellipses.yml')),
  'Grafana/styles/Grafana/GoogleEmDash.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/EmDash.yml')),
  // Replaced by Grafana/Exclamation.yml.
  'Grafana/styles/Grafana/GoogleExclamation.yml': null,
  'Grafana/styles/Grafana/GoogleFirstPerson.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/FirstPerson.yml')),
  'Grafana/styles/Grafana/GoogleGender.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Gender.yml')),
  'Grafana/styles/Grafana/GoogleGenderBias.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/GenderBias.yml')),
  'Grafana/styles/Grafana/GoogleHeadingPunctuation.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/HeadingPunctuation.yml')),
  // Replaced by Grafana/Headings.yml.
  'Grafana/styles/Grafana/GoogleHeadings.yml': null,
  // Replaced by Grafana/Latin.yml.
  'Grafana/styles/Grafana/GoogleLatin.yml': null,
  'Grafana/styles/Grafana/GoogleLyHyphens.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/LyHyphens.yml')),
  'Grafana/styles/Grafana/GoogleOptionalPlurals.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/OptionalPlurals.yml')),
  // Replaced by Grafana/Ordinal.yml.
  'Grafana/styles/Grafana/GoogleOrdinal.yml': null,
  'Grafana/styles/Grafana/GoogleOxfordComma.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/OxfordComma.yml') { level: 'suggestion' }),
  // Replaced by Grafana/Parentheses.yml.
  'Grafana/styles/Grafana/GoogleParens.yml': null,
  'Grafana/styles/Grafana/GooglePassive.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Passive.yml')),
  'Grafana/styles/Grafana/GooglePeriods.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Periods.yml')),
  // Google.Quotes more often a false positive for usage.
  'Grafana/styles/Grafana/GoogleQuotes.yml': null,
  'Grafana/styles/Grafana/GoogleRanges.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Ranges.yml')),
  'Grafana/styles/Grafana/GoogleSemicolons.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Semicolons.yml')),
  'Grafana/styles/Grafana/GoogleSlang.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Slang.yml')),
  'Grafana/styles/Grafana/GoogleSpacing.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Spacing.yml')),
  'Grafana/styles/Grafana/GoogleSpelling.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Spelling.yml')),
  // More often than not, we need to be consistent with Prometheus units or units used Grafana UI which are not represented in the form encouraged by SI.
  'Grafana/styles/Grafana/GoogleUnits.yml': null,
  // Replaced by Grafana/We.yml.
  'Grafana/styles/Grafana/GoogleWe.yml': null,
  'Grafana/styles/Grafana/GoogleWill.yml': std.manifestYamlDoc(std.parseYaml(importstr 'Google/Will.yml') {
    message: super.message + '\n\n' + |||
      Use present tense for statements that describe general behavior that's not associated with a particular time.
    |||,
  }),
  // Replaced by Grafana/WordList.yml.
  'Grafana/styles/Grafana/GoogleWordList.yml': null,
})
