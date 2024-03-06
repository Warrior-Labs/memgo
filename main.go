package main

import (
	"fmt"
	"log"
	"memgo/memgopb"
	"memgo/server"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a new Memgo Server
	srv, err := server.NewMemgoServer()
	if err != nil {
		log.Fatalln(err)
	}

	// Build the Host Connection String
	host := fmt.Sprintf(":%s", srv.Port)

	// Create the TCP Listener
	lis, err := net.Listen("tcp", host)
	if err != nil {
		// Fatal Error
		log.Fatalln("unable to start TCP server: ", err.Error())
	}

	// Set the gRPC Options
	var opts []grpc.ServerOption

	// Create the new gRPC Server
	grpcServer := grpc.NewServer(opts...)

	// Register the gRPC Server
	memgopb.RegisterMemgoServiceServer(grpcServer, srv)

	// Start the Cleanup Routine
	go srv.CleanExpired()

	// Serve the Servants, oh no
	log.Println("Starting Memgo Server on port " + srv.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}
