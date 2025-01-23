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
			--go_out=../gen/go \
			--go-grpc_out=../gen/go \
			--go-grpc_opt=paths=source_relative \
			$$f; \
	done)

	cd protobuf && buf generate

	cd protobuf && buf build -o pbapi/pbapi.binpb
	cd protobuf && npx buf generate --template buf.gen.npx.yaml

	./src/mdgen/main.py

.PHONY: more
more:
	cd protobuf && npx buf generate --template buf.gen.graphql.yaml
	go run github.com/99designs/gqlgen generate
