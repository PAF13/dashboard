package httpserver

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// A thread-safe buffer to store logs.
type LogBuffer struct {
	mu   sync.Mutex
	logs []string
}

func (lb *LogBuffer) Write(p []byte) (n int, err error) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.logs = append(lb.logs, string(p))
	return len(p), nil
}

/*
func (lb *LogBuffer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	initLogSize := len(lb.logs)
	w.Write([]byte("----connected----\n"))
	w.Write([]byte("Nummber of Logs: " + strconv.Itoa(initLogSize) + "\n"))
	lb.mu.Lock()
	defer lb.mu.Unlock()
	for _, logLine := range lb.logs {
		w.Write([]byte(logLine))
	}
}
*/

// Add a method to stream logs as Server-Sent Events (SSE).
func (lb *LogBuffer) ServeSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	initLogSize := len(lb.logs)
	w.Write([]byte("----connected----\n"))
	w.Write([]byte("Nummber of Logs: " + strconv.Itoa(initLogSize) + "\n"))
	// Continuously send logs
	for {
		lb.mu.Lock()
		for _, logLine := range lb.logs {
			w.Write([]byte(logLine))
			//fmt.Fprintf(w, "%s", logLine)
		}
		lb.logs = nil // Clear sent logs
		lb.mu.Unlock()

		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second) // Adjust the frequency as needed
	}
}
func Server(logBuf *LogBuffer) {

	// Start the HTTP server in a separate goroutine.
	log.Println("Starting debug server on :8080")
	http.HandleFunc("/debug", logBuf.ServeSSE)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting debug server: %v", err)
	}

}
