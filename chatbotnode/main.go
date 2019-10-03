package main

import (
	"context"
	"fmt"
	
	chatbotada "github.com/zhs007/adacore/chatbot"
)

func main() {
	err := chatbotada.StartChatBot(context.Background(), "./cfg/adanode.yaml", "./cfg/chatbot.yaml")
	if err != nil {
		fmt.Printf("StartChatBot %v", err)

		return
	}
}
