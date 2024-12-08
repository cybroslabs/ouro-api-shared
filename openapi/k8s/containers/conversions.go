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
	ErrInvalidResourceLimitsValue = fmt.Errorf("invalid limits value, limits cannot be nil")
	ErrInvalidResourceCpuValue    = fmt.Errorf("invalid cpu value")
	ErrInvalidResourceMemoryValue = fmt.Errorf("invalid memory value")
)

// Converts the container resources set - REST to gRPC
func R2GContainerResourcesSet(in *ContainerResourcesSetSchema) (result *pbtaskmaster.ContainerResourceSet, err error) {
	if in == nil {
		return
	}

	var limits *pbtaskmaster.ContainerResources
	var requests *pbtaskmaster.ContainerResources

	if limits, err = R2GContainerResources(&in.Limits); err != nil {
		return
	}
	if requests, err = R2GContainerResources(in.Requests); err != nil {
		return
	}

	result = &pbtaskmaster.ContainerResourceSet{
		Limits:   limits,
		Requests: requests,
	}

	return
}

// Converts the container resources set - gRPC to REST
func G2RContainerResourcesSet(in *pbtaskmaster.ContainerResourceSet) (result *ContainerResourcesSetSchema, err error) {
	if in == nil {
		return
	}
	if in.Limits == nil {
		return nil, ErrInvalidResourceLimitsValue
	}

	var limits *ContainerResourcesSchema
	var requests *ContainerResourcesSchema

	if limits, err = G2RContainerResources(in.Limits); err != nil {
		return
	}
	if requests, err = G2RContainerResources(in.Requests); err != nil {
		return
	}

	result = &ContainerResourcesSetSchema{
		Limits:   *limits,
		Requests: requests,
	}

	return
}

// Converts the container resources - REST to gRPC
func R2GContainerResources(in *ContainerResourcesSchema) (*pbtaskmaster.ContainerResources, error) {
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

// Converts the container resources - gRPC to REST
func G2RContainerResources(in *pbtaskmaster.ContainerResources) (result *ContainerResourcesSchema, err error) {
	if in == nil {
		return
	}

	result = &ContainerResourcesSchema{
		Memory: in.Memory,
	}
	err = result.Cpu.FromContainerResourcesCpuStrSchema(in.Cpu)

	return
}
