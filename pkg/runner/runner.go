package runner

import (
	"github.com/evleria/quiz-cli/pkg/config"
	"math/rand"
)

type Runner struct {
	questions []config.Question
	history []answer

	current int
}

type QuizResult struct {
	Correct int
	Total int
}

type answer struct {
	Indices []int
	Correct bool
}

func NewRunner(questions []config.Question) *Runner {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	return &Runner{
		questions: questions,
		history: make([]answer, len(questions)),
		current: -1,
	}
}

func (r *Runner) Next() bool {
	r.current++

	return r.current < len(r.questions)
}

func (r *Runner) GenerateQuestion() config.Question {
	question := r.questions[r.current]

	rand.Shuffle(len(question.Answers), func(i, j int) {
		question.Answers[i], question.Answers[j] = question.Answers[j], question.Answers[i]
	})

	return question
}

func (r *Runner) MarkAnswer(answered []int) {
	r.history[r.current] = answer{
		Indices: answered,
		Correct: isCorrect(r.questions[r.current], answered),
	}
}

func (r *Runner) Result() QuizResult {
	correct := 0
	for _, h := range r.history {
		if h.Correct {
			correct++
		}
	}

	return QuizResult{
		Correct: correct,
		Total: len(r.questions),
	}
}

func isCorrect(question config.Question, answered []int) bool {
	m := map[int]bool{}
	for i, a := range question.Answers {
		if a.IsCorrect {
			m[i] = true
		}
	}
	for _, n := range answered {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			return false
		}
	}
	return len(m) == 0
}
