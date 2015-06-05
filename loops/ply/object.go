package ply

import "log"

type PlyObject struct {
	Elements map[string]PlyElement
}

type PlyProperty struct {
	Type string
	Name string

	ListCountType string
	ListItemType  string
}

type PlyElement struct {
	Name       string
	Count      int
	Properties []PlyProperty
	Tuples     [][]float64
}

func (self *PlyProperty) isList() bool {
	return self.Type == "list"
}

func (self *PlyElement) isList() bool {
	return len(self.Properties) == 1 && self.Properties[0].isList()
}

func (self *PlyObject) PackF32(names ...string) []float32 {
	element := self.Elements["vertex"]
	return element.PackF32(names...)
}

func (self *PlyElement) PackF32(names ...string) []float32 {
	propertyIdxs := make([]int, 0, len(names))

	for _, name := range names {
		found := false

		for idx, p := range self.Properties {
			if p.Name == name {
				propertyIdxs = append(propertyIdxs, idx)
				found = true
				break
			}
		}

		if !found {
			log.Fatal("failed to find property when packing:", name)
		}
	}

	out := make([]float32, 0)

	for _, t := range self.Tuples {
		for _, idx := range propertyIdxs {
			out = append(out, float32(t[idx]))
		}

	}

	return out
}
