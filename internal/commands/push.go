package commands

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"version-cli/internal/flags"
	"version-cli/internal/gitlab"
	"version-cli/internal/pkg"

	"github.com/urfave/cli/v2"
)

func PushVersion() *cli.Command {
	return pushVersionVarWithFlags(flags.PushFlag{}, "push")
}

func pushVersionVarWithFlags(flags flags.PassedInFlags, name string) *cli.Command {
	flagList := *flags.ListFlags()

	return &cli.Command{
		Name:  name,
		Usage: "Set value a ci env using GitLab's API https://docs.gitlab.com/ee/api/project_level_variables.html",
		Action: func(ctx *cli.Context) error {
			version, err := update(ctx)
			if version.Tag != "" {
				fmt.Println(version.Tag)
				return err
			}

			fmt.Println(version.String())
			return err
		},
		Before:      flags.BeforeHook(flagList),
		Subcommands: nil,
		Flags:       flagList,
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

func update(ctx *cli.Context) (*pkg.Version, error) {
	debug := ctx.Bool(flags.Debug)
	fullUrl := ctx.String(flags.FullURL)
	privateToken := ctx.String(flags.PrivateToken)
	tag := ctx.Bool(flags.Tag)

	if debug {
		log.Printf("url: %s\n", fullUrl)
		log.Printf("token: %s\n", privateToken)
	}

	versionData, err := version(ctx)
	if err != nil {
		return nil, err
	}

	if debug {
		log.Printf("version: %s\n", versionData.String())
	}

	c, err := gitlab.NewGitlabClient(fullUrl, "", privateToken)
	if err != nil {
		log.Fatal(err)
	}

	version := strings.Replace(versionData.String(), "+", "%2B", 1)

	req, err := c.Request(context.Background(), http.MethodPut, version)
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		dreq, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, fmt.Errorf("dump request err: %w", err)
		}
		log.Printf("%s", string(dreq))
	}

	res, err := c.Send(req)
	if err != nil {
		log.Fatal(err)
	}

	if tag {
		versionData.Tag = versionData.GetTag()
	}

	if debug {
		log.Printf("response: %s", string(res))
	}

	return versionData, nil
}
