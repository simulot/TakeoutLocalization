package localization

import (
	"io/fs"
	"sync"

	"github.com/simulot/TakeoutLocalization/localizations"
)

var (
	_products     Products
	_onceProducts sync.Once
)

// GetDefaultLocalizations returns the default localizations embedded into this package.
func GetDefaultLocalizations() Products {
	_onceProducts.Do(func() {
		vfs := localizations.Default

		entries, err := fs.ReadDir(vfs, ".")
		if err != nil {
			panic(err)
		}
		files := make([]string, len(entries))
		for i, entry := range entries {
			files[i] = entry.Name()
		}
		_products, err = LoadLocalizations(vfs, files...)
		if err != nil {
			panic(err)
		}
	})
	return _products
}
