package config

import (
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Config interface {
	ReadConfig() (*Categories, error)
}

func CreateConfigFromPath(path string) Config {
	if strings.HasPrefix(path, "http") {
		return NewRemoteConfig(new(http.Client), path)
	}

	return NewFileConfig(afero.NewOsFs(), path)
}

func readAll(reader io.Reader) (*Categories, error) {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	categories := new(Categories)
	if err := yaml.Unmarshal(content, categories); err != nil {
		return nil, err
	}

	return categories, nil
}
