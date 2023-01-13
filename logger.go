package translate

import (
	"fmt"

	"go.arsenm.dev/logger"
	"golang.org/x/text/language"
)

var _ logger.Logger = &TranslatedLogger{}

type TranslatedLogger struct {
	Logger     logger.Logger
	Translator Translator
	Language   language.Tag
}

func NewLogger(l logger.Logger, t Translator, lang language.Tag) *TranslatedLogger {
	return &TranslatedLogger{l, t, lang}
}

func (tl *TranslatedLogger) NoPanic() {
	tl.Logger.NoPanic()
}

func (tl *TranslatedLogger) NoExit() {
	tl.Logger.NoExit()
}

func (tl *TranslatedLogger) SetLevel(level logger.LogLevel) {
	tl.Logger.SetLevel(level)
}

func (tl *TranslatedLogger) Debug(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Debug(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Debugf(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Debugf(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Info(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Info(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Infof(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Infof(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Warn(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Warn(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Warnf(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Warnf(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Error(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Error(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Errorf(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Errorf(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Fatal(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Fatal(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Fatalf(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Fatalf(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Panic(msg string) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Panic(tl.Translator.TranslateTo(msg, tl.Language)),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

func (tl *TranslatedLogger) Panicf(format string, args ...any) logger.LogBuilder {
	return &TranslatedLogBuilder{
		LogBuilder: tl.Logger.Panicf(tl.Translator.TranslateTo(format, tl.Language), args...),
		Translator: tl.Translator,
		Language:   tl.Language,
	}
}

type TranslatedLogBuilder struct {
	LogBuilder logger.LogBuilder
	Translator Translator
	Language   language.Tag
}

func (tlb *TranslatedLogBuilder) Int(key string, value int) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Int(key, value)
}

func (tlb *TranslatedLogBuilder) Int8(key string, value int8) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Int8(key, value)
}

func (tlb *TranslatedLogBuilder) Int16(key string, value int16) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Int16(key, value)
}

func (tlb *TranslatedLogBuilder) Int32(key string, value int32) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Int32(key, value)
}

func (tlb *TranslatedLogBuilder) Int64(key string, value int64) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Int64(key, value)
}

func (tlb *TranslatedLogBuilder) Uint(key string, value uint) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Uint(key, value)
}

func (tlb *TranslatedLogBuilder) Uint8(key string, value uint8) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Uint8(key, value)
}

func (tlb *TranslatedLogBuilder) Uint16(key string, value uint16) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Uint16(key, value)
}

func (tlb *TranslatedLogBuilder) Uint32(key string, value uint32) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Uint32(key, value)
}

func (tlb *TranslatedLogBuilder) Uint64(key string, value uint64) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Uint64(key, value)
}

func (tlb *TranslatedLogBuilder) Float32(key string, value float32) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Float32(key, value)
}

func (tlb *TranslatedLogBuilder) Float64(key string, value float64) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Float64(key, value)
}

func (tlb *TranslatedLogBuilder) Stringer(key string, value fmt.Stringer) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Stringer(key, value)
}

func (tlb *TranslatedLogBuilder) Bytes(key string, value []byte) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Bytes(key, value)
}

func (tlb *TranslatedLogBuilder) Timestamp() logger.LogBuilder {
	return tlb.LogBuilder.Timestamp()
}

func (tlb *TranslatedLogBuilder) Bool(key string, value bool) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Bool(key, value)
}

func (tlb *TranslatedLogBuilder) Str(key string, value string) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Str(key, value)
}

func (tlb *TranslatedLogBuilder) Any(key string, value any) logger.LogBuilder {
	key = tlb.Translator.TranslateTo(key, tlb.Language)
	return tlb.LogBuilder.Any(key, value)
}

func (tlb *TranslatedLogBuilder) Err(err error) logger.LogBuilder {
	return tlb.LogBuilder.Err(err)
}

func (tlb *TranslatedLogBuilder) Send() {
	tlb.LogBuilder.Send()
}
