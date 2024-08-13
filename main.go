package main

import (
    "log"
    "net/http"
    "time"
)

// Variables to track popup activity
var popupDetected bool
const detectionTimeout = 5 * time.Second

// Utility function to log XSS detection
func logXSSDetection(messageType, message, defaultInput string) {
    log.Printf("\x1b[32mXSS Detected! Type: %s - Message: %s %s\x1b[0m", messageType, message, defaultInput)
    popupDetected = true // Mark that a popup was detected
}

// Utility function to log no XSS detection
func logNoXSSDetection() {
    if !popupDetected {
        log.Printf("\x1b[31mNo XSS Detected.\x1b[0m")
    }
}

// Handler for HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
    // Simulating detection of popups based on URL or request parameters
    // This is a placeholder; you should replace it with actual detection logic
    if r.URL.Path == "/alert" {
        logXSSDetection("Alert", "Alert triggered", "")
    } else if r.URL.Path == "/confirm" {
        logXSSDetection("Confirm", "Confirm triggered", "")
    } else if r.URL.Path == "/prompt" {
        logXSSDetection("Prompt", "Prompt triggered", "Default input")
    }

    // Respond to the request
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Request processed"))
}

func main() {
    // Set up the HTTP server
    http.HandleFunc("/", handler)
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    // Set up a timeout to check for popup detection
    time.AfterFunc(detectionTimeout, func() {
        logNoXSSDetection()
        // Optionally, you could use additional logic to reset the popupDetected flag
        // or continuously check based on more sophisticated conditions.
    })

    // Keep the server running
    select {}
}
