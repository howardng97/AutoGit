package diff

type DiffCommandHandler struct {
	Diff bool
}

func (dh *DiffCommandHandler) Process() (success bool, err error) {
	return false, nil
}
