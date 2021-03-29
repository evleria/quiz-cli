package config

var content = `categories:
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

var fixture = &Categories{
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
