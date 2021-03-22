package main

import (
	"flag"
	"fmt"
	"github.com/chai2010/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	var quality float64 = 100
	var q float32 = 0

	if len(args) < 3 {
		log.Fatal("No arguments specified.")
	}

	quality, _ = strconv.ParseFloat(args[2], 32)

	if 1 <= quality  && quality <= 100 {
		q = float32(quality)
	} else {
		log.Fatal("argument is not Number..")
	}


	f, err := os.Open(args[0]) // 元画像読み込み
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f) // 元画像デコード
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	fso, err := os.Create(args[1]) // 変換後画像作成
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	slice := strings.Split(args[1], ".")

	switch slice[len(slice)-1] { // 出力画像の拡張子によってエンコードを変える
	case "jpeg", "jpg":
		jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		png.Encode(fso, img)
	case "gif":
		gif.Encode(fso, img, nil)
	case "webp":
		webp.Encode(fso, img, &webp.Options{Lossless: false, Quality:q})
	default:
	}
}
