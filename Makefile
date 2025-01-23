.PHONY: all
all: gen-go

.PHONY: setup
setup:
	npm i

.PHONY: gen-go
gen-go:
	npm i

	(rm -rf ./gen/go/* && cd proto && for f in `find . -name '*.proto'`; do \
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
