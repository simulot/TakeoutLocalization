package localization

import (
	"io/fs"

	"github.com/simulot/TakeoutLocalization/localizations"
)

// GetDefaultLocalizations returns the default localizations embedded into this package.
func GetDefaultLocalizations() Products {
	vfs := localizations.Default

	entries, err := fs.ReadDir(vfs, ".")
	if err != nil {
		panic(err)
	}
	files := make([]string, len(entries))
	for i, entry := range entries {
		files[i] = entry.Name()
	}
	p, err := LoadLocalizations(vfs, files...)
	if err != nil {
		panic(err)
	}
	return p
}
