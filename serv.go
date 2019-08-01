package adacore

import (
	"context"
	"net"

	"go.uber.org/zap"

	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/proto"
	"google.golang.org/grpc"
)

// Serv - AdaCore Service
type Serv struct {
	lis      net.Listener
	grpcServ *grpc.Server
}

// NewAdaCoreServ -
func NewAdaCoreServ(cfg *Config) (*Serv, error) {
	lis, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		adacorebase.Error("NewAdaCoreServ", zap.Error(err))

		return nil, err
	}

	adacorebase.Info("Listen", zap.String("addr", cfg.BindAddr))

	grpcServ := grpc.NewServer()

	serv := &Serv{
		lis:      lis,
		grpcServ: grpcServ,
	}

	adacorepb.RegisterAdaCoreServiceServer(grpcServ, serv)

	return serv, nil
}

// Start - start a service
func (serv *Serv) Start(ctx context.Context) error {
	return serv.grpcServ.Serve(serv.lis)
}

// Stop - stop service
func (serv *Serv) Stop() {
	serv.lis.Close()

	return
}

// BuildWithMarkdown - build with markdown
func (serv *Serv) BuildWithMarkdown(s adacorepb.AdaCoreService_BuildWithMarkdownServer) error {
	return nil
}

// // ProcMsg implements jarviscorepb.JarvisCoreServ
// func (s *jarvisServer2) ProcMsg(in *pb.JarvisMsg, stream pb.JarvisCoreServ_ProcMsgServer) error {
// 	if IsSyncMsg(in) {
// 		chanEnd := make(chan int)

// 		s.node.PostMsg(&NormalMsgTaskInfo{
// 			Msg:         in,
// 			ReplyStream: NewJarvisMsgReplyStream(stream),
// 		}, chanEnd)

// 		<-chanEnd

// 		return nil
// 	}

// 	// if isme
// 	if in.SrcAddr == s.node.myinfo.Addr {
// 		jarvisbase.Warn("jarvisServer2.ProcMsg:isme",
// 			JSONMsg2Zap("msg", in))

// 		err := s.replyStream2ProcMsg(in.SrcAddr, in.MsgID,
// 			stream, pb.REPLYTYPE_ISME, "")
// 		if err != nil {
// 			jarvisbase.Warn("jarvisServer2.ProcMsg:replyStream2ProcMsg", zap.Error(err))

// 			return err
// 		}

// 		return nil
// 	}

// 	if in.DestAddr == s.node.myinfo.Addr {
// 		err := s.replyStream2ProcMsg(in.SrcAddr, in.MsgID,
// 			stream, pb.REPLYTYPE_IGOTIT, "")
// 		if err != nil {
// 			jarvisbase.Warn("jarvisServer2.ProcMsg:replyStream2ProcMsg:IGOTIT", zap.Error(err))

// 			return err
// 		}
// 	}

// 	// chanEnd := make(chan int)

// 	s.node.PostMsg(&NormalMsgTaskInfo{
// 		Msg:         in,
// 		ReplyStream: NewJarvisMsgReplyStream(nil),
// 	}, nil)

// 	// <-chanEnd

// 	return nil
// }

// // ProcMsgStream implements jarviscorepb.JarvisCoreServ
// func (s *jarvisServer2) ProcMsgStream(stream pb.JarvisCoreServ_ProcMsgStreamServer) error {

// 	var lstmsgs []JarvisMsgInfo
// 	var firstmsg *pb.JarvisMsg

// 	for {
// 		in, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}

// 		if firstmsg == nil && in != nil {
// 			firstmsg = in
// 		}

// 		if err != nil {
// 			jarvisbase.Warn("jarvisServer2.ProcMsgStream:stream.Recv",
// 				zap.Error(err))

// 			if firstmsg != nil {
// 				err := s.replyStream2ProcMsgStream(firstmsg.SrcAddr, firstmsg.MsgID,
// 					stream, pb.REPLYTYPE_ERROR, err.Error())

// 				if err != nil {
// 					jarvisbase.Warn("jarvisServer2.ProcMsg:replyStream2ProcMsgStream",
// 						zap.Error(err))
// 				}
// 			}

