package args_test

import (
	"reflect"
	"testing"

	tauCLI "github.com/taubyte/tau/cli"
	argsLib "github.com/taubyte/tau/cli/args"
)

func TestPostfix(t *testing.T) {
	app, err := tauCLI.New()
	if err != nil {
		t.Error(err)
		return
	}

	parsedFlags := argsLib.ParseFlags(app.Flags)
	testArgs := []string{"tau", "login", "--env", "someName", "--color", "never"}

	gotArgs := argsLib.MovePostfixOptions(testArgs, parsedFlags)
	expectedArgs := []string{"tau", "--env", "--color", "never", "login", "someName"}
	if reflect.DeepEqual(gotArgs, expectedArgs) == false {
		t.Errorf("\nExpected: %v\nGot     : %v", expectedArgs, gotArgs)
	}
}
