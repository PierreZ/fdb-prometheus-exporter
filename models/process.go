package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	processCPUInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_cpu_usage_cores",
		Help: "process cpu",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoBusy = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_busy",
		Help: "process disk busy",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_free_bytes",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoReadHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_per_second",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoReadTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_reads_total",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_total_bytes",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoWriteHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_per_second",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processDiskInfoWritesTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_writes_total",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processMemoryInfoAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_available_bytes",
		Help: "process memory available_bytes",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processMemoryInfoLimitBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_limit_bytes",
		Help: "process memory info",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processMemoryInfoUnusedMemory = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_unused_allocated_memory",
		Help: "process memory info",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processMemoryInfoUsedBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_used_bytes",
		Help: "process memory info",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkConnectionError = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_errors_per_second",
		Help: "process network errors",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkConnectionsClosed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_closed_per_second",
		Help: "process network closed",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkConnectionsEstablished = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_connection_established_per_second",
		Help: "process network established",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkCurrentConnection = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_current_connections_per_second",
		Help: "process netcowrk current connections",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkReceived = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_received_per_second",
		Help: "process netcowrk mega bit",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	processNetworkSent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_megabits_sent_per_second",
		Help: "process netcowrk current connections",
	}, []string{"process_id", "machine_id", "address", "fault_domain"})

	logRoleDataVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_data_version",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleDurableBytesHZ = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_per_second",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleDurableBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_durable_bytes_total",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_available_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_free_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_total_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleUsedBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_kvstore_used_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleQueueAvailableBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_available_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleQueueFreeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_free_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleQueueTotalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_total_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	logRoleQueueUsedbytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_queue_disk_used_bytes",
		Help: "process log data version",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleBytesQueriedHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleBytesQueriedCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_bytes_queried_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleDataLagSeconds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_seconds",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleDataLagVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_data_lag_versions",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleDurabilityLagSeconds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_seconds",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleDurabilityLagVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durability_lag_versions",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleDurableVersions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_durable_version",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleFinishedQueriesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleFinishedQueriesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_finished_queries_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleInputBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleInputBytesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_input_bytes_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleKeysQueriedCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleKeysQueriedHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_keys_queried_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleMutationsCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleMutationHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleMutationsBytesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleMutationBytesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_mutations_bytes_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleTotalQueriesCounter = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_total",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleTotalQueriesHz = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_queries_per_second",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleQueryMaxQueue = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_query_queue_max",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})

	storageRoleStoredBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_storage_stored_bytes",
		Help: "process storage bytes_queried",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "id", "role"})
)

// ExportProcesses is exporting the configuration
func (s FDBStatus) ExportProcesses() {
	for process, info := range s.Cluster.Processes {
		processCPUInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.CPU.UsageCores))

		processDiskInfoBusy.With(prometheus.Labels{
			"machine_id":   info.Locality.Machineid,
			"process_id":   process,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.Busy))

		processDiskInfoFreeBytes.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.FreeBytes))

		processDiskInfoReadHZ.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.Reads.Hz))

		processDiskInfoReadTotal.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.Reads.Counter))

		processDiskInfoWritesTotal.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.Writes.Counter))

		processDiskInfoWriteHZ.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Disk.Writes.Hz))

		processMemoryInfoAvailableBytes.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Memory.AvailableBytes))

		processMemoryInfoLimitBytes.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Memory.LimitBytes))

		processMemoryInfoUnusedMemory.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Memory.UnusedAllocatedMemory))

		processMemoryInfoUsedBytes.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Memory.UsedBytes))

		// network
		processNetworkConnectionError.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.ConnectionErrors.Hz))

		processNetworkConnectionsClosed.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.ConnectionsClosed.Hz))

		processNetworkConnectionsEstablished.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.ConnectionsEstablished.Hz))

		processNetworkCurrentConnection.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.CurrentConnections))

		processNetworkReceived.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.MegabitsReceived.Hz))

		processNetworkSent.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
		}).Set(float64(info.Network.MegabitsSent.Hz))

		for _, role := range info.Roles {
			switch r := role.Value.(type) {
			case *DynamicLogRole:
				logRoleDataVersion.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DataVersion))
				logRoleDurableBytesHZ.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DurableBytes.Hz))
				logRoleDurableBytesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DurableBytes.Counter))

				logRoleAvailableBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreAvailableBytes))

				logRoleFreeBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreFreeBytes))

				logRoleTotalBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreTotalBytes))

				logRoleUsedBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreUsedBytes))

				logRoleQueueAvailableBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.QueueDiskAvailableBytes))

				logRoleQueueFreeBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.QueueDiskFreeBytes))

				logRoleQueueTotalBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.QueueDiskTotalBytes))

				logRoleQueueUsedbytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.QueueDiskUsedBytes))

			case *DynamicStorageRole:
				storageRoleBytesQueriedCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.BytesQueried.Counter))
				storageRoleBytesQueriedHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.BytesQueried.Hz))

				storageRoleDataLagSeconds.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DataLag.Seconds))

				storageRoleDataLagVersions.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DataLag.Versions))

				logRoleDataVersion.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DataVersion))

				logRoleDurableBytesHZ.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DurableBytes.Hz))

				logRoleDurableBytesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DurableBytes.Counter))

				storageRoleDurableVersions.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.DurableVersion))

				storageRoleFinishedQueriesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.FinishedQueries.Counter))

				storageRoleFinishedQueriesHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.FinishedQueries.Hz))

				storageRoleInputBytesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.InputBytes.Counter))

				storageRoleInputBytesHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.InputBytes.Hz))

				storageRoleKeysQueriedCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KeysQueried.Counter))

				storageRoleKeysQueriedHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KeysQueried.Hz))

				logRoleQueueAvailableBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreAvailableBytes))

				logRoleQueueFreeBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreFreeBytes))

				logRoleQueueTotalBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreTotalBytes))

				logRoleQueueUsedbytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.KvstoreUsedBytes))

				storageRoleMutationBytesHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.MutationBytes.Hz))

				storageRoleMutationsBytesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.MutationBytes.Counter))

				storageRoleMutationHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.Mutations.Hz))

				storageRoleMutationsCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.Mutations.Counter))

				storageRoleQueryMaxQueue.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.QueryQueueMax))

				storageRoleStoredBytes.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.StoredBytes))

				storageRoleTotalQueriesCounter.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.TotalQueries.Counter))

				storageRoleTotalQueriesHz.With(prometheus.Labels{
					"process_id":   process,
					"machine_id":   info.Locality.Machineid,
					"address":      info.Address,
					"fault_domain": info.FaultDomain,
					"id":           r.ID,
					"role":         r.Role,
				}).Set(float64(r.TotalQueries.Hz))
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
