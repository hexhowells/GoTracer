package raytracer

import (
	"math"
)


type Vector struct {
	X float64
	Y float64
	Z float64
}

// type alias
type Colour = Vector


func (v Vector) Minus(v2 Vector) Vector {
	return Vector {
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}


func (v Vector) Add(v2 Vector) Vector {
	return Vector {
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}


func (v Vector) Multiply(v2 Vector) Vector {
	return Vector {
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}


func (v Vector) Divide(v2 Vector) Vector {
	return Vector {
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
		Z: v.Z / v2.Z,
	}
}


func (v Vector) DivideN(n float64) Vector {
	return v.Divide(Vector{X:n, Y:n, Z:n})
}


func (v Vector) MultiplyN(n float64) Vector {
	return v.Multiply(Vector{X:n, Y:n, Z:n})
}


func (v Vector) MinusN(n float64) Vector {
	return v.Minus(Vector{X:n, Y:n, Z:n})
}


func (v Vector) LengthSquared() float64 {
	return (v.X*v.X) + (v.Y*v.Y) + (v.Z*v.Z)
}


func (v Vector) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}


func (v Vector) UnitVector() Vector {
	return v.DivideN(v.Length())
}


func RandomUnitVector() Vector {
	return RandomInUnitSphere().UnitVector()
}


func (v Vector) Dot(v2 Vector) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y) + (v.Z * v2.Z)
}


func (v Vector) Cross(v2 Vector) Vector {
	return Vector{v.Y * v2.Z - v.Z * v2.Y,
				  v.Z * v2.X - v.X * v2.Z ,
				  v.X * v2.Y - v.Y * v2.X}
}


func (v Vector) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}
