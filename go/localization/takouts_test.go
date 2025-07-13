package localization

import (
	"os"
	"testing"
)

func TestGetDirLocalization(t *testing.T) {
	vfs := os.DirFS("../../localizations")
	products, err := LoadLocalizations(vfs, "GooglePhotos.json", "YouTube.json")
	if err != nil {
		t.Fatalf("failed to load localizations: %s", err.Error())
		return
	}

	testCases := []struct {
		dirPath  string
		lang     string
		wantKey  string
		wantKind NodeKind
	}{
		{"YouTube and YouTube Music/channels", "en", "channels", NodeDirectory},
		{"YouTube and YouTube Music/channels/channel.csv", "en", "channel.csv", NodeFile},
		{"YouTube et YouTube Music/chaînes", "fr", "channels", NodeDirectory},
		{"YouTube et YouTube Music/chaînes/chaîne.csv", "fr", "channel.csv", NodeFile},
	}

	for _, tc := range testCases {
		t.Run(tc.dirPath, func(t *testing.T) {
			key, node, err := products.GetKeyAndNode(tc.dirPath)
			if err != nil {
				t.Errorf("failed to get key and node: %s", err.Error())
				return
			}
			if key != tc.wantKey {
				t.Errorf("got key %s, want %s", key, tc.wantKey)
				return
			}
			if node.Kind != tc.wantKind {
				t.Errorf("got kind %s, want %s", node.Kind, tc.wantKind)
				return
			}
		})
	}
}

func TestColumnLocalization(t *testing.T) {
	vfs := os.DirFS("../../localizations")
	products, err := LoadLocalizations(vfs, "GooglePhotos.json", "YouTube.json")
	if err != nil {
		t.Fatalf("failed to load localizations: %s", err.Error())
		return
	}

	testCases := []struct {
		dirPath     string
		localColumn string
		wantKey     string
	}{
		{"YouTube and YouTube Music/channels/channel.csv", "Channel ID", "Channel ID"},
		{"YouTube et YouTube Music/chaînes/chaîne.csv", "ID de la chaîne", "Channel ID"},
	}

	for _, tc := range testCases {
		t.Run(tc.dirPath, func(t *testing.T) {
			_, node, err := products.GetKeyAndNode(tc.dirPath)
			if err != nil {
				t.Errorf("failed to get key and node: %s", err.Error())
				return
			}
			if node.Kind != NodeFile {
				t.Errorf("got kind %s, want %s", node.Kind, NodeFile)
				return
			}

			key, ok := node.GetColumnKey(tc.localColumn)
			if !ok {
				t.Errorf("failed to get column key for %s", tc.localColumn)
				return
			}

			if key != tc.wantKey {
				t.Errorf("got key %s, want %s", key, tc.wantKey)
				return
			}
		})
	}
}
