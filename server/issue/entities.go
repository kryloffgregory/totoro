package issue

import "github.com/google/uuid"

type CommandType int

var (
	CommandTypeInstall CommandType = 1
	CommandTypeRemove CommandType = 2
)

type Decision int

var (
	NoDecision Decision = 0
	DecisionApprove Decision= 1
	DecisionReject Decision= 2
)

type IssueStatus int

var (
	IssueStatusPending IssueStatus = 0
	IssueStatusSuccess IssueStatus = 1
	IssueStatusFail IssueStatus = 2
)
type InstallPayload struct {
	PackageName string `json:"pn"`
	Version string `json:"v"`
}

type RemovePayload struct {
	PackageName string `json:"pn"`
}

type Command struct {
	Type CommandType `json:"ct"`
	InstallPayload *InstallPayload `json:"ip"`
	RemovePayload *RemovePayload `json:"rp"`
}

type Issue struct {
	ID uuid.UUID `json:"id"`
	Command Command `json:"cmd"`
	SrcUser string	`json:"srcUser"`
	Decisions map[string]Decision `json:"decisions"`
	ExecutionResult string `json:"execResult"`
	Status IssueStatus `json:"status"`
}
