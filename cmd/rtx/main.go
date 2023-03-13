package main

import (
	"log"
	"math"
	"os"
	"sergioffpc/rtx/pkg/rtx"

	"github.com/schollz/progressbar/v3"
)

func main() {
	width, height := 1280, 720
	film := rtx.NewFilm(width, height)
	camera := rtx.NewCamera(width, height, math.Pi/3)
	camera.LookAt(rtx.Point3{X: 0, Y: 1.5, Z: -5}, rtx.Point3{X: 0, Y: 1, Z: 0}, rtx.Vector3{X: 0, Y: 1, Z: 0})
	scene := rtx.Scene{
		Geometries: []rtx.GeometricPrimitive{
			{
				Shape: rtx.PlaneShape{},
				Material: rtx.PhongMaterial{
					Ks:    0,
					Kd:    0.9,
					Ka:    0.1,
					Alpha: 200,
					Color: rtx.Spectrum{R: 1, G: 0.9, B: 0.9},
				},
				ObjectToWorld: rtx.IdentityTransform(),
				WorldToObject: rtx.IdentityTransform().Inverse(),
			},
			{
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Color: rtx.Spectrum{R: 0.1, G: 1, B: 0.5},
				},
				ObjectToWorld: rtx.TranslateTransform(-0.5, 1, 0.5),
				WorldToObject: rtx.TranslateTransform(-0.5, 1, 0.5).Inverse(),
			},
			{
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Color: rtx.Spectrum{R: 0.5, G: 1, B: 0.1},
				},
				ObjectToWorld: rtx.ChainTransform(
					rtx.ScaleTransform(0.5, 0.5, 0.5),
					rtx.TranslateTransform(1.5, 0.5, -0.5),
				),
				WorldToObject: rtx.ChainTransform(
					rtx.ScaleTransform(0.5, 0.5, 0.5),
					rtx.TranslateTransform(1.5, 0.5, -0.5),
				).Inverse(),
			},
			{
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Color: rtx.Spectrum{R: 1, G: 0.8, B: 0.1},
				},
				ObjectToWorld: rtx.ChainTransform(
					rtx.ScaleTransform(0.33, 0.33, 0.33),
					rtx.TranslateTransform(-1.5, 0.33, -0.75),
				),
				WorldToObject: rtx.ChainTransform(
					rtx.ScaleTransform(0.33, 0.33, 0.33),
					rtx.TranslateTransform(-1.5, 0.33, -0.75),
				).Inverse(),
			},
		},
		Lights: []rtx.LightPrimitive{
			{
				Light: rtx.PointLight{
					I: rtx.Spectrum{R: 100, G: 100, B: 100},
				},
				LightToWorld: rtx.TranslateTransform(-10, 10, -10),
				WorldToLight: rtx.TranslateTransform(-10, 10, -10).Inverse(),
			},
		},
	}
	integrator := rtx.Whitted{}

	pb := progressbar.Default(int64(height * width))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pb.Add(1)
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
