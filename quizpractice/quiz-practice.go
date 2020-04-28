package quizpractice

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	mds, err := qp.loadMarkdowns("computer-architecture")
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

// LoadMarkdowns loads all markdowns walking in all the sub directories of `path`.
func LoadMarkdowns(path string) ([]*Markdown, error) {
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
	err := filepath.Walk(path, wk)
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

// Outliner represents markdown component.
type Outliner struct {
	Name     string
	Level    int
	Contents []string
}

// ToOutliners creates outliners by parsing the contents.
func (md *Markdown) ToOutliners() []*Outliner {
	isOutline := regexp.MustCompile(`^(#+) +(.*)$`)

	outliners := make([]*Outliner, 0)
	var outliner *Outliner
	for _, line := range md.Contents {
		if line == "" {
			continue
		}
		ss := isOutline.FindStringSubmatch(line)
		if len(ss) == 3 {
			// Outline
			outliner = &Outliner{
				Name:     ss[2],
				Level:    len(ss[1]),
				Contents: []string{},
			}
			outliners = append(outliners, outliner)
			continue
		}
		if outliner == nil {
			// Skip if outliner hasn't been created yet.
			continue
		}
		outliner.Contents = append(outliner.Contents, line)
	}
	return outliners
}

// Node is struct of node.
type Node struct {
	Depth    int
	Name     string
	Contents []string
	Children []*Node
	Parent   *Node
}

// NewNode is constructor of node.
func NewNode(depth int, name string, parent *Node) *Node {
	return &Node{
		Depth:  depth,
		Name:   name,
		Parent: parent,
	}
}
