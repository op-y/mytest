package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {

	a, b, expected := 1, 2, 3
	result := Add(a, b)
	assert.Equal(t, result, expected,
		"Add(%d, %d) is %d, expected is %d", a, b, result, expected)

	a, b, unexpected := 2, 1, 0
	result = Sub(a, b)
	assert.NotEqual(t, result, unexpected,
		"Sub(%d, %d) eq %d, of course not %d", a, b, result, unexpected)

	assert.Nil(t, nil, "nil of course eq nil")

	a, b, expected = 100, 10, 10
	result = Div(a, b)
	if assert.NotNil(t, result, "result of Div must not be nil") {
		assert.Equal(t, expected, result)
	}
}