// 			lstmsgs = append(lstmsgs, JarvisMsgInfo{
// 				JarvisResultType: JarvisResultTypeLocalErrorEnd,
// 				Err:              err,
// 			})

// 			break
// 		}

// 		lstmsgs = append(lstmsgs, JarvisMsgInfo{
// 			JarvisResultType: JarvisResultTypeReply,
// 			Msg:              in,
// 		})

// 		if in.ReplyMsgID > 0 {
// 			s.node.OnReplyProcMsg(stream.Context(), in.SrcAddr, in.ReplyMsgID, JarvisResultTypeReply, in, nil)
// 		}
// 	}

// 	if firstmsg == nil {
// 		jarvisbase.Warn("jarvisServer2.ProcMsg:firstmsg")

// 		return nil
// 	}

// 	// if isme
// 	if firstmsg.SrcAddr == s.node.myinfo.Addr {
// 		jarvisbase.Warn("jarvisServer2.ProcMsg:isme",
// 			JSONMsg2Zap("msg", firstmsg))

// 		err := s.replyStream2ProcMsgStream(firstmsg.SrcAddr, firstmsg.MsgID,
// 			stream, pb.REPLYTYPE_ISME, "")
// 		if err != nil {
// 			jarvisbase.Warn("jarvisServer2.ProcMsgStream:replyStream2ProcMsgStream", zap.Error(err))

// 			return err
// 		}

// 		return nil
// 	}

// 	if firstmsg.DestAddr == s.node.myinfo.Addr {
// 		err := s.replyStream2ProcMsg(firstmsg.SrcAddr, firstmsg.MsgID,
// 			stream, pb.REPLYTYPE_IGOTIT, "")
// 		if err != nil {
// 			jarvisbase.Warn("jarvisServer2.ProcMsg:replyStream2ProcMsgStream:IGOTIT", zap.Error(err))

// 			return err
// 		}
// 	}

// 	// chanEnd := make(chan int)

// 	s.node.PostStreamMsg(&StreamMsgTaskInfo{
// 		Msgs:        lstmsgs,
// 		ReplyStream: NewJarvisMsgReplyStream(nil),
// 	}, nil)

// 	// <-chanEnd

// 	return nil
// }

// // replyStream2ProcMsgStream
// func (s *jarvisServer2) replyStream2ProcMsgStream(addr string, replyMsgID int64,
// 	stream pb.JarvisCoreServ_ProcMsgStreamServer, rt pb.REPLYTYPE, strErr string) error {

// 	sendmsg, err := BuildReply2(s.node, s.node.myinfo.Addr, addr, rt, strErr, replyMsgID)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsgStream:BuildReply2", zap.Error(err))

// 		return err
// 	}

// 	sendmsg.LastMsgID = s.node.GetCoreDB().GetCurRecvMsgID(sendmsg.DestAddr)

// 	err = SignJarvisMsg(s.node.GetCoreDB().GetPrivateKey(), sendmsg)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsgStream:SignJarvisMsg", zap.Error(err))

// 		return err
// 	}

// 	err = stream.Send(sendmsg)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsgStream:SendMsg", zap.Error(err))

// 		return err
// 	}

// 	return nil
// }

// // replyStream2ProcMsg
// func (s *jarvisServer2) replyStream2ProcMsg(addr string, replyMsgID int64,
// 	stream pb.JarvisCoreServ_ProcMsgServer, rt pb.REPLYTYPE, strErr string) error {

// 	sendmsg, err := BuildReply2(s.node, s.node.myinfo.Addr, addr, rt, strErr, replyMsgID)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsg:BuildReply2", zap.Error(err))

// 		return err
// 	}

// 	sendmsg.LastMsgID = s.node.GetCoreDB().GetCurRecvMsgID(sendmsg.DestAddr)

// 	err = SignJarvisMsg(s.node.GetCoreDB().GetPrivateKey(), sendmsg)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsg:SignJarvisMsg", zap.Error(err))

// 		return err
// 	}

// 	err = stream.Send(sendmsg)
// 	if err != nil {
// 		jarvisbase.Warn("jarvisServer2.replyStream2ProcMsg:SendMsg", zap.Error(err))

// 		return err
// 	}

// 	return nil
// }
