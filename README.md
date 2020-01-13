# fdb-prometheus-exporter  [![Docker Pulls](https://img.shields.io/docker/pulls/pierrezemb/fdb-prometheus-exporter.svg?style=plastic)](https://hub.docker.com/r/pierrezemb/fdb-prometheus-exporter/) ![go build](https://github.com/PierreZ/fdb-prometheus-exporter/workflows/Build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/PierreZ/fdb-prometheus-exporter)](https://goreportcard.com/report/github.com/PierreZ/fdb-prometheus-exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter?status.svg)](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter)
A FoundationDB Prometheus metrics exporter

## Deployments

### Docker-Compose

An example using Docker-compose is available:

```bash
cd deployment/docker-compose
docker-compose up --build

# Metrics will be available at
curl localhost:8081/metrics | grep fdb
```

### Kubernetes

An example to deploy `fdb-prometheus-exporter` on Kubernetes is available. It has been tested with the official [FDB Operator](https://github.com/FoundationDB/fdb-kubernetes-operator).

```bash
# After deploying the sample-cluster
kubectl apply -f ./deployment/kubernetes/fdb-metrics-pod.yaml

# You can view metrics through kubctl proxy
kubectl port-forward fdb-prometheus-exporter 8080:8080

# To destroy it
kubectl delete -f ./deployment/kubernetes/fdb-metrics-pod.yaml
```