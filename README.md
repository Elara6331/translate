# Translate

This library is an incredibly simple translation system for Go. It is designed to work well with [go.arsenm.dev/logger](https://pkg.go.dev/go.arsenm.dev/logger), and doesn't have nearly as many features as other solutions, because those features are not needed for all applications, and they increase complexity unnecessarily in use cases that don't need it. All this package does is parses TOML files for translations, and then gets the raw text from the files for whatever language you want to translate to.

If you need advanced functionality, such as handling singular/plural words, gender, etc., use other packages such as [golang.org/x/text/message](https://pkg.go.dev/golang.org/x/text/message) or [github.com/nicksnyder/go-i18n](https://pkg.go.dev/github.com/nicksnyder/go-i18n).