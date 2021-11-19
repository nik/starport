package yaml

import (
	"context"
	"errors"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/parser"
)

var byteSlicePaths = []string{
	"$.content.content.genesisValidator.genTx",
	"$.content.content.genesisValidator.consPubKey",
}

func Parse(ctx context.Context, obj interface{}) (string, error) {
	requestYaml, err := yaml.MarshalContext(ctx, obj)
	if err != nil {
		return "", err
	}
	file, err := parser.ParseBytes(requestYaml, 0)
	if err != nil {
		return "", err
	}

	// normalize the structure converting the byte slice fields to string
	for _, path := range byteSlicePaths {
		pathString, err := yaml.PathString(path)
		if err != nil {
			return "", err
		}
		var obj []byte
		err = pathString.Read(strings.NewReader(string(requestYaml)), &obj)
		if err != nil && !errors.Is(err, yaml.ErrNotFoundNode) {
			return "", err
		}
		if err := pathString.ReplaceWithReader(file, strings.NewReader(string(obj))); err != nil {
			return "", err
		}
	}
	return file.String(), nil
}
