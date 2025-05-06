package main

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var diff bool
	cmd := &cli.Command{
		Authors: []any{
			mail.Address{Name: "Howard Nguyen", Address: "longng977@gmail.com"},
		},
		Version: "v1.0.0",
		Name:    "autogit",
		Usage:   "A fast, lazy-friendly CLI tool for automated git add, commit (with diff), and push",
		Description: `AutoGit is a CLI tool that automates your git workflow:
- Adds all changes (git add .)
- Commits with a summary of the diff (--diff flag)
- Pushes to the current branch

Great for developers who commit often and want cleaner, faster git usage.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "diff",
				Usage:       "Include git diff summary in the commit message",
				Aliases:     []string{"d"},
				Destination: &diff,
			},
		},
		Action: func(cCtx context.Context, cmd *cli.Command) error {
			if cmd.Bool("diff") {
				// handle the logic here
				fmt.Println("test")
			} else {
				fmt.Println("test  false")
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
