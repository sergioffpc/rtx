package rtx

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Film interface {
	Set(x, y int, s Spectrum)
	Write(w io.Writer)
}

type ImageFilm struct {
	width, height int
	pixels        []Spectrum
}

func NewImageFilm(width, height int) ImageFilm {
	return ImageFilm{
		width:  width,
		height: height,
		pixels: make([]Spectrum, width*height),
	}
}

func (f ImageFilm) Set(x, y int, s Spectrum) {
	f.pixels[f.at(x, y)] = s
}

func (f ImageFilm) Write(w io.Writer) error {
	img := image.NewNRGBA(image.Rect(0, 0, f.width, f.height))

	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			s := Spectrum.Clamp(f.get(x, y), 0, 1)
			img.Set(x, y, color.NRGBA{
				R: uint8(s.R * 255),
				G: uint8(s.G * 255),
				B: uint8(s.B * 255),
				A: 255,
			})
		}
	}

	return png.Encode(w, img)
}

func (f ImageFilm) get(x, y int) Spectrum { return f.pixels[f.at(x, y)] }

func (f ImageFilm) at(x, y int) int { return y*f.width + x }
