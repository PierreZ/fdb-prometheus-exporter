package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	jsonRaw := `
	{
    "client" : {
        "cluster_file" : {
            "path" : "/var/dynamic-conf/fdb.cluster",
            "up_to_date" : true
        },
        "coordinators" : {
            "coordinators" : [
                {
                    "address" : "10.2.2.30:4501",
                    "reachable" : true
                },
                {
                    "address" : "10.2.3.30:4501",
                    "reachable" : true
                },
                {
                    "address" : "10.2.4.27:4501",
                    "reachable" : true
                }
            ],
            "quorum_reachable" : true
        },
        "database_status" : {
            "available" : true,
            "healthy" : true
        },
        "messages" : [
        ],
        "timestamp" : 1578668384
    },
    "cluster" : {
        "clients" : {
            "count" : 0,
            "supported_versions" : [
                {
                    "client_version" : "Unknown",
                    "connected_clients" : [
                        {
                            "address" : "10.2.4.26:48308",
                            "log_group" : "default"
                        },
                        {
                            "address" : "10.2.4.26:48384",
                            "log_group" : "default"
                        },
                        {
                            "address" : "10.2.4.26:56916",
                            "log_group" : "default"
                        },
                        {
                            "address" : "10.2.4.26:57056",
                            "log_group" : "default"
                        }
                    ],
                    "count" : 4,
                    "protocol_version" : "Unknown",
                    "source_version" : "Unknown"
                }
            ]
        },
        "cluster_controller_timestamp" : 1578668384,
        "configuration" : {
            "coordinators_count" : 3,
            "excluded_servers" : [
            ],
            "log_spill" : 2,
            "logs" : 3,
            "proxies" : 3,
            "redundancy_mode" : "double",
            "resolvers" : 1,
            "storage_engine" : "ssd-2",
            "usable_regions" : 1
        },
        "connection_string" : "sample_cluster:CZUvWIZHegWxQRZt9ckPWhkT78cafekd@10.2.2.30:4501,10.2.3.30:4501,10.2.4.27:4501",
        "data" : {
            "average_partition_size_bytes" : 400000,
            "least_operating_space_bytes_log_server" : 119895255450,
            "least_operating_space_bytes_storage_server" : 119895701914,
            "moving_data" : {
                "highest_priority" : 0,
                "in_flight_bytes" : 0,
                "in_queue_bytes" : 0,
                "total_written_bytes" : 0
            },
            "partitions_count" : 1,
            "state" : {
                "healthy" : true,
                "name" : "healthy"
            },
            "system_kv_size_bytes" : 0,
            "team_trackers" : [
                {
                    "in_flight_bytes" : 0,
                    "primary" : true,
                    "state" : {
                        "healthy" : true,
                        "name" : "healthy"
                    },
                    "unhealthy_servers" : 0
                }
            ],
            "total_disk_used_bytes" : 629211712,
            "total_kv_size_bytes" : 0
        },
        "database_available" : true,
        "database_locked" : false,
        "datacenter_lag" : {
            "seconds" : 0,
            "versions" : 0
        },
        "degraded_processes" : 1,
        "fault_tolerance" : {
            "max_zone_failures_without_losing_availability" : 1,
            "max_zone_failures_without_losing_data" : 1
        },
        "full_replication" : true,
        "generation" : 8,
        "incompatible_connections" : [
        ],
        "latency_probe" : {
            "batch_priority_transaction_start_seconds" : 0.0024537999999999999,
            "commit_seconds" : 0.0079002399999999993,
            "immediate_priority_transaction_start_seconds" : 0.0026771999999999998,
            "read_seconds" : 0.00055336999999999997,
            "transaction_start_seconds" : 0.0020234599999999999
        },
        "layers" : {
            "_valid" : true
        },
        "machines" : {
            "sample-cluster-log-1" : {
                "address" : "10.2.4.26",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.060161799999999994
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "log-1",
                    "machineid" : "sample-cluster-log-1",
                    "processid" : "3557393b2b9cba82e00acc8a80c42728",
                    "zoneid" : "sample-cluster-log-1"
                },
                "machine_id" : "sample-cluster-log-1",
                "memory" : {
                    "committed_bytes" : 1415610368,
                    "free_bytes" : 60423430144,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.35690699999999997
                    },
                    "megabits_sent" : {
                        "hz" : 0.27507399999999999
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0.39999799999999996
                    }
                }
            },
            "sample-cluster-log-2" : {
                "address" : "10.2.2.31",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.070415400000000003
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "log-2",
                    "machineid" : "sample-cluster-log-2",
                    "processid" : "00dc41f9fac4f1ce5222357eae850975",
                    "zoneid" : "sample-cluster-log-2"
                },
                "machine_id" : "sample-cluster-log-2",
                "memory" : {
                    "committed_bytes" : 1404481536,
                    "free_bytes" : 60434558976,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.133797
                    },
                    "megabits_sent" : {
                        "hz" : 0.110703
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            },
            "sample-cluster-log-3" : {
                "address" : "10.2.4.25",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.065283399999999991
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "log-3",
                    "machineid" : "sample-cluster-log-3",
                    "processid" : "c3cc80b22b816085907438acf1aceb4a",
                    "zoneid" : "sample-cluster-log-3"
                },
                "machine_id" : "sample-cluster-log-3",
                "memory" : {
                    "committed_bytes" : 1431285760,
                    "free_bytes" : 60407754752,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.15018000000000001
                    },
                    "megabits_sent" : {
                        "hz" : 0.11064099999999999
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            },
            "sample-cluster-log-4" : {
                "address" : "10.2.2.29",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.075012700000000002
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "log-4",
                    "machineid" : "sample-cluster-log-4",
                    "processid" : "3b70694250738fb03962edebcc59a188",
                    "zoneid" : "sample-cluster-log-4"
                },
                "machine_id" : "sample-cluster-log-4",
                "memory" : {
                    "committed_bytes" : 1404489728,
                    "free_bytes" : 60434550784,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.28031299999999998
                    },
                    "megabits_sent" : {
                        "hz" : 0.288578
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            },
            "sample-cluster-storage-1" : {
                "address" : "10.2.4.27",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.070886099999999994
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "storage-1",
                    "machineid" : "sample-cluster-storage-1",
                    "processid" : "9a0185ba9488b34d57e1c2a8c5c8e695",
                    "zoneid" : "sample-cluster-storage-1"
                },
                "machine_id" : "sample-cluster-storage-1",
                "memory" : {
                    "committed_bytes" : 1415729152,
                    "free_bytes" : 60423311360,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.089889499999999997
                    },
                    "megabits_sent" : {
                        "hz" : 0.080081699999999992
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            },
            "sample-cluster-storage-2" : {
                "address" : "10.2.2.30",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.076026399999999994
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "storage-2",
                    "machineid" : "sample-cluster-storage-2",
                    "processid" : "20de3b05fb1792d316e8e31bb90c705e",
                    "zoneid" : "sample-cluster-storage-2"
                },
                "machine_id" : "sample-cluster-storage-2",
                "memory" : {
                    "committed_bytes" : 1404776448,
                    "free_bytes" : 60434264064,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.21668000000000001
                    },
                    "megabits_sent" : {
                        "hz" : 0.29699500000000001
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            },
            "sample-cluster-storage-3" : {
                "address" : "10.2.3.30",
                "contributing_workers" : 1,
                "cpu" : {
                    "logical_core_utilization" : 0.0521784
                },
                "excluded" : false,
                "locality" : {
                    "instance_id" : "storage-3",
                    "machineid" : "sample-cluster-storage-3",
                    "processid" : "3e37581b04d8bab7ea5a9db8a96e1c26",
                    "zoneid" : "sample-cluster-storage-3"
                },
                "machine_id" : "sample-cluster-storage-3",
                "memory" : {
                    "committed_bytes" : 1361547264,
                    "free_bytes" : 60477493248,
                    "total_bytes" : 61839040512
                },
                "network" : {
                    "megabits_received" : {
                        "hz" : 0.211727
                    },
                    "megabits_sent" : {
                        "hz" : 0.277111
                    },
                    "tcp_segments_retransmitted" : {
                        "hz" : 0
                    }
                }
            }
        },
        "messages" : [
        ],
        "page_cache" : {
            "log_hit_rate" : 1,
            "storage_hit_rate" : 1
        },
        "processes" : {
            "00dc41f9fac4f1ce5222357eae850975" : {
                "address" : "10.2.2.31:4501",
                "class_source" : "command_line",
                "class_type" : "log",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=log --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=log-2 --locality_machineid=sample-cluster-log-2 --locality_zoneid=sample-cluster-log-2 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.2.31:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.019802699999999999
                },
                "disk" : {
                    "busy" : 0.00079997300000000005,
                    "free_bytes" : 126209888256,
                    "reads" : {
                        "counter" : 299,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 7349,
                        "hz" : 2.7999000000000001,
                        "sectors" : 72
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-log-2",
                "locality" : {
                    "instance_id" : "log-2",
                    "machineid" : "sample-cluster-log-2",
                    "processid" : "00dc41f9fac4f1ce5222357eae850975",
                    "zoneid" : "sample-cluster-log-2"
                },
                "machine_id" : "sample-cluster-log-2",
                "memory" : {
                    "available_bytes" : 60924321792,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 262144,
                    "used_bytes" : 489762816
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 6,
                    "megabits_received" : {
                        "hz" : 0.061450699999999997
                    },
                    "megabits_sent" : {
                        "hz" : 0.059581999999999996
                    }
                },
                "roles" : [
                    {
                        "data_version" : 1498074936,
                        "durable_bytes" : {
                            "counter" : 3208,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "id" : "10d5cff9225c3e5d",
                        "input_bytes" : {
                            "counter" : 3208,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "kvstore_available_bytes" : 126209888256,
                        "kvstore_free_bytes" : 126209888256,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104861752,
                        "queue_disk_available_bytes" : 126209888256,
                        "queue_disk_free_bytes" : 126209888256,
                        "queue_disk_total_bytes" : 126290034688,
                        "queue_disk_used_bytes" : 389120,
                        "role" : "log"
                    }
                ],
                "run_loop_busy" : 0.018677699999999998,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "20de3b05fb1792d316e8e31bb90c705e" : {
                "address" : "10.2.2.30:4501",
                "class_source" : "command_line",
                "class_type" : "storage",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=storage --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=storage-2 --locality_machineid=sample-cluster-storage-2 --locality_zoneid=sample-cluster-storage-2 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.2.30:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.026497199999999999
                },
                "disk" : {
                    "busy" : 0.00059998600000000005,
                    "free_bytes" : 126210203648,
                    "reads" : {
                        "counter" : 346,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 6278,
                        "hz" : 1.59996,
                        "sectors" : 48
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-storage-2",
                "locality" : {
                    "instance_id" : "storage-2",
                    "machineid" : "sample-cluster-storage-2",
                    "processid" : "20de3b05fb1792d316e8e31bb90c705e",
                    "zoneid" : "sample-cluster-storage-2"
                },
                "machine_id" : "sample-cluster-storage-2",
                "memory" : {
                    "available_bytes" : 60924354560,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 393216,
                    "used_bytes" : 490090496
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 7,
                    "megabits_received" : {
                        "hz" : 0.133245
                    },
                    "megabits_sent" : {
                        "hz" : 0.173538
                    }
                },
                "roles" : [
                    {
                        "role" : "coordinator"
                    },
                    {
                        "id" : "39ed477f5f6fb541",
                        "role" : "proxy"
                    },
                    {
                        "bytes_queried" : {
                            "counter" : 4202739,
                            "hz" : 4117.9700000000003,
                            "roughness" : 10777.1
                        },
                        "data_lag" : {
                            "seconds" : 0.13192000000000001,
                            "versions" : 131920
                        },
                        "data_version" : 1498074936,
                        "durability_lag" : {
                            "seconds" : 5,
                            "versions" : 5000000
                        },
                        "durable_bytes" : {
                            "counter" : 326788,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "durable_version" : 1493074936,
                        "finished_queries" : {
                            "counter" : 4811,
                            "hz" : 4.5999699999999999,
                            "roughness" : 7.3227799999999998
                        },
                        "id" : "e33e24a9e1e5bfaf",
                        "input_bytes" : {
                            "counter" : 328792,
                            "hz" : 400.79700000000003,
                            "roughness" : 341.46899999999999
                        },
                        "keys_queried" : {
                            "counter" : 12658,
                            "hz" : 12.1999,
                            "roughness" : 31.9282
                        },
                        "kvstore_available_bytes" : 126210203648,
                        "kvstore_free_bytes" : 126210203648,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104882352,
                        "local_rate" : 100,
                        "mutation_bytes" : {
                            "counter" : 19408,
                            "hz" : 18.799900000000001,
                            "roughness" : 16.016999999999999
                        },
                        "mutations" : {
                            "counter" : 326,
                            "hz" : 0.39999699999999999,
                            "roughness" : 0.34078799999999998
                        },
                        "query_queue_max" : 7,
                        "role" : "storage",
                        "stored_bytes" : 0,
                        "total_queries" : {
                            "counter" : 4811,
                            "hz" : 4.5999699999999999,
                            "roughness" : 7.3241300000000003
                        }
                    }
                ],
                "run_loop_busy" : 0.0279983,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "3557393b2b9cba82e00acc8a80c42728" : {
                "address" : "10.2.4.26:4501",
                "class_source" : "command_line",
                "class_type" : "log",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=log --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=log-1 --locality_machineid=sample-cluster-log-1 --locality_zoneid=sample-cluster-log-1 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.4.26:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.032181899999999999
                },
                "disk" : {
                    "busy" : 0,
                    "free_bytes" : 126210310144,
                    "reads" : {
                        "counter" : 314,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 2265,
                        "hz" : 0,
                        "sectors" : 0
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-log-1",
                "locality" : {
                    "instance_id" : "log-1",
                    "machineid" : "sample-cluster-log-1",
                    "processid" : "3557393b2b9cba82e00acc8a80c42728",
                    "zoneid" : "sample-cluster-log-1"
                },
                "machine_id" : "sample-cluster-log-1",
                "memory" : {
                    "available_bytes" : 60905426944,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 0,
                    "used_bytes" : 481996800
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 6,
                    "megabits_received" : {
                        "hz" : 0.24632199999999999
                    },
                    "megabits_sent" : {
                        "hz" : 0.151001
                    }
                },
                "roles" : [
                    {
                        "id" : "8a0c75bfdf539905",
                        "role" : "master"
                    },
                    {
                        "id" : "5fc32f117d14da3d",
                        "role" : "data_distributor"
                    },
                    {
                        "id" : "42d20f327cac78af",
                        "role" : "ratekeeper"
                    }
                ],
                "run_loop_busy" : 0.032473499999999995,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "3b70694250738fb03962edebcc59a188" : {
                "address" : "10.2.2.29:4501",
                "class_source" : "command_line",
                "class_type" : "log",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=log --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=log-4 --locality_machineid=sample-cluster-log-4 --locality_zoneid=sample-cluster-log-4 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.2.29:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.029500499999999999
                },
                "degraded" : true,
                "disk" : {
                    "busy" : 0.00019999300000000001,
                    "free_bytes" : 126209757184,
                    "reads" : {
                        "counter" : 267,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 5413,
                        "hz" : 2.3999100000000002,
                        "sectors" : 80
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-log-4",
                "locality" : {
                    "instance_id" : "log-4",
                    "machineid" : "sample-cluster-log-4",
                    "processid" : "3b70694250738fb03962edebcc59a188",
                    "zoneid" : "sample-cluster-log-4"
                },
                "machine_id" : "sample-cluster-log-4",
                "memory" : {
                    "available_bytes" : 60934885376,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 262144,
                    "used_bytes" : 500334592
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 6,
                    "megabits_received" : {
                        "hz" : 0.155169
                    },
                    "megabits_sent" : {
                        "hz" : 0.208735
                    }
                },
                "roles" : [
                    {
                        "id" : "82d9b65b21cefbc5",
                        "role" : "cluster_controller"
                    },
                    {
                        "data_version" : 1498074936,
                        "durable_bytes" : {
                            "counter" : 27129,
                            "hz" : 37.798699999999997,
                            "roughness" : 2.7762500000000001
                        },
                        "id" : "c930bde271c36c5d",
                        "input_bytes" : {
                            "counter" : 27318,
                            "hz" : 37.798699999999997,
                            "roughness" : 128.91499999999999
                        },
                        "kvstore_available_bytes" : 126209757184,
                        "kvstore_free_bytes" : 126209757184,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104861752,
                        "queue_disk_available_bytes" : 126209757184,
                        "queue_disk_free_bytes" : 126209757184,
                        "queue_disk_total_bytes" : 126290034688,
                        "queue_disk_used_bytes" : 520192,
                        "role" : "log"
                    }
                ],
                "run_loop_busy" : 0.031006499999999999,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "3e37581b04d8bab7ea5a9db8a96e1c26" : {
                "address" : "10.2.3.30:4501",
                "class_source" : "command_line",
                "class_type" : "storage",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=storage --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=storage-3 --locality_machineid=sample-cluster-storage-3 --locality_zoneid=sample-cluster-storage-3 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.3.30:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.027837199999999999
                },
                "disk" : {
                    "busy" : 0.000399992,
                    "free_bytes" : 126210203648,
                    "reads" : {
                        "counter" : 301,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 6258,
                        "hz" : 1.5999699999999999,
                        "sectors" : 48
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-storage-3",
                "locality" : {
                    "instance_id" : "storage-3",
                    "machineid" : "sample-cluster-storage-3",
                    "processid" : "3e37581b04d8bab7ea5a9db8a96e1c26",
                    "zoneid" : "sample-cluster-storage-3"
                },
                "machine_id" : "sample-cluster-storage-3",
                "memory" : {
                    "available_bytes" : 60966985728,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 393216,
                    "used_bytes" : 489492480
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 7,
                    "megabits_received" : {
                        "hz" : 0.13040399999999999
                    },
                    "megabits_sent" : {
                        "hz" : 0.15809000000000001
                    }
                },
                "roles" : [
                    {
                        "role" : "coordinator"
                    },
                    {
                        "id" : "104e7c19429ac76d",
                        "role" : "proxy"
                    },
                    {
                        "bytes_queried" : {
                            "counter" : 4212196,
                            "hz" : 2743.1700000000001,
                            "roughness" : 7931.04
                        },
                        "data_lag" : {
                            "seconds" : 0.13192000000000001,
                            "versions" : 131920
                        },
                        "data_version" : 1498074936,
                        "durability_lag" : {
                            "seconds" : 5,
                            "versions" : 5000000
                        },
                        "durable_bytes" : {
                            "counter" : 326788,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "durable_version" : 1493074936,
                        "finished_queries" : {
                            "counter" : 4753,
                            "hz" : 3.19997,
                            "roughness" : 6.8238200000000004
                        },
                        "id" : "de71929e80cd9d92",
                        "input_bytes" : {
                            "counter" : 328792,
                            "hz" : 400.79599999999999,
                            "roughness" : 390.90300000000002
                        },
                        "keys_queried" : {
                            "counter" : 12703,
                            "hz" : 7.9999200000000004,
                            "roughness" : 23.129300000000001
                        },
                        "kvstore_available_bytes" : 126210203648,
                        "kvstore_free_bytes" : 126210203648,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104882352,
                        "local_rate" : 100,
                        "mutation_bytes" : {
                            "counter" : 19408,
                            "hz" : 18.799800000000001,
                            "roughness" : 18.335799999999999
                        },
                        "mutations" : {
                            "counter" : 326,
                            "hz" : 0.39999599999999996,
                            "roughness" : 0.390123
                        },
                        "query_queue_max" : 7,
                        "role" : "storage",
                        "stored_bytes" : 0,
                        "total_queries" : {
                            "counter" : 4753,
                            "hz" : 3.19997,
                            "roughness" : 6.82545
                        }
                    }
                ],
                "run_loop_busy" : 0.028872599999999998,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "9a0185ba9488b34d57e1c2a8c5c8e695" : {
                "address" : "10.2.4.27:4501",
                "class_source" : "command_line",
                "class_type" : "storage",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=storage --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=storage-1 --locality_machineid=sample-cluster-storage-1 --locality_zoneid=sample-cluster-storage-1 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.4.27:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.010756
                },
                "disk" : {
                    "busy" : 0.00059998899999999999,
                    "free_bytes" : 126210236416,
                    "reads" : {
                        "counter" : 314,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 6031,
                        "hz" : 1.5999699999999999,
                        "sectors" : 48
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-storage-1",
                "locality" : {
                    "instance_id" : "storage-1",
                    "machineid" : "sample-cluster-storage-1",
                    "processid" : "9a0185ba9488b34d57e1c2a8c5c8e695",
                    "zoneid" : "sample-cluster-storage-1"
                },
                "machine_id" : "sample-cluster-storage-1",
                "memory" : {
                    "available_bytes" : 60845629440,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 131072,
                    "used_bytes" : 422318080
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 7,
                    "megabits_received" : {
                        "hz" : 0.054501399999999998
                    },
                    "megabits_sent" : {
                        "hz" : 0.046488700000000001
                    }
                },
                "roles" : [
                    {
                        "role" : "coordinator"
                    },
                    {
                        "bytes_queried" : {
                            "counter" : 0,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "data_lag" : {
                            "seconds" : 0.040578999999999997,
                            "versions" : 40579
                        },
                        "data_version" : 1501932896,
                        "durability_lag" : {
                            "seconds" : 5.0405800000000003,
                            "versions" : 5040579
                        },
                        "durable_bytes" : {
                            "counter" : 6940,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "durable_version" : 1496892317,
                        "finished_queries" : {
                            "counter" : 0,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "id" : "a9479a3206af3fc2",
                        "input_bytes" : {
                            "counter" : 6940,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "keys_queried" : {
                            "counter" : 0,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "kvstore_available_bytes" : 126210236416,
                        "kvstore_free_bytes" : 126210236416,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104861752,
                        "local_rate" : 100,
                        "mutation_bytes" : {
                            "counter" : 298,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "mutations" : {
                            "counter" : 7,
                            "hz" : 0,
                            "roughness" : 0
                        },
                        "query_queue_max" : 0,
                        "role" : "storage",
                        "stored_bytes" : 0,
                        "total_queries" : {
                            "counter" : 0,
                            "hz" : 0,
                            "roughness" : 0
                        }
                    },
                    {
                        "id" : "5e54228556e15bea",
                        "role" : "resolver"
                    }
                ],
                "run_loop_busy" : 0.010498,
                "uptime_seconds" : 1210.03,
                "version" : "6.2.11"
            },
            "c3cc80b22b816085907438acf1aceb4a" : {
                "address" : "10.2.4.25:4501",
                "class_source" : "command_line",
                "class_type" : "log",
                "command_line" : "/var/dynamic-conf/bin/6.2.11/fdbserver --class=log --cluster_file=/var/fdb/data/fdb.cluster --datadir=/var/fdb/data --knob_disable_posix_kernel_aio=1 --locality_instance_id=log-3 --locality_machineid=sample-cluster-log-3 --locality_zoneid=sample-cluster-log-3 --logdir=/var/log/fdb-trace-logs --loggroup=sample-cluster --public_address=10.2.4.25:4501 --seed_cluster_file=/var/dynamic-conf/fdb.cluster",
                "cpu" : {
                    "usage_cores" : 0.017015300000000001
                },
                "disk" : {
                    "busy" : 0.00079998600000000003,
                    "free_bytes" : 126209888256,
                    "reads" : {
                        "counter" : 267,
                        "hz" : 0,
                        "sectors" : 0
                    },
                    "total_bytes" : 126290034688,
                    "writes" : {
                        "counter" : 5345,
                        "hz" : 1.9999600000000002,
                        "sectors" : 48
                    }
                },
                "excluded" : false,
                "fault_domain" : "sample-cluster-log-3",
                "locality" : {
                    "instance_id" : "log-3",
                    "machineid" : "sample-cluster-log-3",
                    "processid" : "c3cc80b22b816085907438acf1aceb4a",
                    "zoneid" : "sample-cluster-log-3"
                },
                "machine_id" : "sample-cluster-log-3",
                "memory" : {
                    "available_bytes" : 60906913792,
                    "limit_bytes" : 8589934592,
                    "unused_allocated_memory" : 131072,
                    "used_bytes" : 499159040
                },
                "messages" : [
                ],
                "network" : {
                    "connection_errors" : {
                        "hz" : 0
                    },
                    "connections_closed" : {
                        "hz" : 0
                    },
                    "connections_established" : {
                        "hz" : 0
                    },
                    "current_connections" : 6,
                    "megabits_received" : {
                        "hz" : 0.077528299999999994
                    },
                    "megabits_sent" : {
                        "hz" : 0.059954199999999999
                    }
                },
                "roles" : [
                    {
                        "data_version" : 1498074936,
                        "durable_bytes" : {
                            "counter" : 24217,
                            "hz" : 37.798699999999997,
                            "roughness" : 1.76953
                        },
                        "id" : "08a4454d01a54f50",
                        "input_bytes" : {
                            "counter" : 24406,
                            "hz" : 37.798699999999997,
                            "roughness" : 128.51400000000001
                        },
                        "kvstore_available_bytes" : 126209888256,
                        "kvstore_free_bytes" : 126209888256,
                        "kvstore_total_bytes" : 126290034688,
                        "kvstore_used_bytes" : 104861752,
                        "queue_disk_available_bytes" : 126209888256,
                        "queue_disk_free_bytes" : 126209888256,
                        "queue_disk_total_bytes" : 126290034688,
                        "queue_disk_used_bytes" : 389120,
                        "role" : "log"
                    }
                ],
                "run_loop_busy" : 0.017715499999999999,
                "uptime_seconds" : 1210.02,
                "version" : "6.2.11"
            }
        },
        "protocol_version" : "fdb00b062010001",
        "qos" : {
            "batch_performance_limited_by" : {
                "description" : "The database is not being saturated by the workload.",
                "name" : "workload",
                "reason_id" : 2
            },
            "batch_released_transactions_per_second" : 1.95399e-14,
            "batch_transactions_per_second_limit" : 16298200,
            "limiting_data_lag_storage_server" : {
                "seconds" : 0,
                "versions" : 0
            },
            "limiting_durability_lag_storage_server" : {
                "seconds" : 13.8788,
                "versions" : 13878840
            },
            "limiting_queue_bytes_storage_server" : 2004,
            "limiting_version_lag_storage_server" : 0,
            "performance_limited_by" : {
                "description" : "The database is not being saturated by the workload.",
                "name" : "workload",
                "reason_id" : 2
            },
            "released_transactions_per_second" : 2.33127,
            "transactions_per_second_limit" : 22566700,
            "worst_data_lag_storage_server" : {
                "seconds" : 0,
                "versions" : 0
            },
            "worst_durability_lag_storage_server" : {
                "seconds" : 13.9,
                "versions" : 13899982
            },
            "worst_queue_bytes_log_server" : 189,
            "worst_queue_bytes_storage_server" : 2004,
            "worst_version_lag_storage_server" : 0
        },
        "recovery_state" : {
            "description" : "Recovery complete.",
            "name" : "fully_recovered"
        },
        "workload" : {
            "bytes" : {
                "read" : {
                    "counter" : 8414935,
                    "hz" : 6861.1400000000003,
                    "roughness" : 9639.2099999999991
                },
                "written" : {
                    "counter" : 9345,
                    "hz" : 13.999700000000001,
                    "roughness" : 48.8598
                }
            },
            "keys" : {
                "read" : {
                    "counter" : 25361,
                    "hz" : 20.1998,
                    "roughness" : 28.4435
                }
            },
            "operations" : {
                "read_requests" : {
                    "counter" : 9564,
                    "hz" : 7.7999400000000003,
                    "roughness" : 7.1195399999999998
                },
                "reads" : {
                    "counter" : 9564,
                    "hz" : 7.7999400000000003,
                    "roughness" : 7.11808
                },
                "writes" : {
                    "counter" : 248,
                    "hz" : 0.39999299999999999,
                    "roughness" : 1.3959900000000001
                }
            },
            "transactions" : {
                "committed" : {
                    "counter" : 141,
                    "hz" : 0.19999599999999998,
                    "roughness" : 0.700465
                },
                "conflicted" : {
                    "counter" : 0,
                    "hz" : 0,
                    "roughness" : 0
                },
                "started" : {
                    "counter" : 4352,
                    "hz" : 3.3999299999999999,
                    "roughness" : 3.58616
                },
                "started_batch_priority" : {
                    "counter" : 18,
                    "hz" : 0,
                    "roughness" : 0
                },
                "started_default_priority" : {
                    "counter" : 2908,
                    "hz" : 2.1999499999999999,
                    "roughness" : 3.2202600000000001
                },
                "started_immediate_priority" : {
                    "counter" : 1426,
                    "hz" : 1.19997,
                    "roughness" : 1.47241
                }
            }
        }
    }
}
	`
	var status FDBStatus
	err := json.NewDecoder(strings.NewReader(jsonRaw)).Decode(&status)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v\n", status)
}
