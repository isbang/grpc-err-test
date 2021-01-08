package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	grpcErrTest "github.com/isbang/grpc-err-test"
	"github.com/isbang/grpc-err-test/pb"
)

const port = 50021

type codeSvc struct {
	pb.UnimplementedCodeSvcServer
}

func (c codeSvc) GetErrCode(ctx context.Context, req *pb.GetErrCodeReq) (*pb.GetErrCodeResp, error) {
	return nil, grpcErrTest.NewGRPCStatusError(codes.Code(req.Code), req.DetailMessage)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()

	g := grpc.NewServer()
	pb.RegisterCodeSvcServer(g, &codeSvc{})

	if err := g.Serve(lis); err != nil {
		return
	}
}
