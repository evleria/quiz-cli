package config

type Categories struct {
	Categories []Category
}

type Category struct {
	Name       string
	Variations []Variation
}

type Variation struct {
	Name      string
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
