package chatbotada

import (
	"context"

	"github.com/golang/protobuf/proto"
	adacorepb "github.com/zhs007/adacore/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

// DebugExcelColumnType - excel column type
type DebugExcelColumnType struct {
	Name      string
	Type      string
	TypeAuto  string
	Separator string
}

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
		ed, err := ProcExcelMsg(chat)
		if err != nil {
			chatbotbase.Warn("chatbotada.ServiceCore.OnDebug:ProcExcelMsg",
				zap.Error(err))

			return nil, err
		}

		lang := serv.GetChatMsgLang(chat)

		locale, err := serv.MgrText.GetLocalizer(lang)
		if err != nil {
			return nil, err
		}

		mParams, err := serv.BuildBasicParamsMap(chat, ui, lang)
		if err != nil {
			return nil, err
		}

		var lst []*chatbotpb.ChatMsg

		var lstallcts []DebugExcelColumnType

		for i, v := range ed.Columns {
			lstallcts = append(lstallcts, DebugExcelColumnType{
				Name:      v.Name,
				Type:      ExcelColumnType2String(v.Type),
				Separator: v.Separator,
				TypeAuto:  ExcelColumnType2String(ed.ColumnsAuto[i]),
			})
		}

		mParams["Columns"] = lstallcts
		mParams["CurCheet"] = ed.CurSheetName

		msgdebugexcel, err := chatbot.NewChatMsgWithText(locale, "debugexcel", mParams, chat.Uai)
		if err != nil {
			return nil, err
		}

		lst = append(lst, msgdebugexcel)

		return lst, nil
	}

	return nil, nil
}
