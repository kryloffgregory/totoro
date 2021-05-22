package manage

import (
	"errors"
	"fmt"

	"accessModel/server/execute"
	"accessModel/server/issue"
	"accessModel/server/node"

	"github.com/google/uuid"
)

var loginsByUid = map[string]string {
	"501": "vasya2048",
	"502": "kryloffgregory",
}

var emailsByUid = map[string]string {
	"501": "kryloffgv@yahoo.com",
	"502": "kryloff.gv@yandex.ru",
}

func Install(userID string, packageName string, version string) error{
	decisions:=make(map[string]issue.Decision)
	affectedUsers, err:=getAffectedUsersForInstall(userID, packageName)
	if err!=nil {
		return err
	}

	for _, aUser:= range affectedUsers {
		decisions[aUser] = issue.NoDecision
	}

	iss:=&issue.Issue{
		ID:              uuid.New(),
		Command:         issue.Command{
			Type:           issue.CommandTypeInstall,
			InstallPayload: &issue.InstallPayload{
				PackageName: packageName,
				Version:     version,
			},
		},
		SrcUser:         userID,
		Decisions: decisions,
		ExecutionResult: "",
		Status:issue.IssueStatusPending,
	}

	err=issue.UpsertIssue(iss)
	if err!=nil {
		return err
	}

	return tryFinalize(iss)
}

func getAffectedUsersForInstall(userID string, packageName string) ([]string, error) {
	nod, err:=node.GetNode(packageName)
	if err!=nil {
		return nil, err
	}

	if nod == nil || !nod.Critical{
		return nil, nil
	}

	return []string{nod.UserInstalled}, nil
}

func Approve(userID string, issueID uuid.UUID) error {
	iss,err:=issue.GetIssue(issueID)
	if err!=nil {
		return err
	}
	_, ok:=iss.Decisions[userID]
	if !ok {
		return errors.New(fmt.Sprintf("user %v is not interested in issue %v", userID, iss.ID))
	}
	iss.Decisions[userID] = issue.DecisionApprove

	if err:=issue.UpsertIssue(iss); err!=nil {
		return err
	}

	return tryFinalize(iss)
}

func Reject(userID string, issueID uuid.UUID) error {
	iss,err:=issue.GetIssue(issueID)
	if err!=nil {
		return err
	}
	_, ok:=iss.Decisions[userID]
	if !ok {
		return errors.New(fmt.Sprintf("user %v is not interested in issue %v", userID, iss.ID))
	}
	iss.Decisions[userID] = issue.DecisionReject

	if err:=issue.UpsertIssue(iss); err!=nil {
		return err
	}

	return tryFinalize(iss)
}