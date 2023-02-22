package rtx

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Film struct {
	Width, Height int
	pixels        []Spectrum
}

func NewFilm(width, height int) Film {
	return Film{
		Width:  width,
		Height: height,
		pixels: make([]Spectrum, width*height),
	}
}

func (f Film) Get(x, y int) Spectrum { return f.pixels[f.at(x, y)] }

func (f Film) Set(x, y int, s Spectrum) {
	f.pixels[f.at(x, y)] = s
}

func (f Film) Write(w io.Writer) error {
	img := image.NewNRGBA(image.Rect(0, 0, f.Width, f.Height))

	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			s := Spectrum.Clamp(f.Get(x, y), 0, 1)
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

func (f Film) at(x, y int) int { return y*f.Width + x }
