package server

import "time"

func (s *MemgoServer) CleanExpired() {
	// Loop forever
	for {
		// Get the current time
		now := time.Now()

		// Lock the mutex
		s.pageMtx.Lock()

		// Loop through the keys
		for key, page := range s.Pages {
			// Check if the key has expired
			if page.ExpiresAt != nil && page.ExpiresAt.Before(now) {
				// Delete the key
				delete(s.Pages, key)
			}
		}

		// Unlock the mutex
		s.pageMtx.Unlock()

		// Sleep for 10 seconds
		time.Sleep(10 * time.Second)
	}
}
