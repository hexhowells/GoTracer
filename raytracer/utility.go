package raytracer

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)


func SeedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}


func RandomFloat64() float64 {
	return rand.Float64()
}


func RandomFloat64Range(min float64, max float64) float64 {
	return min + rand.Float64() * (max - min)
}


func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}


func Clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}

func RandomVector(min float64, max float64) Vector {
	return Vector{RandomFloat64Range(min, max), RandomFloat64Range(min, max), RandomFloat64Range(min, max)}
}


func RandomInUnitSphere() Vector {
	for {
		p := RandomVector(-1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p.UnitVector()
	}
}


func WritePPMHeader(imageWidth int, imageHeight int) {
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)
}


func WriteColourPPM(pixel Vector, samplesPerPixel int) {
	samples := 1.0 / float64(samplesPerPixel)

	rPixel := math.Sqrt(pixel.X * samples)
	gPixel := math.Sqrt(pixel.Y * samples)
	bPixel := math.Sqrt(pixel.Z * samples)

	fmt.Printf("%d %d %d\n", 
		int(256 * Clamp(rPixel, 0.0, 0.999)),
		int(256 * Clamp(gPixel, 0.0, 0.999)),
		int(256 * Clamp(bPixel, 0.0, 0.999)) )
}


func Min(a float64, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}


func RandomColour() Colour {
	return Colour{RandomFloat64(), RandomFloat64(), RandomFloat64()}
}


func Reflect(v Vector, n Vector) Vector {
	a := n.MultiplyN( (2.0 * v.Dot(n)) )
	return v.Minus(a)
}


func Refract(uv Vector, n Vector, etaiOverEtat float64) Vector {
	cosTheta := Min(-uv.Dot(n), 1.0)
	rOutPerp := (uv.Add(n.MultiplyN(cosTheta))).MultiplyN(etaiOverEtat)
	rOutPara := n.MultiplyN( -math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())) )

	return rOutPerp.Add(rOutPara)
}


func RandomInUnitDisk() Vector {
	for {
		p := Vector{RandomFloat64Range(-1.0, 1.0), RandomFloat64Range(-1.0, 1.0), 0.0}
		if p.LengthSquared() >= 1.0 {
			continue
		}
		return p
	}
}