package node_test

import (
	"fmt"
	"testing"



	"github.com/google/uuid"
	"github.com/kryloffgregory/totoro/server/node"
	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	fmt.Println(uuid.New())
	nod:=&node.Node{
		LibName:        "python3",
		CriticalFor: []string{"vasya", "petya"},
	}

	assert.NoError(t, node.UpsertNode(nod))

	gotNode, err:= node.GetNode("python3")
	assert.NoError(t, err)

	assert.Equal(t, nod, gotNode)

	assert.NoError(t, node.DeleteNode("python3"))

	nod, err= node.GetNode("python3")
	assert.Error(t, err)
}
