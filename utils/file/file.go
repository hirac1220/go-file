package file

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	MaxSize = 10 * 1024 * 1024 // 10MB
)

var (
	ErrFileSize   = errors.New("file size error")
	ErrValidation = errors.New("validation error")
)

var (
	JPG     = "image/jpeg"
	PNG     = "image/png"
	ExtJPG  = ".jpg"
	ExtJPEG = ".jpeg"
	ExtPNG  = ".png"
)

func CheckFileSize(size int64) (int64, error) {
	if size > MaxSize {
		return size, fmt.Errorf("size: %v error: %w", size, ErrFileSize)
	}
	return size, nil
}

func ValidateFile(data *[]byte) (string, error) {
	mine := http.DetectContentType(*data)

	switch {
	case mine == JPG:
		return mine, nil
	case mine == PNG:
		return mine, nil
	default:
		return mine, fmt.Errorf("mine: %v error: %w", mine, ErrValidation)
	}
}

func ExtractExt(name string) string {
	pos := strings.LastIndex(name, ".")

	return name[pos:]
}

func ExtToContentType(ext string) string {
	switch {
	case ext == ExtJPG:
		return JPG
	case ext == ExtJPEG:
		return JPG
	case ext == ExtPNG:
		return PNG
	default:
		return ext
	}
}
