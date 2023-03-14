package main

import (
	"flag"
	"log"
	"math"
	"os"
	"sergioffpc/rtx/pkg/rtx"

	"github.com/schollz/progressbar/v3"
)

func main() {
	var width, height int

	flag.IntVar(&width, "width", 1280, "image width resolution in pixels")
	flag.IntVar(&height, "height", 720, "image height resolution in pixels")
	flag.Parse()

	film := rtx.NewImageFilm(width, height)
	camera := rtx.NewPerspectiveCamera(width, height, math.Pi/3)
	camera.LookAt(rtx.Point3{X: 0, Y: 1.5, Z: -5}, rtx.Point3{X: 0, Y: 1, Z: 0}, rtx.Vector3{X: 0, Y: 1, Z: 0})
	scene := rtx.Scene{
		Geometries: []rtx.GeometricPrimitive{
			{
				Label: "floor",
				Shape: rtx.PlaneShape{},
				Material: rtx.PhongMaterial{
					Ks:    0,
					Kd:    0.9,
					Ka:    0.1,
					Alpha: 200,
					Tex:   rtx.RingTexture{Kd1: rtx.Spectrum{R: 1, G: 0, B: 0}, Kd2: rtx.Spectrum{R: 1, G: 1, B: 1}},
				},
				ObjectToWorld: rtx.IdentityTransform(),
				WorldToObject: rtx.IdentityTransform().Inverse(),
			},
			{
				Label: "ball at center",
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Tex:   rtx.CheckerTexture{Kd1: rtx.Spectrum{R: 1, G: 0, B: 0}, Kd2: rtx.Spectrum{R: 1, G: 1, B: 1}},
				},
				ObjectToWorld: rtx.ChainTransform(
					rtx.RotateZTransform(math.Pi/4),
					rtx.TranslateTransform(-0.5, 1, 0.5),
				),
				WorldToObject: rtx.ChainTransform(
					rtx.RotateZTransform(math.Pi/4),
					rtx.TranslateTransform(-0.5, 1, 0.5),
				).Inverse(),
			},
			{
				Label: "ball at right",
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Tex:   rtx.StripeTexture{Kd1: rtx.Spectrum{R: 1, G: 0, B: 0}, Kd2: rtx.Spectrum{R: 1, G: 1, B: 1}},
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
				Label: "ball at left",
				Shape: rtx.SphereShape{},
				Material: rtx.PhongMaterial{
					Ks:    0.3,
					Kd:    0.7,
					Ka:    0.1,
					Alpha: 200,
					Tex:   rtx.GradientTexture{Kd1: rtx.Spectrum{R: 1, G: 0, B: 0}, Kd2: rtx.Spectrum{R: 1, G: 1, B: 1}},
				},
				ObjectToWorld: rtx.ChainTransform(
					rtx.ScaleTransform(0.33, 0.33, 0.33),
					rtx.RotateZTransform(math.Pi/2),
					rtx.TranslateTransform(-1.5, 0.33, -0.75),
				),
				WorldToObject: rtx.ChainTransform(
					rtx.ScaleTransform(0.33, 0.33, 0.33),
					rtx.RotateZTransform(math.Pi/2),
					rtx.TranslateTransform(-1.5, 0.33, -0.75),
				).Inverse(),
			},
		},
		Lights: []rtx.LightPrimitive{
			{
				Label: "point light",
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
