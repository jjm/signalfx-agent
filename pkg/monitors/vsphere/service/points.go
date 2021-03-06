package service

import (
	"strings"
	"time"

	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors/vsphere/model"
	"github.com/sirupsen/logrus"
	"github.com/vmware/govmomi/vim25/types"
)

type PointsSvc struct {
	log     *logrus.Entry
	gateway IGateway
}

func NewPointsSvc(gateway IGateway, log *logrus.Entry) *PointsSvc {
	return &PointsSvc{gateway: gateway, log: log}
}

// Retrieves datapoints for all of the inventory objects in the passed-in VsphereInfo for the number of 20-second
// intervals indicated by the passed-in numSamplesReqd. Also returns the most recent sample time for the returned points.
func (svc *PointsSvc) RetrievePoints(vsInfo *model.VsphereInfo, numSamplesReqd int32) ([]*datapoint.Datapoint, time.Time) {
	perf, err := svc.gateway.queryPerf(vsInfo.Inv.Objects, numSamplesReqd)
	if err != nil {
		svc.log.WithError(err).Error("queryPerf failed")
		return nil, time.Time{}
	}
	var latestSampleTime time.Time
	var dps []*datapoint.Datapoint
	for _, baseMetric := range perf.Returnval {
		perfEntityMetric, ok := baseMetric.(*types.PerfEntityMetric)
		if !ok {
			svc.log.WithField("baseMetric", baseMetric).Error("Type coersion to PerfEntityMetric failed")
			continue
		}
		if latestSampleTime.IsZero() {
			latestSampleTime = perfEntityMetric.SampleInfo[len(perfEntityMetric.SampleInfo)-1].Timestamp
		}
		for _, metric := range perfEntityMetric.Value {
			intSeries, ok := metric.(*types.PerfMetricIntSeries)
			if !ok {
				svc.log.WithField("metric", metric).Error("Type coersion to PerfMetricIntSeries failed")
				continue
			}

			metricInfo := vsInfo.PerfCounterIndex[intSeries.Id.CounterId]
			metricName := metricInfo.MetricName
			sfxMetricType := statsTypeToMetricType(metricInfo.PerfCounterInfo.StatsType)

			cachedDims, ok := vsInfo.Inv.DimensionMap[perfEntityMetric.Entity.Value]
			var dims map[string]string
			if !ok {
				dims = map[string]string{}
			} else {
				dims = copyMap(cachedDims)
			}

			if intSeries.Id.Instance != "" {
				// the vsphere UI calls this dimension 'Object'
				dims["object"] = intSeries.Id.Instance
			}

			if len(intSeries.Value) > 0 && intSeries.Value[0] > 0 {
				svc.log.Debugf(
					"metric = %s, type = (%s->%s), dims = %v, values = %v",
					metricName,
					metricInfo.PerfCounterInfo.StatsType,
					sfxMetricType,
					dims,
					intSeries.Value,
				)
			}

			for i, value := range intSeries.Value {
				var dpVal datapoint.Value
				if strings.HasSuffix(metricName, "_percent") {
					dpVal = datapoint.NewFloatValue(float64(value) / 100)
				} else {
					dpVal = datapoint.NewIntValue(value)
				}
				dps = append(dps, datapoint.New(
					metricName,
					dims,
					dpVal,
					sfxMetricType,
					perfEntityMetric.SampleInfo[i].Timestamp,
				))
			}
		}
	}
	return dps, latestSampleTime
}

func copyMap(in map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range in {
		out[k] = v
	}
	return out
}

func statsTypeToMetricType(statsType types.PerfStatsType) datapoint.MetricType {
	switch statsType {
	case types.PerfStatsTypeDelta:
		return datapoint.Count
	default:
		return datapoint.Gauge
	}
}
