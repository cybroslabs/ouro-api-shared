package pbauthentication

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// NewGrpcIncomingContext returns a new outgoing context with the subject and issuer in gRPC metadata.
func NewOutgoingContext(subject string, issuer string) context.Context {
	return NewOutgoingContextFromContext(context.Background(), subject, issuer)
}

// NewGrpcIncomingContextFromContext returns outgoing context derived from ctx with the subject and issuer in gRPC metadata.
func NewOutgoingContextFromContext(ctx context.Context, subject string, issuer string) context.Context {
	md := metadata.MD{}
	if len(subject) > 0 {
		md.Append("io-clbs-openhes-auth-sub", subject)
	}
	if len(issuer) > 0 {
		md.Append("io-clbs-openhes-auth-iss", issuer)
	}
	return metadata.NewOutgoingContext(ctx, md)
}

// GetGrpcAuthInfo returns the subject and issuer from incoming gRPC context.
func FromIncomingContext(ctx context.Context) (subject string, issuer string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ""
	}
	subjects := md.Get("io-clbs-openhes-auth-sub")
	if len(subjects) > 0 {
		subject = subjects[0]
	}
	issuers := md.Get("io-clbs-openhes-auth-iss")
	if len(issuers) > 0 {
		issuer = issuers[0]
	}
	return
}
