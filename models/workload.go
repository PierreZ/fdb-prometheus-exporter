package models

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	workloadOperationHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_per_second",
		Help: "operations per second",
	}, []string{"operation"})

	workloadOperationCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_total",
		Help: "Gauge of operations",
	}, []string{"operation"})

	workloadTransactionHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_transactions_per_second",
		Help: "rate of transactions",
	}, []string{"status"})
	workloadTransactionCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_transactions_total",
		Help: "Gauge of operations",
	}, []string{"status"})
)

// ExportWorkload is exporting workloads
func (s FDBStatus) ExportWorkload() {

	fmt.Println("generating workload")

	// writes
	workloadOperationHZ.With(prometheus.Labels{
		"operation": "writes",
	}).Set(s.Cluster.Workload.Operations.Reads.Hz)
	workloadOperationHZ.With(prometheus.Labels{
		"operation": "reads",
	}).Set(s.Cluster.Workload.Operations.Writes.Hz)

	// reads
	workloadOperationCounter.With(prometheus.Labels{
		"operation": "writes",
	}).Set(s.Cluster.Workload.Operations.Reads.Counter)
	workloadOperationCounter.With(prometheus.Labels{
		"operation": "reads",
	}).Set(s.Cluster.Workload.Operations.Writes.Counter)

	// commited transactions
	workloadTransactionHZ.With(prometheus.Labels{
		"status": "committed",
	}).Set(s.Cluster.Workload.Transactions.Committed.Hz)
	workloadTransactionCounter.With(prometheus.Labels{
		"status": "committed",
	}).Set(s.Cluster.Workload.Transactions.Committed.Counter)

	// conflicted transactions
	workloadTransactionHZ.With(prometheus.Labels{
		"status": "conflicted",
	}).Set(s.Cluster.Workload.Transactions.Conflicted.Hz)
	workloadTransactionCounter.With(prometheus.Labels{
		"status": "conflicted",
	}).Set(s.Cluster.Workload.Transactions.Conflicted.Counter)

	// started transactions
	workloadTransactionHZ.With(prometheus.Labels{
		"status": "started",
	}).Set(s.Cluster.Workload.Transactions.Started.Hz)
	workloadTransactionCounter.With(prometheus.Labels{
		"status": "started",
	}).Set(s.Cluster.Workload.Transactions.Started.Counter)
}

func registerWorkload(r *prometheus.Registry) {
	r.MustRegister(workloadOperationHZ)
	r.MustRegister(workloadOperationCounter)
	r.MustRegister(workloadTransactionHZ)
	r.MustRegister(workloadTransactionCounter)
	fmt.Println("registered workload metrics")
}
