package main

import (
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := createRandom(rect)
	save("./newPhoto.png", img)
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	if err != nil {
		log.Println("cannot save file")
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
	imgFile.Close()
}

func createRandom(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]byte, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}
