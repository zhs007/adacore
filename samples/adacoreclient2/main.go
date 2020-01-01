package main

import (
	"context"
	"fmt"

	"github.com/zhs007/adacore"
)

func startClient(cfg *adacore.Config) error {
	client := adacore.NewClient("47.91.209.141:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	md, err := adacore.LoadMarkdownAndFiles("./report/index.md", "./report/*.svg")
	if err != nil {
		fmt.Printf("startClient genMarkdown %v", err)

		return err
	}

	reply, err := client.BuildWithMarkdown(context.Background(), md)
	if err != nil {
		fmt.Printf("startClient BuildWithMarkdownFile %v", err)

		return err
	}

	if reply != nil {
		// fmt.Print(reply.HashName)
		fmt.Print(reply.Url)
	}

	return nil
}

func main() {
	cfg, err := adacore.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("startServ LoadConfig %v", err)

		return
	}

	adacore.InitLogger(cfg)

	startClient(cfg)
}
