# ProbeSim

ProbeSim is a lightweight Go application designed to simulate various failure scenarios for Kubernetes liveness and readiness probes. It helps you test and ensure the resilience of your Kubernetes-deployed services by mimicking conditions such as random failures, time-based failures, dependency issues, network partitions, and latency.

## Features

- **Randomized Failure Mode**: Simulate random failures with a configurable failure rate.
- **Time-Based Failures**: Trigger probe failures after a specified time since the application start.
- **Recoverable Failures**: Introduce temporary failures that recover after a certain number of checks.
- **Dependency Simulation**: Simulate external dependency failures that affect the readiness of the service.
- **Network Partition Simulation**: Mimic network issues that temporarily disconnect the service.
- **Latency Simulation**: Introduce delays in probe responses to test how Kubernetes handles slow responses.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (if running locally)
- Docker (if containerizing the application)
- Kubernetes (for deployment)
