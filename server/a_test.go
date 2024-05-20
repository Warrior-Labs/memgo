package server

import (
	"context"
	"log"
	"memgo/memgopb"
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type Suite struct {
	suite.Suite
	lis *bufconn.Listener
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	bufSize := 1024 * 1024
	s.lis = bufconn.Listen(bufSize)
	grpcSrv := grpc.NewServer()
	memgoSrv, err := NewMemgoServer()
	if err != nil {
		log.Fatalf("unable to create Memgo Server: %v", err)
	}
	memgopb.RegisterMemgoServiceServer(grpcSrv, memgoSrv)
	go func() {
		if err := grpcSrv.Serve(s.lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func (s *Suite) bufDialer(context.Context, string) (net.Conn, error) {
	return s.lis.Dial()
}
