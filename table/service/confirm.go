package serviceTable

import (
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func Confirm(ctx *cli.Context, service *structureSpec.Service, prompt string) bool {
	return prompts.ConfirmData(ctx, prompt, [][]string{
		{"Name", service.Name},
		{"Description", service.Description},
		{"Tags", strings.Join(service.Tags, ", ")},
		{"Protocol", service.Protocol},
	})
}
