package main

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/howardng97/AutoGit/internal/command"
	"github.com/howardng97/AutoGit/internal/command/diff"
)

func main() {
	ch := command.NewCommandHandler()
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
		Commands: []*cli.Command{
			{
				Name:        "diff",
				Category:    "Auto add different to the commit message",
				Usage:       "autogit diff",
				UsageText:   "diff - add different to the commit",
				Description: "no really, there is a lot of dooing to be done",
				ArgsUsage:   "[arrgh]",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "forever", Aliases: []string{"forevvarr"}},
				},
				Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
					ch.RegisterCommand("diff", &diff.DiffCommandHandler{})
					return nil, nil
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Bool("forever") {
						cmd.Run(ctx, nil)
					}
					return nil
				},
				OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
					fmt.Fprintf(cmd.Root().Writer, "for shame\n")
					return err
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
