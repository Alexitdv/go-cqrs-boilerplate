VERSION?=dev

.PHONY: test
test:
	go test ./...

.PHONY:lint
lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

.PHONY: build
build:
	./scripts/docker/build.sh ${VERSION}

.PHONY: push
push:
	./scripts/docker/push.sh ${VERSION}

publish: build push

.PHONY:run
run:
	./scripts/env.sh ${VERSION} "go run ./cmd/boilerplate/*.go"

.PHONY:gen
gen:
	./scripts/gen.sh

.PHONY:mod
mod:
	go mod tidy && go mod vendor

.PHONY: fmt
fmt:
	find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./internal/gen/*" -exec goimports -l -w {} \;
