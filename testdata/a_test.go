package testdata

import (
	"testing"

	v1 "github.com/ccheers/proto-gen-checker/test/v1"
)

func Test123(t *testing.T) {
	a := v1.Test{
		U32: 2,
	}
	err := check(&a)
	if err != nil {
		panic(err)
	}
}

func check(value interface{}) error {
	if checker, ok := value.(interface {
		Check() error
	}); ok {
		return checker.Check()
	}
	return nil
}
