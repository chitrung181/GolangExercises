//draw 3d image write to html file
//run: go run ex3.1.go
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	f, err := os.Create("main.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file: %s", err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, flg := corner(i+1, j)
			if !flg {
				continue
			}
			bx, by, flg := corner(i, j)
			if !flg {
				continue
			}
			cx, cy, flg := corner(i, j+1)
			if !flg {
				continue
			}
			dx, dy, flg := corner(i+1, j+1)
			if !flg {
				continue
			}
			f.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	f.WriteString("</svg>")
	/*
		fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, flg := corner(i+1, j)
				if !flg {
					continue
				}
				bx, by, flg := corner(i, j)
				if !flg {
					continue
				}
				cx, cy, flg := corner(i, j+1)
				if !flg {
					continue
				}
				dx, dy, flg := corner(i+1, j+1)
				if !flg {
					continue
				}
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Println("</svg>")
	*/
}
func corner(i, j int) (sx, sy float64, flag bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		return
	}
	flag = true
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
