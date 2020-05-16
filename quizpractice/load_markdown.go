package quizpractice

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	ast "github.com/gomarkdown/markdown/ast"
)

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
func LoadMarkdowns(paths []string) ([]ast.Node, error) {
	var mds []ast.Node
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
func LoadMarkdown(path string) (ast.Node, error) {
	if !strings.HasSuffix(string(path), ".md") {
		return nil, fmt.Errorf("path should have the suffix `.md`.; path=%s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	node := markdown.Parse(data, nil)
	return node, nil
}
