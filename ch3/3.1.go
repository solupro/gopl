package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
	"io"
	"strconv"
)

const (
	cells = 100
	xyrange = 30.0
	angle = math.Pi / 6
)
var xyscale, zscale float64
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var _width, _height int64
		if _w, ok := r.Form["w"]; ok {
			_width, _ = strconv.ParseInt(_w[0],0, 0)
		}
		if _h, ok := r.Form["h"]; ok {
			_height, _ = strconv.ParseInt(_h[0], 0, 0)
		}


		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w, _width, _height)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svg(out io.Writer, w, h int64) {
	var width, height int64

	if w > 0 {
		width = w
	} else {
		width = 600
	}

	if h > 0 {
		height = h
	} else {
		height = 320
	}

	xyscale = float64(width) / 2.0 / xyrange
	zscale = float64(height) * 0.4


	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey;fill:white;stroke-width:0.7' width='%d'> height='%d'", width, height)

	var minZ, maxZ, avgZ, rangeZ float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z, ok := corner(i, j, width, height)
			if !ok {
				continue
			}

			if z > maxZ {
				maxZ = z
			}
			if z < minZ {
				minZ = z
			}
		}
	}
	avgZ = (minZ + maxZ) / 2
	rangeZ = maxZ - minZ

	var strColor string
	for i := 0; i < cells; i++ {
		for j := 0; j < cells ; j++ {
			ax, ay, _, ok := corner(i+1, j, width, height)
			if !ok {
				continue
			}

			bx, by, z, ok := corner(i, j, width, height)
			if !ok {
				continue
			}

			cx, cy, _, ok := corner(i, j+1, width, height)
			if !ok {
				continue
			}

			dx, dy, _, ok := corner(i+i, j+1, width, height)
			if !ok {
				continue
			}

			fl := z - avgZ
			if fl < 0.0 {
				fl = -fl / (rangeZ / 2.0)
				strColor = fmt.Sprintf("#0000%02x", int(fl * 255))
			} else {
				fl = fl / (rangeZ / 2.0)
				strColor = fmt.Sprintf("#%02x0000", int(fl * 255))
			}

			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%s'/>\n", ax, ay, bx, by, cx ,cy, dx, dy, strColor)
		}
	}

	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, width, height int64) (float64, float64, float64, bool)  {
	ok := true

	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		ok = false
	}

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x-y)*sin30*xyscale - z*zscale

	return sx, sy, z, ok
}

func f(x, y float64) float64  {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
