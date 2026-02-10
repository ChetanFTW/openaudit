package scanner

import (
	"os"
	"path/filepath"
)

type Result struct {
	HasGoMod  bool
	HasCI     bool
	HasReadme bool
}

func Scan(root string) (*Result, error) {
	result := &Result{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		switch filepath.Base(path) {
		case "go.mod":
			result.HasGoMod = true
		case "README.md":
			result.HasReadme = true
		case ".github":
			result.HasCI = true
		}

		return nil
	})

	return result, err
}
