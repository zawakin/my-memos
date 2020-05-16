package quizpractice

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gomarkdown/markdown/ast"
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
	paths, err := ListMarkdowns("computer-architecture")
	if err != nil {
		return err
	}

	var writer io.Writer
	writer = os.Stdout

	nodes, err := LoadMarkdowns(paths)
	for _, node := range nodes {
		fmt.Println(node)
		ast.Print(writer, node)
	}

	return nil
}
