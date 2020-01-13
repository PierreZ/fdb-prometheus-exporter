# fdb-prometheus-exporter  [![Docker Pulls](https://img.shields.io/docker/pulls/pierrezemb/fdb-prometheus-exporter.svg?style=plastic)](https://hub.docker.com/r/pierrezemb/fdb-prometheus-exporter/) ![go build](https://github.com/PierreZ/fdb-prometheus-exporter/workflows/Build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/PierreZ/fdb-prometheus-exporter)](https://goreportcard.com/report/github.com/PierreZ/fdb-prometheus-exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter?status.svg)](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter)
A FoundationDB Prometheus metrics exporter

## Deployments

### Kubernetes

A example to deploy `fdb-prometheus-exporter` on Kubernetes is available. It has been tested with the official [FDB Operator](https://github.com/FoundationDB/fdb-kubernetes-operator).

```bash
# After deploying the sample-cluster
kubectl apply -f ./deployment/kubernetes/fdb-metrics-pod.yaml
```