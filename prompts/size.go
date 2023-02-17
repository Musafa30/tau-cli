package prompts

import (
	"fmt"
	"strings"

	schemaCommon "github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/flags"
	"github.com/taubyte/tau/validate"
	"github.com/urfave/cli/v2"
)

// TODO make generic with memory

func variableMaxSizeValidator(val string) bool {
	var err error
	if val != "" {
		if validate.IsAny(val, validate.IsInt, validate.IsBytes) == false {
			err = fmt.Errorf(InvalidSize, val)
		}
	}
	return ValidateOk(err)
}

func getOrAskForSize(c *cli.Context, prompt string, prev ...string) string {
	var ret string
	for {
		ret = GetOrAskForAStringValue(c, flags.Size.Name, prompt, prev...)
		if variableMaxSizeValidator(ret) == true {
			break
		}

		// Unset the flag to prevent it from circling back into the prompt
		if c.IsSet(flags.Size.Name) == true {
			err := c.Set(flags.Size.Name, "")
			if err != nil {
				panic(err)
			}
		}
	}
	return ret
}

// TODO should take and return expected type
func GetSizeAndType(c *cli.Context, oldSize string, isNew bool) (size string) {
	// Uppercase the relative flags
	flags.ToUpper(c, flags.Size, flags.SizeUnit)

	var memory, unitType string
	if isNew == true {
		memory = RequiredString(c, SizePrompt, getOrAskForSize)
		if _, err := schemaCommon.StringToUnits(memory); err != nil {
			unitType = GetOrAskForSelection(c, flags.SizeUnit.Name, UnitTypePrompt, common.SizeUnitTypes)
		} else {
			return memory
		}
		return memory + unitType
	} else {
		memory = RequiredString(c, SizePrompt, getOrAskForSize, oldSize)
		if _, err := schemaCommon.StringToUnits(memory); err != nil {
			var prevType string
			for _, o := range common.SizeUnitTypes {
				if strings.Contains(strings.ToUpper(oldSize), o) {
					prevType = o
				}
			}
			unitType = GetOrAskForSelection(c, flags.SizeUnit.Name, UnitTypePrompt, common.SizeUnitTypes, strings.ToUpper(prevType))
		} else {
			return memory
		}
		return memory + unitType
	}
}
