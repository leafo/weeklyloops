package main

import (
	"log"

	"github.com/leafo/weeklyloops/loops/ply"
)

func main() {
	// parser, _ := ply.NewParserFromFile("cube.ply")
	parser, _ := ply.NewParserFromString(`
ply
comment eat butt
format dads
element vertex 26
property float x
property float y
property float z
property float nx
property float ny
property float nz
element face 12
property list uchar uint vertex_indices
end_header
`)

	log.Print(parser.ParseHeader())
	log.Print(parser.Elements)
}
