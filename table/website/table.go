package websiteTable

import (
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	websiteLib "github.com/taubyte/tau/lib/website"
)

func getTableData(website *structureSpec.Website, showId bool) (toRender [][]string) {
	if showId == true {
		toRender = [][]string{
			{"ID", website.Id},
		}
	}

	toRender = append(toRender, [][]string{
		{"Name", website.Name},
		{"Description", website.Description},
		{"Tags", strings.Join(website.Tags, ", ")},
		{"Paths", strings.Join(website.Paths, ", ")},
		{"Domains", strings.Join(website.Domains, ", ")},
	}...)

	toRender = append(toRender, [][]string{
		{"Repository", websiteLib.GetRepositoryUrl(website)},
		{"\tName", website.RepoName},
		{"\tID", website.RepoID},
		{"\tProvider", website.Provider},
		{"\tBranch", website.Branch},
	}...)

	return toRender
}
