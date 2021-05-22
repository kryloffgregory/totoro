package issue_test

import (
	"testing"

	"accessModel/server/issue"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpsertIssue(t *testing.T) {
	iss:=&issue.Issue{
		ID:        uuid.New(),
		Command:   issue.Command{
			Type:           issue.CommandTypeInstall,
			InstallPayload: &issue.InstallPayload{
				PackageName: "cowsay",
				Version:     "1.0",
			},
			RemovePayload:  nil,
		},
		SrcUser:   "cheburashka",
		Decisions: map[string]issue.Decision{
			"shapoklyak": issue.DecisionApprove,
			"lariska":issue.NoDecision,
		},
	}

	err:=issue.UpsertIssue(iss)
	assert.NoError(t, err)

	gotIss, err:=issue.GetIssue(iss.ID)
	assert.NoError(t, err)

	assert.Equal(t, iss, gotIss)

	iss.Decisions["lariska"] = issue.DecisionReject

	err=issue.UpsertIssue(iss)
	assert.NoError(t, err)

	gotIss, err=issue.GetIssue(iss.ID)
	assert.NoError(t, err)

	assert.Equal(t, iss, gotIss)
}

