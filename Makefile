.PHONY: all
all: setup gen-go

.PHONY: setup
setup:
	npm i
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: gen-go
gen-go:
	npm i

	find ./gen/go -type f ! -name "docs.go" -delete
	find ./gen/go -type d -empty -delete

	(cd proto && for f in `find . -name '*.proto'`; do \
		export dn=`dirname $$f`; \
		buf generate --template buf.gen.grpc.yaml --path $$f; \
	done)

	cd proto && npx buf generate --template buf.gen.api.yaml
	cd proto && npx buf generate --template buf.gen.grpc.yaml
	cd proto && buf build -o ../gen/go/services/api/api.binpb

	./src/mdgen/main.py

.PHONY: more
more:
	cd proto && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate
