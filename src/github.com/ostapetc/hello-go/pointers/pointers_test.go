package pointers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestChangeIntValueTo(t *testing.T) {
	n := 100
	ChangeIntValueTo(&n, 200)
	assert.Equal(t, 200, n)
}
