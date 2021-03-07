package config

import (
	"github.com/spf13/afero"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConfig_ReadConfig(t *testing.T) {
	content :=
		`categories:
  git:
  - variation: git init
    links:
    - https://git-scm.com/docs/git-init
    questions:
    - question: "How to initialize local git repository?"
      links:
      - https://git-scm.com/docs/git-init
      answers:
      - answer: "git init"
        correct: true
      - answer: "git create"
      - answer: "git start"
      - answer: "git new"`
	fixture := &Categories{
		Categories: map[string][]Variation{
			"git": {
				{
					Variation: "git init",
					Links:     []string{"https://git-scm.com/docs/git-init"},
					Questions: []Question{
						{
							Question: "How to initialize local git repository?",
							Links:    []string{"https://git-scm.com/docs/git-init"},
							Answers: []Answer{
								{Answer: "git init", IsCorrect: true},
								{Answer: "git create"},
								{Answer: "git start"},
								{Answer: "git new"},
							},
						},
					},
				},
			},
		},
	}

	fs := afero.NewMemMapFs()
	file, _ := fs.Create(ExpandPath(DefaultPath))
	_, _ = file.WriteString(content)

	config := NewConfig(fs, DefaultPath)
	categories, err := config.ReadConfig()

	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual(categories, fixture))
}

func TestConfig_ReadConfig_FileNotFound(t *testing.T) {
	mapFs := afero.NewMemMapFs()

	config := NewConfig(mapFs, DefaultPath)
	_, err := config.ReadConfig()

	assert.ErrorContains(t, err, "file does not exist")
}

func TestConfig_ReadConfig_InvalidFormat(t *testing.T) {
	fs := afero.NewMemMapFs()
	file, _ := fs.Create(ExpandPath(DefaultPath))
	_, _ = file.WriteString("ERROR")

	config := NewConfig(fs, DefaultPath)
	_, err := config.ReadConfig()

	assert.ErrorContains(t, err, "yaml: unmarshal errors")
}
