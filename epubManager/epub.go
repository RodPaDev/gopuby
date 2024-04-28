package epubManager

import (
	"fmt"
	"gopuby/utils"
	"os"
	"path/filepath"
)

const (
	BOOKS_DIR           = "books"
	TMP_DIR             = BOOKS_DIR + "/tmp"
	TMP_UNPROCESSED_DIR = TMP_DIR + "/unprocessed"
)

func Open(path string) string {
	// Uncompress the epub file
	unzipEpub(path, TMP_UNPROCESSED_DIR)

	// Get epub metadata
	metadata, err := GetMetadata(TMP_UNPROCESSED_DIR)
	if err != nil {
		fmt.Print(err)
	}

	// check if there is already a processed directory for the epub
	processedDir := filepath.Join(TMP_DIR, metadata.ID)
	if _, err := os.Stat(processedDir); err == nil {
		return processedDir
	}

	// Ensure the books directory exists
	if err := utils.EnsureDir(BOOKS_DIR); err != nil {
		fmt.Print(err)
	}

	// Copy the epub file to the books directory
	dstEpubPath := filepath.Join(BOOKS_DIR, fmt.Sprintf("%s.epub", metadata.ID))
	if err := utils.CopyExec(path, dstEpubPath); err != nil {
		fmt.Print(err)
	}

	// Move the unprocessed unzipped epub directory to the processed directory with the epub ID as the directory name
	dstTmpPath := filepath.Join(TMP_DIR, metadata.ID)
	if err := utils.MoveExec(TMP_UNPROCESSED_DIR, dstTmpPath); err != nil {
		fmt.Print(err)
	}

	return dstTmpPath
}
