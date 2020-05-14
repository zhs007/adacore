package adacore

import (
	"context"
	"io"
	"io/ioutil"

	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/adacorepb"
	"google.golang.org/grpc"
)

// Client - AdaRenderServClient
type Client struct {
	servAddr string
	token    string
	conn     *grpc.ClientConn
	client   adacorepb.AdaCoreServiceClient
}

// NewClient - new AdaRenderClient
func NewClient(servAddr string, token string) *Client {
	return &Client{
		servAddr: servAddr,
		token:    token,
	}
}

// isValid - check myself
func (client *Client) isValid() error {
	if client.servAddr == "" {
		return adacorebase.ErrAdaCoreClientNoServAddr
	}

	if client.token == "" {
		return adacorebase.ErrAdaCoreClientNoToken
	}

	return nil
}

// reset - reset
func (client *Client) reset() {
	if client.conn != nil {
		client.conn.Close()
	}

	client.conn = nil
	client.client = nil
}

// BuildWithMarkdown - MarkdownData => ReplyMarkdown
func (client *Client) BuildWithMarkdown(ctx context.Context, mddata *adacorepb.MarkdownData) (
	*adacorepb.ReplyMarkdown, error) {

	err := client.isValid()
	if err != nil {
		return nil, err
	}

	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}

		client.conn = conn
		client.client = adacorepb.NewAdaCoreServiceClient(conn)
	}

	stream, err := client.client.BuildWithMarkdown(ctx)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	lst, err := BuildMarkdownStream(mddata, client.token)
	if err != nil {
		// if error, close
		stream.CloseSend()

		// if error, reset
		client.reset()

		return nil, err
	}

	for _, cn := range lst {
		curerr := stream.Send(cn)
		if curerr != nil {
			// if error, close
			stream.CloseSend()

			// if error, reset
			client.reset()

			return nil, curerr
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil && err != io.EOF {

		// if error, reset
		client.reset()

		return nil, err
	}

	return reply, nil
}

// BuildWithMarkdownFile - markdown file => ReplyMarkdown
func (client *Client) BuildWithMarkdownFile(ctx context.Context, fn string, tempname string) (*adacorepb.ReplyMarkdown, error) {
	fd, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	md := &adacorepb.MarkdownData{
		StrData:      string(fd),
		TemplateName: tempname,
	}

	return client.BuildWithMarkdown(ctx, md)
}
