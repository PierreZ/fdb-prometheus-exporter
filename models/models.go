package models

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
)

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
	Timestamp float64 `json:"timestamp"`
}

// FDBClusterStatus represents a cluster status
type FDBClusterStatus struct {
	Clients                    FDBClusterClientsStatus       `json:"clients"`
	ClusterControllerTimestamp float64                       `json:"cluster_controller_timestamp"`
	Configuration              FDBClusterConfigurationStatus `json:"configuration"`
	ConnectionString           string                        `json:"connection_string"`
	Data                       FDBClusterDataStatus          `json:"data"`
	DatabaseAvailable          bool                          `json:"database_available"`
	DatabaseLocked             bool                          `json:"database_locked"`
	DatacenterLag              struct {
		Seconds  float64 `json:"seconds"`
		Versions float64 `json:"versions"`
	} `json:"datacenter_lag"`
	DegradedProcesses float64 `json:"degraded_processes"`
	FaultTolerance    struct {
		MaxZoneFailuresWithoutLosingAvailability float64 `json:"max_zone_failures_without_losing_availability"`
		MaxZoneFailuresWithoutLosingData         float64 `json:"max_zone_failures_without_losing_data"`
	} `json:"fault_tolerance"`
	FullReplication         bool          `json:"full_replication"`
	Generation              float64       `json:"generation"`
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
		LogHitRate     float64 `json:"log_hit_rate"`
		StorageHitRate float64 `json:"storage_hit_rate"`
	} `json:"page_cache"`
	Processes       map[string]FDBClusterProcessStatus `json:"processes"`
	ProtocolVersion string                             `json:"protocol_version"`
	Qos             struct {
		BatchPerformanceLimitedBy struct {
			Description string  `json:"description"`
			Name        string  `json:"name"`
			ReasonID    float64 `json:"reason_id"`
		} `json:"batch_performance_limited_by"`
		BatchReleasedTransactionsPerSecond float64 `json:"batch_released_transactions_per_second"`
		BatchTransactionsPerSecondLimit    float64 `json:"batch_transactions_per_second_limit"`
		LimitingDataLagStorageServer       struct {
			Seconds  float64 `json:"seconds"`
			Versions float64 `json:"versions"`
		} `json:"limiting_data_lag_storage_server"`
		LimitingDurabilityLagStorageServer struct {
			Seconds  float64 `json:"seconds"`
			Versions float64 `json:"versions"`
		} `json:"limiting_durability_lag_storage_server"`
		LimitingQueueBytesStorageServer float64 `json:"limiting_queue_bytes_storage_server"`
		LimitingVersionLagStorageServer float64 `json:"limiting_version_lag_storage_server"`
		PerformanceLimitedBy            struct {
			Description string  `json:"description"`
			Name        string  `json:"name"`
			ReasonID    float64 `json:"reason_id"`
		} `json:"performance_limited_by"`
		ReleasedTransactionsPerSecond float64 `json:"released_transactions_per_second"`
		TransactionsPerSecondLimit    float64 `json:"transactions_per_second_limit"`
		WorstDataLagStorageServer     struct {
			Seconds  float64 `json:"seconds"`
			Versions float64 `json:"versions"`
		} `json:"worst_data_lag_storage_server"`
		WorstDurabilityLagStorageServer struct {
			Seconds  float64 `json:"seconds"`
			Versions float64 `json:"versions"`
		} `json:"worst_durability_lag_storage_server"`
		WorstQueueBytesLogServer     float64 `json:"worst_queue_bytes_log_server"`
		WorstQueueBytesStorageServer float64 `json:"worst_queue_bytes_storage_server"`
		WorstVersionLagStorageServer float64 `json:"worst_version_lag_storage_server"`
	} `json:"qos"`
	RecoveryState struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"recovery_state"`
	Workload struct {
		Bytes struct {
			Read struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"read"`
			Written struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"written"`
		} `json:"bytes"`
		Keys struct {
			Read struct {
				Counter   float64 `json:"counter"`
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
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_batch_priority"`
			StartedDefaultPriority struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_default_priority"`
			StartedImmediatePriority struct {
				Counter   float64 `json:"counter"`
				Hz        float64 `json:"hz"`
				Roughness float64 `json:"roughness"`
			} `json:"started_immediate_priority"`
		} `json:"transactions"`
	} `json:"workload"`
}

