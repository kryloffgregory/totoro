package git_test

import (
	"testing"

	"accessModel/server/git"

	"github.com/stretchr/testify/assert"
)

func TestCreatePR(t *testing.T) {
	git.CreatePR("vasya2048", "brew install cowsay", []string{"kryloffgregory"})
}

func TestProcessPRs(t *testing.T) {
	err := git.ProcessPRs()
	assert.NoError(t, err)
}
