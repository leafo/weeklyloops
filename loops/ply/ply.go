package ply

import (
	"io/ioutil"
	"log"
	"regexp"
)

type Parser struct {
	Buffer []byte
	Pos    int
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
		if '\n' == self.Buffer[i] {
			break
		}
		stop = i
	}

	log.Print("LOCATION: ", string(self.Buffer[self.Pos:stop]))
}

// read from head of buffer, advance position
func (self *Parser) match(pat string) []byte {
	pattern := regexp.MustCompile("^" + pat)
	match := pattern.Find(self.Buffer[self.Pos:])
	if match != nil {
		log.Print("advancing head: ", len(match))
		self.Pos += len(match)
	}
	return match
}

func (self *Parser) matchBool(pat string) bool {
	return nil != self.match(pat)
}

func (self *Parser) matchLine(str string) bool {
	self.eatWhite()
	return nil != self.match(regexp.QuoteMeta(str)+"(?:\n|$)")
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
	return self.matchBool("comment.*\n")
}

func (self *Parser) parseFormat() bool {
	return self.matchBool("format.*\n")
}

func (self *Parser) parseElement() bool {
	return false
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

			if self.matchLine("end_header") {
				return true
			}

			return false
		}

		return true
	})
}
