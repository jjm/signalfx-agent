// Code generated by monitor-code-gen. DO NOT EDIT.

package metadata

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/signalfx-metadata"

var groupSet = map[string]bool{}

const (
	cpuUtilization         = "cpu.utilization"
	cpuUtilizationPerCore  = "cpu.utilization_per_core"
	diskSummaryUtilization = "disk.summary_utilization"
	diskUtilization        = "disk.utilization"
	diskOpsTotal           = "disk_ops.total"
	memoryUtilization      = "memory.utilization"
	networkTotal           = "network.total"
)

var metricSet = map[string]monitors.MetricInfo{
	cpuUtilization:         {Type: datapoint.Gauge},
	cpuUtilizationPerCore:  {Type: datapoint.Gauge},
	diskSummaryUtilization: {Type: datapoint.Gauge},
	diskUtilization:        {Type: datapoint.Gauge},
	diskOpsTotal:           {Type: datapoint.Counter},
	memoryUtilization:      {Type: datapoint.Gauge},
	networkTotal:           {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	cpuUtilization:         true,
	diskSummaryUtilization: true,
	diskUtilization:        true,
	diskOpsTotal:           true,
	memoryUtilization:      true,
	networkTotal:           true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/signalfx-metadata",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
