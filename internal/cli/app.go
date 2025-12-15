package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"

	"version-cli/internal/flags"

	"version-cli/internal/commands"
)

// New creates cli.App for the release CLI
func New(author, version string) *cli.App {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show help",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the version"}

	return &cli.App{
		Name:     "version-cli",
		Usage:    "A CLI tool that interacts with GitLab's variables API",
		Version:  version,
		HelpName: "help",
		Description: `
CLI tool that interacts with GitLab's variables API.
Get started with version-cli https://gitlab.com/gitlab-org/version-cli.`,
		Before: func(context *cli.Context) error {
			debug := context.Bool(flags.Debug)
			if debug {
				log.Println("Debug ON")
			}
			return nil
		},
		Commands: []*cli.Command{
			commands.Version(),
			commands.PushVersion(),
			commands.Tag(),
		},
		CommandNotFound: func(context *cli.Context, cmd string) {
			fmt.Errorf("Command not found: %q", cmd)
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return cli.ShowAppHelp(context)
		},
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: author,
			},
		},
		Flags: flags.BaseFlags(),
		//EnableBashCompletion: true,
	}
}
