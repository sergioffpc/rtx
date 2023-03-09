package main

import (
	"log"
	"os"
	"sergioffpc/rtx/pkg/rtx"
)

func main() {
	width, height := 1280, 720
	film := rtx.NewFilm(width, height)
	camera := rtx.NewCamera(width, height)
	scene := rtx.Scene{
		Geometries: []rtx.GeometricPrimitive{
			{
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.9,
					Kd:    0.9,
					Ka:    0.1,
					Alpha: 200,
					Color: rtx.Spectrum{R: 0, G: 0, B: 1},
				},
				ObjectToWorld: rtx.TranslateTransform(0, 0, 5),
				WorldToObject: rtx.TranslateTransform(0, 0, 5).Inverse(),
			},
		},
		Lights: []rtx.LightPrimitive{
			{
				Light: rtx.PointLight{
					Ia: rtx.Spectrum{R: 1, G: 1, B: 1},
				},
				ObjectToWorld: rtx.TranslateTransform(5, 0, 0),
				WorldToObject: rtx.TranslateTransform(5, 0, 0).Inverse(),
			},
		},
	}
	integrator := rtx.Whitted{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := camera.GenerateRay(x, y)
			l := integrator.Render(&scene, r)
			film.Set(x, y, l)
		}
	}

	w, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	film.Write(w)
}
