package models

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

var processLabels = []string{"process_id", "machine_id", "address", "fault_domain", "role"}
var roleLabels = []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"}
var allProcessMetrics []*prometheus.GaugeVec
var allRoleMetrics []*prometheus.GaugeVec

func ClearAll() {
	for _, metric := range allProcessMetrics {
		metric.Reset()
	}

	for _, metric := range allRoleMetrics {
		metric.Reset()
	}
}

func NewProcessGaugeVec(opts prometheus.GaugeOpts) *prometheus.GaugeVec {
	vec := prometheus.NewGaugeVec(opts, processLabels)
	allProcessMetrics = append(allProcessMetrics, vec)
	return vec
}

func NewRoleGaugeVec(opts prometheus.GaugeOpts) *prometheus.GaugeVec {
	vec := prometheus.NewGaugeVec(opts, roleLabels)
	allRoleMetrics = append(allRoleMetrics, vec)
	return vec
}

var (
	processCPUInfo = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_cpu_usage_cores",
		Help: "process cpu",
	})

	processDiskInfoBusy = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_busy",
		Help: "process disk busy",
	})

	processDiskInfoFreeBytes = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_free_bytes",
		Help: "process disk",
	})

	processDiskInfoReadHZ = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_per_second",
		Help: "process disk",
	})

	processDiskInfoReadTotal = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_total",
		Help: "process disk",
	})

	processDiskInfoReadSectors = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_sectors",
		Help: "process disk reads sectors",
	})

	processDiskInfoTotalBytes = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_total_bytes",
		Help: "process disk",
	})

	processDiskInfoWriteHZ = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_per_second",
		Help: "process disk",
	})

	processDiskInfoWriteSectors = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_sectors",
		Help: "process disk writes sectors",
	})

	processDiskInfoWritesTotal = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_total",
		Help: "process disk",
	})

	processMemoryInfoAvailableBytes = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_available_bytes",
		Help: "process memory available_bytes",
	})

	processMemoryInfoLimitBytes = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_limit_bytes",
		Help: "process memory info",
	})

	processMemoryInfoUnusedMemory = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_unused_allocated_memory",
		Help: "process memory info",
	})

	processMemoryInfoUsedBytes = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_used_bytes",
		Help: "process memory info",
	})

	processNetworkConnectionError = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_errors_per_second",
		Help: "process network errors",
	})

	processNetworkConnectionsClosed = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_closed_per_second",
		Help: "process network closed",
	})

	processNetworkConnectionsEstablished = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_established_per_second",
		Help: "process network established",
	})

	processNetworkCurrentConnections = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_current_connections",
		Help: "process network current connections",
	})

	processNetworkReceived = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_received_per_second",
		Help: "process network mega bit",
	})

	processNetworkSent = NewProcessGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_sent_per_second",
		Help: "process network current connections",
	})

	logRoleDataVersion = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_data_version",
		Help: "process log data version",
	})

	logRoleDurableBytesHZ = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_per_second",
		Help: "process log data version",
	})

	logRoleDurableBytesCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_total",
		Help: "process log data version",
	})

	logRoleAvailableBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_available_bytes",
		Help: "process log data version",
	})

	logRoleFreeBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_free_bytes",
		Help: "process log data version",
	})

	logRoleTotalBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_total_bytes",
		Help: "process log data version",
	})

	logRoleUsedBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_used_bytes",
		Help: "process log data version",
	})

	logRoleQueueAvailableBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_available_bytes",
		Help: "process log data version",
	})

	logRoleQueueFreeBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_free_bytes",
		Help: "process log data version",
	})

	logRoleQueueTotalBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_total_bytes",
		Help: "process log data version",
	})

	logRoleQueueUsedbytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_used_bytes",
		Help: "process log data version",
	})

	storageRoleBytesQueriedHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleBytesQueriedCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_total",
		Help: "process storage bytes_queried",
	})

	storageRoleDataLagSeconds = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_seconds",
		Help: "process storage bytes_queried",
	})

	storageRoleDataLagVersions = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_versions",
		Help: "process storage bytes_queried",
	})

	storageRoleDurabilityLagSeconds = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_seconds",
		Help: "process storage bytes_queried",
	})

	storageRoleDurabilityLagVersions = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_versions",
		Help: "process storage bytes_queried",
	})

	storageRoleDurableVersions = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durable_version",
		Help: "process storage bytes_queried",
	})

	storageRoleFinishedQueriesCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_total",
		Help: "process storage bytes_queried",
	})

	storageRoleFinishedQueriesHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleInputBytesCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_total",
		Help: "process storage bytes_queried",
	})

	storageRoleInputBytesHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleKeysQueriedCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_total",
		Help: "process storage bytes_queried",
	})

	storageRoleKeysQueriedHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleMutationsCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_total",
		Help: "process storage bytes_queried",
	})

	storageRoleMutationHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleMutationsBytesCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_total",
		Help: "process storage bytes_queried",
	})

	storageRoleMutationBytesHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleTotalQueriesCounter = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_total",
		Help: "process storage bytes_queried",
	})

	storageRoleTotalQueriesHz = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_per_second",
		Help: "process storage bytes_queried",
	})

	storageRoleQueryMaxQueue = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_query_queue_max",
		Help: "process storage bytes_queried",
	})

	storageRoleStoredBytes = NewRoleGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_stored_bytes",
		Help: "process storage bytes_queried",
	})
)

