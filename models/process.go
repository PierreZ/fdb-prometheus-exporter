package models

import "github.com/prometheus/client_golang/prometheus"

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

	processMemoryInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_info",
		Help: "process memory info",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "class_type", "info"})

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
)

// ExportProcesses is exporting the configuration
func (s FDBStatus) ExportProcesses() {
	for process, info := range s.Cluster.Processes {
		processCPUInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
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
			"class_type":   info.ClassType,
			"info":         "reads_per_second",
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

		processMemoryInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "available_bytes",
		}).Set(float64(info.Memory.AvailableBytes))
		processMemoryInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "limit_bytes",
		}).Set(float64(info.Memory.LimitBytes))
		processMemoryInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "unused_allocated_memory",
		}).Set(float64(info.Memory.UnusedAllocatedMemory))
		processMemoryInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "used_bytes",
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
			case DynamicStorageRole:
				exposeStorageRoleMetrics(&r)
			case DynamicLogRole:
				exposeLogRoleMetrics(&r)
			default: // nothing to expose
			}
		}
	}
}

func exposeStorageRoleMetrics(r *DynamicStorageRole) {}
func exposeLogRoleMetrics(r *DynamicLogRole)         {}

func registerProcesses(r *prometheus.Registry) {
	r.MustRegister(processCPUInfo)
	r.MustRegister(processDiskInfoBusy)
	r.MustRegister(processDiskInfoFreeBytes)
	r.MustRegister(processDiskInfoReadHZ)
	r.MustRegister(processDiskInfoReadTotal)
	r.MustRegister(processDiskInfoTotalBytes)
	r.MustRegister(processDiskInfoWriteHZ)
	r.MustRegister(processDiskInfoWritesTotal)
	r.MustRegister(processMemoryInfo)
	r.MustRegister(processNetworkConnectionError)
	r.MustRegister(processNetworkConnectionsClosed)
	r.MustRegister(processNetworkConnectionsEstablished)
	r.MustRegister(processNetworkCurrentConnection)
	r.MustRegister(processNetworkReceived)
	r.MustRegister(processNetworkSent)
}
