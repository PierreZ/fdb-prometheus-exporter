package models

import "github.com/prometheus/client_golang/prometheus"

// FDBStatus represent the status of a Cluster
type FDBStatus struct {
	Client  FDBClientStatus  `json:"client"`
	Cluster FDBClusterStatus `json:"cluster"`
}

// FDBClientStatus represent status of the clients
type FDBClientStatus struct {
	ClusterFile struct {
		Path     string `json:"path"`
		UpToDate bool   `json:"up_to_date"`
	} `json:"cluster_file"`
	Coordinators struct {
		Coordinators []struct {
			Address   string `json:"address"`
			Reachable bool   `json:"reachable"`
		} `json:"coordinators"`
		QuorumReachable bool `json:"quorum_reachable"`
	} `json:"coordinators"`
	DatabaseStatus struct {
		Available bool `json:"available"`
		Healthy   bool `json:"healthy"`
	} `json:"database_status"`
	Timestamp int `json:"timestamp"`
}

// FDBClusterStatus represents a cluster status
type FDBClusterStatus struct {
	Clients                    FDBClusterClientsStatus       `json:"clients"`
	ClusterControllerTimestamp int                           `json:"cluster_controller_timestamp"`
	Configuration              FDBClusterConfigurationStatus `json:"configuration"`
	ConnectionString           string                        `json:"connection_string"`
	Data                       FDBClusterDataStatus          `json:"data"`
	DatabaseAvailable          bool                          `json:"database_available"`
	DatabaseLocked             bool                          `json:"database_locked"`
	DatacenterLag              struct {
		Seconds  int `json:"seconds"`
		Versions int `json:"versions"`
	} `json:"datacenter_lag"`
	DegradedProcesses int `json:"degraded_processes"`
	FaultTolerance    struct {
		MaxZoneFailuresWithoutLosingAvailability int `json:"max_zone_failures_without_losing_availability"`
		MaxZoneFailuresWithoutLosingData         int `json:"max_zone_failures_without_losing_data"`
	} `json:"fault_tolerance"`
	FullReplication         bool          `json:"full_replication"`
	Generation              int           `json:"generation"`
	IncompatibleConnections []interface{} `json:"incompatible_connections"`
	LatencyProbe            struct {
		BatchPriorityTransactionStartSeconds     float64 `json:"batch_priority_transaction_start_seconds"`
		CommitSeconds                            float64 `json:"commit_seconds"`
		ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
		ReadSeconds                              float64 `json:"read_seconds"`
		TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
	} `json:"latency_probe"`
	Layers struct {
		Valid bool `json:"_valid"`
	} `json:"layers"`
	Machines  map[string]FDBClusterMachineStatus `json:"machines"`
	PageCache struct {
		LogHitRate     int `json:"log_hit_rate"`
		StorageHitRate int `json:"storage_hit_rate"`
	} `json:"page_cache"`
	Processes       map[string]FDBClusterProcessStatus `json:"processes"`
	ProtocolVersion string                             `json:"protocol_version"`
	Qos             struct {
		BatchPerformanceLimitedBy struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			ReasonID    int    `json:"reason_id"`
		} `json:"batch_performance_limited_by"`
		BatchReleasedTransactionsPerSecond float64 `json:"batch_released_transactions_per_second"`
		BatchTransactionsPerSecondLimit    int     `json:"batch_transactions_per_second_limit"`
		LimitingDataLagStorageServer       struct {
			Seconds  int `json:"seconds"`
			Versions int `json:"versions"`
		} `json:"limiting_data_lag_storage_server"`
		LimitingDurabilityLagStorageServer struct {
			Seconds  float64 `json:"seconds"`
			Versions int     `json:"versions"`
		} `json:"limiting_durability_lag_storage_server"`
		LimitingQueueBytesStorageServer int `json:"limiting_queue_bytes_storage_server"`
		LimitingVersionLagStorageServer int `json:"limiting_version_lag_storage_server"`
		PerformanceLimitedBy            struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			ReasonID    int    `json:"reason_id"`
		} `json:"performance_limited_by"`
		ReleasedTransactionsPerSecond float64 `json:"released_transactions_per_second"`
		TransactionsPerSecondLimit    int     `json:"transactions_per_second_limit"`
		WorstDataLagStorageServer     struct {
			Seconds  int `json:"seconds"`
			Versions int `json:"versions"`
		} `json:"worst_data_lag_storage_server"`
		WorstDurabilityLagStorageServer struct {
			Seconds  float64 `json:"seconds"`
			Versions int     `json:"versions"`
		} `json:"worst_durability_lag_storage_server"`
		WorstQueueBytesLogServer     int `json:"worst_queue_bytes_log_server"`
		WorstQueueBytesStorageServer int `json:"worst_queue_bytes_storage_server"`
		WorstVersionLagStorageServer int `json:"worst_version_lag_storage_server"`
	} `json:"qos"`
	RecoveryState struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"recovery_state"`
	Workload struct {
		Bytes struct {
			Read struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"read"`
			Written struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"written"`
		} `json:"bytes"`
		Keys struct {
			Read struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"read"`
		} `json:"keys"`
		Operations struct {
			ReadRequests struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"read_requests"`
			Reads struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"reads"`
			Writes struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"writes"`
		} `json:"operations"`
		Transactions struct {
			Committed struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"committed"`
			Conflicted struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"conflicted"`
			Started struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started"`
			StartedBatchPriority struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_batch_priority"`
			StartedDefaultPriority struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_default_priority"`
			StartedImmediatePriority struct {
				Counter   int     `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_immediate_priority"`
		} `json:"transactions"`
	} `json:"workload"`
}

