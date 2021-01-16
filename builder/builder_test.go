package builder_test

import (
	. "github.com/GUTINGLIAO/design-pattern-go/builder"
	"testing"
)

func TestUserBuilder_Build(t *testing.T) {
	userBuilder := NewUserBuilder()
	userBuilder.
		Name("Derek").
		Father("Cloud").
		Son("Huston").
		Age(22).
		Sex(0).
		Height(20).
		Build()
}
