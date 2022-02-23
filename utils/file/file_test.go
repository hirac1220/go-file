package file

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckFileSize(t *testing.T) {
	// arrange
	f, _ := os.Open("../../sample/image.jpg")
	defer func() {
		_ = f.Close()
	}()

	// action
	fi, _ := f.Stat()
	_, actual := CheckFileSize(fi.Size())

	// assert
	assert.ErrorIs(t, nil, actual)
}

func TestValidateJPEG(t *testing.T) {
	// arrange
	expected := JPEG
	bytes, _ := os.ReadFile("../../sample/image.jpeg")

	// action
	actual, err := ValidateFile(&bytes)
	if err != nil {
		log.Println(err)
	}

	// assert
	assert.Equal(t, expected, actual)
}

func TestValidateJPG(t *testing.T) {
	// arrange
	expected := JPEG
	bytes, _ := os.ReadFile("../../sample/image.jpg")

	// action
	actual, err := ValidateFile(&bytes)
	if err != nil {
		log.Println(err)
	}

	// assert
	assert.Equal(t, expected, actual)
}

func TestValidatePNG(t *testing.T) {
	// arrange
	expected := PNG
	bytes, _ := os.ReadFile("../../sample/image.png")

	// action
	actual, err := ValidateFile(&bytes)
	if err != nil {
		log.Println(err)
	}

	// assert
	assert.Equal(t, expected, actual)
}

func TestJPGExtToContentType(t *testing.T) {
	// arrange
	expected := JPEG
	f, _ := os.Open("../../sample/image.jpg")
	defer func() {
		_ = f.Close()
	}()

	// action
	var ext string
	if fi, err := f.Stat(); err == nil {
		ext = ExtractExt(fi.Name())
	}
	actual, _ := ExtToContentType(ext)

	// assert
	assert.Equal(t, expected, actual)
}

func TestJPEGExtToContentType(t *testing.T) {
	// arrange
	expected := JPEG
	f, _ := os.Open("../../sample/image.jpeg")
	defer func() {
		_ = f.Close()
	}()

	// action
	var ext string
	if fi, err := f.Stat(); err == nil {
		ext = ExtractExt(fi.Name())
	}
	actual, _ := ExtToContentType(ext)

	// assert
	assert.Equal(t, expected, actual)
}

func TestPNGExtToContentType(t *testing.T) {
	// arrange
	expected := PNG
	f, _ := os.Open("../../sample/image.png")
	defer func() {
		_ = f.Close()
	}()

	// action
	var ext string
	if fi, err := f.Stat(); err == nil {
		ext = ExtractExt(fi.Name())
	}
	actual, _ := ExtToContentType(ext)

	// assert
	assert.Equal(t, expected, actual)
}
