package quizpractice

import (
	"context"
	"fmt"
	"os"
	"strings"

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

	for {
		for i, path := range paths {
			fmt.Printf("%d) %s\n", i, path)
		}
		fmt.Print("Please choose path you want to look into: ")
		var index int
		_, err := fmt.Scanf("%d", &index)
		if err != nil || index < 0 || index >= len(paths) {
			fmt.Println("Invalid input")
			continue
		}
		path := paths[index]
		outliners := outlinersMap[path]
		for i, outline := range outliners {
			fmt.Printf("%d) Header %d %s\n", i, outline.Level, outline.Name)
		}

		fmt.Print("Please choose outline you want to read: ")
		_, err = fmt.Scanf("%d", &index)
		if err != nil || index < 0 || index >= len(outliners) {
			fmt.Println("Invalid input")
			continue
		}
		outline := outliners[index]
		fmt.Printf("Header %d %s\n", outline.Level, outline.Name)
		for _, content := range outline.Contents {
			fmt.Println(content)
		}

		var s string
		fmt.Printf("Do you continue? (y/n): ")
		fmt.Scanf("%s", &s)
		if sl := strings.ToLower(s); sl == "n" || sl == "no" {
			break
		}
	}
	return nil
}
