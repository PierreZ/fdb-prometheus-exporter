package models

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

var processLabels = []string{"process_id", "machine_id", "address", "fault_domain", "role"}
var roleLabels = []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"}

var (
	processCPUInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_cpu_usage_cores",
		Help: "process cpu",
	}, processLabels)

	processDiskInfoBusy = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_busy",
		Help: "process disk busy",
	}, processLabels)

	processDiskInfoFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_free_bytes",
		Help: "process disk",
	}, processLabels)

	processDiskInfoReadHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_per_second",
		Help: "process disk",
	}, processLabels)

	processDiskInfoReadTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_total",
		Help: "process disk",
	}, processLabels)

	processDiskInfoTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_total_bytes",
		Help: "process disk",
	}, processLabels)

	processDiskInfoWriteHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_per_second",
		Help: "process disk",
	}, processLabels)

	processDiskInfoWritesTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_total",
		Help: "process disk",
	}, processLabels)

	processMemoryInfoAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_available_bytes",
		Help: "process memory available_bytes",
	}, processLabels)

	processMemoryInfoLimitBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_limit_bytes",
		Help: "process memory info",
	}, processLabels)

	processMemoryInfoUnusedMemory = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_unused_allocated_memory",
		Help: "process memory info",
	}, processLabels)

	processMemoryInfoUsedBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_used_bytes",
		Help: "process memory info",
	}, processLabels)

	processNetworkConnectionError = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_errors_per_second",
		Help: "process network errors",
	}, processLabels)

	processNetworkConnectionsClosed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_closed_per_second",
		Help: "process network closed",
	}, processLabels)

	processNetworkConnectionsEstablished = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_established_per_second",
		Help: "process network established",
	}, processLabels)

	processNetworkCurrentConnection = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_current_connections_per_second",
		Help: "process network current connections",
	}, processLabels)

	processNetworkReceived = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_received_per_second",
		Help: "process network mega bit",
	}, processLabels)

	processNetworkSent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_sent_per_second",
		Help: "process network current connections",
	}, processLabels)

	logRoleDataVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_data_version",
		Help: "process log data version",
	}, roleLabels)

	logRoleDurableBytesHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_per_second",
		Help: "process log data version",
	}, roleLabels)

	logRoleDurableBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_total",
		Help: "process log data version",
	}, roleLabels)

	logRoleAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_available_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_free_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_total_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleUsedBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_used_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleQueueAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_available_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleQueueFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_free_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleQueueTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_total_bytes",
		Help: "process log data version",
	}, roleLabels)

	logRoleQueueUsedbytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_used_bytes",
		Help: "process log data version",
	}, roleLabels)

	storageRoleBytesQueriedHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleBytesQueriedCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleDataLagSeconds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_seconds",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleDataLagVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_versions",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleDurabilityLagSeconds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_seconds",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleDurabilityLagVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_versions",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleDurableVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durable_version",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleFinishedQueriesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleFinishedQueriesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleInputBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleInputBytesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleKeysQueriedCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleKeysQueriedHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleMutationsCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleMutationHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleMutationsBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleMutationBytesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleTotalQueriesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_total",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleTotalQueriesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_per_second",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleQueryMaxQueue = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_query_queue_max",
		Help: "process storage bytes_queried",
	}, roleLabels)

	storageRoleStoredBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_stored_bytes",
		Help: "process storage bytes_queried",
	}, roleLabels)
)

