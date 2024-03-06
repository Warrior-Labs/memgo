package server

import (
	"runtime"
)

func (s *MemgoServer) getMemUsage() uint64 {
	// Get the Heap Allocation
	runtime.ReadMemStats(&s.m)
	return s.m.HeapAlloc
}
