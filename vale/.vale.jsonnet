{
  configuration: {
    main: {
      StylesPath: '/etc/vale/styles',
      MinAlertLevel: 'suggestion',

      Packages: 'https://github.com/jdbaldry/Hugo/releases/download/v0.2.0-rc.1/Hugo.zip',
    },
    sections: {
      '*.md': {
        BasedOnStyles: 'Google, Grafana',

        'Google.Headings': 'NO',
        'Google.Quotes': 'NO',
        // More often than not, we need to be consistent with Prometheus units or units used Grafana UI which are not represented in the form encouraged by SI.
        'Google.Units': 'NO',
        'Google.WordList': 'NO',

        // https://github.com/errata-ai/vale/issues/288
        TokenIgnores: @'(<http[^\n]+>+?), \*\*[^\n]+\*\*',
      },
    },
  },

  container: std.manifestIni(self.configuration),

  repository: std.manifestIni(self.configuration {
    main+: {
      StylesPath: 'vale',
      Packages: 'Google, ' + super.Packages,
    },
  }),
}
