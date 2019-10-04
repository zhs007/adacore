package chatbotada

import (
	"bytes"
	"context"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/golang/protobuf/proto"
	adacorepb "github.com/zhs007/adacore/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

// ServiceCore - chatbot service core
type ServiceCore struct {
}

// UnmarshalAppData - unmarshal
func (core *ServiceCore) UnmarshalAppData(buf []byte) (proto.Message, error) {
	ad := &adacorepb.ChatBotData{}

	err := proto.Unmarshal(buf, ad)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

// NewAppData - new a app data
func (core *ServiceCore) NewAppData() (proto.Message, error) {
	return &adacorepb.ChatBotData{}, nil
}

// UnmarshalUserData - unmarshal
func (core *ServiceCore) UnmarshalUserData(buf []byte) (proto.Message, error) {
	ud := &adacorepb.UserData{}

	err := proto.Unmarshal(buf, ud)
	if err != nil {
		return nil, err
	}

	return ud, nil
}

// NewUserData - new a userdata
func (core *ServiceCore) NewUserData(ui *chatbotpb.UserInfo) (proto.Message, error) {
	return &adacorepb.UserData{}, nil
}

// OnDebug - call in plugin.debug
func (core *ServiceCore) OnDebug(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if isExcelFile(chat) {
		r := bytes.NewReader(chat.File.FileData)
		f, err := excelize.OpenReader(r)
		if err != nil {
			chatbotbase.Warn("chatbotada.ServiceCore.OnDebug:OpenReader",
				zap.Error(err))

			return nil, err
		}

		mapSheet := f.GetSheetMap()
		curSheet := ""

		for _, v := range mapSheet {
			curSheet = v

			break
		}

		arr, err := f.GetRows(curSheet)
		if err != nil {
			chatbotbase.Warn("chatbotada.ServiceCore.OnDebug:GetRows",
				zap.Error(err))

			return nil, err
		}

		arr = ProcHead(arr)

		var lst []*chatbotpb.ChatMsg

		lstct := AnalysisColumnsType(arr)
		mapct := map[string]string{}

		for i, v := range arr[0] {
			mapct[v] = ExcelColumnType2String(lstct[i])
		}

		strct, err := chatbotbase.JSONFormat(mapct)
		if err != nil {
			chatbotbase.Warn("chatbotada.ServiceCore.OnDebug:JSONFormat",
				zap.Error(err))

			return nil, err
		}

		msgct := &chatbotpb.ChatMsg{
			Msg: strct,
			Uai: chat.Uai,
		}

		lst = append(lst, msgct)

		return lst, nil
	}

	return nil, nil
}
