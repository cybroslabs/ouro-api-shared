package containers

import (
	"fmt"
	"math"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

var (
	ErrInvalidResourceCpuValue = fmt.Errorf("invalid cpu value")
)

// Converts the container resources - gRPC to REST
func G2RContainerResources(in *ContainerResourcesSchema) (*pbdriver.ContainerResources, error) {
	if in == nil {
		return nil, nil
	}

	cpu := "1"
	if cpu_f, err := in.Cpu.AsContainerResourcesCpuIntSchema(); err == nil {
		if math.Mod(float64(cpu_f), 10) != 0 {
			return nil, ErrInvalidResourceCpuValue
		}
		if math.Mod(float64(cpu_f), 1) != 0 {
			cpu = fmt.Sprintf("%fm", cpu_f*1000)
		} else {
			cpu = fmt.Sprintf("%f", cpu_f)
		}
	} else if cpu_s, err := in.Cpu.AsContainerResourcesCpuStrSchema(); err == nil {
		cpu = cpu_s
	}
	return &pbdriver.ContainerResources{
		Cpu:    cpu,
		Memory: in.Memory,
	}, nil
}
