package diff

type DiffCommandHandler struct {
	Diff bool
}

func (dh *DiffCommandHandler) Process() (string, error) {
	return "You call --diff", nil
}
