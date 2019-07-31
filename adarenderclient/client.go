package adarenderclient

import (
	"context"
	"io"

	adarender "github.com/zhs007/adacore/adarenderpb"
	adacoredef "github.com/zhs007/adacore/basedef"
	"google.golang.org/grpc"
)

// Client - AdaRenderServClient
type Client struct {
	servAddr string
	token    string
	conn     *grpc.ClientConn
	client   adarender.AdaRenderServiceClient
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
		return adacoredef.ErrAdaRenderClientNoServAddr
	}

	if client.token == "" {
		return adacoredef.ErrAdaRenderClientNoToken
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

// Render - MarkdownData => HTMLData
func (client *Client) Render(ctx context.Context, mddata *adarender.MarkdownData) (*adarender.HTMLData, error) {
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
		client.client = adarender.NewAdaRenderServiceClient(conn)
	}

	stream, err := client.client.Render(ctx)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	var recverr error
	var lstrect []*adarender.HTMLStream
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)

				return
			}

			if err != nil {
				recverr = err
			}

			lstrect = append(lstrect, in)
		}
	}()

	lst, err := BuildMarkdownStream(mddata)
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

	stream.CloseSend()
	<-waitc

	if recverr != nil {

		// if error, reset
		client.reset()

		return nil, recverr
	}

	htmldata, htmlerr := BuildHTMLData(lstrect)
	if htmlerr != nil {
		// if error, reset
		client.reset()
	}

	return htmldata, nil
}
