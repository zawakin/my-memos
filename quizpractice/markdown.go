package quizpractice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Outliner represents markdown component.
type Outliner struct {
	Name     string
	Level    int
	Contents []string
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
				Name:     strings.TrimSpace(ss[2]),
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

// ListMarkdowns returns list of markdowns walking in all the sub directories of `rootDir`.
func ListMarkdowns(rootDir string) ([]string, error) {
	var paths []string
	wk := func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" {
			paths = append(paths, path)
		}
		return nil
	}
	err := filepath.Walk(rootDir, wk)
	if err != nil {
		return nil, err
	}
	return paths, nil

}

// LoadMarkdowns loads all markdowns.
func LoadMarkdowns(paths []string) ([]*Markdown, error) {
	var mds []*Markdown
	for _, path := range paths {
		md, err := LoadMarkdown(path)
		if err != nil {
			return nil, err
		}
		mds = append(mds, md)
	}
	return mds, nil
}

// LoadMarkdown loads markdown file.
func LoadMarkdown(path string) (*Markdown, error) {
	if !strings.HasSuffix(path, ".md") {
		return nil, fmt.Errorf("path should have the suffix `.md`.; path=%s", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	md, err := loadMarkdown(file, path)
	if err != nil {
		return nil, err
	}
	return md, nil
}

func loadMarkdown(reader io.Reader, path string) (*Markdown, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var contents []string

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Markdown{
		Path:     path,
		Contents: contents,
	}, nil
}
