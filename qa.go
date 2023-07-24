package goarc

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"io"
)

type Selection int

const (
	Single Selection = iota
	Multiple
)

type qas []*QA

func (q *qas) LoadQAs() {
	var questions []Question
	err := viper.UnmarshalKey("questions", &questions)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal questions: %v", err))
	}
	for _, question := range questions {
		*q = append(*q, &QA{Question: question, Answer: []string{}})
	}
}

var QAs = &qas{}

type Question struct {
	Selection `yaml:"selection"`
	Component string   `yaml:"component"`
	ID        int      `yaml:"id"`
	Prompt    string   `yaml:"prompt"`
	Choices   []string `yaml:"choices"`
}

type QA struct {
	Question
	outerWriter io.Writer
	Answer      []string
}

// Ask ask the question, use survey to get the answer
func (q *QA) Ask() error {
	var prompt survey.Prompt

	switch q.Selection {
	case Single:
		tmpAnswer := ""
		prompt = &survey.Select{
			Message: q.Prompt,
			Options: q.Choices,
		}
		if err := survey.AskOne(prompt, &tmpAnswer); err != nil {
			return err
		}
		q.Answer = append(q.Answer, tmpAnswer)
		return nil

	case Multiple:
		prompt = &survey.MultiSelect{
			Message: q.Prompt,
			Options: q.Choices,
		}
		return survey.AskOne(prompt, &q.Answer)
	default:
		return fmt.Errorf("invalid selection type")
	}

}
