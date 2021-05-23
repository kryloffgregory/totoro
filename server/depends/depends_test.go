package depends

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRDepends(t *testing.T) {
	rdeps, err := GetRDepends("cowsay")
	assert.NoError(t, err)

	fmt.Println(rdeps)
}
