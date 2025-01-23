.PHONY: all
all: gen-go

.PHONY: setup
setup:
	npm i

.PHONY: gen-go
gen-go:
	npm i

	(rm -rf ./gen/go/* && cd protobuf && for f in `find . -name '*.proto'`; do \
		export dn=`dirname $$f`; \
		protoc \
			--go_opt=default_api_level=API_OPAQUE \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			$$f; \
	done)

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=../gen/go/services/deviceregistry \
		--go-grpc_out=../gen/go/services/deviceregistry \
		--go-grpc_opt=paths=source_relative \
		services/deviceregistry.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=../gen/go/services/driver \
		--go-grpc_out=../gen/go/services/driver \
		--go-grpc_opt=paths=source_relative \
		services/driver.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=../gen/go/services/taskmaster \
		--go-grpc_out=../gen/go/services/taskmaster \
		--go-grpc_opt=paths=source_relative \
		services/taskmaster.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=../gen/go/services/dataproxy \
		--go-grpc_out=../gen/go/services/dataproxy \
		--go-grpc_opt=paths=source_relative \
		services/dataproxy.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=../gen/go/services/driveroperator \
		--go-grpc_out=../gen/go/services/driveroperator \
		--go-grpc_opt=paths=source_relative \
		services/driveroperator.proto

	cd protobuf && buf generate

	cd protobuf && buf build -o pbapi/pbapi.binpb
	cd protobuf && npx buf generate --template buf.gen.npx.yaml

	./src/mdgen/main.py

.PHONY: more
more:
	cd protobuf && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate
