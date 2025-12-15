package flags

import (
	"github.com/urfave/cli/v2"
)

// VersionFlag is a struct used when instantiating PassedInFlags upon passing release data as Versions.
type VersionFlag struct{}

// This is the BeforeHook implementation for VersionFlag.
func (VersionFlag) BeforeHook(flags []cli.Flag) cli.BeforeFunc {
	return nil
}

// ListFlags implementation for VersionFlag.
func (VersionFlag) ListFlags() *[]cli.Flag {
	return &[]cli.Flag{
		&cli.BoolFlag{
			Name:     Release,
			Usage:    "The minor increment",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     Patch,
			Usage:    "The path increment",
			Required: false,
		},
		&cli.StringFlag{
			Name:        Build,
			Usage:       "The name for build. Example: Dev, Test, Alpha etc",
			Required:    false,
			DefaultText: "Empty",
		},
		&cli.StringFlag{
			Name:     Revision,
			Usage:    "The build revision. If it is empty will use UTC nano and based on crc32 or use short commit sha",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     AutoRevision,
			Usage:    "Generate revision",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     Tag,
			Usage:    "Get tag from version",
			Required: false,
		},
	}
}
