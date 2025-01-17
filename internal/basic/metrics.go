// Code generated by go generate; DO NOT EDIT.
package basic

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = []Metric{
	{
		Name: "ActiveTransactions",
		Desc: prometheus.NewDesc(
			"aws_rds_active_transactions_average",
			"ActiveTransactions",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "BinLogDiskUsage",
		Desc: prometheus.NewDesc(
			"aws_rds_bin_log_disk_usage_average",
			"The amount of disk space occupied by binary logs on the master. Applies to MySQL read replicas. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "BlockedTransactions",
		Desc: prometheus.NewDesc(
			"aws_rds_blocked_transactions_average",
			"BlockedTransactions",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "BufferCacheHitRatio",
		Desc: prometheus.NewDesc(
			"aws_rds_buffer_cache_hit_ratio_average",
			"BufferCacheHitRatio",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "BurstBalance",
		Desc: prometheus.NewDesc(
			"aws_rds_burst_balance_average",
			"The percent of General Purpose SSD (gp2) burst-bucket I/O credits available. Units: Percent",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUCreditBalance",
		Desc: prometheus.NewDesc(
			"aws_rds_cpu_credit_balance_average",
			"[T2 instances] The number of CPU credits available for the instance to burst beyond its base CPU utilization. Credits are stored in the credit balance after they are earned and removed from the credit balance after they expire. Credits expire 24 hours after they are earned. CPU credit metrics are available only at a 5 minute frequency. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUCreditUsage",
		Desc: prometheus.NewDesc(
			"aws_rds_cpu_credit_usage_average",
			"[T2 instances] The number of CPU credits consumed by the instance. One CPU credit equals one vCPU running at 100% utilization for one minute or an equivalent combination of vCPUs, utilization, and time (for example, one vCPU running at 50% utilization for two minutes or two vCPUs running at 25% utilization for two minutes). CPU credit metrics are available only at a 5 minute frequency. If you specify a period greater than five minutes, use the Sum statistic instead of the Average statistic. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUUtilization",
		Desc: prometheus.NewDesc(
			"node_cpu_average",
			"The percentage of CPU utilization. Units: Percent",
			[]string{"instance", "region"},
			map[string]string{"cpu": "All", "mode": "total"},
		),
	},
	{
		Name: "CommitLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_commit_latency_average",
			"CommitLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CommitThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_commit_throughput_average",
			"CommitThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DDLLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_ddl_latency_average",
			"DDLLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DDLThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_ddl_throughput_average",
			"DDLThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DMLLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_dml_latency_average",
			"DMLLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DMLThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_dml_throughput_average",
			"DMLThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DatabaseConnections",
		Desc: prometheus.NewDesc(
			"aws_rds_database_connections_average",
			"The number of database connections in use. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "Deadlocks",
		Desc: prometheus.NewDesc(
			"aws_rds_deadlocks_average",
			"Deadlocks",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DeleteLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_delete_latency_average",
			"DeleteLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DeleteThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_delete_throughput_average",
			"DeleteThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "DiskQueueDepth",
		Desc: prometheus.NewDesc(
			"aws_rds_disk_queue_depth_average",
			"The number of outstanding IOs (read/write requests) waiting to access the disk. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "EngineUptime",
		Desc: prometheus.NewDesc(
			"node_boot_time",
			"EngineUptime",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "FreeLocalStorage",
		Desc: prometheus.NewDesc(
			"aws_rds_free_local_storage_average",
			"FreeLocalStorage",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "FreeStorageSpace",
		Desc: prometheus.NewDesc(
			"node_filesystem_free",
			"The amount of available storage space. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "FreeableMemory",
		Desc: prometheus.NewDesc(
			"node_memory_Cached",
			"The amount of available random access memory. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "InsertLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_insert_latency_average",
			"InsertLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "InsertThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_insert_throughput_average",
			"InsertThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "LoginFailures",
		Desc: prometheus.NewDesc(
			"aws_rds_login_failures_average",
			"LoginFailures",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "NetworkReceiveThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_network_receive_throughput_average",
			"The incoming (Receive) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication. Units: Bytes/second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "NetworkThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_network_throughput_average",
			"NetworkThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "NetworkTransmitThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_network_transmit_throughput_average",
			"The outgoing (Transmit) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication. Units: Bytes/second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "Queries",
		Desc: prometheus.NewDesc(
			"aws_rds_queries_average",
			"Queries",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReadIOPS",
		Desc: prometheus.NewDesc(
			"aws_rds_read_iops_average",
			"The average number of disk I/O operations per second. Units: Count/Second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReadLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_read_latency_average",
			"The average amount of time taken per disk I/O operation. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReadThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_read_throughput_average",
			"The average number of bytes read from disk per second. Units: Bytes/Second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReplicaLag",
		Desc: prometheus.NewDesc(
			"aws_rds_replica_lag_average",
			"The amount of time a Read Replica DB instance lags behind the source DB instance. Applies to MySQL, MariaDB, and PostgreSQL Read Replicas. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ResultSetCacheHitRatio",
		Desc: prometheus.NewDesc(
			"aws_rds_result_set_cache_hit_ratio_average",
			"ResultSetCacheHitRatio",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "SelectLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_select_latency_average",
			"SelectLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "SelectThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_select_throughput_average",
			"SelectThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "SwapUsage",
		Desc: prometheus.NewDesc(
			"aws_rds_swap_usage_average",
			"The amount of swap space used on the DB instance. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "UpdateLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_update_latency_average",
			"UpdateLatency",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "UpdateThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_update_throughput_average",
			"UpdateThroughput",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "VolumeBytesUsed",
		Desc: prometheus.NewDesc(
			"aws_rds_volume_bytes_used_average",
			"VolumeBytesUsed",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "VolumeReadIOPs",
		Desc: prometheus.NewDesc(
			"aws_rds_volume_read_io_ps_average",
			"VolumeReadIOPs",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "VolumeWriteIOPs",
		Desc: prometheus.NewDesc(
			"aws_rds_volume_write_io_ps_average",
			"VolumeWriteIOPs",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "WriteIOPS",
		Desc: prometheus.NewDesc(
			"aws_rds_write_iops_average",
			"The average number of disk I/O operations per second. Units: Count/Second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "WriteLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_write_latency_average",
			"The average amount of time taken per disk I/O operation. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "WriteThroughput",
		Desc: prometheus.NewDesc(
			"aws_rds_write_throughput_average",
			"The average number of bytes written to disk per second. Units: Bytes/Second",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
}
