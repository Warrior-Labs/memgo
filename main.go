package main

import (
	"fmt"
	"log"
	"memgo/memgopb"
	"memgo/server"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// Check that /x509/certificate.pem and /x509/key.pem files are present
	tlsEnabled := true
	if _, err := os.Stat("/x509/certificate.pem"); os.IsNotExist(err) {
		tlsEnabled = false
	}
	if _, err := os.Stat("/x509/key.pem"); os.IsNotExist(err) {
		tlsEnabled = false
	}

	if tlsEnabled {
		// Load the TLS Certificates
		creds, err := credentials.NewServerTLSFromFile("/x509/certificate.pem", "/x509/key.pem")
		if err != nil {
			log.Fatalln("unable to load TLS certificates: ", err.Error())
		}
		log.Println("TLS Enabled")
		opts = append(opts, grpc.Creds(creds))
	}

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
