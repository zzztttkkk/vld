package internal

import (
	"fmt"
	"reflect"
	"testing"
)

type Address struct {
	City string `json:"city"`
}

type User struct {
	Name    string   `json:"name" path:"users.$1.name"`
	Address *Address `json:"address"`
	age     int32
}

func TestReflect(t *testing.T) {
	tpy := reflect.TypeOf(User{})

	num := tpy.NumField()
	for i := 0; i < num; i++ {
		f := tpy.Field(i)
		if !f.IsExported() {
			continue
		}

		fmt.Println(f.Name)
	}
}
