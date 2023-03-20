package film

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"sync"

	"sergioffpc/rtx/pkg/rtx/camera"
	crtx "sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/integrator"
	"sergioffpc/rtx/pkg/rtx/sampler"
	"sergioffpc/rtx/pkg/rtx/scene"

	"github.com/schollz/progressbar/v3"
)

type ImageFilm struct {
	width, height int
	msaa          int
	pixels        []crtx.Spectrum
}

func NewImageFilm(width, height, msaa int) ImageFilm {
	return ImageFilm{
		width:  width,
		height: height,
		msaa:   msaa,
		pixels: make([]crtx.Spectrum, width*height),
	}
}

func (f ImageFilm) Render(scene *scene.Scene, integrator integrator.Integrator, camera camera.Camera) error {
	pb := progressbar.Default(int64(f.height * f.width * f.msaa))
	var wg sync.WaitGroup
	for y := 0; y < f.height; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := 0; x < f.width; x++ {
				pb.Add(1)
				var li crtx.Spectrum
				for i := 0; i < f.msaa; i++ {
					ray := camera.GenerateRay(x, y, sampler.Get2D())
					li.AddAssign(integrator.Li(scene, ray))
				}
				f.set(x, y, li)
			}
		}(y)
	}
	wg.Wait()
	pb.Finish()

	w, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	return f.write(w)
}

func (f ImageFilm) at(x, y int) int { return y*f.width + x }

func (f ImageFilm) get(x, y int) crtx.Spectrum { return f.pixels[f.at(x, y)] }

func (f ImageFilm) set(x, y int, s crtx.Spectrum) {
	f.pixels[f.at(x, y)] = s.DivFloat(float64(f.msaa))
}

func (f ImageFilm) write(w io.Writer) error {
	img := image.NewNRGBA(image.Rect(0, 0, f.width, f.height))

	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			s := crtx.Spectrum.Clamp(f.get(x, y), 0, 1)
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