// ExportProcesses is exporting the configuration
func (s FDBStatus) ExportProcesses() {
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

		processCPUInfo.With(labels).Set(float64(info.CPU.UsageCores))

		processDiskInfoBusy.With(labels).Set(float64(info.Disk.Busy))

		processDiskInfoFreeBytes.With(labels).Set(float64(info.Disk.FreeBytes))

		processDiskInfoReadHZ.With(labels).Set(float64(info.Disk.Reads.Hz))

		processDiskInfoReadTotal.With(labels).Set(float64(info.Disk.Reads.Counter))

		processDiskInfoWritesTotal.With(labels).Set(float64(info.Disk.Writes.Counter))

		processDiskInfoWriteHZ.With(labels).Set(float64(info.Disk.Writes.Hz))

		processMemoryInfoAvailableBytes.With(labels).Set(float64(info.Memory.AvailableBytes))

		processMemoryInfoLimitBytes.With(labels).Set(float64(info.Memory.LimitBytes))

		processMemoryInfoUnusedMemory.With(labels).Set(float64(info.Memory.UnusedAllocatedMemory))

		processMemoryInfoUsedBytes.With(labels).Set(float64(info.Memory.UsedBytes))

		// network
		processNetworkConnectionError.With(labels).Set(float64(info.Network.ConnectionErrors.Hz))

		processNetworkConnectionsClosed.With(labels).Set(float64(info.Network.ConnectionsClosed.Hz))

		processNetworkConnectionsEstablished.With(labels).Set(float64(info.Network.ConnectionsEstablished.Hz))

		processNetworkCurrentConnection.With(labels).Set(float64(info.Network.CurrentConnections))

		processNetworkReceived.With(labels).Set(float64(info.Network.MegabitsReceived.Hz))

		processNetworkSent.With(labels).Set(float64(info.Network.MegabitsSent.Hz))

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
				logRoleDataVersion.With(labels).Set(float64(r.DataVersion))
				logRoleDurableBytesHZ.With(labels).Set(float64(r.DurableBytes.Hz))
				logRoleDurableBytesCounter.With(labels).Set(float64(r.DurableBytes.Counter))
				logRoleAvailableBytes.With(labels).Set(float64(r.KvstoreAvailableBytes))
				logRoleFreeBytes.With(labels).Set(float64(r.KvstoreFreeBytes))
				logRoleTotalBytes.With(labels).Set(float64(r.KvstoreTotalBytes))
				logRoleUsedBytes.With(labels).Set(float64(r.KvstoreUsedBytes))
				logRoleQueueAvailableBytes.With(labels).Set(float64(r.QueueDiskAvailableBytes))
				logRoleQueueFreeBytes.With(labels).Set(float64(r.QueueDiskFreeBytes))
				logRoleQueueTotalBytes.With(labels).Set(float64(r.QueueDiskTotalBytes))
				logRoleQueueUsedbytes.With(labels).Set(float64(r.QueueDiskUsedBytes))

			case *DynamicStorageRole:
				labels := prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         "," + r.Role + ",",
				}

				storageRoleBytesQueriedCounter.With(labels).Set(float64(r.BytesQueried.Counter))
				storageRoleBytesQueriedHz.With(labels).Set(float64(r.BytesQueried.Hz))
				storageRoleDataLagSeconds.With(labels).Set(float64(r.DataLag.Seconds))
				storageRoleDataLagVersions.With(labels).Set(float64(r.DataLag.Versions))
				logRoleDataVersion.With(labels).Set(float64(r.DataVersion))
				logRoleDurableBytesHZ.With(labels).Set(float64(r.DurableBytes.Hz))
				logRoleDurableBytesCounter.With(labels).Set(float64(r.DurableBytes.Counter))
				storageRoleDurableVersions.With(labels).Set(float64(r.DurableVersion))
				storageRoleFinishedQueriesCounter.With(labels).Set(float64(r.FinishedQueries.Counter))
				storageRoleFinishedQueriesHz.With(labels).Set(float64(r.FinishedQueries.Hz))
				storageRoleInputBytesCounter.With(labels).Set(float64(r.InputBytes.Counter))
				storageRoleInputBytesHz.With(labels).Set(float64(r.InputBytes.Hz))
				storageRoleKeysQueriedCounter.With(labels).Set(float64(r.KeysQueried.Counter))
				storageRoleKeysQueriedHz.With(labels).Set(float64(r.KeysQueried.Hz))
				logRoleQueueAvailableBytes.With(labels).Set(float64(r.KvstoreAvailableBytes))
				logRoleQueueFreeBytes.With(labels).Set(float64(r.KvstoreFreeBytes))
				logRoleQueueTotalBytes.With(labels).Set(float64(r.KvstoreTotalBytes))
				logRoleQueueUsedbytes.With(labels).Set(float64(r.KvstoreUsedBytes))
				storageRoleMutationBytesHz.With(labels).Set(float64(r.MutationBytes.Hz))
				storageRoleMutationsBytesCounter.With(labels).Set(float64(r.MutationBytes.Counter))
				storageRoleMutationHz.With(labels).Set(float64(r.Mutations.Hz))
				storageRoleMutationsCounter.With(labels).Set(float64(r.Mutations.Counter))
				storageRoleQueryMaxQueue.With(labels).Set(float64(r.QueryQueueMax))
				storageRoleStoredBytes.With(labels).Set(float64(r.StoredBytes))
				storageRoleTotalQueriesCounter.With(labels).Set(float64(r.TotalQueries.Counter))
				storageRoleTotalQueriesHz.With(labels).Set(float64(r.TotalQueries.Hz))
			default: // nothing to expose
			}
		}
	}
}

