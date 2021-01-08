package main

import (
	"context"
	"fmt"
	"log"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/isbang/grpc-err-test/pb"
)

func main() {
	cc, err := grpc.Dial("localhost:50021", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer cc.Close()

	svc := pb.NewCodeSvcClient(cc)
	{
		_, err := svc.GetErrCode(context.Background(), &pb.GetErrCodeReq{
			Code: uint32(codes.NotFound),
			DetailMessage: []string{
				"foo", "bar", "this", "is", "detail", "message",
			},
		})

		if err == nil {
			log.Panic("should be error but not error")
		}

		if st, ok := status.FromError(err); ok {
			fmt.Printf("err: %v\n"+
				"code: %v\n"+
				"detail: %+v\n", err, st.Code(), st.Details())
		} else {
			fmt.Println("err: ", err)
		}
	}
}
