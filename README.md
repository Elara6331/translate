# Translate

This library is an incredibly simple translation system for Go. It is designed to work well with [go.arsenm.dev/logger](https://pkg.go.dev/go.arsenm.dev/logger), and doesn't have nearly as many features as other solutions, because those features are not needed for all applications, and they increase complexity unnecessarily in use cases that don't need it. All this package does is parses TOML files for translations, and then gets the raw text from the files for whatever language you want to translate to.

If you need advanced functionality, such as handling singular/plural words, gender, etc., use other packages such as [golang.org/x/text/message](https://pkg.go.dev/golang.org/x/text/message) or [github.com/nicksnyder/go-i18n](https://pkg.go.dev/github.com/nicksnyder/go-i18n).

### Translation files

Translation files for this package are incredibly simple. They are just TOML files with sections that look like the following:

```toml
[[translation]]
id = 458405366
value = 'First test'

[[translation]]
id = 2703638123
value = 'Second test'
```

The `value` field contains the string you want to use for the given language, and the `id` field is a crc32 checksum of the string that you will input, most commonly the same as the `value` field of the fallback language. For example, if I wanted to add a Russian translation, I'd do something like this:

```toml
[[translation]]
id = 458405366
value = 'Первая проверка'
```

Notice that the ID does not correspond to the checksum of the actual string. Instead, it corresponds to the English string above, because that's the fallback and that's what I'll be using when calling `Translator.TranslateTo()`. To use this translation, you'd do the following:

```go
tr := translate.New(cat)
out := tr.TranslateTo("First test", language.Russian) // out == "Первая проверка"
```

The checksum is calculated from `"First test"`.