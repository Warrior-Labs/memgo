package server

import (
	"context"
	"memgo/memgopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Suite) TestReadWrite_OK() {
	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(s.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		s.T().Fatal(err)
	}
	defer conn.Close()

	client := memgopb.NewMemgoServiceClient(conn)
	_, err = client.Write(context.Background(), &memgopb.WriteRequest{
		Key:   "key",
		Value: []byte("value"),
	})
	s.NoError(err)

	rtn, err := client.Read(context.Background(), &memgopb.ReadRequest{
		Key: "key",
	})
	s.NoError(err)
	s.Equal([]byte("value"), rtn.Value)
}

func (s *Suite) TestWrite_NoKeyError() {
	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(s.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		s.T().Fatal(err)
	}
	defer conn.Close()

	client := memgopb.NewMemgoServiceClient(conn)
	_, err = client.Write(context.Background(), &memgopb.WriteRequest{
		Value: []byte("value"),
	})

	s.Error(err)
}
