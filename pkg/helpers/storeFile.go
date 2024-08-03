package helpers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func StoreFile(w http.ResponseWriter, file multipart.File, handler *multipart.FileHeader) (string, error) {
	// Generate a new UUID for the image name
	uniqueID := uuid.New()
	imageExtension := filepath.Ext(handler.Filename)
	imageName := uniqueID.String() + imageExtension

	// Create file in the public/uploads/ directory
	uploadDir := "public/uploads/"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	dstPath := filepath.Join(uploadDir, imageName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return dstPath, nil
}