// FDBClusterClientsStatus represents the status of clients
type FDBClusterClientsStatus struct {
	Count             int `json:"count"`
	SupportedVersions []struct {
		ClientVersion    string `json:"client_version"`
		ConnectedClients []struct {
			Address  string `json:"address"`
			LogGroup string `json:"log_group"`
		} `json:"connected_clients"`
		Count           int    `json:"count"`
		ProtocolVersion string `json:"protocol_version"`
		SourceVersion   string `json:"source_version"`
	} `json:"supported_versions"`
}

// FDBClusterConfigurationStatus represents a configuration status
type FDBClusterConfigurationStatus struct {
	CoordinatorsCount int           `json:"coordinators_count"`
	ExcludedServers   []interface{} `json:"excluded_servers"`
	LogSpill          int           `json:"log_spill"`
	Logs              int           `json:"logs"`
	Proxies           int           `json:"proxies"`
	RedundancyMode    string        `json:"redundancy_mode"`
	Resolvers         int           `json:"resolvers"`
	StorageEngine     string        `json:"storage_engine"`
	UsableRegions     int           `json:"usable_regions"`
}

// FDBClusterDataStatus represents data status
type FDBClusterDataStatus struct {
	AveragePartitionSizeBytes             int   `json:"average_partition_size_bytes"`
	LeastOperatingSpaceBytesLogServer     int64 `json:"least_operating_space_bytes_log_server"`
	LeastOperatingSpaceBytesStorageServer int64 `json:"least_operating_space_bytes_storage_server"`
	MovingData                            struct {
		HighestPriority   int `json:"highest_priority"`
		InFlightBytes     int `json:"in_flight_bytes"`
		InQueueBytes      int `json:"in_queue_bytes"`
		TotalWrittenBytes int `json:"total_written_bytes"`
	} `json:"moving_data"`
	PartitionsCount int `json:"partitions_count"`
	State           struct {
		Healthy bool   `json:"healthy"`
		Name    string `json:"name"`
	} `json:"state"`
	SystemKvSizeBytes int `json:"system_kv_size_bytes"`
	TeamTrackers      []struct {
		InFlightBytes int  `json:"in_flight_bytes"`
		Primary       bool `json:"primary"`
		State         struct {
			Healthy bool   `json:"healthy"`
			Name    string `json:"name"`
		} `json:"state"`
		UnhealthyServers int `json:"unhealthy_servers"`
	} `json:"team_trackers"`
	TotalDiskUsedBytes int `json:"total_disk_used_bytes"`
	TotalKvSizeBytes   int `json:"total_kv_size_bytes"`
}

