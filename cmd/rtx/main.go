package main

import (
	"flag"
	"log"
	"math"

	"sergioffpc/rtx/pkg/rtx/camera"
	"sergioffpc/rtx/pkg/rtx/cgmath"
	"sergioffpc/rtx/pkg/rtx/color"
	"sergioffpc/rtx/pkg/rtx/film"
	"sergioffpc/rtx/pkg/rtx/integrator"
	"sergioffpc/rtx/pkg/rtx/light"
	"sergioffpc/rtx/pkg/rtx/material"
	"sergioffpc/rtx/pkg/rtx/scene"
	"sergioffpc/rtx/pkg/rtx/shape"
	"sergioffpc/rtx/pkg/rtx/texture"
)

func main() {
	var width, height int
	var msaa int

	flag.IntVar(&width, "width", 1280, "image width resolution in pixels")
	flag.IntVar(&height, "height", 720, "image height resolution in pixels")
	flag.IntVar(&msaa, "msaa", 8, "number of multisample anti-aliasing (MSAA) samples")
	flag.Parse()

	film := film.NewImageFilm(width, height, msaa)
	camera := camera.NewPerspectiveCamera(width, height, math.Pi/3)
	camera.LookAt(cgmath.Point3{X: 0, Y: 1.5, Z: -5}, cgmath.Point3{X: 0, Y: 1, Z: 0}, cgmath.Vector3{X: 0, Y: 1, Z: 0})
	scene := scene.Scene{
		Geometries: []scene.GeometricPrimitive{
			{
				Label: "floor",
				Shape: shape.PlaneShape{},
				Material: material.PhongMaterial{
					Ks:        0,
					Kd:        0.9,
					Ka:        0.1,
					Kr:        0.5,
					Kt:        1,
					Shininess: 200,
					Tex:       texture.RingTexture{Kd1: color.Spectrum{R: 1, G: 0, B: 0}, Kd2: color.Spectrum{R: 1, G: 1, B: 1}},
				},
				ObjectToWorld: cgmath.IdentityTransform(),
				WorldToObject: cgmath.IdentityTransform().Inverse(),
			},
			{
				Label: "glass ball",
				Shape: shape.SphereShape{},
				Material: material.PhongMaterial{
					Ks:           0.3,
					Kd:           0.7,
					Ka:           0.1,
					Kr:           0.5,
					Kt:           1.52,
					Shininess:    200,
					Transparency: 1,
					Tex:          texture.SolidTexture{Kd: color.Spectrum{R: 1, G: 0, B: 0}},
				},
				ObjectToWorld: cgmath.ChainTransform(
					cgmath.RotateZTransform(math.Pi/4),
					cgmath.TranslateTransform(-0.5, 1, 0.5),
				),
				WorldToObject: cgmath.ChainTransform(
					cgmath.RotateZTransform(math.Pi/4),
					cgmath.TranslateTransform(-0.5, 1, 0.5),
				).Inverse(),
			},
			{
				Label: "ball at right",
				Shape: shape.SphereShape{},
				Material: material.PhongMaterial{
					Ks:        0.3,
					Kd:        0.7,
					Ka:        0.1,
					Kr:        0.5,
					Kt:        1,
					Shininess: 200,
					Tex:       texture.SolidTexture{Kd: color.Spectrum{R: 1, G: 0, B: 0}},
				},
				ObjectToWorld: cgmath.ChainTransform(
					cgmath.ScaleTransform(0.5, 0.5, 0.5),
					cgmath.TranslateTransform(1.5, 0.5, -0.5),
				),
				WorldToObject: cgmath.ChainTransform(
					cgmath.ScaleTransform(0.5, 0.5, 0.5),
					cgmath.TranslateTransform(1.5, 0.5, -0.5),
				).Inverse(),
			},
			{
				Label: "ball at left",
				Shape: shape.SphereShape{},
				Material: material.PhongMaterial{
					Ks:        0.3,
					Kd:        0.7,
					Ka:        0.1,
					Kr:        0.5,
					Kt:        1,
					Shininess: 200,
					Tex:       texture.SolidTexture{Kd: color.Spectrum{R: 1, G: 0, B: 0}},
				},
				ObjectToWorld: cgmath.ChainTransform(
					cgmath.ScaleTransform(0.33, 0.33, 0.33),
					cgmath.RotateZTransform(math.Pi/2),
					cgmath.TranslateTransform(-1.5, 0.33, -0.75),
				),
				WorldToObject: cgmath.ChainTransform(
					cgmath.ScaleTransform(0.33, 0.33, 0.33),
					cgmath.RotateZTransform(math.Pi/2),
					cgmath.TranslateTransform(-1.5, 0.33, -0.75),
				).Inverse(),
			},
		},
		Lights: []scene.LightPrimitive{
			{
				Label: "point light",
				Light: light.PointLight{
					I: color.Spectrum{R: 100, G: 100, B: 100},
				},
				LightToWorld: cgmath.TranslateTransform(-10, 10, -10),
				WorldToLight: cgmath.TranslateTransform(-10, 10, -10).Inverse(),
			},
		},
	}
	integrator := integrator.Whitted{MaxDepth: 4}

	if err := film.Render(&scene, &integrator, &camera); err != nil {
		log.Fatal(err)
	}
}
