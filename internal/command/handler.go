package command

type CommandHanlder struct {
	Flag string
}

func (ch *CommandHanlder) Process() (success bool, err error) {
	return false, nil
}
