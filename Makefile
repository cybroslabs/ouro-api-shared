.PHONY: all
all: gen-go

.PHONY: setup
setup:
	npm install --save-dev @bufbuild/protoc-gen-es
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: gen-go
gen-go:
	npm i

	rm -rf ./gen/ts
	find ./gen/go -type f -name "*.pb.go" -delete
	find ./gen/go -type d -name "docs" -delete
	find ./gen/go -type d -empty -delete

	cd proto && buf dep update
	cd proto && npx buf generate --template buf.gen.grpc.yaml
	cd proto && npx buf generate --template buf.gen.api.yaml
	cd proto && buf build -o ../gen/go/services/svcapi/raw.binpb

	./src/mdgen/main.py

.PHONY: more
more:
	cd proto && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate
