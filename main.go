package main

import (
	"github.com/hirac1220/go/go-file-validate/utils/file"
	"log"
	"os"
)

func main() {
	// check file size
	f, _ := os.Open("./sample/image.jpg")
	defer func() {
		_ = f.Close()
	}()

	fi, _ := f.Stat()
	size, err := file.CheckFileSize(fi.Size())
	if err != nil {
		log.Println(err)
	}
	log.Printf("size: %d", size)

	// extension to content type
	ext := file.ExtractExt(fi.Name())
	c := file.ExtToContentType(ext)
	log.Printf("content type: %s", c)

	// validate file
	bytes, _ := os.ReadFile("./sample/image.jpeg")
	mine, err := file.ValidateFile(&bytes)
	if err != nil {
		log.Println(err)
	}
	log.Printf("mine: %s", mine)
}
