package services

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryInterceptor adds namespace metadata to all outgoing unary RPC requests
func grpcNamespaceInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Add namespace to metadata
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
		// Invoke the RPC
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
