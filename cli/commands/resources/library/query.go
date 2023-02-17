package library

import (
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	libraryLib "github.com/taubyte/tau/lib/library"
	libraryPrompts "github.com/taubyte/tau/prompts/library"
	libraryTable "github.com/taubyte/tau/table/library"
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

	library, err := libraryPrompts.GetOrSelect(ctx)
	if err != nil {
		return err
	}
	libraryTable.Query(library)

	return nil
}

func list(ctx *cli.Context) error {
	libraries, err := libraryLib.ListResources()
	if err != nil {
		return err
	}

	libraryTable.List(libraries)
	return nil
}
