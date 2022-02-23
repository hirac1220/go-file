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
	ErrNotFound   = errors.New("not found")
)

var (
	JPEG    = "image/jpeg"
	PNG     = "image/png"
	ExtJPG  = ".jpg"
	ExtJPEG = ".jpeg"
	ExtPNG  = ".png"
)

func CheckFileSize(size int64) (int64, error) {
	if size > MaxSize { // MaxSizeより大きい場合はエラー
		return size, fmt.Errorf("size: %v error: %w", size, ErrFileSize)
	}
	return size, nil
}

func ValidateFile(data *[]byte) (string, error) {
	mine := http.DetectContentType(*data)

	switch {
	case mine == JPEG: // MINEタイプimage/jpeg
		return mine, nil
	case mine == PNG: // MINEタイプがimage/png
		return mine, nil
	default:
		return mine, fmt.Errorf("mine: %v error: %w", mine, ErrValidation)
	}
}

func ExtractExt(name string) string {
	pos := strings.LastIndex(name, ".")

	// .から後ろをスライスで取得
	return name[pos:]
}

func ExtToContentType(ext string) (string, error) {
	switch {
	case ext == ExtJPG: // 拡張子が".jpg"の場合 ContentType：image/jpeg
		return JPEG, nil
	case ext == ExtJPEG: // 拡張子が".jpeg"の場合 ContentType：image/jpeg
		return JPEG, nil
	case ext == ExtPNG: // 拡張子が".png"の場合 ContentType：image/png
		return PNG, nil
	default:
		return ext, fmt.Errorf("extension: %v error: %w", ext, ErrNotFound)
	}
}
