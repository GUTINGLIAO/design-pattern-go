/**
factory包实现了工厂模式及其一系列变种

工厂模式
作用
1. 解耦，调用方不需要考虑产品的生产过程，只需要传入相应的类型，也可以说是将创建实例的控制权交给了工厂
2. 通过接口约定了产品需要实现的功能

*/

package factory

import (
	"errors"
	"fmt"
)

const (
	PencilType = iota
	PenType
)

type Product interface {
	Detail()
}

type Pencil struct {
	Width int
	Depth int
}

func (p *Pencil) Detail() {
	fmt.Printf("Depth:%v, Width:%v", p.Depth, p.Width)
}

type Pen struct {
	Width  int
	Length int
}

func (p *Pen) Detail() {
	fmt.Printf("Length:%v, Width:%v", p.Length, p.Width)
}

func NewProduct(productType int) (Product, error) {
	switch productType {
	case PencilType:
		return &Pencil{}, nil
	case PenType:
		return &Pen{}, nil
	default:
		return nil, errors.New("wrong type")
	}
}