// FDBClusterMachineStatus represents status of the machine
type FDBClusterMachineStatus struct {
	Address             string `json:"address"`
	ContributingWorkers int    `json:"contributing_workers"`
	CPU                 struct {
		LogicalCoreUtilization float64 `json:"logical_core_utilization"`
	} `json:"cpu"`
	Excluded bool `json:"excluded"`
	Locality struct {
		InstanceID string `json:"instance_id"`
		Machineid  string `json:"machineid"`
		Processid  string `json:"processid"`
		Zoneid     string `json:"zoneid"`
	} `json:"locality"`
	MachineID string `json:"machine_id"`
	Memory    struct {
		CommittedBytes int   `json:"committed_bytes"`
		FreeBytes      int64 `json:"free_bytes"`
		TotalBytes     int64 `json:"total_bytes"`
	} `json:"memory"`
	Network struct {
		MegabitsReceived struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_received"`
		MegabitsSent struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_sent"`
		TCPSegmentsRetransmitted struct {
			Hz float64 `json:"hz"`
		} `json:"tcp_segments_retransmitted"`
	} `json:"network"`
}

// FDBClusterProcessStatus represents a Process status
type FDBClusterProcessStatus struct {
	Address     string `json:"address"`
	ClassSource string `json:"class_source"`
	ClassType   string `json:"class_type"`
	CommandLine string `json:"command_line"`
	CPU         struct {
		UsageCores float64 `json:"usage_cores"`
	} `json:"cpu"`
	Disk struct {
		Busy      float64 `json:"busy"`
		FreeBytes int64   `json:"free_bytes"`
		Reads     struct {
			Counter int `json:"counter"`
			Hz      int `json:"hz"`
			Sectors int `json:"sectors"`
		} `json:"reads"`
		TotalBytes int64 `json:"total_bytes"`
		Writes     struct {
			Counter int     `json:"counter"`
			Hz      float64 `json:"hz"`
			Sectors int     `json:"sectors"`
		} `json:"writes"`
	} `json:"disk"`
	Excluded    bool   `json:"excluded"`
	FaultDomain string `json:"fault_domain"`
	Locality    struct {
		InstanceID string `json:"instance_id"`
		Machineid  string `json:"machineid"`
		Processid  string `json:"processid"`
		Zoneid     string `json:"zoneid"`
	} `json:"locality"`
	MachineID string `json:"machine_id"`
	Memory    struct {
		AvailableBytes        int64 `json:"available_bytes"`
		LimitBytes            int64 `json:"limit_bytes"`
		UnusedAllocatedMemory int   `json:"unused_allocated_memory"`
		UsedBytes             int   `json:"used_bytes"`
	} `json:"memory"`
	Messages []interface{} `json:"messages"`
	Network  struct {
		ConnectionErrors struct {
			Hz int `json:"hz"`
		} `json:"connection_errors"`
		ConnectionsClosed struct {
			Hz int `json:"hz"`
		} `json:"connections_closed"`
		ConnectionsEstablished struct {
			Hz int `json:"hz"`
		} `json:"connections_established"`
		CurrentConnections int `json:"current_connections"`
		MegabitsReceived   struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_received"`
		MegabitsSent struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_sent"`
	} `json:"network"`
	Roles []struct {
		DataVersion  int `json:"data_version"`
		DurableBytes struct {
			Counter   int     `json:"counter"`
			Hz        float64 `json:"hz"`
			Roughness float64 `json:"roughness"`
		} `json:"durable_bytes"`
		ID         string `json:"id"`
		InputBytes struct {
			Counter   int     `json:"counter"`
			Hz        float64 `json:"hz"`
			Roughness float64 `json:"roughness"`
		} `json:"input_bytes"`
		KvstoreAvailableBytes   int64  `json:"kvstore_available_bytes"`
		KvstoreFreeBytes        int64  `json:"kvstore_free_bytes"`
		KvstoreTotalBytes       int64  `json:"kvstore_total_bytes"`
		KvstoreUsedBytes        int    `json:"kvstore_used_bytes"`
		QueueDiskAvailableBytes int64  `json:"queue_disk_available_bytes"`
		QueueDiskFreeBytes      int64  `json:"queue_disk_free_bytes"`
		QueueDiskTotalBytes     int64  `json:"queue_disk_total_bytes"`
		QueueDiskUsedBytes      int    `json:"queue_disk_used_bytes"`
		Role                    string `json:"role"`
	} `json:"roles"`
	RunLoopBusy   float64 `json:"run_loop_busy"`
	UptimeSeconds float64 `json:"uptime_seconds"`
	Version       string  `json:"version"`
}

func boolToNumber(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

// Register is registering all metrics
func Register(r *prometheus.Registry) {
	registerWorkload(r)
	registerDatabaseStatus(r)
	registerConfiguration(r)
	registerProcesses(r)
}
