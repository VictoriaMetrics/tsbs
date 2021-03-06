package devops

import (
	"math/rand"
	"time"

	"github.com/hagen1778/tsbs/cmd/tsbs_generate_data/common"
	"github.com/hagen1778/tsbs/cmd/tsbs_generate_data/serialize"
)

var (
	labelCPU  = []byte("cpu") // heap optimization
	CPUFields = []labeledDistributionMaker{
		{[]byte("usage_user"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_system"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_idle"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_nice"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_iowait"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_irq"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_softirq"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_steal"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_guest"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
		{[]byte("usage_guest_nice"), func() common.Distribution { return common.CWD(cpuND, 0.0, 100.0, rand.Float64()*100.0) }},
	}
)

// Reuse NormalDistributions as arguments to other distributions. This is
// safe to do because the higher-level distribution advances the ND and
// immediately uses its value and saves the state
var cpuND = common.ND(0.0, 1.0)

type CPUMeasurement struct {
	*subsystemMeasurement
}

func NewCPUMeasurement(start time.Time) *CPUMeasurement {
	return newCPUMeasurementNumDistributions(start, len(CPUFields))
}

func newSingleCPUMeasurement(start time.Time) *CPUMeasurement {
	return newCPUMeasurementNumDistributions(start, 1)
}

func newCPUMeasurementNumDistributions(start time.Time, numDistributions int) *CPUMeasurement {
	sub := newSubsystemMeasurementWithDistributionMakers(start, CPUFields[:numDistributions])
	return &CPUMeasurement{sub}
}

func (m *CPUMeasurement) ToPoint(p *serialize.Point) {
	m.toPointAllInt64(p, labelCPU, CPUFields)
}
