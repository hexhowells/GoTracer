package raytracer

import ( "math" )

type ScatterFunc func(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool


type Material struct {
 	Scatter ScatterFunc
 	Colour Colour
    Fuzz float64
    Light bool
}


func LambertianScatter(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool {
 	scatterDirection := rec.Normal.Add(RandomUnitVector())

 	if scatterDirection.NearZero() {
 		scatterDirection = rec.Normal
 	}

 	*scattered = Ray{rec.P, scatterDirection}
 	*attenuation = mat.Colour

 	return true
}


func MetalScatter(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool {
    reflected := Reflect(rIn.Direction.UnitVector(), rec.Normal)

    fuzzedReflection := RandomInUnitSphere().MultiplyN(mat.Fuzz)
    *scattered = Ray{rec.P, reflected.Add(fuzzedReflection)}
    *attenuation = mat.Colour

    return rec.Normal.Dot(scattered.Direction) > 0
}


func LightScatter(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool {
    return false
}


func DielectricScatter(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool {
    *attenuation = Vector{1.0, 1.0, 1.0}
    refractionIndex := 1.5
    var refractionRatio float64
    
    if rec.FrontFace {
        refractionRatio = 1.0 / refractionIndex
    } else {
        refractionRatio = refractionIndex
    }

    unitDirection := rIn.Direction.UnitVector()
    cosTheta := Min(rec.Normal.Dot(unitDirection.MultiplyN(-1.0)), 1.0)
    sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

    cannotRefract := refractionRatio * sinTheta > 1.0
    var direction Vector

    if cannotRefract || (Reflectance(cosTheta, refractionRatio) > RandomFloat64()) {
        direction = Reflect(unitDirection, rec.Normal)
    } else {
        direction = Refract(unitDirection, rec.Normal, refractionRatio)
    }

    *scattered = Ray{rec.P, direction}

    return true
}


func IsotropicScatter(rIn Ray, rec *HitRecord, attenuation *Vector, scattered *Ray, mat Material) bool {
    *scattered = Ray{rec.P, RandomInUnitSphere()}
    *attenuation = mat.Colour
    return true
}


func Reflectance(cosine float64, refIdx float64) float64 {
    // Schlick's approximation for reflectance
    r0 := (1.0 - refIdx) / (1.0 + refIdx)
    r0 = r0 * r0
    return r0 + (1.0-r0) * math.Pow((1.0 - cosine), 5)
}