// FDBClusterClientsStatus represents the status of clients
type FDBClusterClientsStatus struct {
	Count             float64 `json:"count"`
	SupportedVersions []struct {
		ClientVersion    string `json:"client_version"`
		ConnectedClients []struct {
			Address  string `json:"address"`
			LogGroup string `json:"log_group"`
		} `json:"connected_clients"`
		Count           float64 `json:"count"`
		ProtocolVersion string  `json:"protocol_version"`
		SourceVersion   string  `json:"source_version"`
	} `json:"supported_versions"`
}

// FDBClusterConfigurationStatus represents a configuration status
type FDBClusterConfigurationStatus struct {
	CoordinatorsCount float64       `json:"coordinators_count"`
	ExcludedServers   []interface{} `json:"excluded_servers"`
	LogSpill          float64       `json:"log_spill"`
	Logs              float64       `json:"logs"`
	Proxies           float64       `json:"proxies"`
	RedundancyMode    string        `json:"redundancy_mode"`
	Resolvers         float64       `json:"resolvers"`
	StorageEngine     string        `json:"storage_engine"`
	UsableRegions     float64       `json:"usable_regions"`
}

// FDBClusterDataStatus represents data status
type FDBClusterDataStatus struct {
	AveragePartitionSizeBytes             float64 `json:"average_partition_size_bytes"`
	LeastOperatingSpaceBytesLogServer     int64   `json:"least_operating_space_bytes_log_server"`
	LeastOperatingSpaceBytesStorageServer int64   `json:"least_operating_space_bytes_storage_server"`
	MovingData                            struct {
		HighestPriority   float64 `json:"highest_priority"`
		InFlightBytes     float64 `json:"in_flight_bytes"`
		InQueueBytes      float64 `json:"in_queue_bytes"`
		TotalWrittenBytes float64 `json:"total_written_bytes"`
	} `json:"moving_data"`
	PartitionsCount float64 `json:"partitions_count"`
	State           struct {
		Healthy bool   `json:"healthy"`
		Name    string `json:"name"`
	} `json:"state"`
	SystemKvSizeBytes float64 `json:"system_kv_size_bytes"`
	TeamTrackers      []struct {
		InFlightBytes float64 `json:"in_flight_bytes"`
		Primary       bool    `json:"primary"`
		State         struct {
			Healthy bool   `json:"healthy"`
			Name    string `json:"name"`
		} `json:"state"`
		UnhealthyServers float64 `json:"unhealthy_servers"`
	} `json:"team_trackers"`
	TotalDiskUsedBytes float64 `json:"total_disk_used_bytes"`
	TotalKvSizeBytes   float64 `json:"total_kv_size_bytes"`
}

// FDBClusterMachineStatus represents status of the machine
type FDBClusterMachineStatus struct {
	Address             string  `json:"address"`
	ContributingWorkers float64 `json:"contributing_workers"`
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
		CommittedBytes float64 `json:"committed_bytes"`
		FreeBytes      int64   `json:"free_bytes"`
		TotalBytes     int64   `json:"total_bytes"`
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
			Counter float64 `json:"counter"`
			Hz      float64 `json:"hz"`
			Sectors float64 `json:"sectors"`
		} `json:"reads"`
		TotalBytes int64 `json:"total_bytes"`
		Writes     struct {
			Counter float64 `json:"counter"`
			Hz      float64 `json:"hz"`
			Sectors float64 `json:"sectors"`
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
		AvailableBytes        int64   `json:"available_bytes"`
		LimitBytes            int64   `json:"limit_bytes"`
		UnusedAllocatedMemory float64 `json:"unused_allocated_memory"`
		UsedBytes             float64 `json:"used_bytes"`
	} `json:"memory"`
	Messages []interface{} `json:"messages"`
	Network  struct {
		ConnectionErrors struct {
			Hz float64 `json:"hz"`
		} `json:"connection_errors"`
		ConnectionsClosed struct {
			Hz float64 `json:"hz"`
		} `json:"connections_closed"`
		ConnectionsEstablished struct {
			Hz float64 `json:"hz"`
		} `json:"connections_established"`
		CurrentConnections float64 `json:"current_connections"`
		MegabitsReceived   struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_received"`
		MegabitsSent struct {
			Hz float64 `json:"hz"`
		} `json:"megabits_sent"`
	} `json:"network"`
	Roles         []DynamicRole `json:"roles"`
	RunLoopBusy   float64       `json:"run_loop_busy"`
	UptimeSeconds float64       `json:"uptime_seconds"`
	Version       string        `json:"version"`
}

