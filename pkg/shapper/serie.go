package shapper

// Serie is a serie
type Serie struct {
	Name string
}

// NewSerie creates series
func NewSerie(name string) *Serie {
	return &Serie{Name: name}
}
