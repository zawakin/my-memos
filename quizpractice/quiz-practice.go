package quizpractice

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	mds, err := qp.loadMarkdowns()
	if err != nil {
		return err
	}

	for _, md := range mds {
		fmt.Println(md.Path)
		for i, line := range md.Contents {
			if i < 5 {
				fmt.Println(line)
			}
		}

	}

	return nil
}

func (qp *quizPractice) loadMarkdowns() ([]*Markdown, error) {
	var mds []*Markdown
	wk := func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" {
			md := NewMarkdown(path)
			err := md.Load()
			if err != nil {
				return err
			}
			mds = append(mds, md)
		}
		return nil
	}
	err := filepath.Walk(".", wk)
	if err != nil {
		return nil, err
	}

	return mds, nil
}

// Markdown is struct for markdown(.md) file.
type Markdown struct {
	Path     string
	Contents []string
}

// NewMarkdown is constructor of Markdown interface.
func NewMarkdown(path string) *Markdown {
	return &Markdown{
		Path: path,
	}
}

// Load sets `Markdown.Contents` to its contents by reading the file.
func (md *Markdown) Load() error {
	file, err := os.Open(md.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var contents []string

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file; path=%s; err=%#v", md.Path, err)
		return err
	}

	md.Contents = contents
	return nil
}
