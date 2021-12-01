package plugin

import (
	"github.com/tendermint/starport/starport/pkg/giturl"
	"golang.org/x/tools/go/vcs"
)

type plugin struct {
	Name    string
	RepoUrl string
	InstalledLocation string
	Cmd 	*vcs.Cmd
}

/*
NewPlugin returns a struct with a well-formed URL <e.g. https://github.com/place/repo>
as well as access to a built-in go VCS cmd helper which exposes VCS logic (clone, pull, etc)
 */
func NewPlugin(url, name string) (*plugin, error) {
	rr, err := parseVCS(url)
	if err != nil {
		return nil, err
	}

	return &plugin{
		Name: name,
		RepoUrl: rr.Repo,
		Cmd: rr.VCS,
	}, nil
}

func (p *plugin) ValidRepo() bool {
	_, err := giturl.Parse(p.RepoUrl)

	if err != nil {
		return false
	}

	return true
}

func parseVCS(url string) (*vcs.RepoRoot, error) {
	rr, err := vcs.RepoRootForImportPath(url, false)
	if err != nil {
		return nil, err
	}
	return rr, nil
}