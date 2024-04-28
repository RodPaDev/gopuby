package epubManager

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unzipEpub(path string, extractPath string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %s", err)
	}
	reader, err := zip.OpenReader(absPath)

	if err != nil {
		return fmt.Errorf("impossible to open zip reader: %s", err)

	}
	defer reader.Close()

	for k, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("impossible to open file n°%d in archive: %s", k, err)
		}
		defer rc.Close()

		newFilePath := filepath.Join(extractPath, f.Name)
		if f.FileInfo().IsDir() {
			// Create directory if it is a directory
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				return fmt.Errorf("impossible to MkdirAll: %s", err)
			}
			continue
		}

		// Ensure the directory for the file exists
		dirPath := filepath.Dir(newFilePath)
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return fmt.Errorf("failed to create directory %s: %s", dirPath, err)
		}

		uncompressedFile, err := os.Create(newFilePath)
		if err != nil {
			return fmt.Errorf("impossible to create file %s: %s", newFilePath, err)
		}
		_, err = io.Copy(uncompressedFile, rc)
		if err != nil {
			return fmt.Errorf("impossible to copy file n°%d: %s", k, err)
		}
	}

	return nil
}
