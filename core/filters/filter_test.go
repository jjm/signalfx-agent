package filters

import (
	"testing"

	"github.com/signalfx/golib/datapoint"
	"github.com/stretchr/testify/assert"
)

func TestFilters(t *testing.T) {
	t.Run("Exclude based on simple metric name", func(t *testing.T) {
		f := New("", []string{"cpu.utilization"}, nil)
		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "cpu.utilization"}))
		assert.False(t, f.Matches(&datapoint.Datapoint{Metric: "memory.utilization"}))
	})

	t.Run("Excludes based on multiple metric names", func(t *testing.T) {
		f := New("", []string{"cpu.utilization", "memory.utilization"}, nil)
		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "cpu.utilization"}))

		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "memory.utilization"}))

		assert.False(t, f.Matches(&datapoint.Datapoint{Metric: "disk.utilization"}))
	})

	t.Run("Excludes based on regex metric name", func(t *testing.T) {
		f := New("", []string{`/cpu\..*/`}, nil)
		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "cpu.utilization"}))

		assert.False(t, f.Matches(&datapoint.Datapoint{Metric: "disk.utilization"}))
	})

	t.Run("Excludes based on glob metric name", func(t *testing.T) {
		f := New("", []string{`cpu.util*`, "memor*"}, nil)
		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "cpu.utilization"}))
		assert.True(t, f.Matches(&datapoint.Datapoint{Metric: "memory.utilization"}))

		assert.False(t, f.Matches(&datapoint.Datapoint{Metric: "disk.utilization"}))
	})

	t.Run("Excludes based on dimension name", func(t *testing.T) {
		f := New("", nil, map[string][]string{
			"container_name": []string{"PO"},
		})

		assert.True(t, f.Matches(&datapoint.Datapoint{
			Metric: "cpu.utilization",
			Dimensions: map[string]string{
				"container_name": "PO",
			},
		}))

		assert.False(t, f.Matches(&datapoint.Datapoint{
			Metric: "disk.utilization",
			Dimensions: map[string]string{
				"container_name": "test",
			},
		}))
	})

	t.Run("Excludes based on dimension name regex", func(t *testing.T) {
		f := New("", nil, map[string][]string{
			"container_name": []string{`/^[A-Z][A-Z]$/`},
		})

		assert.True(t, f.Matches(&datapoint.Datapoint{
			Metric: "cpu.utilization",
			Dimensions: map[string]string{
				"container_name": "PO",
			},
		}))

		assert.False(t, f.Matches(&datapoint.Datapoint{
			Metric: "disk.utilization",
			Dimensions: map[string]string{
				"container_name": "test",
			},
		}))
	})

	t.Run("Excludes based on dimension name glob", func(t *testing.T) {
		f := New("", nil, map[string][]string{
			"container_name": []string{"mycontainer", `*O*`},
		})

		assert.True(t, f.Matches(&datapoint.Datapoint{
			Metric: "cpu.utilization",
			Dimensions: map[string]string{
				"container_name": "POD",
			},
		}))

		assert.False(t, f.Matches(&datapoint.Datapoint{
			Metric: "disk.utilization",
			Dimensions: map[string]string{
				"container_name": "test",
			},
		}))
	})

	t.Run("Excludes based on conjunction of both dimensions and metric name", func(t *testing.T) {
		f := New("", []string{"cpu.utilization"}, map[string][]string{
			"container_name": []string{"mycontainer", `*O*`},
		})

		assert.False(t, f.Matches(&datapoint.Datapoint{
			Metric: "cpu.utilization",
			Dimensions: map[string]string{
				"container_name": "not matching",
			},
		}))

		assert.False(t, f.Matches(&datapoint.Datapoint{
			Metric: "disk.utilization",
			Dimensions: map[string]string{
				"container_name": "test",
			},
		}))
	})
}