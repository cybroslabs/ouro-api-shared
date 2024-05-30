package containers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmaster"

	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	ErrInvalidResourceCpuValue    = fmt.Errorf("invalid cpu value")
	ErrInvalidResourceMemoryValue = fmt.Errorf("invalid memory value")
)

// Converts the container resources - gRPC to REST
func G2RContainerResources(in *ContainerResourcesSchema) (*pbtaskmaster.ContainerResources, error) {
	if in == nil {
		return nil, nil
	}

	var cpu string

	var cpu_f ContainerResourcesCpuIntSchema
	d := json.NewDecoder(bytes.NewReader(in.Cpu.union))
	d.UseNumber()
	err := d.Decode(&cpu_f)
	if err == nil {
		// 0.01 is the minimum scale for CPU in our case
		if math.Mod(float64(cpu_f)*100, 1) != 0 {
			return nil, ErrInvalidResourceCpuValue
		}
		if math.Mod(float64(cpu_f), 1) != 0 {
			cpu = fmt.Sprintf("%fm", cpu_f*1000)
		} else {
			cpu = fmt.Sprintf("%f", cpu_f)
		}
	} else {
		var cpu_s ContainerResourcesCpuStrSchema
		d := json.NewDecoder(bytes.NewReader(in.Cpu.union))
		d.UseNumber()
		err := d.Decode(&cpu_s)
		if err != nil {
			return nil, ErrInvalidResourceCpuValue
		}
		_, err = resource.ParseQuantity(cpu_s)
		if err != nil {
			return nil, ErrInvalidResourceCpuValue
		}
		cpu = cpu_s
	}

	if _, err := resource.ParseQuantity(in.Memory); err != nil {
		return nil, ErrInvalidResourceMemoryValue
	}

	return &pbtaskmaster.ContainerResources{
		Cpu:    cpu,
		Memory: in.Memory,
	}, nil
}
