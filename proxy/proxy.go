package proxy

import "fmt"

type Person interface {
	Eat()
	Drink()
}

type Cloud struct {
}

func (c *Cloud) Eat() {
	fmt.Printf("cloud is eating\n")
}

func (c *Cloud) Drink() {
	fmt.Printf("cloud is drinking\n")
}

func NewCloudProxy() *CloudProxy {
	return &CloudProxy{
		cloud: Cloud{},
	}
}

type CloudProxy struct {
	cloud Cloud
}

func (c *CloudProxy) Eat() {
	fmt.Printf("cloud always sleep before eat\n")
	c.cloud.Eat()
}

func (c *CloudProxy) Drink() {
	fmt.Printf("cloud always sleep before drink\n")
	c.cloud.Drink()
}
