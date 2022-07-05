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

	latencyProbe = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_latency_probe",
		Help: "Latency probe",
	}, []string{"operation"})

	// Measures Transactions limit
	qosTransactionsLimit = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_qos_transactions_per_second_limit",
		Help: "Transactions limit"},
		[]string{"kind"})

	qosLimitedBy = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_qos_performance_limited_by",
		Help: "Performance limit reason"})

	qosQueueBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_qos_worst_queue_bytes",
		Help: "Worst queue bytes"},
		[]string{"kind"})

	qosLimitByMessageToCode = map[string]int{
		"workload":                            0,
		"storage_server_write_queue_size":     1,
		"storage_server_write_bandwidth_mvcc": 2,
		"storage_server_readable_behind":      3,
		"log_server_mvcc_write_bandwidth":     4,
		"log_server_write_queue":              5,
		"storage_server_min_free_space":       6,
		"storage_server_min_free_space_ratio": 7,
		"log_server_min_free_space":           8,
		"log_server_min_free_space_ratio":     9,
		"storage_server_durability_lag":       10,
		"storage_server_list_fetch_failed":    11}
)

// ExportWorkload is exporting workloads
func (s FDBStatus) ExportWorkload() {

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

	// Latency probes
	latencyProbe.With(prometheus.Labels{
		"operation": "commit",
	}).Set(s.Cluster.LatencyProbe.CommitSeconds)

	latencyProbe.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.LatencyProbe.ReadSeconds)

	latencyProbe.With(prometheus.Labels{
		"operation": "transaction start",
	}).Set(s.Cluster.LatencyProbe.TransactionStartSeconds)

	latencyProbe.With(prometheus.Labels{
		"operation": "batch priority transaction start",
	}).Set(s.Cluster.LatencyProbe.BatchPriorityTransactionStartSeconds)

	latencyProbe.With(prometheus.Labels{
		"operation": "immediate priority transaction start",
	}).Set(s.Cluster.LatencyProbe.ImmediatePriorityTransactionStartSeconds)

	// QOS metrics ratekeeper vs actual
	qosTransactionsLimit.With(prometheus.Labels{
		"kind": "ratekeeper",
	}).Set(s.Cluster.Qos.TransactionsPerSecondLimit)

	qosTransactionsLimit.With(prometheus.Labels{
		"kind": "released",
	}).Set(s.Cluster.Qos.ReleasedTransactionsPerSecond)

	// QOS performance limit reason. We are doing ugly workaround here as for some reason
	// code id is not always accessible
	if reasonID, ok := qosLimitByMessageToCode[s.Cluster.Qos.PerformanceLimitedBy.Name]; ok {
		qosLimitedBy.Set(float64(reasonID))
	} else {
		qosLimitedBy.Set(-1)
	}

	// Queued bytes on storage/log server
	qosQueueBytes.With(prometheus.Labels{
		"kind": "log_server",
	}).Set(s.Cluster.Qos.WorstQueueBytesLogServer)

	qosQueueBytes.With(prometheus.Labels{
		"kind": "storage_server",
	}).Set(s.Cluster.Qos.WorstQueueBytesStorageServer)

}

func registerWorkload(r *prometheus.Registry) {
	r.MustRegister(workloadOperationHZ)
	r.MustRegister(workloadOperationCounter)
	r.MustRegister(workloadTransactionHZ)
	r.MustRegister(workloadTransactionCounter)
	r.MustRegister(latencyProbe)
	r.MustRegister(qosTransactionsLimit)
	r.MustRegister(qosLimitedBy)
	r.MustRegister(qosQueueBytes)
	fmt.Println("registered workload metrics")
}
