package chatbotada

import (
	"bytes"
	"context"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/golang/protobuf/proto"
	adacore "github.com/zhs007/adacore"
	chatbot "github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// excelFP - file processor for excel
type excelFP struct {
	servAda *adacore.Serv
}

// Proc - process
func (fp *excelFP) Proc(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if chat.File != nil && chat.File.FileData != nil {
		r := bytes.NewReader(chat.File.FileData)
		f, err := excelize.OpenReader(r)
		if err != nil {
			return nil, err
		}

		cs := f.GetActiveSheetIndex()
		curSheet := f.GetSheetName(cs)

		arr, err := f.GetRows(curSheet)
		if err != nil {
			return nil, err
		}

		arr = ProcHead(arr)

		var lst []*chatbotpb.ChatMsg

		// msghashname, err := chatbot.NewChatMsgWithText(locale, "iprocok", map[string]interface{}{
		// 	"Url": adacorebase.AppendString(fp.serv.Cfg.BaseURL, hashname),
		// }, chat.Uai)
		// if err != nil {
		// 	return nil, err
		// }

		// lst = append(lst, msghashname)

		return lst, nil
	}

	return nil, nil
}

// IsMyFile - is my file
func (fp *excelFP) IsMyFile(chat *chatbotpb.ChatMsg) bool {
	return isExcelFile(chat)
}

func isExcelFile(chat *chatbotpb.ChatMsg) bool {
	if chat.File != nil && chat.File.FileData != nil {

		arr := strings.Split(chat.File.Filename, ".")
		if len(arr) > 1 &&
			(arr[len(arr)-1] == "xls" || arr[len(arr)-1] == "xlsx") {

			return true
		}
	}

	return false
}
