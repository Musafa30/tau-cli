package prompts

import (
	"github.com/taubyte/tau/flags"
	"github.com/urfave/cli/v2"
)

func GetClone(ctx *cli.Context, prev ...bool) bool {
	return GetOrAskForBoolDefaultTrue(ctx, flags.Clone.Name, ClonePrompt, prev...)
}
