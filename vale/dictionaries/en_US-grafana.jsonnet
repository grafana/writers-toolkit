local defs = (import '../dictionary.jsonnet').words;
local entries = std.map(function(word) '%s/%s po:%s' % [word.word, word.affixes, word.po], defs);
'%d\n%s' % [std.length(entries), std.join('\n', entries)]
