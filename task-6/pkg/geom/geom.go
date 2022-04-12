package geom

import (
	"math"
)

// Считает дистанцию между двумя точками. Точки описаны координатами (lat и long)
func Distance(lat1, long1, lat2, long2 float64) float64 {
	return math.Sqrt(math.Pow(lat2-lat1, 2) + math.Pow(long2-long1, 2))
}
