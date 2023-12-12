# Why Typescript for Part 2?

[Go's Regex does not support positive lookahead](https://pkg.go.dev/regexp/syntax), which would be necessary to detect numbers inside others, such as "twone". Since Regex consumes the letters it matches, if it matches the first "two", what remains is "ne" instead of "one". This way, we need positive lookahead, which does not consume characters, to detect those.

This is why I'm using Typescrit for day 1 part 2.