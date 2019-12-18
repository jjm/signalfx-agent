// Code generated by monitor-code-gen. DO NOT EDIT.

package dotnet

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "dotnet"

var groupSet = map[string]bool{}

const (
	netClrExceptionsNumExcepsThrownSec               = "net_clr_exceptions.num_exceps_thrown_sec"
	netClrLocksandthreadsContentionRateSec           = "net_clr_locksandthreads.contention_rate_sec"
	netClrLocksandthreadsCurrentQueueLength          = "net_clr_locksandthreads.current_queue_length"
	netClrLocksandthreadsNumOfCurrentLogicalThreads  = "net_clr_locksandthreads.num_of_current_logical_threads"
	netClrLocksandthreadsNumOfCurrentPhysicalThreads = "net_clr_locksandthreads.num_of_current_physical_threads"
	netClrMemoryNumBytesInAllHeaps                   = "net_clr_memory.num_bytes_in_all_heaps"
	netClrMemoryNumGcHandles                         = "net_clr_memory.num_gc_handles"
	netClrMemoryNumOfPinnedObjects                   = "net_clr_memory.num_of_pinned_objects"
	netClrMemoryNumTotalCommittedBytes               = "net_clr_memory.num_total_committed_bytes"
	netClrMemoryNumTotalReservedBytes                = "net_clr_memory.num_total_reserved_bytes"
	netClrMemoryPctTimeInGc                          = "net_clr_memory.pct_time_in_gc"
)

var metricSet = map[string]monitors.MetricInfo{
	netClrExceptionsNumExcepsThrownSec:               {Type: datapoint.Gauge},
	netClrLocksandthreadsContentionRateSec:           {Type: datapoint.Gauge},
	netClrLocksandthreadsCurrentQueueLength:          {Type: datapoint.Gauge},
	netClrLocksandthreadsNumOfCurrentLogicalThreads:  {Type: datapoint.Gauge},
	netClrLocksandthreadsNumOfCurrentPhysicalThreads: {Type: datapoint.Gauge},
	netClrMemoryNumBytesInAllHeaps:                   {Type: datapoint.Gauge},
	netClrMemoryNumGcHandles:                         {Type: datapoint.Gauge},
	netClrMemoryNumOfPinnedObjects:                   {Type: datapoint.Gauge},
	netClrMemoryNumTotalCommittedBytes:               {Type: datapoint.Gauge},
	netClrMemoryNumTotalReservedBytes:                {Type: datapoint.Gauge},
	netClrMemoryPctTimeInGc:                          {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "dotnet",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           true,
}
