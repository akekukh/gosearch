package main

import (
	"fmt"

	"github.com/akekukh/gosearch/task-6/pkg/geom"
)

func main() {
	var lat1, long1, lat2, long2 float64 = 1, 1, 4, 5
	if lat1 < 0 || lat2 < 0 || long1 < 0 || long2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return
	}
	distance := geom.Distance(lat1, long1, lat2, long2)
	fmt.Println(distance)
}
