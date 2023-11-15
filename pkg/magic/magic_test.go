package magic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/pkg/magic"
)

func Test_Simple(t *testing.T) {
	foo := magic.Foo{}
	res := foo.Process(5)

	assert.Equal(t, 50, res)
}
