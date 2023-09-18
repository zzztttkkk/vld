package internal

import (
	"fmt"
	"reflect"
)

type Field struct {
	Name string
	Info reflect.StructField
}

type Type struct {
	Tpy         reflect.Type
	Fields      map[string]*Field
	Implemented bool
}

type Storage struct {
	TagName   string
	Interface reflect.Type
	Types     map[reflect.Type]*Type
}

func NewStorage(tag string, inf reflect.Type) *Storage {
	return &Storage{
		TagName:   tag,
		Interface: inf,
		Types:     map[reflect.Type]*Type{},
	}
}

func (s *Storage) TypeOf(v any) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("`%+v` is not a struct", v))
	}

	tobj := &Type{
		Tpy:         t,
		Fields:      map[string]*Field{},
		Implemented: s.Interface != nil && t.Implements(s.Interface),
	}

	s.Types[t] = tobj
	if tobj.Implemented {
		return
	}

}
