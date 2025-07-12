package imageformatter

import (
	"os"
	"path/filepath"
	"strings"
)

// GetImageFiles scans a directory and returns supported image files
func GetImageFiles(dir string) ([]string, error) {
	var files []string
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if extensions[ext] {
				files = append(files, entry.Name())
			}
		}
	}

	return files, nil
}
