package execute

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteString(t *testing.T) {
	res, err := ExecuteString("brew info go")
	assert.NoError(t, err)
	fmt.Println(res)
}
