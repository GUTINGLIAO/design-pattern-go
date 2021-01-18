package factory_test

import (
	. "github.com/GUTINGLIAO/design-pattern-go/factory"
	"reflect"
	"testing"
)

func TestNewProductPositive(t *testing.T) {
	inputs := []int{PenType, PencilType, 3}
	expected := []reflect.Type{reflect.TypeOf(&Pen{}), reflect.TypeOf(&Pencil{})}
	for i := 0; i < len(inputs); i++ {
		p, err := NewProduct(inputs[i])
		if err != nil {
			t.Errorf("test failed, %v", err.Error())
		}

		if reflect.TypeOf(p) != expected[i] {
			t.Errorf("test Failed, wrong type")
		}
	}
}

func TestNewProductNegative(t *testing.T) {
	inputs := []int{3, 4, 5, 6}
	expected := "wrong type"
	for i := 0; i < len(inputs); i++ {
		_, err := NewProduct(inputs[i])
		if err == nil {
			t.Fatal("test Failed, an error is expected")
		}
		if err.Error() != expected {
			t.Fatal("test Failed, wrong err message")
		}
	}
}
