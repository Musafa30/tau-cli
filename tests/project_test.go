package tests

import (
	"testing"
)

func TestProjectAll(t *testing.T) {
	runTests(t, createProjectMonkey(), true)
}

func createProjectMonkey() *testSpider {
	// userName := "taf-test"

	beforeEach := func(tt testMonkey) [][]string {
		return [][]string{
			{
				"login", "-new",
				"someProfile",
				"--token", "123456",
				"--provider", "github",
			},
		}
	}

	basicNew := func(name string) []string {
		return []string{
			"new", "project", "-y",
			"--name", name,
			"--description", "somedesc",
			"--private",
			"--no-embed-token",

			// disable color
			"--color", "never",
		}
	}

	projectName := "someProject"
	tests := []testMonkey{
		{
			mock: true,
			name: "basic new",
			args: basicNew(projectName),
			wantOut: []string{
				"Created project: someProject",
				"Selected project: someProject",
			},
			evaluateSession: expectSelectedProject(projectName),
		},
		{
			name: "query project",
			args: []string{
				"query", "project", projectName,
			},
			preRun: [][]string{
				basicNew(projectName),
			},
			wantOut: []string{
				"test_user/tb_code_someProject",
				"test_user/tb_someProject",
				"Code", "Config", "ID", "Name",
			},
			mock:            true,
			evaluateSession: expectSelectedProject(projectName),
		},
		{
			name: "query projects",
			args: []string{
				"query", "project", "--list",
			},
			preRun: [][]string{
				basicNew(projectName + "1"),
				basicNew(projectName + "2"),
				basicNew(projectName + "3"),
				basicNew(projectName + "4"),
				basicNew(projectName + "5"),
			},
			wantOut: []string{
				projectName + "1", projectName + "2",
				projectName + "3", projectName + "4",
				projectName + "5",
			},
			mock:            true,
			evaluateSession: expectSelectedProject(projectName + "5"),
		},
		{
			name: "--env select project",
			args: []string{
				"select", "--env", "project", "--name", projectName,

				// disable color
				"--color", "never",
			},
			wantOut: []string{
				"export TAUBYTE_PROJECT=",
			},
			preRun: [][]string{
				basicNew(projectName),
			},
			mock:            true,
			evaluateSession: expectSelectedProject(projectName),
		},
	}
	return &testSpider{"some_project", tests, beforeEach, nil, "project"}
}
