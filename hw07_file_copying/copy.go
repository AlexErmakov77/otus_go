package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
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
		log.Fatal(" offset > fileSize")
	}

	source, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err = source.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	barSize := limit
	if barSize == 0 || limit > fileSize-offset {
		barSize = fileSize - offset
	}
	destFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	bar := pb.Full.Start64(barSize)
	defer bar.Finish()
	barReader := bar.NewProxyReader(source)
	_, err = io.CopyN(destFile, barReader, barSize)
	if err != nil {
		return err
	}

	return nil
}
