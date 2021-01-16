package proxy_test

import (
	. "github.com/GUTINGLIAO/design-pattern-go/proxy"
	"testing"
)

func TestNewCloudProxy(t *testing.T) {
	cloudProxy := NewCloudProxy()
	cloudProxy.Eat()
}
