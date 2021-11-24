package plugin

import (
	"fmt"
	"github.com/tendermint/starport/starport/pkg/xexec"
	"github.com/tendermint/starport/starport/services/chain"
)

var plugs []plugin

func Init(c *chain.Chain) ([]plugin, error) {
	if c.ConfigPath() == "" {
		return nil, fmt.Errorf("%s %s", ErrPrefix, MissingConfigFileError)
	}

	conf, err := c.Config()
	if err != nil {
		return nil, err
	}

	if len(conf.Plugins) == 0 {
		return nil, fmt.Errorf("%s %s", ErrPrefix, BadlyFormedConfigError)
	}

	if !xexec.IsCommandAvailable("git") {
		return nil, fmt.Errorf("%s: only git is supported for now, please install it", ErrPrefix)
	}

	for _, cp := range conf.Plugins {
		p, err := NewPlugin(cp.Repo, cp.Name)
		if err != nil {
			fmt.Printf("%s: %s\n", ErrPrefix, err.Error())
			continue // we can safely ignore malformed url
		}

		if p.ValidRepo() {
			plugs = append(plugs, *p)
		}

		if err != nil {
			return nil, fmt.Errorf("%s: %s", ErrPrefix, err.Error())
		}
	}

	return plugs, nil
}