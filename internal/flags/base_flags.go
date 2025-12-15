package flags

import "github.com/urfave/cli/v2"

// PassedInFlags abstracts flags and before hook to specify in cli.Command,
// when creating a new release either with a file or with Versions.
type PassedInFlags interface {
	ListFlags() *[]cli.Flag
	BeforeHook([]cli.Flag) cli.BeforeFunc
}

func BaseFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     Debug,
			Usage:    "Set to true if you want extra debug output when running version-cli",
			Required: false,
			EnvVars:  []string{"CI_DEBUG_VERSION"},
		},
		&cli.StringFlag{
			Name:     Value,
			Usage:    "The value for ci env",
			Required: true,
			EnvVars:  []string{"CI_VERSION"},
		},
	}
}
