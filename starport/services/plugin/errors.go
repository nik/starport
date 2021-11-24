package plugin

import (
	"errors"
)

const (
	ErrPrefix = "❌️ oops!"
)

var (
	MissingConfigFileError = errors.New("couldn't find a config.yml")
	BadlyFormedConfigError = errors.New("no plugins detected in config.yml")
)