// DynamicRole is wrapping the fact that there is different roles inside a process
type DynamicRole struct {
	Value interface{}
}

// UnmarshalJSON is implementing unmarshal for DynamicRole
func (d *DynamicRole) UnmarshalJSON(data []byte) error {
	var typ struct {
		Role string `json:"role"`
	}
	if err := json.Unmarshal(data, &typ); err != nil {
		return err
	}
	switch typ.Role {
	case "log":
		d.Value = new(DynamicLogRole)
	// used by master, ratekeeper, data_distributor, proxy
	default:
		d.Value = new(DynamicEmptyStructRole)
	case "storage":
		d.Value = new(DynamicStorageRole)
	}
	return json.Unmarshal(data, d.Value)

}

// DynamicEmptyStructRole represents usess struct for metrics
type DynamicEmptyStructRole struct{}

// DynamicStorageRole represents metrics for storage role
type DynamicStorageRole struct {
	BytesQueried struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"bytes_queried"`
	DataLag struct {
		Seconds  float64 `json:"seconds"`
		Versions float64 `json:"versions"`
	} `json:"data_lag"`
	DataVersion   float64 `json:"data_version"`
	DurabilityLag struct {
		Seconds  float64 `json:"seconds"`
		Versions float64 `json:"versions"`
	} `json:"durability_lag"`
	DurableBytes struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"durable_bytes"`
	DurableVersion  float64 `json:"durable_version"`
	FinishedQueries struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"finished_queries"`
	ID         string `json:"id"`
	InputBytes struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"input_bytes"`
	KeysQueried struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"keys_queried"`
	KvstoreAvailableBytes int64   `json:"kvstore_available_bytes"`
	KvstoreFreeBytes      int64   `json:"kvstore_free_bytes"`
	KvstoreTotalBytes     int64   `json:"kvstore_total_bytes"`
	KvstoreUsedBytes      float64 `json:"kvstore_used_bytes"`
	LocalRate             float64 `json:"local_rate"`
	MutationBytes         struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"mutation_bytes"`
	Mutations struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"mutations"`
	QueryQueueMax float64 `json:"query_queue_max"`
	Role          string  `json:"role"`
	StoredBytes   float64 `json:"stored_bytes"`
	TotalQueries  struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"total_queries"`
}

// DynamicLogRole contains details about the Role log
type DynamicLogRole struct {
	Role         string  `json:"role"`
	DataVersion  float64 `json:"data_version"`
	DurableBytes struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"durable_bytes"`
	ID         string `json:"id"`
	InputBytes struct {
		Counter   float64 `json:"counter"`
		Hz        float64 `json:"hz"`
		Roughness float64 `json:"roughness"`
	} `json:"input_bytes"`
	KvstoreAvailableBytes   int64   `json:"kvstore_available_bytes"`
	KvstoreFreeBytes        int64   `json:"kvstore_free_bytes"`
	KvstoreTotalBytes       int64   `json:"kvstore_total_bytes"`
	KvstoreUsedBytes        float64 `json:"kvstore_used_bytes"`
	QueueDiskAvailableBytes int64   `json:"queue_disk_available_bytes"`
	QueueDiskFreeBytes      int64   `json:"queue_disk_free_bytes"`
	QueueDiskTotalBytes     int64   `json:"queue_disk_total_bytes"`
	QueueDiskUsedBytes      float64 `json:"queue_disk_used_bytes"`
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
