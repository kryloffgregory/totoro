package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig (t *testing.T) {
	conf, err:= GetUserMapping()
	assert.NoError(t, err)
	fmt.Println(conf)
}