// ExportProcesses is exporting the configuration
func (s FDBStatus) ExportProcesses() {
	ClearAll()

	for process, info := range s.Cluster.Processes {

		var roleNames []string

		for _, role := range info.Roles {
			switch r := role.Value.(type) {
			case *DynamicLogRole:
				roleNames = append(roleNames, r.Role)
			case *DynamicStorageRole:
				roleNames = append(roleNames, r.Role)
			case *DynamicEmptyStructRole:
				roleNames = append(roleNames, r.Role)
			default: // nothing to expose
			}
		}

		labels := prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"role":         "," + strings.Join(roleNames, ",") + ",",
		}

		processCPUInfo.With(labels).Set(info.CPU.UsageCores)

		processDiskInfoBusy.With(labels).Set(info.Disk.Busy)

		processDiskInfoTotalBytes.With(labels).Set(float64(info.Disk.TotalBytes))

		processDiskInfoFreeBytes.With(labels).Set(float64(info.Disk.FreeBytes))

		processDiskInfoReadHZ.With(labels).Set(info.Disk.Reads.Hz)

		processDiskInfoReadTotal.With(labels).Set(info.Disk.Reads.Counter)

		processDiskInfoReadSectors.With(labels).Set(info.Disk.Reads.Sectors)

		processDiskInfoWritesTotal.With(labels).Set(info.Disk.Writes.Counter)

		processDiskInfoWriteHZ.With(labels).Set(info.Disk.Writes.Hz)

		processDiskInfoWriteSectors.With(labels).Set(info.Disk.Writes.Sectors)

		processMemoryInfoAvailableBytes.With(labels).Set(float64(info.Memory.AvailableBytes))

		processMemoryInfoLimitBytes.With(labels).Set(float64(info.Memory.LimitBytes))

		processMemoryInfoUnusedMemory.With(labels).Set(info.Memory.UnusedAllocatedMemory)

		processMemoryInfoUsedBytes.With(labels).Set(info.Memory.UsedBytes)

		// network
		processNetworkConnectionError.With(labels).Set(info.Network.ConnectionErrors.Hz)

		processNetworkConnectionsClosed.With(labels).Set(info.Network.ConnectionsClosed.Hz)

		processNetworkConnectionsEstablished.With(labels).Set(info.Network.ConnectionsEstablished.Hz)

		processNetworkCurrentConnections.With(labels).Set(info.Network.CurrentConnections)

		processNetworkReceived.With(labels).Set(info.Network.MegabitsReceived.Hz)

		processNetworkSent.With(labels).Set(info.Network.MegabitsSent.Hz)

		for _, role := range info.Roles {
			switch r := role.Value.(type) {
			case *DynamicLogRole:
				labels := prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         "," + r.Role + ",",
				}
				logRoleDataVersion.With(labels).Set(r.DataVersion)
				logRoleDurableBytesHZ.With(labels).Set(r.DurableBytes.Hz)
				logRoleDurableBytesCounter.With(labels).Set(r.DurableBytes.Counter)
				logRoleAvailableBytes.With(labels).Set(float64(r.KvstoreAvailableBytes))
				logRoleFreeBytes.With(labels).Set(float64(r.KvstoreFreeBytes))
				logRoleTotalBytes.With(labels).Set(float64(r.KvstoreTotalBytes))
				logRoleUsedBytes.With(labels).Set(r.KvstoreUsedBytes)
				logRoleQueueAvailableBytes.With(labels).Set(float64(r.QueueDiskAvailableBytes))
				logRoleQueueFreeBytes.With(labels).Set(float64(r.QueueDiskFreeBytes))
				logRoleQueueTotalBytes.With(labels).Set(float64(r.QueueDiskTotalBytes))
				logRoleQueueUsedbytes.With(labels).Set(r.QueueDiskUsedBytes)

			case *DynamicStorageRole:
				labels := prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         "," + r.Role + ",",
				}

				storageRoleBytesQueriedCounter.With(labels).Set(r.BytesQueried.Counter)
				storageRoleBytesQueriedHz.With(labels).Set(r.BytesQueried.Hz)
				storageRoleDataLagSeconds.With(labels).Set(r.DataLag.Seconds)
				storageRoleDataLagVersions.With(labels).Set(r.DataLag.Versions)
				storageRoleDurabilityLagSeconds.With(labels).Set(r.DurabilityLag.Seconds)
				storageRoleDurabilityLagVersions.With(labels).Set(r.DurabilityLag.Versions)

				logRoleDataVersion.With(labels).Set(r.DataVersion)
				logRoleDurableBytesHZ.With(labels).Set(r.DurableBytes.Hz)
				logRoleDurableBytesCounter.With(labels).Set(r.DurableBytes.Counter)
				storageRoleDurableVersions.With(labels).Set(r.DurableVersion)
				storageRoleFinishedQueriesCounter.With(labels).Set(r.FinishedQueries.Counter)
				storageRoleFinishedQueriesHz.With(labels).Set(r.FinishedQueries.Hz)
				storageRoleInputBytesCounter.With(labels).Set(r.InputBytes.Counter)
				storageRoleInputBytesHz.With(labels).Set(r.InputBytes.Hz)
				storageRoleKeysQueriedCounter.With(labels).Set(r.KeysQueried.Counter)
				storageRoleKeysQueriedHz.With(labels).Set(r.KeysQueried.Hz)
				logRoleQueueAvailableBytes.With(labels).Set(float64(r.KvstoreAvailableBytes))
				logRoleQueueFreeBytes.With(labels).Set(float64(r.KvstoreFreeBytes))
				logRoleQueueTotalBytes.With(labels).Set(float64(r.KvstoreTotalBytes))
				logRoleQueueUsedbytes.With(labels).Set(r.KvstoreUsedBytes)
				storageRoleMutationBytesHz.With(labels).Set(r.MutationBytes.Hz)
				storageRoleMutationsBytesCounter.With(labels).Set(r.MutationBytes.Counter)
				storageRoleMutationHz.With(labels).Set(r.Mutations.Hz)
				storageRoleMutationsCounter.With(labels).Set(r.Mutations.Counter)
				storageRoleQueryMaxQueue.With(labels).Set(r.QueryQueueMax)
				storageRoleStoredBytes.With(labels).Set(r.StoredBytes)
				storageRoleTotalQueriesCounter.With(labels).Set(r.TotalQueries.Counter)
				storageRoleTotalQueriesHz.With(labels).Set(r.TotalQueries.Hz)
			default: // nothing to expose
			}
		}
	}
}

func registerProcesses(r *prometheus.Registry) {
	for _, metric := range allProcessMetrics {
		r.MustRegister(metric)
	}

	for _, metric := range allRoleMetrics {
		r.MustRegister(metric)
	}
}
