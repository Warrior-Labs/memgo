package grpcerrors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewNotFoundError(businessMsg string, args ...any) error {
	msg := fmt.Sprintf(businessMsg, args...)
	st := status.New(codes.NotFound, msg)
	return st.Err()
}
