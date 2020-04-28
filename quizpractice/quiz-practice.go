package quizpractice

import (
	"context"
	"fmt"
	"os"

	"github.com/zawawahoge/my-memos/quizpractice/markdown"
)

// QuizPractice is interface for quiz-practice.
type QuizPractice interface {
	Run(ctx context.Context) error
}

type quizPractice struct {
	Args []string
}

// New is constructor of quiz-practice.
func New() QuizPractice {
	return &quizPractice{
		Args: os.Args,
	}
}

func (qp *quizPractice) Run(cxt context.Context) error {
	paths, err := markdown.ListMarkdowns("computer-architecture")
	if err != nil {
		return err
	}
	mds, err := markdown.LoadMarkdowns(paths)
	if err != nil {
		return err
	}
	outlinersMap := mds.ToOutlinersMap()
	for _, path := range paths {
		outliners := outlinersMap[path]
		for j, outline := range outliners {
			fmt.Println(j, outline)
		}
	}
	return nil
}
