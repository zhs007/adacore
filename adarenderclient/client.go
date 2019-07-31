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
		return nil, err
	}

	var recverr error
	waitc := make(chan struct{})
	go func() {
		for {
			// in, err := stream.Recv()
			_, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				recverr = err
			}
		}
	}()

	// for _, note := range notes {
	// 	if err := stream.Send(note); err != nil {
	// 		log.Fatalf("Failed to send a note: %v", err)
	// 	}
	// }
	stream.CloseSend()
	<-waitc

	if recverr != nil {
		// jarvisbase.Warn("crawlerClient.getArticles:GetArticles", zap.Error(err))

		// if error, close connect
		client.conn.Close()

		client.conn = nil
		client.client = nil

		return nil, recverr
	}

	return nil, nil
}
