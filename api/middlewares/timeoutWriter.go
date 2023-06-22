package middlewares

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type timeoutWriter struct {
	gin.ResponseWriter
	isExpired bool
	mu        sync.Mutex
}

func newTimeoutWriter(w gin.ResponseWriter) *timeoutWriter {
	return &timeoutWriter{ResponseWriter: w, isExpired: false}
}

// Timeout the writer
func (tw *timeoutWriter) Expire() {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	tw.isExpired = true
}

// Return whehter the writer has timed out
func (tw *timeoutWriter) IsExpired() bool {
	return tw.isExpired
}

// Write the data to the body
func (tw *timeoutWriter) Write(data []byte) (int, error) {
	if !tw.isExpired {
		return tw.ResponseWriter.Write(data)
	}

	return 0, nil
}

// Write the status code to the header
func (tw *timeoutWriter) WriteHeader(code int) {
	if !tw.isExpired {
		tw.ResponseWriter.WriteHeader(code)
	}
}

// Write string to the body
func (tw *timeoutWriter) WriteString(s string) (int, error) {
	if !tw.isExpired {
		return tw.ResponseWriter.WriteString(s)
	}
	return 0, nil
}
