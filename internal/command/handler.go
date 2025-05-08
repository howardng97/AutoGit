package command

import "fmt"

type CommandHandler struct {
	Flag    string
	Handler map[string]CommandProcess
}

// NewCommandHandler initializes a new handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Handler: make(map[string]CommandProcess),
	}
}

// RegisterCommand adds a new command to the handler
func (ch *CommandHandler) RegisterCommand(name string, handler CommandProcess) {
	ch.Handler[name] = handler
}

// Process executes the command based on the flag
func (ch *CommandHandler) Process() {
	handler, exists := ch.Handler[ch.Flag]
	if !exists {
		fmt.Printf("Error: Command '%s' not found\n", ch.Flag)
		return
	}

	info, err := handler.Process()
	if err != nil {
		fmt.Printf("Error executing '%s': %v\n", ch.Flag, err)
		return
	}
	fmt.Println("Success:", info)
}
