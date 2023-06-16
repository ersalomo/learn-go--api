package test

import "testing"
import "github.com/stretchr/testify/assert"

func Add(a, b int) int {
	return a + b

}
func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	assert.Equal(t, expected, result, "should add  two number ")
}
