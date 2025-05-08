package command

type CommandProcess interface {
	Process() (info string, err error)
}
