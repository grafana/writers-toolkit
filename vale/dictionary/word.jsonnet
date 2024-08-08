// The structure of the word object is documented in https://grafana.com/docs/writers-toolkit/review/lint-prose/dictionary/#word-metadata.
{
  _prototype: {
    abbreviation: false,
    affixes: [],
    Amazon: false,
    Apache: false,
    Google: false,
    description: null,
    established_abbreviation: false,
    po: '',
    product: false,
    word: '',
  },

  new(word, affixes, po):: self._prototype {
    word: word,
    affixes: affixes,
    po: po,
  },
}
