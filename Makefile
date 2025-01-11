.PHONY: build api gen clean all

all: build api

build:
	go build -o ./bin/api-first-svc .

api:
	./bin/api-first-svc api

gen: gen_go gen_java

gen_go:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api-first-svc.yaml \
      --skip-validate-spec \
      -g go \
      -o /local/out/go

gen_java:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api-first-svc.yaml \
      --skip-validate-spec \
      -g java \
      -o /local/out/java

clean:
	rm -rf ./bin
	rm -rf ./out