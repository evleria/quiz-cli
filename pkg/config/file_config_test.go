package config

import (
	"github.com/spf13/afero"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFileConfig_ReadConfig(t *testing.T) {
	fs := afero.NewMemMapFs()
	file, _ := fs.Create(ExpandPath(DefaultPath))
	_, _ = file.WriteString(content)

	config := NewFileConfig(fs, DefaultPath)
	categories, err := config.ReadConfig()

	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual(categories, fixture))
}

func TestFileConfig_ReadConfig_FileNotFound(t *testing.T) {
	mapFs := afero.NewMemMapFs()

	config := NewFileConfig(mapFs, DefaultPath)
	_, err := config.ReadConfig()

	assert.ErrorContains(t, err, "file does not exist")
}

func TestFileConfig_ReadConfig_InvalidFormat(t *testing.T) {
	fs := afero.NewMemMapFs()
	file, _ := fs.Create(ExpandPath(DefaultPath))
	_, _ = file.WriteString("ERROR")

	config := NewFileConfig(fs, DefaultPath)
	_, err := config.ReadConfig()

	assert.ErrorContains(t, err, "yaml: unmarshal errors")
}
