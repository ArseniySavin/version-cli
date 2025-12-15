package flags

import (
	"github.com/urfave/cli/v2"
)

// PushFlag is a struct used when instantiating PassedInFlags upon passing release data as Pushs.
type PushFlag struct{}

// This is the BeforeHook implementation for PushFlag.
func (PushFlag) BeforeHook(flags []cli.Flag) cli.BeforeFunc {
	return nil
}

// ListFlags implementation for PushFlag.
func (PushFlag) ListFlags() *[]cli.Flag {
	return &[]cli.Flag{
		&cli.StringFlag{
			Name:     FullURL,
			Usage:    "The full URL of the GitLab instance, including protocol and port and variable path, for example https://gitlab.com/api/v4/projects/{project-id}/variables/{variable-name}",
			Required: true,
		},
		&cli.StringFlag{
			Name:     JobToken,
			Usage:    "Job token used for authenticating with the GitLab Variable API",
			Required: false,
			EnvVars:  []string{"CI_JOB_TOKEN"},
		},
		&cli.StringFlag{
			Name:     PrivateToken,
			Usage:    "Private token used for authenticating with the GitLab Variable API, requires api scope https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html, overrides job-token",
			Required: true,
			EnvVars:  []string{"CI_PRIVATE_TOKEN"},
		},
		&cli.BoolFlag{
			Name:     Tag,
			Usage:    "Get tag from version",
			Required: false,
		},
	}
}
