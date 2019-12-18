// Code generated by monitor-code-gen. DO NOT EDIT.

package jaegergrpc

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "jaeger-grpc"

var groupSet = map[string]bool{}

var metricSet = map[string]monitors.MetricInfo{}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "jaeger-grpc",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
