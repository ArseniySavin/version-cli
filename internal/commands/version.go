package commands

import (
	"fmt"
	"log"
	"version-cli/internal/flags"
	"version-cli/internal/pkg"

	"github.com/urfave/cli/v2"
)

func Version() *cli.Command {
	return versionVarWithFlags(flags.VersionFlag{}, "version")
}

func versionVarWithFlags(flags flags.PassedInFlags, name string) *cli.Command {
	flagList := *flags.ListFlags()

	return &cli.Command{
		Name:  name,
		Usage: "Make version",
		Action: func(ctx *cli.Context) error {
			version, err := version(ctx)
			if version.Tag != "" {
				fmt.Println(version.Tag)
				return err
			}

			fmt.Println(version.String())
			return err

		},
		Before: flags.BeforeHook(flagList),
		Subcommands: []*cli.Command{
			PushVersion(),
		},
		Flags: flagList,
		CustomHelpTemplate: `
		USAGE:
				{{.Usage}}
		COMMANDS:
				{{range .Subcommands}}{{.Name}}: {{.Usage}}
				{{end}}
		OPTIONS:
				{{range .VisibleFlags}}{{.}}
				{{end}}`,
	}
}

func version(ctx *cli.Context) (*pkg.Version, error) {
	debug := ctx.Bool(flags.Debug)
	value := ctx.String(flags.Value)
	release := ctx.Bool(flags.Release)
	patch := ctx.Bool(flags.Patch)
	build := ctx.String(flags.Build)
	revision := ctx.String(flags.Revision)
	autoRevision := ctx.Bool(flags.AutoRevision)
	tag := ctx.Bool(flags.Tag)

	if debug {
		log.Printf("incoming value: %s\n", value)
		log.Printf("build name: %s\n", build)
		log.Printf("revision: %s\n", revision)
		log.Printf("is increment minor: %t\n", release)
		log.Printf("is increment path: %t\n", patch)
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
	if autoRevision {
		version.InitRevision()
	}

	if release {
		version.UpRelease()
	} else if patch {
		version.UpPatch()
	}

	if tag {
		version.Tag = version.GetTag()
	}

	if debug {
		log.Printf("out put version: %s\n", version.String())
	}

	return &version, nil
}
