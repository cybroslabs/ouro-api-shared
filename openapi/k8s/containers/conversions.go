package containers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"regexp"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmaster"
)

var (
	ErrInvalidResourceCpuValue    = fmt.Errorf("invalid cpu value")
	ErrInvalidResourceMemoryValue = fmt.Errorf("invalid memory value")

	reCpuResource    = regexp.MustCompile(`^[0-9]+m?$`)
	reMemoryResource = regexp.MustCompile(`^[0-9]+[EPTGMk]i?$`)
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
		if math.Mod(float64(cpu_f), 10) != 0 {
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
		if !reCpuResource.MatchString(cpu_s) {
			return nil, ErrInvalidResourceCpuValue
		}
		cpu = cpu_s
	}

	if reMemoryResource.MatchString(in.Memory) {
		return nil, ErrInvalidResourceMemoryValue
	}

	return &pbtaskmaster.ContainerResources{
		Cpu:    cpu,
		Memory: in.Memory,
	}, nil
}
