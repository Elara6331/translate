package translate

import (
	"hash/crc32"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

type Translations struct {
	Items []Translation `toml:"translation"`
}

type Translation struct {
	ID    uint32 `toml:"id"`
	Value string `toml:"value"`
}

type Dictionary map[uint32]string
type Catalog map[language.Tag]Dictionary

type Translator struct {
	cat      Catalog
	Fallback language.Tag
}

func New(cat Catalog) Translator {
	return Translator{cat, language.English}
}

func NewFromDir(dir string) (Translator, error) {
	return NewFromFS(os.DirFS(dir))
}

func NewFromFS(fsys fs.FS) (Translator, error) {
	cat := Catalog{}
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".toml" {
			return nil
		}

		lang := filepath.Ext(strings.TrimSuffix(path, ext))
		lang = strings.TrimPrefix(lang, ".")

		tag, err := language.Parse(lang)
		if err != nil {
			return err
		}

		fl, err := fsys.Open(path)
		if err != nil {
			return err
		}

		var tr Translations
		err = toml.NewDecoder(fl).Decode(&tr)
		if err != nil {
			return err
		}
		fl.Close()

		cat[tag] = toDict(tr)
		return nil
	})
	if err != nil {
		return Translator{}, err
	}

	return Translator{cat, language.English}, nil
}

func (t *Translator) TranslateTo(s string, tag language.Tag) string {
	dict, ok := t.cat[tag]
	if !ok {
		dict, ok = t.cat[t.Fallback]
		if !ok {
			return s
		}
	}
	id := hashString(s)
	str, ok := dict[id]
	if !ok {
		str = s
	}
	return str
}

func toDict(tr Translations) Dictionary {
	out := Dictionary{}
	for _, item := range tr.Items {
		out[item.ID] = item.Value
	}
	return out
}

func hashString(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}
