package server

import (
	"fmt"
	"memgo/memgopb"
	"memgo/types"
	"os"
	"runtime"
	"sync"
)

type MemgoServer struct {
	// Private fields for the server
	size uint64
	m    runtime.MemStats

	// Port for the server
	Port string

	// Page Array
	pageMtx sync.Mutex
	Pages   map[string]types.Page

	// gRPC Server Definition
	memgopb.UnimplementedMemgoServiceServer
}

func NewMemgoServer() (*MemgoServer, error) {
	// Check if MEMGO_PORT is set, and revert to 8000 if not
	port := "8000"

	// Check if MEMGO_MAX_SIZE is set, and revert to 128M if not
	size := "128M"
	if s := os.Getenv("MEMGO_MAX_SIZE"); s != "" {
		size = s
	}
	sizeBytes := parseMemString(size)

	// Create a new MemgoServer
	s := &MemgoServer{
		size:  sizeBytes,
		Port:  port,
		Pages: make(map[string]types.Page),
	}

	return s, nil
}

func parseMemString(mem string) uint64 {
	// Extract the number and the unit from the string
	var num uint64
	var unit string
	fmt.Sscanf(mem, "%d%s", &num, &unit)

	// Convert the number to bytes
	switch unit {
	case "K":
		return num * 1024
	case "M":
		return num * 1024 * 1024
	case "G":
		return num * 1024 * 1024 * 1024
	default:
		return num
	}
}
