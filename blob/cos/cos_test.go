package cos

import (
	"context"
	"fmt"
	blobpb "github.com/shenxiang11/coolcar/blob/gen/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreate(t *testing.T) {
	conn, err := grpc.Dial("localhost:10003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := blobpb.NewBlobServiceClient(conn)

	ctx := context.Background()
	res, err := c.CreateBlob(ctx, &blobpb.CreateBlobRequest{
		AccountId:           "account2",
		UploadUrlTimeoutSec: 1000,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
}
