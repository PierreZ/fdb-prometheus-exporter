# fdb-prometheus-exporter  [![Docker Pulls](https://img.shields.io/docker/pulls/pierrezemb/fdb-prometheus-exporter.svg?style=plastic)](https://hub.docker.com/r/pierrezemb/fdb-prometheus-exporter/) ![go build](https://github.com/PierreZ/fdb-prometheus-exporter/workflows/Build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/PierreZ/fdb-prometheus-exporter)](https://goreportcard.com/report/github.com/PierreZ/fdb-prometheus-exporter) [![GoDoc](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter?status.svg)](https://godoc.org/github.com/PierreZ/fdb-prometheus-exporter)
A FoundationDB Prometheus metrics exporter

## Building

### Requirements

* go >= 1.13

```bash
git clone git@github.com:PierreZ/fdb-prometheus-exporter.git

go install
```

## Using it

`fdb-prometheus-exporter` is env-var driven, you can customize:

* `FDB_API_VERSION`
* `FDB_CLUSTER_FILE`
* `FDB_CREATE_CLUSTER_FILE`
* `FDB_EXPORT_WORKLOAD`
* `FDB_METRICS_LISTEN`
* `FDB_METRICS_EVERY`

## Deployments

### Docker-Compose

An example using Docker-compose is available:

```bash
cd deployment/docker-compose
docker-compose up --build

# Metrics will be available at
curl localhost:8081/metrics | grep fdb | grep -v "#"
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

You can also find a `Batch` job for [go-ycsb](https://github.com/pingcap/go-ycsb/) to spawn some workloads.

```bash
# spawn ycsbn workload a to f as a K8S Batch
kubectl apply -f ./deployment/kubernetes/go-ycsb-batch.yaml
```