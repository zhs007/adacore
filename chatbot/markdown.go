package chatbotada

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	adacore "github.com/zhs007/adacore"
	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorebase "github.com/zhs007/adacore/base"
	chatbot "github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// markdownFP - file processor for markdown
type markdownFP struct {
	servAda *adacore.Serv
}

// Proc - process
func (fp *markdownFP) Proc(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if chat.File != nil && chat.File.FileData != nil {
		rendermd := &adarender.MarkdownData{
			StrData:      string(chat.File.FileData),
			TemplateName: fp.servAda.Cfg.Templates[0],
		}

		htmldata, err := fp.servAda.ClientRender.Render(ctx, rendermd)
		if err != nil {
			return nil, err
		}

		hashname, err := adacore.SaveHTMLData(htmldata, fp.servAda.Cfg)
		if err != nil {
			return nil, err
		}

		lang := serv.GetChatMsgLang(chat)

		locale, err := serv.MgrText.GetLocalizer(lang)
		if err != nil {
			return nil, err
		}

		var lst []*chatbotpb.ChatMsg

		msghashname, err := chatbot.NewChatMsgWithText(locale, "iprocok", map[string]interface{}{
			"Url": adacorebase.AppendString(fp.servAda.Cfg.BaseURL, hashname),
		}, chat.Uai)
		if err != nil {
			return nil, err
		}

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
