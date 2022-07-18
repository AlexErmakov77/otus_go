package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.

	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	if fileSize == 0 {
		return ErrUnsupportedFile
	}
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}
	if limit > fileSize {
		limit = fileSize
	}
	if offset+limit > fileSize {
		limit = fileSize - offset
	}

	source, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err = source.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	barSize := fileSize
	if limit != 0 {
		barSize = limit
	}

	destFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	bar := pb.New64(barSize)
	defer bar.Finish()
	barReader := bar.NewProxyReader(source)
	_, err = io.CopyN(destFile, barReader, barSize)
	if err != nil {
		return err
	}

	return nil
}
