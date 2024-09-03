# kube-probesim

## Overview

`kube-probesim` is a lightweight HTTP server designed to simulate Kubernetes liveness and readiness probes. This tool is ideal for testing and debugging how your Kubernetes cluster handles various failure scenarios such as random probe failures, response latency, and dependency outages.
Its main purpose is to demonstrate how live and readyness probes work, as well as being able to demonstrate it

## Features

- **Liveness Probe Simulation**: Simulate failure based on time elapsed since start or randomly.
- **Readiness Probe Simulation**: Mimic readiness probe behavior with configurable failure modes.
- **External Dependency Simulation**: Test scenarios where external service dependencies may fail.
- **Latency Injection**: Simulate network latency by adding a delay to responses.
- **Network Partition Simulation**: Simulate conditions where the service becomes unreachable after a certain time.

## Configuration

Customize the simulation using the following environment variables:

- **`FAILURE_RATE`** (int, 0-100): Probability percentage for random failure of probes.
- **`LATENCY`** (int, milliseconds): Simulated delay for each HTTP response.
- **`LIVENESS_FAIL_AFTER_TIME`** (int, seconds): Time after which the liveness probe will start failing.
- **`READINESS_FAIL_AFTER_TIME`** (int, seconds): Time after which the readiness probe will start failing.
- **`DISCONNECT_AFTER`** (int, minutes): Time after which the service will simulate a network partition.
- **`RECONNECT_AFTER`** (int, minutes): Duration of the network partition.
- **`FAIL_DEPENDENCY`** (bool): Forces the `/dependency` endpoint to return failure.

## Deployment

### Kubernetes Manifest

To deploy `kube-probesim` in your Kubernetes cluster, you can create a simple deployment YAML:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-probesim
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-probesim
  template:
    metadata:
      labels:
        app: kube-probesim
    spec:
      containers:
      - name: kube-probesim
        image: ghcr.io/jhoelzel/kube-probesim:latest
        ports:
        - containerPort: 8080
        env:
        - name: FAILURE_RATE
          value: "20"
        - name: LATENCY
          value: "50"
        - name: LIVENESS_FAIL_AFTER_TIME
          value: "60"
        - name: READINESS_FAIL_AFTER_TIME
          value: "120"
        - name: DISCONNECT_AFTER
          value: "30"
        - name: RECONNECT_AFTER
          value: "10"
```

### Running Locally

You can also run the server locally for testing:

```bash
FAILURE_RATE=10 LATENCY=100 LIVENESS_FAIL_AFTER_TIME=60 ./kube-probesim
```

This will start the server on port 8080 with the specified configurations.

## Endpoints

- **`/liveness`**: Simulates the liveness probe.
- **`/readiness`**: Simulates the readiness probe.
- **`/dependency`**: Simulates an external dependency, which can be forced to fail with `FAIL_DEPENDENCY=true`.

## Example Scenarios

### Random Failures

To simulate a 10% chance of random failures in both liveness and readiness probes:

```bash
FAILURE_RATE=10 ./kube-probesim
```

### Liveness Probe Failures After a Delay

Simulate a liveness probe that fails after 60 seconds:

```bash
LIVENESS_FAIL_AFTER_TIME=60 ./kube-probesim
```

## Monitoring and Debugging

You can integrate `kube-probesim` with your logging and monitoring stack in Kubernetes, allowing you to visualize how your cluster responds to simulated failures.

## License

`kube-probesim` is licensed under the MIT License.

## Contributions

Contributions are welcome! Feel free to open an issue or submit a pull request on [GitHub](https://github.com/jhoelzel/kube-probesim).

## Contact

For any questions or issues, please reach out via the [GitHub repository](https://github.com/jhoelzel/kube-probesim).

---

This README reflects the actual capabilities and intended deployment scenarios for `kube-probesim`, making it clearer for users how to deploy and configure the tool within their Kubernetes environments.
