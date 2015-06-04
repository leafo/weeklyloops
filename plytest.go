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
element face 12
end_header
`)

	log.Print(parser.ParseHeader())
	log.Print(parser.Elements)
}
