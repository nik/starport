package plugin

import (
	"github.com/tendermint/starport/starport/pkg/giturl"
	"golang.org/x/tools/go/vcs"
)

type plugin struct {
	Name    string
	RepoUrl string
	Cmd 	*vcs.Cmd
}

func NewPlugin(url, name string) (*plugin, error) {
	cmd, err := detectVCS(url)
	if err != nil {
		return nil, err
	}

	return &plugin{
		Name: name,
		RepoUrl: url,
		Cmd: cmd,
	}, nil
}

func (p *plugin) ValidRepo() bool {
	_, err := giturl.Parse(p.RepoUrl)

	if err != nil {
		return false
	}

	return true
}

func detectVCS(url string) (*vcs.Cmd, error) {
	rr, err := vcs.RepoRootForImportPath(url, false)
	if err != nil {
		return nil, err
	}
	return rr.VCS, nil
}