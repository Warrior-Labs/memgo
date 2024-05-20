package server

import (
	"context"
	"errors"
	"memgo/grpcerrors"
	"memgo/memgopb"
	"memgo/ptr"
	"memgo/types"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// gRPC Write Method
func (s *MemgoServer) Write(ctx context.Context, in *memgopb.WriteRequest) (*emptypb.Empty, error) {
	// Validate the key
	if err := s.validateKey(in.Key); err != nil {
		return nil, grpcerrors.NewBadRequestError(
			"key",
			"invalid key",
			err.Error(),
		)
	}

	// Set the page value
	page := types.Page{
		Value: in.Value,
	}

	// Set the expiration time if it is not nil
	if in.ExpiresAt != nil {
		// Make sure the time is at least 1 second in the future
		if in.ExpiresAt.AsTime().Before(time.Now().Add(time.Second)) {
			return nil, grpcerrors.NewBadRequestError(
				"expires_at",
				"invalid expiration time",
				"expiration time must be at least 1 second in the future",
			)
		}
		page.ExpiresAt = ptr.To(in.ExpiresAt.AsTime())
	}

	// Prevent concurrent access
	s.pageMtx.Lock()
	defer s.pageMtx.Unlock()

	// Get the Memory Usage
	total := s.getMemUsage()

	// Check if the total memory usage is greater than the allowed size
	if total > s.size {
		return nil, grpcerrors.NewInternalServerError(
			errors.New("memory usage is greater than the allowed size"),
		)
	}

	// Set the page in the map
	s.Pages[in.Key] = page

	// Return an empty response
	return &emptypb.Empty{}, nil
}
