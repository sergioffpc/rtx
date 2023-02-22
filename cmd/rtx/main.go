package main

import (
	"log"
	"os"
	"sergioffpc/rtx/pkg/rtx"
)

func main() {
	width, height := 1280, 720
	f := rtx.NewFilm(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			f.Set(x, y, rtx.Spectrum{R: 0, G: 0, B: 1})
		}
	}

	w, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	f.Write(w)
}
