package gclient

import (
	"context"
	"fmt"
	"os"

	"github.com/zhenhuanlee/protos/gproto"
	"google.golang.org/grpc"
)

var (
	sourceurl    string
	sourceConn   *grpc.ClientConn
	sourceClient gproto.GouClient
)

func init() {
	sourceurl = os.Getenv("SOURCEURL")
}

// NewSourceClient new a source client
func NewSourceClient(url string) (gproto.GouClient, error) {
	if sourceClient != nil {
		return sourceClient, nil
	}
	var err error
	sourceConn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	sourceClient = gproto.NewGouClient(sourceConn)
	return sourceClient, nil
}

// CheckIn send msg
func CheckIn(uuid, kind, body string) (*gproto.Resp, error) {
	if sourceurl == "" {
		sourceurl = os.Getenv("SOURCEURL")
	}
	fmt.Println(sourceurl)
	c, err := NewSourceClient(sourceurl)
	if err != nil {
		return nil, err
	}

	r, err := c.CheckIn(context.Background(), &gproto.Req{
		Uuid: uuid,
		Kind: kind,
		Body: body,
	})
	return r, err
}

/*-----------------------------------------------------------*/
