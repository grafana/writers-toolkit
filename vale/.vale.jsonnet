{
  configuration: {
    main: {
      StylesPath: '/etc/vale/styles',
      MinAlertLevel: 'suggestion',

      Packages: 'Google, https://github.com/errata-ai/Hugo/releases/download/v0.2.0/Hugo.zip, Readability',
    },
    sections: {
      '*': {
        BasedOnStyles: 'Grafana',

        // https://github.com/errata-ai/vale/issues/288
        TokenIgnores: @'(<http[^\n]+>+?), \*\*[^\n]+\*\*',
      },
    },
  },

  container: std.manifestIni(self.configuration),

  repository: std.manifestIni(self.configuration {
    main+: {
      StylesPath: 'vale',
    },
  }),
}
