package chatbotada

import (
	"context"

	adacore "github.com/zhs007/adacore"
	adacorebase "github.com/zhs007/adacore/base"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	basicchatbot "github.com/zhs007/chatbot/basicchatbot"
	chatbotusermgr "github.com/zhs007/chatbot/usermgr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func startAdaServ(ctx context.Context, cfg *adacore.Config, endchan chan int) (
	*adacore.Serv, error) {

	serv, err := adacore.NewAdaCoreServ(cfg)
	if err != nil {
		adacorebase.Warn("startAdaServ:NewAdaCoreServ",
			zap.Error(err))

		return nil, err
	}

	go func() {
		err := serv.Start(context.Background())
		if err != nil {
			adacorebase.Warn("startAdaServ:Start",
				zap.Error(err))
		}

		endchan <- 0
	}()

	return serv, nil
}

func startChatBot(ctx context.Context, servAda *adacore.Serv, adacorecfg *adacore.Config,
	chatbotcfg *chatbot.Config, endchan chan int) error {

	chatbotbase.InitLogger(zapcore.InfoLevel, true, adacorecfg.Log.LogPath)

	mgr, err := chatbotusermgr.NewUserMgr(chatbotcfg.DBPath,
		"", chatbotcfg.DBEngine, nil)
	if err != nil {
		adacorebase.Warn("startChatBot:NewUserMgr",
			zap.Error(err))

		return err
	}

	serv, err := chatbot.NewChatBotServ(chatbotcfg, mgr, &ServiceCore{})
	if err != nil {
		adacorebase.Warn("startChatBot:NewChatBotServ",
			zap.Error(err))

		return err
	}

	serv.MgrFile.RegisterFileProcessor(&markdownFP{
		serv: servAda,
	})

	serv.Init(context.Background())

	go func() {
		serv.Start(context.Background())
		if err != nil {
			adacorebase.Warn("startChatBot:Start",
				zap.Error(err))
		}

		endchan <- 0
	}()

	return nil
}

// StartChatBot - start chatbot
func StartChatBot(ctx context.Context, adacorecfgfn string, chatbotcfgfn string) error {
	adacorecfg, err := adacore.LoadConfig(adacorecfgfn)
	if err != nil {
		return err
	}

	adacore.InitLogger(adacorecfg)

	err = basicchatbot.InitBasicChatBot()
	if err != nil {
		adacorebase.Warn("StartChatBot:InitBasicChatBot",
			zap.Error(err))

		return err
	}

	chatbotcfg, err := chatbot.LoadConfig(chatbotcfgfn)
	if err != nil {
		adacorebase.Warn("StartChatBot:LoadConfig",
			zap.Error(err))

		return err
	}

	servchan := make(chan int)

	adaServ, err := startAdaServ(context.Background(), adacorecfg, servchan)
	if err != nil {
		adacorebase.Warn("StartChatBot:startAdaServ",
			zap.Error(err))

		return err
	}

	err = startChatBot(context.Background(), adaServ, adacorecfg, chatbotcfg, servchan)
	if err != nil {
		adacorebase.Warn("StartChatBot:startChatBot",
			zap.Error(err))

		return err
	}

	endi := 0
	for {
		<-servchan
		endi++
		if endi >= 2 {
			break
		}
	}

	return nil
}
