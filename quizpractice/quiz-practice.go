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
	fmt.Println(paths)
	mds, err := markdown.LoadMarkdowns(paths)
	if err != nil {
		return err
	}

	for _, md := range mds {
		fmt.Println(md.Path)
		outliners := md.ToOutliners()
		for _, o := range outliners {
			fmt.Println(o.Level, o.Name, len(o.Contents))
			for i, content := range o.Contents {
				if i > 0 {
					continue
				}
				fmt.Println(i, content)
			}
		}
	}

	return nil
}
