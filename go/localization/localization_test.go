package localization

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
	"testing/fstest"
)

func TestNodesStruct(t *testing.T) {
	products := map[string]*Node{
		"Product1": {
			Kind: "directory",
			Localizations: Localizations{
				"fr": "Produit1",
				"de": "Produkt1",
			},
			Directories: map[string]*Node{
				"directory1": {
					Kind: "directory",
					Localizations: Localizations{
						"fr": "Répertoire1",
						"de": "Verzeichnis1",
					},
					Files: map[string]*Node{
						"file1.csv": {
							Kind: "file",
							Localizations: Localizations{
								"fr": "Fichier1",
								"de": "Datei1",
							},
							Columns: map[string]*Column{
								"column1": {
									Localizations: Localizations{
										"fr": "Colonne1",
										"de": "Spalte1",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	err := enc.Encode(products)
	if err != nil {
		t.Fatalf("Failed to encode JSON: %v", err)
	}
	fmt.Println("Encode:")
	fmt.Println(buf.String())
	dec := json.NewDecoder(bytes.NewReader(buf.Bytes()))
	err = dec.Decode(&products)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	buf.Reset()
	enc.Encode(products)
	fmt.Println("Decode:")
	fmt.Println(buf.String())
}

func TestLoadJSON_Success(t *testing.T) {
	jsonContent := `{
	"root": {
		"kind": "directory",
		"localizations": {
			"fr": "Racine",
			"de": "Wurzel"
		},
		"directories": {
			"SubChild": {
				"kind": "directory",
				"localizations": {
					"fr": "Sous-enfant",
					"de": "Unterkind"
				}
			}
		},
		"files": {
			"file1": {
				"kind": "file",
				"localizations": {
					"fr": "Fichier1",
					"de": "Datei1"
				},
				"columns": {
					"col1": {
						"localizations": {
							"fr": "Colonne1",
							"de": "Spalte1"
						}
					},
					"col2": {
						"localizations": {
							"fr": "Colonne2",
							"de": "Spalte2"
						}
					}
				}
			}
		}
	}
}`

	vfs := fstest.MapFS{
		"test.json": &fstest.MapFile{Data: []byte(jsonContent)},
	}

	n, err := LoadJSON(vfs, "test.json")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(n)
}

func TestLoadJSON_FileNotFound(t *testing.T) {
	vfs := fstest.MapFS{}
	_, err := LoadJSON(vfs, "missing.json")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("expected fs.ErrNotExist, got %v", err)
	}
}

func TestLoadJSON_InvalidJSON(t *testing.T) {
	vfs := fstest.MapFS{
		"bad.json": &fstest.MapFile{Data: []byte("{invalid json")},
	}
	_, err := LoadJSON(vfs, "bad.json")
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
	if !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected JSON error, got %v", err)
	}
}

func TestLocalJSONFile(t *testing.T) {
	vfs := os.DirFS("../../localizations")
	n, err := LoadJSON(vfs, "YouTube.json")
	if err != nil {
		t.Fatalf("failed to load JSON: %s", err.Error())
		return
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(n)
}
