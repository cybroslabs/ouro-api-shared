.PHONY: gen-go
gen-go:
	protoc \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		protobuf/pbdriver/pbdriver.proto

	protoc \
		--go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		protobuf/pbdataproxy/pbdataproxy.proto

.PHONY: generate
generate:
	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/k8s/containers/containers.go -package containers \
		./openapi/k8s/containers/containers.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/driver/driverdata/driverdata.go -package driverdata \
		./openapi/openhes/driver/driverdata/driverdata.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/job/job.go -package job \
		-import-mapping '"../driver/driverdata/driverdata.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"' \
		./openapi/openhes/job/job.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/driver/driver.go -package driver \
		-import-mapping '"../../openhes/driver/driverdata/driverdata.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata","../../k8s/containers/containers.yaml":"github.com/cybroslabs/hes-2-apis/openapi/k8s/containers","../job/job.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/job"' \
		./openapi/openhes/driver/driver.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/rfc/rfc7807/rfc7807.go -package rfc7807 \
		./openapi/rfc/rfc7807/rfc7807.yaml