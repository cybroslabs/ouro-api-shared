package helpers

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcdataproxy"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcdeviceregistry"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svcourooperator"
	"github.com/cybroslabs/ouro-api-shared/gen/go/services/svctaskmaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ConnectorsOpts contains the options for creating a Connectors instance.
// It includes the hostnames for the various services and custom gRPC options.
type ConnectorsOpts struct {
	// Hostname for the API service. If the value is empty, it will default to the environment variable OURO_API_GRPC_HOST.
	ApiHost string
	// Hostname for the taskmaster service. If the value is empty, it will default to the environment variable OURO_CORE_TASKMASTER_GRPC_HOST.
	TaskmasterHost string
	// Hostname for the data-proxy service. If the value is empty, it will default to the environment variable OURO_CORE_DATAPROXY_GRPC_HOST.
	DataproxyHost string
	// Hostname for the device registry service. If the value is empty, it will default to the environment variable OURO_CORE_DEVICEREGISTRY_GRPC_HOST.
	DeviceRegistryHost string
	// Hostname for the driver operator service. If the value is empty, it will default to the environment variable OURO_OPERATOR_GRPC_HOST.
	OuroOperatorHost string

	// Custom gRPC options for the data-proxy service.
	GrpcOptionsDataproxy []grpc.DialOption
	// Custom gRPC options for the taskmaster service.
	GrpcOptionsTaskmaster []grpc.DialOption
	// Custom gRPC options for the device registry service.
	GrpcOptionsDeviceRegistry []grpc.DialOption
	// Custom gRPC options for the driver operator service.
	GrpcOptionsOuroOperator []grpc.DialOption
}

// Connectors is an interface that provides methods to open gRPC connections to various services.
type Connectors interface {
	// OpenTaskmasterServiceClient opens a new gRPC connection to the taskmaster service.
	OpenTaskmasterServiceClient() (svctaskmaster.TaskmasterServiceClient, context.CancelFunc, error)
	// OpenDataproxyServiceClient opens a new gRPC connection to the data-proxy service.
	OpenDataproxyServiceClient() (svcdataproxy.DataproxyServiceClient, context.CancelFunc, error)
	// OpenDeviceRegistryServiceClient opens a new gRPC connection to the device registry service.
	OpenDeviceRegistryServiceClient() (svcdeviceregistry.DeviceRegistryServiceClient, context.CancelFunc, error)
	// OpenOuroOperatorServiceClient opens a new gRPC connection to the OuroOperator service.
	OpenOuroOperatorServiceClient() (svcourooperator.OuroOperatorServiceClient, context.CancelFunc, error)
}

type connectors struct {
	// Implements the Connectors interface.
	Connectors

	apiHost            string
	taskmasterHost     string
	dataproxyHost      string
	deviceRegistryHost string
	ouroOperatorHost   string

	grpcOptionsDataproxy      []grpc.DialOption
	grpcOptionsTaskmaster     []grpc.DialOption
	grpcOptionsDeviceRegistry []grpc.DialOption
	grpcOptionsOuroOperator   []grpc.DialOption
}

func NewConnectors(opts *ConnectorsOpts) (Connectors, error) {
	tokenBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return nil, fmt.Errorf("error reading service account token: %w", err)
	}

	grpcOptionsDataproxy := initGrpcOptions(opts.GrpcOptionsDataproxy)
	grpcOptionsTaskmaster := initGrpcOptions(opts.GrpcOptionsTaskmaster)
	grpcOptionsDeviceRegistry := initGrpcOptions(opts.GrpcOptionsDeviceRegistry)
	grpcOptionsOuroOperator := initGrpcOptions(opts.GrpcOptionsOuroOperator)

	// Attach our 'namespace' to all outgoing unary RPC requests for the data-proxy service
	grpcOptionsDataproxy = append(grpcOptionsDataproxy, grpc.WithUnaryInterceptor(grpcNamespaceInterceptor(string(tokenBytes))))
	grpcOptionsTaskmaster = append(grpcOptionsTaskmaster, grpc.WithUnaryInterceptor(grpcNamespaceInterceptor(string(tokenBytes))))
	grpcOptionsDeviceRegistry = append(grpcOptionsDeviceRegistry, grpc.WithUnaryInterceptor(grpcNamespaceInterceptor(string(tokenBytes))))
	grpcOptionsOuroOperator = append(grpcOptionsOuroOperator, grpc.WithUnaryInterceptor(grpcNamespaceInterceptor(string(tokenBytes))))

	return &connectors{
		apiHost:            emptyOr(opts.ApiHost, os.Getenv("OURO_API_GRPC_HOST")),
		taskmasterHost:     emptyOr(opts.TaskmasterHost, os.Getenv("OURO_CORE_TASKMASTER_GRPC_HOST")),
		dataproxyHost:      emptyOr(opts.DataproxyHost, os.Getenv("OURO_CORE_DATAPROXY_GRPC_HOST")),
		deviceRegistryHost: emptyOr(opts.DeviceRegistryHost, os.Getenv("OURO_CORE_DEVICEREGISTRY_GRPC_HOST")),
		ouroOperatorHost:   emptyOr(opts.OuroOperatorHost, os.Getenv("OURO_OPERATOR_GRPC_HOST")),

		grpcOptionsDataproxy:      grpcOptionsDataproxy,
		grpcOptionsTaskmaster:     grpcOptionsTaskmaster,
		grpcOptionsDeviceRegistry: grpcOptionsDeviceRegistry,
		grpcOptionsOuroOperator:   grpcOptionsOuroOperator,
	}, nil
}

func emptyOr(value string, valueFallback string) string {
	if value == "" {
		return valueFallback
	}
	return value
}

func initGrpcOptions(opts []grpc.DialOption) []grpc.DialOption {
	if opts == nil {
		return []grpc.DialOption{
			// Use insecure credentials by default if no options are provided.
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
	}
	return slices.Clone(opts)
}

// Open a new gRPC connection to the taskmaster service. The connection must be closed by calling func in the second return value.
func (ra *connectors) OpenTaskmasterServiceClient() (svctaskmaster.TaskmasterServiceClient, context.CancelFunc, error) {
	conn, err := grpc.NewClient(ra.taskmasterHost, ra.grpcOptionsTaskmaster...)
	if err != nil {
		return nil, nil, err
	}

	client := svctaskmaster.NewTaskmasterServiceClient(conn)
	return client, func() { _ = conn.Close() }, nil
}

// Open a new gRPC connection to the dataproxy service. The connection must be closed by calling func in the second return value.
func (ra *connectors) OpenDataproxyServiceClient() (svcdataproxy.DataproxyServiceClient, context.CancelFunc, error) {
	conn, err := grpc.NewClient(ra.dataproxyHost, ra.grpcOptionsDataproxy...)
	if err != nil {
		return nil, nil, err
	}

	client := svcdataproxy.NewDataproxyServiceClient(conn)
	return client, func() { _ = conn.Close() }, nil
}

// Open a new gRPC connection to the deviceregistry service. The connection must be closed by calling func in the second return value.
func (ra *connectors) OpenDeviceRegistryServiceClient() (svcdeviceregistry.DeviceRegistryServiceClient, context.CancelFunc, error) {
	conn, err := grpc.NewClient(ra.deviceRegistryHost, ra.grpcOptionsDeviceRegistry...)
	if err != nil {
		return nil, nil, err
	}

	client := svcdeviceregistry.NewDeviceRegistryServiceClient(conn)
	return client, func() { _ = conn.Close() }, nil
}

// Open a new gRPC connection to the OuroOperator service. The connection must be closed by calling func in the second return value.
func (ra *connectors) OpenOuroOperatorServiceClient() (svcourooperator.OuroOperatorServiceClient, context.CancelFunc, error) {
	conn, err := grpc.NewClient(ra.ouroOperatorHost, ra.grpcOptionsOuroOperator...)
	if err != nil {
		return nil, nil, err
	}

	client := svcourooperator.NewOuroOperatorServiceClient(conn)
	return client, func() { _ = conn.Close() }, nil
}
