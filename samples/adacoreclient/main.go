package main

import (
	"context"
	"fmt"

	"github.com/zhs007/adacore"
)

func startServ(cfg *adacore.Config, endchan chan int) (*adacore.Serv, error) {

	serv, err := adacore.NewAdaCoreServ(cfg)
	if err != nil {
		fmt.Printf("startServ NewAdaCoreServ %v", err)

		return nil, err
	}

	go func() {
		err := serv.Start(context.Background())
		if err != nil {
			fmt.Printf("startServ Start %v", err)
		}

		endchan <- 0
	}()

	return serv, nil
}

func startClient(cfg *adacore.Config) error {
	client := adacore.NewClient("47.91.209.141:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	_, err := client.BuildWithMarkdownFile(context.Background(), "../../unittest/sample001.md", "default")
	if err != nil {
		fmt.Printf("startClient BuildWithMarkdownFile %v", err)

		return err
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

	// servchan := make(chan int)
	// serv, err := startServ(cfg, servchan)
	// if err != nil {
	// 	return
	// }

	// time.Sleep(1000 * 3)

	startClient(cfg)

	// serv.Stop()

	// <-servchan

	// serv.Stop()
}
