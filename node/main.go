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

func main() {
	cfg, err := adacore.LoadConfig("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("startServ LoadConfig %v", err)

		return
	}

	adacore.InitLogger(cfg)

	servchan := make(chan int)
	_, err = startServ(cfg, servchan)
	if err != nil {
		return
	}

	<-servchan
}
