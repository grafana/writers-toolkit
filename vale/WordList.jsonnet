local defs = (import './dictionary.libsonnet').words;
std.manifestYamlDoc({
  extends: 'substitution',
  message: "Use '%s' instead of '%s'.",
  link: 'https://grafana.com/docs/writers-toolkit/write/style-guide/word-list/',
  level: 'warning',
  ignorecase: false,
  action: {
    name: 'replace',
  },

  swap: std.foldl(
    function(acc, def) acc + def.swaps,
    defs,
    {
      'ad[- ]?hoc': 'free-form|user-written',
      'the Grafana Agent': 'Grafana Agent',
      'network IP address': 'internal IP address',
      'left[- ]hand[- ]side': 'left-side',
      'fewer data': 'less data',
      '(?:hamburger menu|kebab menu)': 'menu icon',
      '(?:cell ?phone|smart ?phone)': 'phone|mobile phone',
      'right[- ]hand[- ]side': 'right-side',
      'sign into': 'sign in to',
      '(?:kill|terminate|abort)': 'stop|exit|cancel|end',
      'in order to': 'to',
      timeseries: 'time series|time-series',
      'grayed-out': 'unavailable',
    },
  ),
})
