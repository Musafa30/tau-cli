package website

import (
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	websiteLib "github.com/taubyte/tau/lib/website"
	websitePrompts "github.com/taubyte/tau/prompts/website"
	websiteTable "github.com/taubyte/tau/table/website"
	"github.com/urfave/cli/v2"
)

func (link) Query() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				flags.List,
			},
			Action: query,
		},
	)
}

func (link) List() common.Command {
	return common.Create(
		&cli.Command{
			Action: list,
		},
	)
}

func query(ctx *cli.Context) error {
	if ctx.Bool(flags.List.Name) == true {
		return list(ctx)
	}

	website, err := websitePrompts.GetOrSelect(ctx)
	if err != nil {
		return err
	}
	websiteTable.Query(website)

	return nil
}

func list(ctx *cli.Context) error {
	websites, err := websiteLib.ListResources()
	if err != nil {
		return err
	}

	websiteTable.List(websites)
	return nil
}
