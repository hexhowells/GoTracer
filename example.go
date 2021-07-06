package main

import (
	"math"
	"./raytracer"
	"pb/v3"
)


func rayColour(r raytracer.Ray, scene raytracer.Scene, depth int) raytracer.Colour {
	rec := raytracer.HitRecord{}

	if depth <= 0 {
		return raytracer.Colour{0.0, 0.0, 0.0}
	}

	if scene.Hit(r, 0.001, math.Inf(1), &rec) {
		scattered := raytracer.Ray{}
		attenuation := raytracer.Vector{}
		if rec.Mat.Scatter(r, &rec, &attenuation, &scattered, rec.Mat) {
			return attenuation.Multiply(rayColour(scattered, scene, depth-1))
		}
		if rec.Mat.Light {
			return rec.Mat.Colour
		}
	}
	return raytracer.Colour{0.0, 0.0, 0.0}	
}


func createScene() raytracer.Scene {
	scene := raytracer.Scene{}
	
	// Materials
	red := raytracer.Material{raytracer.LambertianScatter, raytracer.Colour{0.65, 0.05, 0.05}, 0.0, false}
	white := raytracer.Material{raytracer.LambertianScatter, raytracer.Colour{0.73, 0.73, 0.73}, 0.0, false}
	green := raytracer.Material{raytracer.LambertianScatter, raytracer.Colour{0.12, 0.45, 0.15}, 0.0, false}
	light := raytracer.Material{raytracer.LightScatter, raytracer.Colour{20, 20, 20}, 0.0, true}
	glass := raytracer.Material{raytracer.DielectricScatter, raytracer.Colour{0.0, 0.0, 0.0}, 0.0, false}
	metal := raytracer.Material{raytracer.MetalScatter, raytracer.Colour{1, 1, 1}, 0.0, false}

	// Room
	scene.AddObject(raytracer.YZRect{0, 555, 0, 555, 705, green})  // left wall
	scene.AddObject(raytracer.YZRect{0, 555, 0, 555, -150, red})  // right wall
	scene.AddObject(raytracer.XZRect{193, 363, 207, 352, 554, light})  // light
	scene.AddObject(raytracer.XZRect{-150, 705, 0, 555, 0, white})  // floor
	scene.AddObject(raytracer.XZRect{-150, 705, 0, 555, 555, white})  // ceiling
	scene.AddObject(raytracer.XYRect{-150, 705, 0, 555, 555, white})  // back wall

	// Objects in the Cornell Box
	newBox1 := raytracer.NewBox(raytracer.Vector{265, 0, 295}, raytracer.Vector{430, 330, 460}, white)
	rotatedBox1 := raytracer.NewRotateY(newBox1, 15)
	translatedBox1 := raytracer.Translate{rotatedBox1, raytracer.Vector{-60, 0, 70}}
	scene.AddObject(translatedBox1)
	
	newBox2 := raytracer.NewBox(raytracer.Vector{130, 0, 65}, raytracer.Vector{295, 165, 230}, white)
	rotatedBox2 := raytracer.NewRotateY(newBox2, -18)
	translatedBox2 := raytracer.Translate{rotatedBox2, raytracer.Vector{25, 0, -30}}
	scene.AddObject(raytracer.ConstantMedium{translatedBox2, 0.01, raytracer.Colour{1,1,1}})

	scene.AddObject(raytracer.Sphere{raytracer.Vector{30, 340, 160}, 90, glass})
	scene.AddObject(raytracer.Sphere{raytracer.Vector{550, 90, 220}, 90, metal})

	return scene
}


func main() {
	// Image
	aspectRatio := 3.0 / 2.0
	imageWidth := 800
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 300
	maxDepth := 50
	img := raytracer.NewImage(imageWidth, imageHeight, samplesPerPixel)

	// Scene
	scene := createScene()

	// Camera
	lookFrom := raytracer.Vector{278, 278, -800}
	lookAt := raytracer.Vector{278, 278, 0}
	viewUp := raytracer.Vector{0.0, 1.0, 0.0}
	distToFocus := (lookFrom.Minus(lookAt)).Length()
	aperture := 0.0
	vfov := 40.0

	cam := raytracer.NewCamera(lookFrom, lookAt, viewUp, vfov, aspectRatio, aperture, distToFocus)

	// Render
	progressBar := pb.StartNew(imageHeight)

	for j:=imageHeight-1; j >= 0; j-- {
		progressBar.Increment()
		for i:=0; i < imageWidth; i++ {
			pixel := raytracer.Colour{0.0, 0.0, 0.0}
			for s:=0; s < samplesPerPixel; s++ {
				u := ( float64(i) + raytracer.RandomFloat64() ) / float64(imageWidth-1)
				v := ( float64(j) + raytracer.RandomFloat64() ) / float64(imageHeight-1)

				r := cam.GetRay(u, v)

				pixel = pixel.Add(rayColour(r, scene, maxDepth))
			}
			
			y := -1 * j + imageHeight
			img.WriteColour(i, y, pixel)
		}
	}
	progressBar.Finish()
	img.SaveAsPng("render.png")
}