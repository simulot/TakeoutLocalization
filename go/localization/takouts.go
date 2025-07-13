package localization

import (
	"encoding/json"
	"errors"
	"io/fs"
	"maps"
)

func LoadLocalizations(vfs fs.FS, jsonNames ...string) (Products, error) {
	var errs error
	products := map[string]*Node{}
	for _, jsonName := range jsonNames {
		d, err := LoadJSON(vfs, jsonName)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		maps.Copy(products, d)
	}
	return products, nil
}

func LoadJSON(vfs fs.FS, path string) (Products, error) {
	f, err := vfs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	root := make(Products)
	if err := json.NewDecoder(f).Decode(&root); err != nil {
		return nil, err
	}

	for key, node := range root {
		if err := node.validate(key); err != nil {
			return nil, err
		}
	}
	return root, nil
}
