package models

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	workloadBytesHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_bytes_per_second",
		Help: "Workload bytes per second",
	}, []string{"operation"})

	workloadBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_bytes_total",
		Help: "Workload bytes total",
	}, []string{"operation"})

	workloadBytesRoughness = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_bytes_roughness",
		Help: "Workload bytes roughness",
	}, []string{"operation"})

	workloadKeysHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_keys_per_second",
		Help: "Workload keys per second",
	}, []string{"operation"})

	workloadKeysCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_keys_total",
		Help: "Workload keys total",
	}, []string{"operation"})

	workloadKeysRoughness = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_keys_roughness",
		Help: "Workload keys roughness",
	}, []string{"operation"})
	
	workloadOperationsHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_per_second",
		Help: "Workload operations per second",
	}, []string{"operation"})

	workloadOperationsCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_total",
		Help: "Workload operations total",
	}, []string{"operation"})

	workloadOperationsRoughness = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_operations_roughness",
		Help: "Workload operations roughness",
	}, []string{"operation"})

	workloadTransactionsHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_transactions_per_second",
		Help: "Workload transactions per second",
	}, []string{"transaction"})

	workloadTransactionsCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_transactions_total",
		Help: "Workload transactions total",
	}, []string{"transaction"})

	workloadTransactionsRoughness = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_workload_transactions_roughness",
		Help: "Workload transactions roughness",
	}, []string{"transaction"})
)

// ExportWorkload is exporting workloads
func (s FDBStatus) ExportWorkload() {
	workloadBytesHZ.With(prometheus.Labels{
		"operation": "written",
	}).Set(s.Cluster.Workload.Bytes.Written.Hz)

	workloadBytesCounter.With(prometheus.Labels{
		"operation": "written",
	}).Set(s.Cluster.Workload.Bytes.Written.Counter)

	workloadBytesRoughness.With(prometheus.Labels{
		"operation": "written",
	}).Set(s.Cluster.Workload.Bytes.Written.Roughness)

	workloadBytesHZ.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Bytes.Read.Hz)

	workloadBytesCounter.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Bytes.Read.Counter)

	workloadBytesRoughness.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Bytes.Read.Roughness)

	workloadKeysHZ.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Keys.Read.Hz)

	workloadKeysCounter.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Keys.Read.Counter)

	workloadKeysRoughness.With(prometheus.Labels{
		"operation": "read",
	}).Set(s.Cluster.Workload.Keys.Read.Roughness)

	workloadOperationsHZ.With(prometheus.Labels{
		"operation": "read_requests",
	}).Set(s.Cluster.Workload.Operations.ReadRequests.Hz)

	workloadOperationsCounter.With(prometheus.Labels{
		"operation": "read_requests",
	}).Set(s.Cluster.Workload.Operations.ReadRequests.Counter)

	workloadOperationsRoughness.With(prometheus.Labels{
		"operation": "read_requests",
	}).Set(s.Cluster.Workload.Operations.ReadRequests.Roughness)

	workloadOperationsHZ.With(prometheus.Labels{
		"operation": "reads",
	}).Set(s.Cluster.Workload.Operations.Reads.Hz)

	workloadOperationsCounter.With(prometheus.Labels{
		"operation": "reads",
	}).Set(s.Cluster.Workload.Operations.Reads.Counter)

	workloadOperationsRoughness.With(prometheus.Labels{
		"operation": "reads",
	}).Set(s.Cluster.Workload.Operations.Reads.Roughness)

	workloadOperationsHZ.With(prometheus.Labels{
		"operation": "writes",
	}).Set(s.Cluster.Workload.Operations.Writes.Hz)

	workloadOperationsCounter.With(prometheus.Labels{
		"operation": "writes",
	}).Set(s.Cluster.Workload.Operations.Writes.Counter)

	workloadOperationsRoughness.With(prometheus.Labels{
		"operation": "writes",
	}).Set(s.Cluster.Workload.Operations.Writes.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "committed",
	}).Set(s.Cluster.Workload.Transactions.Committed.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "committed",
	}).Set(s.Cluster.Workload.Transactions.Committed.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "committed",
	}).Set(s.Cluster.Workload.Transactions.Committed.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "conflicted",
	}).Set(s.Cluster.Workload.Transactions.Conflicted.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "conflicted",
	}).Set(s.Cluster.Workload.Transactions.Conflicted.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "conflicted",
	}).Set(s.Cluster.Workload.Transactions.Conflicted.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "started",
	}).Set(s.Cluster.Workload.Transactions.Started.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "started",
	}).Set(s.Cluster.Workload.Transactions.Started.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "started",
	}).Set(s.Cluster.Workload.Transactions.Started.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "started_batch_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedBatchPriority.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "started_batch_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedBatchPriority.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "started_batch_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedBatchPriority.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "started_default_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedDefaultPriority.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "started_default_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedDefaultPriority.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "started_default_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedDefaultPriority.Roughness)

	workloadTransactionsHZ.With(prometheus.Labels{
		"transaction": "started_immediate_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedImmediatePriority.Hz)

	workloadTransactionsCounter.With(prometheus.Labels{
		"transaction": "started_immediate_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedImmediatePriority.Counter)

	workloadTransactionsRoughness.With(prometheus.Labels{
		"transaction": "started_immediate_priority",
	}).Set(s.Cluster.Workload.Transactions.StartedImmediatePriority.Roughness)
}

func registerWorkload(r *prometheus.Registry) {
	r.MustRegister(workloadBytesHZ)
	r.MustRegister(workloadBytesCounter)
	r.MustRegister(workloadBytesRoughness)
	r.MustRegister(workloadKeysHZ)
	r.MustRegister(workloadKeysCounter)
	r.MustRegister(workloadKeysRoughness)
	r.MustRegister(workloadOperationsHZ)
	r.MustRegister(workloadOperationsCounter)
	r.MustRegister(workloadOperationsRoughness)
	r.MustRegister(workloadTransactionsHZ)
	r.MustRegister(workloadTransactionsCounter)
	r.MustRegister(workloadTransactionsRoughness)

	fmt.Println("registered workload metrics")
}
