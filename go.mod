module github.com/cybroslabs/ouro-api-shared

go 1.23.8

toolchain go1.24.5

retract (
	v0.0.0-20240512071833-41b886c985c4
	v0.0.0-20240511215858-4449dad5324b
)

require (
	connectrpc.com/connect v1.18.1
	github.com/99designs/gqlgen v0.17.76
	github.com/google/uuid v1.6.0
	github.com/rmg/iso4217 v1.0.1
	github.com/vektah/gqlparser/v2 v2.5.30
	go.uber.org/zap v1.27.0
	google.golang.org/genproto v0.0.0-20250721164621-a45f3dfb1074
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
	k8s.io/utils v0.0.0-20250604170112-4c0f3b243397
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/urfave/cli/v2 v2.27.7 // indirect
	github.com/xrash/smetrics v0.0.0-20250705151800-55b8f293f342 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.26.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/tools v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250721164621-a45f3dfb1074 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
