# GoTracer - Volumetric Path Tracing in Go
GoTracer is a Go implimentation of the ray tracer in [Ray Tracing in One Weekend](https://raytracing.github.io/).

See ```/images``` for example images rendered using the ray tracer.

```example.go``` uses the package [pb](https://github.com/cheggaaa/pb) to display a progress bar but isn't required for the raytracer.

![Example Render Showcasing the Ray Tracer's Features](https://github.com/hexhowells/GoTracer/blob/main/images/GithubExampleImage.png)

Features / Notes
-----
- Includes multiple surfaces (lambertian, metallic, dielectric, isotropic)
- Includes emissive materials
- Can output in PPM or PNG format
- Supports Spheres, Rectangles, and Boxes
- Supports instance rotation around the Y axis (shouldn't be hard to impliment rotation around X and Z)
- Adjustable depth of field and viewing angle
- Has ambient occlusion, soft shadows, colour bleeding, global illumination
- Anti-Aliasing by taking multiple samples per pixel
- Images are gamma corrected before being saved


Future Additions / Improvements
-----
- Support for cylinders and triangle meshes
- Motion blur
- Procedural textures and image texture mapping
- Importance sampling
- Bidirectional path tracing
- Add parallelism via goroutines
