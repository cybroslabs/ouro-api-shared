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
	find ./gen/go -depth -type d -name "docs" -exec rm -r {} \;
	find ./gen/go -type d -empty -delete
	find ./gen/python -depth -mindepth 1 -type d -delete -exec rm -r {} \;

	cd proto && buf dep update
	cd proto && npx buf generate --template buf.gen.grpc.yaml
	cd proto && npx buf generate --template buf.gen.api.yaml
	cd proto && buf build -o ../gen/go/services/svcapi/raw.binpb

	find ./gen/python -type d -mindepth 1 -maxdepth 1 -exec touch {}/__init__.py \;

	./src/mdgen/main.py

.PHONY: more
more:
	cd proto && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate

update:
	GOPROXY=direct GOPRIVATE=github.com go get -u ./...
	sed -i '' 's|github.com/google/cel-go v[^ ]*|github.com/google/cel-go v0.22.1|g' go.mod
	sed -i '' 's|github.com/bufbuild/protovalidate-go v[^ ]*|github.com/bufbuild/protovalidate-go v0.8.0|g' go.mod
	go mod tidy
