package models

import "github.com/prometheus/client_golang/prometheus"

var (
	clientsCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_clients_count",
		Help: "Clients count",
	})

	clusterControllerTimestamp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_controller_timestamp",
		Help: "Cluster controller timestamp",
	})

	ClusterDataAveragePartitionSizeBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_data_average_partition_size_bytes",
		Help: "Cluster average partition size bytes",
	})

	clusterLeastOperatingSpaceBytesLogServer = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_data_least_operating_space_bytes_log_server",
		Help: "Cluster least operating space bytes log server",
	})

	clusterLeastOperatingSpaceBytesStorageServer = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_data_least_operating_space_bytes_storage_server",
		Help: "Cluster least operating space bytes storage server",
	})

	clusterMovingDataHighestPriority = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_moving_data_highest_priority",
		Help: "Cluster moving data highest priority",
	})

	clusterMovingDataInQueueBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_moving_data_in_flight_bytes",
		Help: "Cluster moving data in flight bytes",
	})

	clusterMovingDataInFlightBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_moving_data_in_flight_bytes",
		Help: "Cluster moving data in flight bytes",
	})

	clusterMovingDataTotalWrittenBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_moving_data_total_written_bytes",
		Help: "Cluster moving data total written bytes",
	})

	clusterPartitionsCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_partitions_count",
		Help: "Cluster partitions count",
	})

	clusterSystemKvSizeBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_system_kv_size_bytes",
		Help: "Cluster system kv size bytes",
	})

	clusterTotalDiskUsedBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_total_disk_used_bytes",
		Help: "Cluster total disk used bytes",
	})

	clusterTotalKvSizeBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_total_kv_size_bytes",
		Help: "Cluster total kv size bytes",
	})

	clusterDatabaseLocked = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_database_locked",
		Help: "Cluster database locked",
	})

	clusterDegradedProcesses = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_degraded_processes",
		Help: "Cluster degraded processes",
	})

	clusterLatencyProbeBatchPriorityTransactionStartSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_latency_probe_batch_priority_transaction_start_seconds",
		Help: "Cluster latency probe batch priority transaction start seconds",
	})

	clusterLatencyProbeCommitSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_latency_probe_commit_seconds",
		Help: "Cluster latency probe commit seconds",
	})

	clusterLatencyProbeImmediatePriorityTransactionStartSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_latency_probe_immediate_priority_transaction_start_seconds",
		Help: "Cluster latency probe immediate priority transaction start seconds",
	})

	clusterLatencyProbeReadSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_latency_probe_read_seconds",
		Help: "Cluster latency probe read seconds",
	})

	clusterLatencyProbeTransactionStartSeconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_latency_probe_transaction_start_seconds",
		Help: "Cluster fdb cluster latency probe transaction start seconds",
	})

	configurationStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_cluster_configuration",
		Help: "Configuration",
	}, []string{"component"})

	generation = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_cluster_generation",
		Help: "Generation of the fdb",
	})

)

// ExportConfiguration is exporting the configuration
func (s FDBStatus) ExportConfiguration() {

	clientsCount.Set(s.Cluster.Clients.Count)
	clusterControllerTimestamp.Set(s.Cluster.ClusterControllerTimestamp)
	ClusterDataAveragePartitionSizeBytes.Set(s.Cluster.Data.AveragePartitionSizeBytes)
	clusterLeastOperatingSpaceBytesLogServer.Set(float64(s.Cluster.Data.LeastOperatingSpaceBytesLogServer))
	clusterLeastOperatingSpaceBytesStorageServer.Set(float64(s.Cluster.Data.LeastOperatingSpaceBytesStorageServer))
	clusterMovingDataHighestPriority.Set(s.Cluster.Data.MovingData.HighestPriority)

	clusterMovingDataInQueueBytes.Set(s.Cluster.Data.MovingData.InQueueBytes)
	clusterMovingDataInFlightBytes.Set(s.Cluster.Data.MovingData.InFlightBytes)
	clusterMovingDataTotalWrittenBytes.Set(s.Cluster.Data.MovingData.TotalWrittenBytes)
	clusterPartitionsCount.Set(s.Cluster.Data.PartitionsCount)
	clusterSystemKvSizeBytes.Set(s.Cluster.Data.SystemKvSizeBytes)

	clusterTotalDiskUsedBytes.Set(s.Cluster.Data.TotalDiskUsedBytes)
	clusterTotalKvSizeBytes.Set(s.Cluster.Data.TotalKvSizeBytes)

	if s.Cluster.DatabaseLocked {
		clusterDatabaseLocked.Set(1.0)
	} else {
		clusterDatabaseLocked.Set(0.0)
	}

	clusterDegradedProcesses.Set(s.Cluster.DegradedProcesses)
	clusterLatencyProbeBatchPriorityTransactionStartSeconds.Set(s.Cluster.LatencyProbe.BatchPriorityTransactionStartSeconds)
	clusterLatencyProbeCommitSeconds.Set(s.Cluster.LatencyProbe.CommitSeconds)
	clusterLatencyProbeImmediatePriorityTransactionStartSeconds.Set(s.Cluster.LatencyProbe.ImmediatePriorityTransactionStartSeconds)
	clusterLatencyProbeReadSeconds.Set(s.Cluster.LatencyProbe.ReadSeconds)
	clusterLatencyProbeTransactionStartSeconds.Set(s.Cluster.LatencyProbe.TransactionStartSeconds)

	configurationStatus.With(prometheus.Labels{
		"component": "coordinators_count",
	}).Set(s.Cluster.Configuration.CoordinatorsCount)

	configurationStatus.With(prometheus.Labels{
		"component": "excluded_servers",
	}).Set(float64(len(s.Cluster.Configuration.ExcludedServers)))

	configurationStatus.With(prometheus.Labels{
		"component": "proxies",
	}).Set(s.Cluster.Configuration.Proxies)

	configurationStatus.With(prometheus.Labels{
		"component": "logs",
	}).Set(s.Cluster.Configuration.Logs)

	configurationStatus.With(prometheus.Labels{
		"component": "log_spill",
	}).Set(s.Cluster.Configuration.LogSpill)

	configurationStatus.With(prometheus.Labels{
		"component": "resolvers",
	}).Set(s.Cluster.Configuration.Resolvers)

	configurationStatus.With(prometheus.Labels{
		"component": "usable_regions",
	}).Set(s.Cluster.Configuration.UsableRegions)

	generation.Set(s.Cluster.Generation)
}

func registerConfiguration(r *prometheus.Registry) {
	r.MustRegister(clientsCount)
	r.MustRegister(clusterControllerTimestamp)
	r.MustRegister(ClusterDataAveragePartitionSizeBytes)
	r.MustRegister(clusterLeastOperatingSpaceBytesLogServer)
	r.MustRegister(clusterLeastOperatingSpaceBytesStorageServer)
	r.MustRegister(clusterMovingDataHighestPriority)
	r.MustRegister(clusterMovingDataInQueueBytes)
	r.MustRegister(clusterMovingDataInFlightBytes)
	r.MustRegister(clusterMovingDataTotalWrittenBytes)
	r.MustRegister(clusterPartitionsCount)
	r.MustRegister(clusterSystemKvSizeBytes)
	r.MustRegister(clusterTotalDiskUsedBytes)
	r.MustRegister(clusterTotalKvSizeBytes)
	r.MustRegister(clusterDatabaseLocked)
	r.MustRegister(clusterDegradedProcesses)
	r.MustRegister(clusterLatencyProbeBatchPriorityTransactionStartSeconds)
	r.MustRegister(clusterLatencyProbeCommitSeconds)
	r.MustRegister(clusterLatencyProbeImmediatePriorityTransactionStartSeconds)
	r.MustRegister(clusterLatencyProbeReadSeconds)
	r.MustRegister(clusterLatencyProbeTransactionStartSeconds)

	r.MustRegister(configurationStatus)
	r.MustRegister(generation)
}
