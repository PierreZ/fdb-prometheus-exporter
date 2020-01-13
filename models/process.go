package models

import "github.com/prometheus/client_golang/prometheus"

var (
	processCPUInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_cpu_usage_cores",
		Help: "process cpu",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "class_type"})

	processDiskInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_disk_info",
		Help: "process disk",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "class_type", "info"})

	processMemoryInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_memory_info",
		Help: "process memory info",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "class_type", "info"})

	processNetworkInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_processes_network_info",
		Help: "process network info",
	}, []string{"process_id", "machine_id", "address", "fault_domain", "class_type", "info"})
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

		processDiskInfo.With(prometheus.Labels{
			"machine_id":   info.Locality.Machineid,
			"process_id":   process,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "busy",
		}).Set(float64(info.Disk.Busy))
		processDiskInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "free_bytes",
		}).Set(float64(info.Disk.FreeBytes))
		processDiskInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "reads_total",
		}).Set(float64(info.Disk.Reads.Counter))
		processDiskInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "reads_per_second",
		}).Set(float64(info.Disk.Reads.Hz))
		processDiskInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "writes_total",
		}).Set(float64(info.Disk.Writes.Counter))
		processDiskInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "writes_per_second",
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

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "connection_errors_per_second",
		}).Set(float64(info.Network.ConnectionErrors.Hz))

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "connections_closed",
		}).Set(float64(info.Network.ConnectionsClosed.Hz))

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "connections_established",
		}).Set(float64(info.Network.ConnectionsEstablished.Hz))

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "current_connections",
		}).Set(float64(info.Network.CurrentConnections))

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "megabits_received_per_second",
		}).Set(float64(info.Network.MegabitsReceived.Hz))

		processNetworkInfo.With(prometheus.Labels{
			"process_id":   process,
			"machine_id":   info.Locality.Machineid,
			"address":      info.Address,
			"fault_domain": info.FaultDomain,
			"class_type":   info.ClassType,
			"info":         "megabits_send_per_second",
		}).Set(float64(info.Network.MegabitsSent.Hz))
	}
}

func registerProcesses(r *prometheus.Registry) {
	r.MustRegister(processCPUInfo)
	r.MustRegister(processDiskInfo)
	r.MustRegister(processMemoryInfo)
	r.MustRegister(processNetworkInfo)
}
