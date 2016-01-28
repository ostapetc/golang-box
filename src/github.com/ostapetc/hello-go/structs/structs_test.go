package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildStaticPerson(t *testing.T) {
	p := BuildStaticPerson()

	assert.Equal(t, "Artem", p.name)
	assert.Equal(t, 30, p.age)
	assert.Equal(t, "man", p.sex)
}

func TestBuildEmptyPerson(t *testing.T) {
	p := Person{}
	assert.Equal(t, "", p.name)
	assert.Equal(t, 0, p.age)
	assert.Equal(t, "", p.sex)
}
