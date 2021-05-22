package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {
	assert.NoError(t, AddNode("ansible"))
	assert.NoError(t, AddAffected("ansible", "505"))
	assert.NoError(t, AddAffected("ansible", "506"))
	expected:=&Node{
		LibName:     "ansible",
		CriticalFor: []string{"505", "506"},
	}

	actual, err:=GetNode("ansible")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
