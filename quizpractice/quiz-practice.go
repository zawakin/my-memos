package quizpractice

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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

	var s string

	for {
		ClearScreen()
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
		fmt.Printf("Path = %s\n", path)

		outliners := outlinersMap[path]
		index = 0
		ClearScreen()
		for {
			outline := outliners[index]
			PrintOutliner(outline)

			fmt.Printf("===\n")
			for i, outline := range outliners {
				if i%5 == 0 {
					fmt.Println()
				}
				fmt.Printf("(%d) %s ", i, outline.Name)
			}

			fmt.Print("Please choose outline you want to read (-1 if you quit): ")
			_, err = fmt.Scanf("%s", &s)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			if s == "f" {
				index++
				ClearScreen()
			} else if s == "b" {
				index--
				ClearScreen()
			} else {
				i, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println("Invalid input")
				}
				if i < -1 || i >= len(outliners) {
					fmt.Println("Invalid input")
					continue
				}
				if i == -1 {
					break
				}
				index = i
				ClearScreen()
			}
		}
	}
}

// ClearScreen crears screen.
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// PrintOutliner describes outliner.
func PrintOutliner(outline *markdown.Outliner) {
	fmt.Printf("\n%s %s\n\n===\n", strings.Repeat("#", outline.Level), outline.Name)
	for _, content := range outline.Contents {
		fmt.Println(content)
	}

	fmt.Println()

	re := regexp.MustCompile(`^(.+)は、(.+)です`)
	for _, content := range outline.Contents {
		for _, s := range strings.Split(content, "。") {
			ss := re.FindStringSubmatch(s)
			if len(ss) > 0 {
				fmt.Println(ss[1])
				fmt.Println(ss[2])
			}
		}
	}

}
