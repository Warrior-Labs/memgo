package grpcerrors

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBadRequestError(fieldName, internalMsg, businessMsg string) error {
	st := status.New(codes.InvalidArgument, internalMsg)
	v := &errdetails.BadRequest_FieldViolation{
		Field:       fieldName,
		Description: businessMsg,
	}
	br := &errdetails.BadRequest{}
	br.FieldViolations = append(br.FieldViolations, v)
	st, err := st.WithDetails(br)
	if err != nil {
		// If this errored, it will always error
		// here, so better panic so we can figure
		// out why than have this silently passing.
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}
	return st.Err()
}
