# Grafana/Headings

## Thresholds

The following headings test the behavior of the `thresholds` parameter.
Each heading is sentence cased and each word represents the word's index in the sentence.

### One Two Three Four

Given one correctly sentence cased words, you need three additional incorrectly cased words to trigger the rule.

The measured ratio is 0.25.

### One two Three Four Five Six Seven Eight Nine Ten

Given two correctly sentence cased words, you need eight additional incorrectly cased words to trigger the rule.

The measured ratio is 0.2.

### One two three Four Five Six Seven Eight Nine Ten Eleven

Given three correctly sentence cased words, you need eight additional incorrectly cased words to trigger the rule.

The measured ratio is approximately 0.273.

### One two three four Five Six Seven Eight Nine Ten Eleven Twelve Thirteen Fourteen

Given four correctly sentence cased words, you need ten additional incorrectly cased words to trigger the rule.

The measured ratio is approximately 0.286.

### One two three four five Six Seven Eight Nine Ten Eleven Twelve Thirteen Fourteen Fifteen Sixteen Seventeen

Given four correctly sentence cased words, you need twelve additional incorrectly cased words to trigger the rule.

The measured ratio is approximately 0.294.

### One Two Three Beyla

Since this heading doesn't trigger the rule, Vale must consider exceptions to be correctly cased.
Without an exception for Beyla, the ratio is 0.25.
With the exception, the ratio is 0.5.

### Grafana Enterprise Metrics Four Five Six

Vale must consider multi-word exceptions as a single word in the ratio.
If Vale considered the exception _Grafana Enterprise Metrics_ as three separate words, the ratio would be 0.5.
Since Vale considers the exception as one word, the ratio is 0.25.
