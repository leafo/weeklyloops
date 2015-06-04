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
end_header
`)

	log.Print(parser.ParseHeader())
}
