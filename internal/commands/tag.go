package commands

import (
	"fmt"
	"log"
	"version-cli/internal/flags"
	"version-cli/internal/pkg"

	"github.com/urfave/cli/v2"
)

func Tag() *cli.Command {
	return tagVarWithFlags(flags.VersionFlag{}, "tag")
}

func tagVarWithFlags(flags flags.PassedInFlags, name string) *cli.Command {
	flagList := *flags.ListFlags()

	return &cli.Command{
		Name:  name,
		Usage: "Get tag from version",
		Action: func(ctx *cli.Context) error {
			version, err := tag(ctx)
			fmt.Println(version.GetTag())

			return err

		},
		Before:      flags.BeforeHook(flagList),
		Subcommands: nil,
		Flags:       flagList,
	}
}

func tag(ctx *cli.Context) (*pkg.Version, error) {
	build := ctx.String(flags.Build)
	revision := ctx.String(flags.Revision)
	value := ctx.String(flags.Value)
	debug := ctx.Bool(flags.Debug)

	if debug {
		log.Printf("incoming value: %s\n", value)
		log.Printf("build name: %s\n", build)
		log.Printf("revision: %s\n", revision)
	}

	version := pkg.Version{}
	err := version.Parse(value)
	if err != nil {
		return nil, err
	}
	version.ParseMetadata()

	if build != "" {
		version.Build = build
	}

	if revision != "" {
		version.Revision = revision
	}

	if debug {
		log.Printf("version: %s\n", version.String())
	}

	return &version, nil
}
