package adacore

import (
	"context"
	"testing"
	"time"
)

func startServ(cfg *Config, endchan chan int) (*Serv, error) {

	serv, err := NewAdaCoreServ(cfg)
	if err != nil {
		// fmt.Printf("startServ NewAdaCoreServ %v", err)

		return nil, err
	}

	go func() {
		err := serv.Start(context.Background())
		if err != nil {
			// fmt.Printf("startServ Start %v", err)
		}

		endchan <- 0
	}()

	return serv, nil
}

func startClient(cfg *Config) error {
	client := NewClient("127.0.0.1:7201", "x7sSGGHgmKwUMoa5S4VZlr9bUF2lCCzF")

	_, err := client.BuildWithMarkdownFile(context.Background(), "./unittest/sample001.md", "default")
	if err != nil {
		// fmt.Printf("startClient BuildWithMarkdownFile %v", err)

		return err
	}

	return nil
}

func TestServ(t *testing.T) {
	cfg, err := LoadConfig("./unittest/servtestcfrg.yaml")
	if err != nil {
		t.Fatalf("TestServ LoadConfig %v err is %v", "./unittest/servtestcfrg.yaml", err)

		return
	}

	InitLogger(cfg)

	servchan := make(chan int)
	serv, err := startServ(cfg, servchan)
	if err != nil {
		t.Fatalf("TestServ startServ err is %v", err)
	}

	time.Sleep(1000 * 3)

	err = startClient(cfg)
	if err != nil {
		t.Fatalf("TestServ startClient err is %v", err)
	}

	serv.Stop()

	<-servchan

	// serv.Stop()
}
