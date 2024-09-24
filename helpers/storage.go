package helpers

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := validateFileSize(fileHeader); err != nil {
		return "", err
	}

	fileExt := filepath.Ext(fileHeader.Filename)
	fileName := CreateId() + fileExt
	savePath := filepath.Join("/app/uploads", fileName)

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	imageUrl := os.Getenv("IMAGE_URL")
	return imageUrl +"/uploads"+ fileName, nil
}

func validateFileSize(fileHeader *multipart.FileHeader) error {
	const (
		maxPhotoSize = 5 * 1024 * 1024
		maxVideoSize = 15 * 1024 * 1024
	)

	fileType := fileHeader.Header.Get("Content-Type")
	fileSize := fileHeader.Size

	switch {
	case isPhoto(fileType) && fileSize > maxPhotoSize:
		return errors.New("photo size cannot exceed 5 MB")
	case isVideo(fileType) && fileSize > maxVideoSize:
		return errors.New("video size cannot exceed 15 MB")
	}

	return nil
}

func isPhoto(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}

func isVideo(contentType string) bool {
	return strings.HasPrefix(contentType, "video/")
}