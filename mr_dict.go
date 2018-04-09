package main

type Dict struct {
	List    []string
	Mapping map[string]struct{}
}

func NewDictFromFile(filename string) Dict {
	lines := ReadFileAsLines(filename)
	d := Dict{
		lines,
		make(map[string]struct{}, len(lines)),
	}
	for _, word := range d.List {
		d.Mapping[word] = struct{}{}
	}
	return d
}

func (d Dict) Has(s string) bool {
	if _, ok := d.Mapping[s]; ok {
		return true
	}
	return false
}
