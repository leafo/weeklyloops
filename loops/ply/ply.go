package ply

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type PlyProperty struct {
	Type string
	Name string

	ListCountType string
	ListItemType  string
}

func (self *PlyProperty) isList() bool {
	return self.Type == "list"
}

type PlyElement struct {
	Name       string
	Count      int
	Properties []PlyProperty
}

type Parser struct {
	Buffer   []byte
	Pos      int
	Last     map[string]string
	Elements []PlyElement
}

func NewParserFromFile(fname string) (*Parser, error) {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	parser := &Parser{
		Buffer: bytes,
	}

	return parser, nil
}

func NewParserFromString(str string) (*Parser, error) {
	return &Parser{
		Buffer: []byte(str),
	}, nil
}

// show what's on the head of the parser
func (self *Parser) show() {
	stop := self.Pos
	for i := self.Pos + 1; i < len(self.Buffer)-1; i += 1 {
		stop = i
		if '\n' == self.Buffer[i] {
			break
		}
	}

	log.Print("LOCATION: ", string(self.Buffer[self.Pos:stop]))
}

// read from head of buffer, advance position
func (self *Parser) match(pat string) bool {
	re := regexp.MustCompile("^" + pat)
	matches := re.FindSubmatch(self.Buffer[self.Pos:])

	if matches != nil && len(matches) > 0 {
		self.Pos += len(matches[0])

		// extract matches
		if len(re.SubexpNames()) > 1 {
			self.Last = make(map[string]string)
			for idx, key := range re.SubexpNames() {
				if key == "" {
					continue
				}

				self.Last[key] = string(matches[idx])
			}

			log.Print("Set last:", self.Last)
		} else {
			self.Last = nil
		}

		return true
	}

	return false
}

func (self *Parser) matchLine(str string) bool {
	self.eatWhite()
	return self.match(regexp.QuoteMeta(str) + "(?:\n|$)")
}

func (self *Parser) eatWhite() {
	self.match(`\s*`)
}

func (self *Parser) group(wrapped func() bool) bool {
	pos := self.Pos
	passed := wrapped()

	if !passed {
		self.Pos = pos
	}

	return passed
}

func (self *Parser) parseComment() bool {
	return self.match(`comment\b.*` + "\n")
}

func (self *Parser) parseFormat() bool {
	return self.match(`format\b.*` + "\n")
}

func (self *Parser) parseElement() bool {
	return self.group(func() bool {
		if self.match(`element\s+(?P<name>\w+)\s+(?P<count>\d+)\s*` + "\n") {
			count, err := strconv.Atoi(self.Last["count"])

			if err != nil {
				log.Fatal(err)
			}

			element := PlyElement{
				Name:  self.Last["name"],
				Count: count,
			}

			for {
				prop := self.parseProperty()
				if prop == nil {
					break
				} else {
					element.Properties = append(element.Properties, *prop)
				}
			}

			self.Elements = append(self.Elements, element)
			return true
		}

		return false
	})
}

func (self *Parser) parseProperty() *PlyProperty {
	// property float x
	if self.match(`property\s+(?P<type>\w+)\s+(?P<name>\w+)\s*` + "\n") {
		return &PlyProperty{
			Name: self.Last["name"],
			Type: self.Last["type"],
		}
	}

	// property list uchar uint vertex_indices
	if self.match(`property\s+list\s+(?P<count_type>\w+)\s+(?P<item_type>\w+)\s+(?P<name>\w+)\s*` + "\n") {
		return &PlyProperty{
			Name:          self.Last["name"],
			Type:          "list",
			ListCountType: self.Last["count_type"],
			ListItemType:  self.Last["item_type"],
		}
	}

	return nil
}

func (self *Parser) ParseHeader() bool {
	return self.group(func() bool {
		if !self.matchLine("ply") {
			return false
		}

		for {
			self.show()

			if self.parseComment() || self.parseFormat() {
				continue
			}

			if self.parseElement() {
				continue
			}

			if self.matchLine("end_header") {
				return true
			}

			return false
		}

		return true
	})
}
