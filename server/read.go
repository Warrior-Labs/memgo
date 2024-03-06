package server

import (
	"context"
	"errors"
	"memgo/memgopb"
	"memgo/ptr"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// gRPC Read Method
func (s *MemgoServer) Read(ctx context.Context, in *memgopb.ReadRequest) (*memgopb.ReadResponse, error) {
	// Validate the key
	if err := s.validateKey(in.Key); err != nil {
		return nil, err
	}

	// Prevent concurrent access
	s.pageMtx.Lock()
	defer s.pageMtx.Unlock()

	// Get the page from the map
	page, ok := s.Pages[in.Key]
	if !ok {
		return nil, errors.New("key not found")
	}

	// Create the response
	res := &memgopb.ReadResponse{
		Value: page.Value,
	}
	if page.ExpiresAt != nil {
		// Check if the page has expired
		if page.ExpiresAt.Before(time.Now()) {
			// Delete the page
			delete(s.Pages, in.Key)
			// Return an error
			return nil, errors.New("key has expired")
		}
		res.ExpiresAt = timestamppb.New(ptr.From(page.ExpiresAt))
	}

	// Return the value
	return res, nil
}
