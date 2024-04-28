package epubManager

import (
	"gopuby/htmlParse"
	"os"
	"path/filepath"
	"strings"
)

type Book struct {
	TableOfContents     Toc
	Metadata            Metadata
	CurrentText         string
	CurrentTextPage     *int
	CurrentChapter      string
	CurrentChapterIndex *int
}

func (book *Book) LoadBook(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		println("0")
		return book.LoadUnzippedBook(path)
	} else {
		println("0")
		dirPath := Open(path)
		println("1")
		return book.LoadUnzippedBook(dirPath)
	}

}

func (book *Book) LoadUnzippedBook(dir string) error {

	toc, err := GetTableOfContents(dir)
	if err != nil {
		return err
	}
	book.TableOfContents = *toc

	metadata, err := GetMetadata(dir)
	if err != nil {
		return err
	}
	book.Metadata = *metadata

	// load the first chapterPath from table of contents
	chapterPath := book.TableOfContents.NavPoints[*book.CurrentChapterIndex].Content.Link
	// find the # in the link and cut
	chapterPath = removeAnchor(chapterPath)
	str, err := htmlParse.ParseHtml(filepath.Join(dir, "EPUB", chapterPath))
	if err != nil {
		return err
	}

	book.CurrentText = str

	return nil
}

func removeAnchor(chapterPath string) string {
	anchorIndex := strings.Index(chapterPath, "#")
	if anchorIndex != -1 {
		return chapterPath[:anchorIndex]
	}
	return chapterPath
}

func (book *Book) MoveChapter(direction int) error {

	if book.CurrentChapterIndex == nil {
		book.CurrentChapterIndex = new(int)
		*book.CurrentChapterIndex = 0
	} else {
		*book.CurrentChapterIndex += direction
	}
	if *book.CurrentChapterIndex >= len(book.TableOfContents.NavPoints) {
		*book.CurrentChapterIndex = len(book.TableOfContents.NavPoints) - 1
	}

	if *book.CurrentChapterIndex < 0 {
		*book.CurrentChapterIndex = 0
	}

	dir := filepath.Join(TMP_DIR, book.Metadata.ID)
	chapterPath := book.TableOfContents.NavPoints[*book.CurrentChapterIndex].Content.Link
	chapterPath = removeAnchor(chapterPath)
	str, err := htmlParse.ParseHtml(filepath.Join(dir, "EPUB", chapterPath))
	if err != nil {
		return err
	}

	book.CurrentText = str

	*book.CurrentTextPage = 1

	return nil
}
