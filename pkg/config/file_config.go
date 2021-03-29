package config

import (
	"github.com/spf13/afero"
)

const DefaultPath = "~/.quiz.yaml"

type FileConfig struct {
	fs   afero.Fs
	path string
}

func NewFileConfig(fs afero.Fs, path string) *FileConfig {
	return &FileConfig{
		fs:   fs,
		path: ExpandPath(path),
	}
}

func (c *FileConfig) ReadConfig() (*Categories, error) {
	file, err := c.fs.Open(c.path)
	if err != nil {
		return nil, err
	}

	return readAll(file)
}
