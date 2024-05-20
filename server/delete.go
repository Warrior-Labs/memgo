package server

import (
	"context"
	"memgo/memgopb"

	"google.golang.org/protobuf/types/known/emptypb"
)

// gRPC Delete Method
func (s *MemgoServer) Delete(ctx context.Context, in *memgopb.DeleteRequest) (*emptypb.Empty, error) {
	// Validate the key
	if err := s.validateKey(in.Key); err != nil {
		return nil, err
	}

	// Prevent concurrent access
	s.pageMtx.Lock()
	defer s.pageMtx.Unlock()

	// Delete the page
	delete(s.Pages, in.Key)

	// Return an empty response
	return &emptypb.Empty{}, nil
}
