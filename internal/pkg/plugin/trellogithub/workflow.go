package trellogithub

import (
	"github.com/merico-dev/stream/internal/pkg/plugin/trellogithub/trello"
	"github.com/merico-dev/stream/pkg/util/github"
)

const (
	defaultCommitMessage = "builder by DevStream"
	BuilderYmlTrello     = "trello-github-integ.yml"
)

var trelloWorkflow = &github.Workflow{
	CommitMessage:    defaultCommitMessage,
	WorkflowFileName: BuilderYmlTrello,
	WorkflowContent:  trello.IssuesBuilder,
}

// Options is the struct for configurations of the trellogithub plugin.
type Options struct {
	Owner       string
	Repo        string
	Branch      string
	BoardId     string
	TodoListId  string
	DoingListId string
	DoneListId  string
}
