package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	startTime      time.Time
	livenessFails  int
	readinessFails int
)

func init() {
	startTime = time.Now()
	livenessFails = 0
	readinessFails = 0
}

// Randomized failure logic based on failure rate
func randomizedFailure() bool {
	failureRate, _ := strconv.Atoi(os.Getenv("FAILURE_RATE"))
	return failureRate > 0 && rand.Intn(100) < failureRate
}

// Liveness probe handler
func livenessHandler(w http.ResponseWriter, r *http.Request) {
	// Randomized failure
	if randomizedFailure() {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Randomized liveness probe failure")
		return
	}

	// Simulate latency
	latency, _ := strconv.Atoi(os.Getenv("LATENCY"))
	time.Sleep(time.Duration(latency) * time.Millisecond)

	// Simulate failure after a certain time
	failAfterTime, _ := strconv.Atoi(os.Getenv("LIVENESS_FAIL_AFTER_TIME"))
	if failAfterTime > 0 && int(time.Since(startTime).Seconds()) >= failAfterTime {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Liveness probe failed due to time threshold")
		return
	}

	// Recoverable failures
	recoveryPeriod, _ := strconv.Atoi(os.Getenv("RECOVERY_PERIOD"))
	failCount, _ := strconv.Atoi(os.Getenv("FAIL_COUNT"))

	if livenessFails >= failCount && livenessFails < failCount+recoveryPeriod {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Recoverable liveness probe failure")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Liveness probe successful")
	}

	livenessFails++
}

// Readiness probe handler
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Randomized failure
	if randomizedFailure() {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Randomized readiness probe failure")
		return
	}

	// Simulate latency
	latency, _ := strconv.Atoi(os.Getenv("LATENCY"))
	time.Sleep(time.Duration(latency) * time.Millisecond)

	// Simulate network partition
	disconnectAfter, _ := strconv.Atoi(os.Getenv("DISCONNECT_AFTER"))
	reconnectAfter, _ := strconv.Atoi(os.Getenv("RECONNECT_AFTER"))

	if disconnectAfter > 0 && int(time.Since(startTime).Minutes()) >= disconnectAfter &&
		int(time.Since(startTime).Minutes()) < disconnectAfter+reconnectAfter {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Network partition simulated")
		return
	}

	// Check dependency
	resp, err := http.Get("http://localhost:8080/dependency")
	if err != nil || resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Readiness probe failed due to dependency")
		return
	}

	// Simulate failure after a certain time
	failAfterTime, _ := strconv.Atoi(os.Getenv("READINESS_FAIL_AFTER_TIME"))
	if failAfterTime > 0 && int(time.Since(startTime).Seconds()) >= failAfterTime {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Readiness probe failed due to time threshold")
		return
	}

	// Recoverable failures
	recoveryPeriod, _ := strconv.Atoi(os.Getenv("RECOVERY_PERIOD"))
	failCount, _ := strconv.Atoi(os.Getenv("FAIL_COUNT"))

	if readinessFails >= failCount && readinessFails < failCount+recoveryPeriod {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Recoverable readiness probe failure")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Readiness probe successful")
	}

	readinessFails++
}

// Dependency simulation handler
func dependencyHandler(w http.ResponseWriter, r *http.Request) {
	failDependency, _ := strconv.ParseBool(os.Getenv("FAIL_DEPENDENCY"))
	if failDependency {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Dependency failure")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Dependency is healthy")
	}
}

func main() {
	http.HandleFunc("/liveness", livenessHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/dependency", dependencyHandler) // Dependency handler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
