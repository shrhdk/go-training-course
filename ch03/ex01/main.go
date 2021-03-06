package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 30deg
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: green; fill: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}

			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}

			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}

			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64, err error) {
	x := realize(i)
	y := realize(j)
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 1) || math.IsInf(z, -1) {
		err = errors.New("Invalid Result")
		return
	}

	sx, sy = project(x, y, z)
	return
}

func realize(i int) float64 {
	return xyrange * (float64(i)/cells - 0.5)
}

func project(x, y, z float64) (sx, sy float64) {
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
