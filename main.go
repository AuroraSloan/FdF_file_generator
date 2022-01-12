package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		prnt_usage()
	}
	jpgimg, wid, hght := get_imgdata()
	newimg, err := os.Create(strings.TrimRight(filepath.Base(os.Args[1]), ".jpg") + ".fdf")
	if err != nil {
		log.Fatal(err)
	}
	for y := 0; y < hght; y++ {
		for x := 0; x < wid; x++ {
			col := rgbToColor(jpgimg.At(x, y).RGBA()) / 4
			newimg.WriteString(fmt.Sprint(col))
			if x < wid-1 {
				newimg.Write([]byte(" "))
			}
		}
		newimg.WriteString("\n")
	}
	newimg.Close()
}

func rgbToColor(r uint32, g uint32, b uint32, a uint32) int8 {
	a = 0
	return (int8(r>>16) | int8(g>>8) | int8(b))
}

func prnt_usage() {
	fmt.Println("Usage: ./fdfgen <filename.jpg>")
	os.Exit(1)
}

func get_imgdata() (image.Image, int, int) {
	imgbytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	jpgimg, err := jpeg.Decode(bytes.NewReader(imgbytes))
	if err != nil {
		log.Fatal(err)
	}
	bounds := jpgimg.Bounds()
	wid, hght := bounds.Max.X, bounds.Max.Y
	return jpgimg, wid, hght
}
