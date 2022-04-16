package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum ...float64) float64 {
	if len(sidesNum) == 0 {
		return math.Pi * sideLen * sideLen
	}
	if len(sidesNum) == 1 {
		return sideLen * sidesNum[0]
	}
	if len(sidesNum) == 2 {
		b := sidesNum[0]
		c := sidesNum[1]
		p := (sideLen + b + c) / 2
		return math.Sqrt(p * (p - sideLen) * (p - b) * (p - c))
	}
	if len(sidesNum) == 3 {
		b := sidesNum[0]
		c := sidesNum[1]
		d := sidesNum[2]
		p := (sideLen + b + c + d) / 2
		return math.Sqrt((p - d) * (p - sideLen) * (p - b) * (p - c))
	}
	p := sideLen
	for _, d := range sidesNum {
		p *= d
	}
	return p
}
