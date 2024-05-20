package grpcerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewUnauthorizedError(businessMsg string) error {
	st := status.New(codes.Unauthenticated, businessMsg)
	return st.Err()
}
