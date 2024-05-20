package grpcerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewInternalServerError(err error) error {
	st := status.New(codes.Internal, err.Error())
	return st.Err()
}
