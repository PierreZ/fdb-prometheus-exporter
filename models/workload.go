package models

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	workloadOperationHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_per_second",
		Help: "operations per second",
	}, []string{"operation", "unit"})

	workloadOperationCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_total",
		Help: "Gauge of operations",
	}, []string{"operation", "unit"})
)

// ExportWorkload is exporting workloads
func (s FDBStatus) ExportWorkload() {

	fmt.Println("generating workload")

	workloadOperationHZ.With(prometheus.Labels{
		"operation": "writes",
		"unit":      "hz",
	}).Set(s.Cluster.Workload.Operations.Reads.Hz)
	workloadOperationHZ.With(prometheus.Labels{
		"operation": "reads",
		"unit":      "hz",
	}).Set(s.Cluster.Workload.Operations.Writes.Hz)

	workloadOperationCounter.With(prometheus.Labels{
		"operation": "writes",
		"unit":      "hz",
	}).Set(s.Cluster.Workload.Operations.Reads.Counter)
	workloadOperationCounter.With(prometheus.Labels{
		"operation": "reads",
		"unit":      "hz",
	}).Set(s.Cluster.Workload.Operations.Writes.Counter)
}

// Register is registering metrics
func Register(r *prometheus.Registry) {
	r.MustRegister(workloadOperationHZ)
	r.MustRegister(workloadOperationCounter)
	fmt.Println("registered workload metrics")
}
