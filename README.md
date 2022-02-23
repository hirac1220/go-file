# Go File Validate

## Overview
* Check File Size
* Validate File
* Extension to ContentType

## Usage

```Go
import (
	"log"
	"os"

	"github.com/hirac1220/go/go-file-validate/utils/file"
)

func main() {
	// ファイルサイズチェック
	f, _ := os.Open("./sample/image.jpg") // ファイルオープン
	defer func() {
		_ = f.Close()
	}()

	fi, _ := f.Stat()
	size, err := file.CheckFileSize(fi.Size()) // ファイルサイズチェック
	if err != nil {
		log.Println(err)
	}
	log.Printf("size: %d", size)

	// ファイル拡張子からメディアタイプへ変換
	ext := file.ExtractExt(fi.Name())    // ファイル拡張子取得
	c, err := file.ExtToContentType(ext) // ファイル拡張子からメディアタイプへ変換
	if err != nil {
		log.Println(err)
	}
	log.Printf("content type: %s", c)

	// ファイルValidation
	bytes, _ := os.ReadFile("./sample/image.jpg") // ファイル読み込み
	mine, err := file.ValidateFile(&bytes)        // ファイルバリデーション
	if err != nil {
		log.Println(err)
	}
	log.Printf("mine: %s", mine)
}
```

## Results:
```
$ go run main.go
2022/02/23 17:51:29 size: 2604
2022/02/23 17:51:29 content type: image/jpeg
2022/02/23 17:51:29 mine: image/jpeg
```
