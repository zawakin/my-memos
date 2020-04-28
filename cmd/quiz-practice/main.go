package main

import (
	"context"
	"log"
	"os"

	"github.com/zawawahoge/my-memos/quizpractice"
)

func main() {
	ctx := context.Background()

	cli := quizpractice.New()
	if err := cli.Run(ctx); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
