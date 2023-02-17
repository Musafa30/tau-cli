package libraryTable

import (
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	repositoryLib "github.com/taubyte/tau/lib/repository"
)

func getTableData(library *structureSpec.Library, showId bool) (toRender [][]string) {
	if showId == true {
		toRender = [][]string{
			{"ID", library.Id},
		}
	}

	toRender = append(toRender, [][]string{
		{"Name", library.Name},
		{"Description", library.Description},
		{"Tags", strings.Join(library.Tags, ", ")},
		{"Path", library.Path},
	}...)

	toRender = append(toRender, [][]string{
		{"Repository", repositoryLib.GetRepositoryUrl(library.Provider, library.RepoName)},
		{"\tName", library.RepoName},
		{"\tID", library.RepoID},
		{"\tProvider", library.Provider},
		{"\tBranch", library.Branch},
	}...)

	return toRender
}
