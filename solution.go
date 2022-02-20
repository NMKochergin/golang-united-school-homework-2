package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
import (
	"math"
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	if sidesNum == 2 || sidesNum > 4 || sidesNum == 1 {
		return 0
	} else {
		if sidesNum == 4 {
			return sideLen * sideLen
		} else if sidesNum == 3 {
			return math.Sqrt(3) / 4 * sideLen * sideLen
		} else if sidesNum == 0 {
			pi := math.Pi
			return pi * sideLen * sideLen
		}
	}
	return 0
}
