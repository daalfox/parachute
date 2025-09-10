package parachute_test

import (
	"testing"

	"github.com/daalfox/parachute"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	pch := parachute.Group[int]{}

	expected := 71
	actual, err, shared := pch.Do("k1", func() (int, error) {
		return expected, nil
	})
	assert.NoError(t, err)
	assert.False(t, shared)
	assert.Equal(t, expected, actual)

	expected = 178
	ch := pch.DoChan("k2", func() (int, error) {
		return expected, nil
	})

	r := <-ch

	assert.NoError(t, r.Err)
	assert.False(t, r.Shared)
	assert.Equal(t, expected, r.Val)
}
