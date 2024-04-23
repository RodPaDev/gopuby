package epubManager

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unzipEpub(path string, extractPath string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("error getting absolute path: %s", err)
	}
	log.Printf("Unzipping epub file: %s", absPath)
	reader, err := zip.OpenReader(absPath)

	if err != nil {
		log.Fatalf("impossible to open zip reader: %s", err)
	}
	defer reader.Close()

	for k, f := range reader.File {
		log.Printf("Unzipping %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatalf("impossible to open file n°%d in archive: %s", k, err)
		}
		defer rc.Close()

		newFilePath := filepath.Join(extractPath, f.Name)
		if f.FileInfo().IsDir() {
			// Create directory if it is a directory
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				log.Fatalf("impossible to MkdirAll: %s", err)
			}
			continue
		}

		// Ensure the directory for the file exists
		dirPath := filepath.Dir(newFilePath)
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			log.Fatalf("Failed to create directory %s: %s", dirPath, err)
		}

		uncompressedFile, err := os.Create(newFilePath)
		if err != nil {
			log.Fatalf("impossible to create file %s: %s", newFilePath, err)
		}
		_, err = io.Copy(uncompressedFile, rc)
		if err != nil {
			log.Fatalf("impossible to copy file n°%d: %s", k, err)
		}
	}
}
