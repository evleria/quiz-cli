package config

import (
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const DefaultPath = "~/.quiz.yaml"

type Config struct {
	fs   afero.Fs
	path string
}

func NewConfig(fs afero.Fs, path string) *Config {
	return &Config{
		fs:   fs,
		path: ExpandPath(path),
	}
}

func (c *Config) ReadConfig() (*Categories, error) {
	file, err := c.fs.Open(c.path)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	categories := new(Categories)
	if err := yaml.Unmarshal(content, categories); err != nil {
		return nil, err
	}

	return categories, nil
}
