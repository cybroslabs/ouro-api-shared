.PHONY: all
all: gen-go

.PHONY: setup
setup:
	npm i

.PHONY: gen-go
gen-go:
	npm i

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdeviceregistrymodels \
		--go-grpc_out=./pbdeviceregistrymodels \
		--go-grpc_opt=paths=source_relative \
		pbdeviceregistry-models.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdrivermodels \
		--go-grpc_out=./pbdrivermodels \
		--go-grpc_opt=paths=source_relative \
		pbdriver-models.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbtaskmastermodels \
		--go-grpc_out=./pbtaskmastermodels \
		--go-grpc_opt=paths=source_relative \
		pbtaskmaster-models.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdataproxymodels \
		--go-grpc_out=./pbdataproxymodels \
		--go-grpc_opt=paths=source_relative \
		pbdataproxy-models.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdriveroperatormodels \
		--go-grpc_out=./pbdriveroperatormodels \
		--go-grpc_opt=paths=source_relative \
		pbdriveroperator-models.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdeviceregistry \
		--go-grpc_out=./pbdeviceregistry \
		--go-grpc_opt=paths=source_relative \
		pbdeviceregistry.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdriver \
		--go-grpc_out=./pbdriver \
		--go-grpc_opt=paths=source_relative \
		pbdriver.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbtaskmaster \
		--go-grpc_out=./pbtaskmaster \
		--go-grpc_opt=paths=source_relative \
		pbtaskmaster.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdataproxy \
		--go-grpc_out=./pbdataproxy \
		--go-grpc_opt=paths=source_relative \
		pbdataproxy.proto

	cd protobuf && protoc \
		--go_opt=default_api_level=API_OPAQUE \
		--go_opt=paths=source_relative \
		--go_out=./pbdriveroperator \
		--go-grpc_out=./pbdriveroperator \
		--go-grpc_opt=paths=source_relative \
		pbdriveroperator.proto

	cd protobuf && buf generate

	cd protobuf && buf build -o pbapi/pbapi.binpb
	cd protobuf && npx buf generate --template buf.gen.npx.yaml

	./src/mdgen/main.py

.PHONY: more
more:
	cd protobuf && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate
