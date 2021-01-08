package grpcErrTest

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatusError interface {
	error
	GRPCStatus() *status.Status
}

type errGRPCStatusError struct {
	code    codes.Code
	details []string
}

func (e errGRPCStatusError) Error() string {
	return "errstatus:" +
		" code=" + e.code.String() +
		" details=[" + strings.Join(e.details, ",") + "]"
}

func (e errGRPCStatusError) GRPCStatus() *status.Status {
	details := make([]proto.Message, len(e.details))
	for i, detail := range e.details {
		details[i] = &errdetails.DebugInfo{
			Detail: detail,
		}
	}

	st, _ := status.New(e.code, e.code.String()).WithDetails(details...)
	return st
}

func NewGRPCStatusError(code codes.Code, details []string) GRPCStatusError {
	return &errGRPCStatusError{
		code:    code,
		details: details,
	}
}
