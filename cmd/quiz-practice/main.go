package main

import (
	"context"
	"os"

	"github.com/zawawahoge/my-memos/quizpractice"
)

func main() {
	ctx := context.Background()

	cli := quizpractice.New()
	if err := cli.Run(ctx); err != nil {
		os.Exit(1)
	}
}
