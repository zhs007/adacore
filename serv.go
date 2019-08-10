package adacore

import (
	"context"
	"io"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/zhs007/adacore/adarenderclient"
	adarender "github.com/zhs007/adacore/adarenderpb"
	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/proto"
)

// Serv - AdaCore Service
type Serv struct {
	cfg          *Config
	lis          net.Listener
	grpcServ     *grpc.Server
	renderClient *adarenderclient.Client
}

// NewAdaCoreServ -
func NewAdaCoreServ(cfg *Config) (*Serv, error) {
	if cfg == nil {
		return nil, adacorebase.ErrServNoConfig
	}

	lis, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		adacorebase.Error("NewAdaCoreServ", zap.Error(err))

		return nil, err
	}

	adacorebase.Info("Listen", zap.String("addr", cfg.BindAddr))

	renderClient := adarenderclient.NewClient(cfg.AdaRenderServAddr, cfg.AdaRenderToken)

	grpcServ := grpc.NewServer()

	serv := &Serv{
		cfg:          cfg,
		lis:          lis,
		grpcServ:     grpcServ,
		renderClient: renderClient,
	}

	adacorepb.RegisterAdaCoreServiceServer(grpcServ, serv)

	adacorebase.Info("NewAdaCoreServ OK.")

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
func (serv *Serv) BuildWithMarkdown(stream adacorepb.AdaCoreService_BuildWithMarkdownServer) error {
	var lst []*adacorepb.MarkdownStream

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			serv.replyErr(stream, err)

			return err
		}

		if !serv.cfg.HasToken(in.Token) {
			return adacorebase.ErrServInvalidToken
		}

		lst = append(lst, in)
	}

	md, err := BuildMarkdownData(lst)
	if err != nil {
		serv.replyErr(stream, err)

		return err
	}

	rendermd := &adarender.MarkdownData{
		StrData:      md.StrData,
		BinaryData:   md.BinaryData,
		TemplateName: md.TemplateName,
		TemplateData: md.TemplateData,
	}

	if len(md.BinaryData) > 0 {
		rendermd.BinaryData = make(map[string][]byte)

		for k, v := range md.BinaryData {
			rendermd.BinaryData[k] = v
		}
	}

	htmldata, err := serv.renderClient.Render(stream.Context(), rendermd)
	if err != nil {
		serv.replyErr(stream, err)

		return err
	}

	hashname, err := SaveHTMLData(htmldata, serv.cfg)
	if err != nil {
		serv.replyErr(stream, err)

		return err
	}

	serv.replyResult(stream, hashname)

	return nil
}

// replyErr - reply a error
func (serv *Serv) replyErr(stream adacorepb.AdaCoreService_BuildWithMarkdownServer, err error) error {
	if err == nil {
		return serv.replyErr(stream, adacorebase.ErrServInvalidErrString)
	}

	reply := &adacorepb.ReplyMarkdown{
		Err: err.Error(),
	}

	return stream.SendAndClose(reply)
}

// replyResult - reply result
func (serv *Serv) replyResult(stream adacorepb.AdaCoreService_BuildWithMarkdownServer, hashname string) error {
	if hashname == "" {
		return serv.replyErr(stream, adacorebase.ErrServInvalidResult)
	}

	reply := &adacorepb.ReplyMarkdown{
		HashName: hashname,
		Url:      adacorebase.AppendString(serv.cfg.BaseURL, hashname),
	}

	return stream.SendAndClose(reply)
}
