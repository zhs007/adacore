package chatbotada

import (
	"github.com/golang/protobuf/proto"
	adacorepb "github.com/zhs007/adacore/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
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
