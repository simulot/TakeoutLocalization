package localization

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Products map[string]*Node

func (prod Products) GetKeyAndNode(localizedPath string) (string, *Node, error) {
	var (
		dir string
		p   int
	)
	p = strings.Index(localizedPath, "/")
	if p > 0 {
		dir = localizedPath[:p]
		localizedPath = localizedPath[p+1:]
	} else {
		dir = localizedPath
		localizedPath = ""
	}

	for k, n := range prod {
		if _, ok := n.HasLocalization(dir); ok {
			if localizedPath == "" {
				return k, n, nil
			}
			if childK, childN, err := n.walkNodes(localizedPath); err == nil {
				return childK, childN, nil
			}
		}
	}
	return "", nil, errors.New("not found")
}

// Node represents a structure used for organizing and managing localiszations
// and their associated metadata. It can represent directories, files, or
// other hierarchical elements in a localization system.
//
// Fields:
//   - Kind: Specifies the type of the node (e.g., directory, file, etc.).
//   - localization: Contains the localizations data associated with this node.
//   - Directories: A map of subdirectory names to their corresponding Node
//     structures, representing the hierarchical structure of directories.
//   - Files: A map of file names to their corresponding Node structures,
//     representing the files contained within this node.
//   - Columns: A map of column names to their corresponding Column structures,
//     used for managing additional metadata or attributes related to the file node.
type Node struct {
	Kind          NodeKind `json:"kind"`
	Localizations `json:"localizations,omitempty"`

	Directories map[string]*Node   `json:"directories,omitempty"`
	Files       map[string]*Node   `json:"files,omitempty"`
	Columns     map[string]*Column `json:"columns,omitempty"`
}

type NodeKind string

const (
	NodeDirectory NodeKind = "directory"
	NodeFile      NodeKind = "file"
)

// Localizations represents a base structure for managing translation keys and their corresponding localization in multiple languages.
type Localizations map[string]string

// validate checks the validity of a Node based on its kind and structure.
// It ensures that:
// - The Node's kind is either "directory" or "file".
// - A Node of kind "directory" does not have any columns.
// - A Node of kind "file" does not have subdirectories or files.
// The function recursively validates all subdirectories and files within the Node.
// If any validation rule is violated, an error is returned with a descriptive message.
//
// Parameters:
//   - p: A string representing the path of the current Node.
//
// Returns:
//   - An error if the Node or any of its children are invalid, otherwise nil.
func (n *Node) validate(p string) error {
	if n.Kind != NodeDirectory && n.Kind != NodeFile {
		return fmt.Errorf("invalid node kind at %s: %s", p, n.Kind)
	}
	if n.Kind == NodeDirectory && len(n.Columns) > 0 {
		return fmt.Errorf("directory node cannot have columns at %s", p)
	}
	if n.Kind == NodeFile && len(n.Directories) > 0 {
		return fmt.Errorf("file node cannot have subdirectories at %s", p)
	}
	if n.Kind == NodeFile && len(n.Files) > 0 {
		return fmt.Errorf("file node cannot have files at %s", p)
	}
	for name, subDir := range n.Directories {
		if err := subDir.validate(p + "/" + name); err != nil {
			return err
		}
	}
	for name, file := range n.Files {
		if err := file.validate(p + "/" + name); err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) HasLocalization(local string) (string, bool) {
	for k, v := range n.Localizations {
		if v == local {
			return k, true
		}
	}
	return "", false
}

func (n *Node) GetColumnKey(local string) (string, bool) {
	if n.Kind == NodeFile {
		for k, c := range n.Columns {
			for _, v := range c.Localizations {
				if v == local {
					return k, true
				}
			}
		}
	}
	return "", false
}

func (n *Node) walkNodes(localizedPath string) (string, *Node, error) {
	var (
		dir string
		p   int
	)
	p = strings.Index(localizedPath, "/")
	if p >= 0 {
		dir = localizedPath[:p]
		localizedPath = localizedPath[p+1:]
	} else {
		dir = localizedPath
		localizedPath = ""
	}

	for k, n := range n.Directories {
		if _, ok := n.HasLocalization(dir); ok {
			if localizedPath == "" {
				return k, n, nil
			}
			return n.walkNodes(localizedPath)
		}
	}
	for k, n := range n.Files {
		if _, ok := n.HasLocalization(dir); ok {
			return k, n, nil
		}
	}
	return "", nil, errors.New("not found")
}

// Column represents a column in a file node.
// It embeds the localization struct.
type Column struct {
	Localizations
}

// MarshalJSON customizes the JSON output for the Column struct.
// It directly attaches the Localizations map to the JSON output.
func (c Column) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Localizations)
}

// UnmarshalJSON customizes the JSON input for the Column struct.
// It directly parses the JSON into the Localizations map.
func (c *Column) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.Localizations)
}
