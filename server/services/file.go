package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveFile(file multipart.File, header *multipart.FileHeader, uploadDir string) (string, error) {
    defer file.Close()

    ext := filepath.Ext(header.Filename)
    filename := uuid.New().String() + ext

    path := filepath.Join(uploadDir, filepath.Base(filename))

    dst, err := os.Create(path)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        return "", err
    }

    return filename, nil
}

func DeleteFile(filename string, uploadDir string) error {
    path := filepath.Join(uploadDir, filename)
    return os.Remove(path)
}
