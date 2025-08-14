.PHONY: all
all: gen-go

.PHONY: setup
setup:
	npm install --save-dev @bufbuild/protoc-gen-es
	npm update --save
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go mod tidy

.PHONY: gen-go
gen-go:
	npm i

	rm -rf ./gen/ts
	find ./gen/go -type f -name "*.pb.go" -delete
	find ./gen/go -depth -type d -name "docs" -exec rm -r {} \;
	find ./gen/go -type d -empty -delete
	find ./gen/python/platform_api -depth -mindepth 1 -type d -exec rm -r {} \;

	cd proto && buf dep update
	cd proto && npx buf generate --template buf.gen.grpc.yaml
	cd proto && npx buf generate --template buf.gen.api.yaml
	cd proto && buf build -o ../gen/go/services/svcapi/raw.binpb

	./src/mdgen/main.py --mode html markdown

.PHONY: more
more:
	rm -rf ./graph
	rm -rf ./gen/markdown
	./src/mdgen/main.py --mode markdown
	cd proto && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate

update:
	GOPROXY=direct GOPRIVATE=github.com go get -u ./... && go mod tidy
