.PHONY: build test unittest lint clean update fmt docker run

GO=CGO_ENABLED=1 GO111MODULE=on go
ARCH=$(shell uname -m)

MICROSERVICE=martian

.PHONY: build test clean fmt docker run


build:
	$(GO) build $(GOFLAGS) -o $(MICROSERVICE)

tidy:
	go mod tidy

t:
	[ -z "$$(gofmt -p -l . || echo 'err')" ]

unittest:
	$(GO) test ./... -coverprofile=coverage.out ./...

lint:
	@which golangci-lint >/dev/null || echo "WARNING: go linter not installed. To install, run\n  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.42.1"
	@if [ "z${ARCH}" = "zx86_64" ] && which golangci-lint >/dev/null ; then golangci-lint run --config .golangci.yml ; else echo "WARNING: Linting skipped (not on x86_64 or linter not installed)"; fi

test: unittest lint
	$(GO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]

clean:
	rm -f $(MICROSERVICE)

update:
	$(GO) mod download

fmt:
	$(GO) fmt ./...

#docker:
#	docker build \
#		--rm \
#		--build-arg http_proxy \
#		--build-arg https_proxy \
#			--label "git_sha=$(GIT_SHA)" \
#			-t edgexfoundry/misc:$(GIT_SHA) \
#			-t edgexfoundry/misc:$(APPVERSION)-dev \
#			.

run: build
	./$(MICROSERVICE)

vendor:
	$(GO) mod vendor