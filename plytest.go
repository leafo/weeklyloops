package main

import (
	"log"

	"github.com/leafo/weeklyloops/loops/ply"
)

func main() {
	//parser, _ := ply.NewParserFromFile("cube.ply")
	parser, _ := ply.NewParserFromString(`
ply
comment eat butt
format dads
element vertex 3
property float x
property float y
property float z
property float nx
property float ny
property float nz
property uchar red
property uchar green
property uchar blue
element face 2
property list uchar uint vertex_indices
end_header
0.100000 -0.910000 -0.500000 0.000000 -0.000000 -1.000000 255 255 255
-0.200000 -0.920000 -0.500000 0.000000 -0.000000 -1.000000 255 255 255
-0.300000 0.930000 -0.500000 0.000000 -0.000000 -1.000000 255 255 255
3 0 1 2
3 3 4 5
	`)

	if parser.ParseHeader() {
		object := parser.ParseBody()
		// object.Pack("hello", "world")
		v := object.Elements["vertex"]
		log.Print(v.PackF32("x", "y"))
	}
}