func registerProcesses(r *prometheus.Registry) {
	r.MustRegister(processCPUInfo)
	r.MustRegister(processDiskInfoBusy)
	r.MustRegister(processDiskInfoFreeBytes)
	r.MustRegister(processDiskInfoReadHZ)
	r.MustRegister(processDiskInfoReadTotal)
	r.MustRegister(processDiskInfoTotalBytes)
	r.MustRegister(processDiskInfoWriteHZ)
	r.MustRegister(processDiskInfoWritesTotal)
	r.MustRegister(processMemoryInfoAvailableBytes)
	r.MustRegister(processMemoryInfoLimitBytes)
	r.MustRegister(processMemoryInfoUsedBytes)
	r.MustRegister(processMemoryInfoUnusedMemory)
	r.MustRegister(processNetworkConnectionError)
	r.MustRegister(processNetworkConnectionsClosed)
	r.MustRegister(processNetworkConnectionsEstablished)
	r.MustRegister(processNetworkCurrentConnection)
	r.MustRegister(processNetworkReceived)
	r.MustRegister(processNetworkSent)
	r.MustRegister(logRoleAvailableBytes)
	r.MustRegister(logRoleDataVersion)
	r.MustRegister(logRoleDurableBytesCounter)
	r.MustRegister(logRoleDurableBytesHZ)
	r.MustRegister(logRoleFreeBytes)
	r.MustRegister(logRoleQueueAvailableBytes)
	r.MustRegister(logRoleQueueFreeBytes)
	r.MustRegister(logRoleQueueTotalBytes)
	r.MustRegister(logRoleQueueUsedbytes)
	r.MustRegister(logRoleTotalBytes)
	r.MustRegister(logRoleUsedBytes)

	r.MustRegister(storageRoleBytesQueriedCounter)
	r.MustRegister(storageRoleBytesQueriedHz)
	r.MustRegister(storageRoleDataLagSeconds)
	r.MustRegister(storageRoleDataLagVersions)
	r.MustRegister(storageRoleDurabilityLagSeconds)
	r.MustRegister(storageRoleDurabilityLagVersions)
	r.MustRegister(storageRoleDurableVersions)
	r.MustRegister(storageRoleFinishedQueriesCounter)
	r.MustRegister(storageRoleFinishedQueriesHz)
	r.MustRegister(storageRoleInputBytesCounter)
	r.MustRegister(storageRoleInputBytesHz)
	r.MustRegister(storageRoleKeysQueriedCounter)
	r.MustRegister(storageRoleKeysQueriedHz)
	r.MustRegister(storageRoleMutationBytesHz)
	r.MustRegister(storageRoleMutationHz)
	r.MustRegister(storageRoleMutationsBytesCounter)
	r.MustRegister(storageRoleMutationsCounter)
	r.MustRegister(storageRoleQueryMaxQueue)
	r.MustRegister(storageRoleStoredBytes)
	r.MustRegister(storageRoleTotalQueriesCounter)
	r.MustRegister(storageRoleTotalQueriesHz)
}
