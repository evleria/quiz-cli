package config

type Categories struct {
	Categories map[string][]Variation
}

type Variation struct {
	Variation string
	Questions []Question
}

type Question struct {
	Question string
	Answers  []Answer
}

type Answer struct {
	Answer    string
	IsCorrect bool `yaml:"correct"`
}
