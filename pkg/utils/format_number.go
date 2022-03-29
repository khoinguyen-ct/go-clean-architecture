package utils

import "math"

func RoundNearest(x float64, precision float64) (y float64) {
	return math.Round(x*math.Pow(10, precision)) / math.Pow(10, precision)
}

func RoundDown(x float64, precision float64) (y float64) {
	return math.Floor(x*math.Pow(10, precision)) / math.Pow(10, precision)
}

func RoundUp(x float64, precision float64) (y float64) {
	return math.Ceil(x*math.Pow(10, precision)) / math.Pow(10, precision)
}
