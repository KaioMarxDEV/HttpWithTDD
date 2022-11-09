package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	myImg := image.NewRGBA(image.Rect(0, 0, 400, 200))

	out, err := os.Create("new.png")
	if err != nil {
		log.Fatal("failed creating new.png file")
	}
	defer out.Close()
	err = png.Encode(out, myImg)
	if err != nil {
		log.Fatal("failed to encode the file")
	}
}
