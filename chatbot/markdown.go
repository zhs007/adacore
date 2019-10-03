package chatbotada

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	adacore "github.com/zhs007/adacore"
	adarender "github.com/zhs007/adacore/adarenderpb"
	chatbot "github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// markdownFP - file processor markdown
type markdownFP struct {
	serv *adacore.Serv
}

// Proc - process
func (fp *markdownFP) Proc(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if chat.File != nil && chat.File.FileData != nil {
		rendermd := &adarender.MarkdownData{
			StrData:      string(chat.File.FileData),
			TemplateName: fp.serv.Cfg.Templates[0],
		}

		htmldata, err := fp.serv.ClientRender.Render(ctx, rendermd)
		if err != nil {
			return nil, err
		}

		hashname, err := adacore.SaveHTMLData(htmldata, fp.serv.Cfg)
		if err != nil {
			return nil, err
		}

		var lst []*chatbotpb.ChatMsg

		msghashname := chatbot.BuildTextChatMsg(hashname, chat.Uai,
			chat.Token, chat.SessionID)

		lst = append(lst, msghashname)

		return lst, nil
	}

	return nil, nil
}

// IsMyFile - is my file
func (fp *markdownFP) IsMyFile(chat *chatbotpb.ChatMsg) bool {

	if chat.File != nil {
		arr := strings.Split(chat.File.Filename, ".")
		if len(arr) > 1 && arr[len(arr)-1] == "md" {
			return true
		}
	}

	return false
}
