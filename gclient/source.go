package gclient

import (
	"context"
	"os"

	"github.com/zhenhuanlee/protos/proto"
	"google.golang.org/grpc"
)

var (
	sourceurl    string
	sourceConn   *grpc.ClientConn
	sourceClient proto.GouClient
)

func init() {
	sourceurl = os.Getenv("SOURCEURL")
}

// NewSourceClient new a source client
func NewSourceClient(url string) (proto.GouClient, error) {
	if sourceClient != nil {
		return sourceClient, nil
	}
	var err error
	sourceConn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	sourceClient = proto.NewGouClient(sourceConn)
	return sourceClient, nil
}

// CheckIn send msg
func CheckIn(uuid, kind, body string) (*proto.Resp, error) {
	if sourceurl == "" {
		sourceurl = os.Getenv("SOURCEURL")
	}

	c, err := NewSourceClient(sourceurl)
	if err != nil {
		return nil, err
	}

	r, err := c.CheckIn(context.Background(), &proto.Req{
		Uuid: uuid,
		Kind: kind,
		Body: body,
	})
	return r, err
}

// Close close
func Close() {
	sourceConn.Close()
}

/*-----------------------------------------------------------*/
