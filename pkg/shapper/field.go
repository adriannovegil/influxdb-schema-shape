package shapper

import (
	"fmt"
)

// Field is a field
type Field struct {
	Name string
}

// NewField creates fields
func NewField(args []interface{}) *Field {
	return &Field{
		Name: args[0].(string),
	}
}

func (f *Field) String() string {
	return fmt.Sprintf("    F %v", f.Name)
}

// // Use when the type return is consistent
//
// // NewField creates fields
// func NewField(args []interface{}) *Field {
// 	return &Field{
// 		Name: args[0].(string),
// 		Type: args[1].(string),
// 	}
// }
//
// // Field is a field
// type Field struct {
// 	Name string
// 	Type string
// }
//
// func (f *Field) String() string {
// 	return fmt.Sprintf("    F %v -> %v", f.Name, f.Type)
// }
//
// // NewSeries creates series
// func NewSeries(name string) *Series {
// 	return &Series{Name: name}
// }
