package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func SaveAttachment(folder string, file *multipart.FileHeader) (string, error) {
	// open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// ensure upload folder exists
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", err
	}

	// generate filename (avoid collisions)
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := folder + filename

	// create destination file
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// copy content
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}
