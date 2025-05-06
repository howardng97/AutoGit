package command

import "github.com/howardng97/AutoGit/internal/command/diff"

type CommandHanlder struct {
	Flag string
}

func (ch *CommandHanlder) Process() (success bool, err error) {
	if ch.Flag == "--diff" || ch.Flag == "-d" {
		handler := diff.DiffCommandHandler{Diff: true}
		handler.Process()
	}
	return false, nil
}
