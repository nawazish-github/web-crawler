package model

//Generation encapsulates site-map for a
//particular generation in the generation tree
type Generation struct {
	GenMap map[string][]string
	Next   *Generation
}

//NewGeneration returns a newly initialized
//Generation struct
func NewGeneration() *Generation {
	return &Generation{
		GenMap: make(map[string][]string),
	}
}
