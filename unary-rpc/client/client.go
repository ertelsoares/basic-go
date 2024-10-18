package client

import (
	"github.com/ertelsoares/grcp-project/tree/main/unary-rpc/pb"
	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	pb.NewUserServiceClient(dial)
}
